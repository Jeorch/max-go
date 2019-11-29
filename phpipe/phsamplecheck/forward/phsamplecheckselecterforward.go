package samplecheckforward

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/Jeorch/max-go/phmodel/samplecheck"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmredis"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHSampleCheckSelecterForwardBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHSampleCheckSelecterForwardBrick) Exec() error {
	tmp := b.BrickInstance().Pr.(samplecheck.SampleCheckSelecter)
	tmp = getAllYms(tmp)
	tmp = getAllMktForSingleJob(tmp)
	//tmp.GetAllYms().GetAllMkt()
	b.BrickInstance().Pr = tmp
	return nil
}

func (b *PHSampleCheckSelecterForwardBrick) Prepare(pr interface{}) error {
	req := pr.(samplecheck.SampleCheckSelecter)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHSampleCheckSelecterForwardBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHSampleCheckSelecterForwardBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHSampleCheckSelecterForwardBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(samplecheck.SampleCheckSelecter)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHSampleCheckSelecterForwardBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval samplecheck.SampleCheckSelecter = b.BrickInstance().Pr.(samplecheck.SampleCheckSelecter)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func getAllYms(scs samplecheck.SampleCheckSelecter) samplecheck.SampleCheckSelecter {
	client := bmredis.GetRedisClient()
	defer client.Close()
	ymLst, err := client.SMembers(scs.JobID + "ym").Result()
	if err != nil {
		panic("no yms found")
	}
	rst := make([]interface{}, 0)
	for _, v := range ymLst {
		rst = append(rst, v)
	}
	scs.YmList = rst
	return scs

}

func getAllMktForSingleJob(scs samplecheck.SampleCheckSelecter) samplecheck.SampleCheckSelecter {

	var err error
	rst := make([]interface{}, 0)

	client := bmredis.GetRedisClient()
	defer client.Close()
	mktLst, err := client.SMembers(scs.JobID + "mkt").Result()
	if err == nil && len(mktLst) != 0 {
		for _, m := range mktLst {
			rst = append(rst, m)
		}
		scs.MktList = rst
		return scs
	}

	req := request.Request{}
	req.Res = "PhCalcConf"
	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = scs.CompanyID
	var condi1 []interface{}
	condi1 = append(condi1, eq)
	req = req.SetConnect("Eqcond", condi1).(request.Request)

	var reval []max.PhCalcConf
	err = bmmodel.FindMutil(req, &reval)
	if err != nil {
		panic("no PhCalcConf found")
	}

	for _, v := range reval {
		rst = append(rst, v.Mkt)
	}
	scs.MktList = rst
	return scs
}
