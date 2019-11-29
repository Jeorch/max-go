package samplecheckforward

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/Jeorch/max-go/phmodel/samplecheck"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PhSelecterForMarketBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhSelecterForMarketBrick) Exec() error {
	tmp := b.BrickInstance().Pr.(samplecheck.SampleCheckSelecter)
	tmp = getAllMkt(tmp)
	b.BrickInstance().Pr = tmp
	return nil
}

func (b *PhSelecterForMarketBrick) Prepare(pr interface{}) error {
	req := pr.(samplecheck.SampleCheckSelecter)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhSelecterForMarketBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhSelecterForMarketBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhSelecterForMarketBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(samplecheck.SampleCheckSelecter)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhSelecterForMarketBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval samplecheck.SampleCheckSelecter = b.BrickInstance().Pr.(samplecheck.SampleCheckSelecter)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func getAllMkt(scs samplecheck.SampleCheckSelecter) samplecheck.SampleCheckSelecter {

	var err error
	rst := make([]interface{}, 0)

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
