package maxactionfind

import (
	"github.com/Jeorch/max-go/phmodel/max"
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

type PhMaxActionCalcBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhMaxActionCalcBrick) Exec() error {
	var tmp max.PhAction = b.bk.Pr.(max.PhAction)
	var err error

	tmp, err = generateCalcConf(tmp)

	b.BrickInstance().Pr = tmp
	return err
}

func (b *PhMaxActionCalcBrick) Prepare(pr interface{}) error {
	req := pr.(max.PhAction)

	b.BrickInstance().Pr = req
	return nil
}

func (b *PhMaxActionCalcBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhMaxActionCalcBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhMaxActionCalcBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PhAction)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhMaxActionCalcBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PhAction = b.BrickInstance().Pr.(max.PhAction)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func generateCalcConf(paction max.PhAction) (max.PhAction, error) {

	req := request.Request{}
	req.Res = "PhCalcConf"

	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = paction.CompanyId

	incond := request.Incond{}
	incond.Ky = "ym"
	//incond.Vy = paction.Yms

	var condi1 []interface{}
	var condi2 []interface{}
	condi1 = append(condi1, eq)
	condi2 = append(condi2, incond)

	req.SetConnect("Eqcond", condi1)
	req.SetConnect("Incond", condi2)

	var reval []max.PhCalcConf
	err := bmmodel.FindMutil(req, &reval)
	paction.CalcConf = reval

	return paction, err
}