package maxactionpush

import (
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/PharbersDeveloper/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PhCompanyProdPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhCompanyProdPushBrick) Exec() error {
	tmp := b.bk.Pr.(max.PhCompanyProd)
	err := tmp.InsertBMObject()
	b.bk.Pr = tmp
	return err
}

func (b *PhCompanyProdPushBrick) Prepare(pr interface{}) error {
	req := pr.(max.PhCompanyProd)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhCompanyProdPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhCompanyProdPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhCompanyProdPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PhCompanyProd)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhCompanyProdPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PhCompanyProd = b.BrickInstance().Pr.(max.PhCompanyProd)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
