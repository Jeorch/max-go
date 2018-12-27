package companypush

import (
	"github.com/Jeorch/max-go/phmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHAuthCompanyPush struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthCompanyPush) Exec() error {
	var tmp auth.PhAuth = b.bk.Pr.(auth.PhAuth)
	profile_tmp := tmp.Profile
	company_tmp := profile_tmp.Company
	if company_tmp.Id != "" && company_tmp.Id_.Valid() {
		if company_tmp.Valid() && company_tmp.IsCompanyRegisted() {
			//b.bk.Err = -101
		} else {
			company_tmp.InsertBMObject()
		}
	}
	b.bk.Pr = tmp
	return nil
}

func (b *PHAuthCompanyPush) Prepare(pr interface{}) error {
	req := pr.(auth.PhAuth)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHAuthCompanyPush) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthCompanyPush) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthCompanyPush) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PhAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHAuthCompanyPush) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PhAuth = b.BrickInstance().Pr.(auth.PhAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
