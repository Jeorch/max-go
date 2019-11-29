package maxjobpush

import (
	"fmt"
	"github.com/Jeorch/max-go/phmodel/config"
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

type PHMaxJobPushPanelBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHMaxJobPushPanelBrick) Exec() error {
	var tmp max.Phmaxjob = b.bk.Pr.(max.Phmaxjob)
	var err error

	fmt.Println(tmp)
	panelcsv, err := panel2csv(tmp.Panel)
	panelhdfs, err := push2hdfs(panelcsv)
	if err != nil {
		return err
	}
	tmp.Panel = panelhdfs
	err = push2redis(tmp)

	b.BrickInstance().Pr = tmp
	return err
}

func (b *PHMaxJobPushPanelBrick) Prepare(pr interface{}) error {
	req := pr.(max.Phmaxjob)

	b.BrickInstance().Pr = req
	return nil
}

func (b *PHMaxJobPushPanelBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHMaxJobPushPanelBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHMaxJobPushPanelBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.Phmaxjob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobPushPanelBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.Phmaxjob = b.BrickInstance().Pr.(max.Phmaxjob)
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

func push2hdfs(localFile string) (string, error) {
	var maxHdfs maxconfig.PhHdfsConfig
	maxHdfs.GenerateConfig()
	hdfsAddress := maxHdfs.Host + ":" + maxHdfs.Port
	client, _ := hdfs.New(hdfsAddress)
	localDir := localFile
	fileDesName, _ := uuid.GenerateUUID()
	fmt.Println(fileDesName)
	fileDesPath := "/workData/Panel/" + fileDesName
	err := client.CopyToRemote(localDir, fileDesPath)
	os.Remove(localDir)
	return fileDesName, err
}

func push2redis(maxjob max.Phmaxjob) error {

	client := bmredis.GetRedisClient()
	defer client.Close()

	client.SAdd(maxjob.JobID, maxjob.Panel)
	client.Expire(maxjob.JobID, 24*time.Hour)
	client.SAdd(maxjob.JobID+"ym", maxjob.Yms)
	client.Expire(maxjob.JobID+"ym", 24*time.Hour)
	client.SAdd(maxjob.JobID+"mkt", maxjob.PanelMkt)
	client.Expire(maxjob.JobID+"mkt", 24*time.Hour)
	client.HSet(maxjob.Panel, "ym", maxjob.Yms)
	client.HSet(maxjob.Panel, "mkt", maxjob.PanelMkt)
	client.Expire(maxjob.Panel, 24*time.Hour)

	return nil
}
