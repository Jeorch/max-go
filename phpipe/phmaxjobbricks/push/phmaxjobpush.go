package maxjobpush

import (
	"fmt"
	"github.com/Jeorch/max-go/phmodel/maxjob"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmexcelhandle"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/colinmarc/hdfs"
	"github.com/hashicorp/go-uuid"
	"io"
	"net/http"
	"os"
)

type PHMaxJobPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHMaxJobPushBrick) Exec() error {
	var tmp maxjob.PHMaxJob = b.bk.Pr.(maxjob.PHMaxJob)
	var err error
	var cpaCsv string
	var gycCsv string
	var cpaDesName string
	var gycDesName string
	var notArrivalHospCsv string
	var notArrivalHospDesName string
	cpa := tmp.Args["cpa"].(string)
	gyc := tmp.Args["gyc"].(string)
	if cpa != "" {
		cpaCsv, notArrivalHospCsv, err = cpa2csv(cpa)
		cpaDesName, err = push2hdfs(cpaCsv)
		notArrivalHospDesName, err = push2hdfs(notArrivalHospCsv)
		tmp.Args["cpa"] = cpaDesName
		tmp.Args["not_arrival_hosp_file"] = notArrivalHospDesName
	}
	if err!=nil {
		return err
	}
	if gyc != "" {
		gycCsv, err = gyc2csv(gyc)
		gycDesName, err = push2hdfs(gycCsv)
		tmp.Args["gyc"] = gycDesName
	}
	b.BrickInstance().Pr = tmp
 	return err
}

func (b *PHMaxJobPushBrick) Prepare(pr interface{}) error {
	req := pr.(maxjob.PHMaxJob)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHMaxJobPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHMaxJobPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHMaxJobPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(maxjob.PHMaxJob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval maxjob.PHMaxJob = b.BrickInstance().Pr.(maxjob.PHMaxJob)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func cpa2csv(cpaFile string) (string, string, error) {
	var err error
	var cpa string
	var notArrivalHosp string
	localCpa := "resource/" + cpaFile
	cpa, err = bmexcelhandle.GenerateCSVFromXLSXFile(localCpa, 0)
	notArrivalHosp, err = bmexcelhandle.GenerateCSVFromXLSXFile(localCpa, 1)
	os.Remove(localCpa)
	return cpa, notArrivalHosp, err
}

func gyc2csv(gycFile string) (string, error) {
	var err error
	var gyc string
	localGyc := "resource/" + gycFile
	gyc, err = bmexcelhandle.GenerateCSVFromXLSXFile(localGyc, 0)
	os.Remove(localGyc)
	return gyc, err
}

func push2hdfs(localFile string) (string, error) {
	//localDir := "resource/" + localFile
	localDir := localFile
	fileDesName, _ := uuid.GenerateUUID()
	fileDesPath := "/workData/Client/"+ fileDesName
	fmt.Println(fileDesName)
	client,_ := hdfs.New("192.168.100.137:9000")
	err := client.CopyToRemote(localDir, fileDesPath)
	os.Remove(localDir)
	return fileDesName, err
}
