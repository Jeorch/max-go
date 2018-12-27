package maxactionpush

import (
	"fmt"
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmexcelhandle"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmredis"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/colinmarc/hdfs"
	"github.com/hashicorp/go-uuid"
	"io"
	"net/http"
	"os"
	"time"
)

type PhMaxActionPushPanelBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhMaxActionPushPanelBrick) Exec() error {
	var tmp max.PhAction = b.bk.Pr.(max.PhAction)
	var err error

	fmt.Println(tmp)
	panelcsv, err := panel2csv(tmp.PanelPath)
	panelhdfs, err := pushpanel2hdfs(panelcsv)
	if err != nil {
		return err
	}
	tmp.PanelPath = panelhdfs
	panelConf := tmp.PanelConf
	panelConf[0].PanelName = panelhdfs
	err = push2redis(tmp)

	b.BrickInstance().Pr = tmp
	return err
}

func (b *PhMaxActionPushPanelBrick) Prepare(pr interface{}) error {
	req := pr.(max.PhAction)

	b.BrickInstance().Pr = req
	return nil
}

func (b *PhMaxActionPushPanelBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhMaxActionPushPanelBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhMaxActionPushPanelBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PhAction)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhMaxActionPushPanelBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PhAction = b.BrickInstance().Pr.(max.PhAction)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func panel2csv(panelFile string) (string, error) {
	var err error
	var panel string
	var bmRouter bmconfig.BMRouterConfig
	bmRouter.GenerateConfig()
	localFile := bmRouter.TmpDir + "/" + panelFile
	panel, err = bmexcelhandle.GenerateCSVFromXLSXFile(localFile, 0)
	os.Remove(localFile)
	return panel, err
}

func pushpanel2hdfs(localFile string) (string, error) {
	localDir := localFile
	fileDesName, _ := uuid.GenerateUUID()
	fmt.Println(fileDesName)
	fileDesPath := "/workData/Panel/" + fileDesName
	client, _ := hdfs.New("192.168.100.137:9000")
	err := client.CopyToRemote(localDir, fileDesPath)
	os.Remove(localDir)
	return fileDesName, err
}

func push2redis(maxjob max.PhAction) error {

	client := bmredis.GetRedisClient()
	defer client.Close()

	client.SAdd(maxjob.JobId, maxjob.PanelPath)
	client.Expire(maxjob.JobId, 24*time.Hour)
	//TODO：目前仅支持单月份单市场的panel文件.
	client.SAdd(maxjob.JobId+"ym", maxjob.PanelConf[0].Ym)
	client.Expire(maxjob.JobId+"ym", 24*time.Hour)
	client.HSet(maxjob.PanelPath, "ym", maxjob.PanelConf[0].Ym)
	client.HSet(maxjob.PanelPath, "mkt", maxjob.PanelConf[0].Mkt)
	client.Expire(maxjob.PanelPath, 24*time.Hour)

	return nil
}
