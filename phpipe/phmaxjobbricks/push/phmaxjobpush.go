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
	"time"
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

	//err = tmp.CheckJobIdCall()
	//if err == nil {
	//	return errors.New("Duplicated push job!")
	//}

	cpa := tmp.Cpa
	gyc := tmp.Gycx
	if cpa != "" {
		cpaCsv, notArrivalHospCsv, err = cpa2csv(cpa)
		cpaDesName, err = push2hdfs(cpaCsv)
		notArrivalHospDesName, err = push2hdfs(notArrivalHospCsv)
		tmp.Cpa = cpaDesName
		tmp.NotArrivalHospFile = notArrivalHospDesName
	}
	if err != nil {
		return err
	}
	if gyc != "" {
		gycCsv, err = gyc2csv(gyc)
		gycDesName, err = push2hdfs(gycCsv)
		tmp.Gycx = gycDesName
	}
	//tmp.Id = tmp.JobID
	tmp.Date = time.Now().String()
	b.BrickInstance().Pr = tmp

	return err
}

func (b *PHMaxJobPushBrick) Prepare(pr interface{}) error {
	req := pr.(maxjob.PHMaxJob)

	//err := req.CheckJobIdCall()
	//if err == nil {
	//	req.PushJobIdCall()
	//}

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
	localDir := localFile
	fileDesName, _ := uuid.GenerateUUID()
	fmt.Println(fileDesName)
	fileDesPath := "/workData/Client/" + fileDesName
	client, _ := hdfs.New("192.168.100.137:9000")
	err := client.CopyToRemote(localDir, fileDesPath)
	os.Remove(localDir)
	return fileDesName, err
}

//func (b *PHMaxJobPushBrick) pushSync(wg *sync.WaitGroup, m *sync.Mutex) error {
//	m.Lock()
//
//	var tmp maxjob.PHMaxJob = b.bk.Pr.(maxjob.PHMaxJob)
//	var err error
//	var cpaCsv string
//	var gycCsv string
//	var cpaDesName string
//	var gycDesName string
//	var notArrivalHospCsv string
//	var notArrivalHospDesName string
//
//	err = tmp.CheckJobIdCall()
//	if err != nil {
//		return err
//	}
//	cpa := tmp.Cpa
//	gyc := tmp.Gycx
//	if cpa != "" {
//		cpaCsv, notArrivalHospCsv, err = cpa2csv(cpa)
//		cpaDesName, err = push2hdfs(cpaCsv)
//		notArrivalHospDesName, err = push2hdfs(notArrivalHospCsv)
//		tmp.Cpa = cpaDesName
//		tmp.NotArrivalHospFile = notArrivalHospDesName
//	}
//	if err!=nil {
//		return err
//	}
//	if gyc != "" {
//		gycCsv, err = gyc2csv(gyc)
//		gycDesName, err = push2hdfs(gycCsv)
//		tmp.Gycx = gycDesName
//	}
//	//tmp.Id = tmp.JobID
//	tmp.Date = time.Now().String()
//	b.BrickInstance().Pr = tmp
//	tmp.PushJobIdCall()
//	m.Unlock()
//	wg.Done()
//	return err
//}
