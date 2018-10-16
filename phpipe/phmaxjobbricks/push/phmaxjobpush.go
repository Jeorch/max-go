package maxjobpush

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHMaxJobPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHMaxJobPushBrick) Exec() error {
	var tmp max.PHMaxJob = b.bk.Pr.(max.PHMaxJob)
	var err error

	//var cpaCsv string
	//var gycCsv string
	//var cpaDesName string
	//var gycDesName string
	//var notArrivalHospCsv string
	//var notArrivalHospDesName string
	//
	//cpa := tmp.Cpa
	//gyc := tmp.Gycx
	//if cpa != "" {
	//	cpaCsv, notArrivalHospCsv, err = cpa2csv(cpa)
	//	cpaDesName, err = push2hdfs(cpaCsv)
	//	notArrivalHospDesName, err = push2hdfs(notArrivalHospCsv)
	//	tmp.Cpa = cpaDesName
	//	tmp.NotArrivalHospFile = notArrivalHospDesName
	//}
	//if err != nil {
	//	return err
	//}
	//if gyc != "" {
	//	gycCsv, err = gyc2csv(gyc)
	//	gycDesName, err = push2hdfs(gycCsv)
	//	tmp.Gycx = gycDesName
	//}
	////TODO：JobID的问题
	//tmp.Id = tmp.JobID
	//tmp.Date = time.Now().String()

	err = tmp.CheckJobIdCall()
	b.BrickInstance().Pr = tmp
	return err
}

func (b *PHMaxJobPushBrick) Prepare(pr interface{}) error {
	req := pr.(max.PHMaxJob)

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
	tmp := pr.(max.PHMaxJob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PHMaxJob = b.BrickInstance().Pr.(max.PHMaxJob)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

//func (b *PHMaxJobPushBrick) pushSync(wg *sync.WaitGroup, m *sync.Mutex) error {
//	m.Lock()
//
//	var tmp max.PHMaxJob = b.bk.Pr.(max.PHMaxJob)
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
