package profilepush

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

type PHAuthProfilePush struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthProfilePush) Exec() error {
	var tmp auth.PHAuth = b.bk.Pr.(auth.PHAuth)
	profile_tmp := tmp.Profile
	if profile_tmp.Id != "" && profile_tmp.Id_.Valid() {
		if profile_tmp.Valid() && profile_tmp.IsUserRegisted() {
			b.bk.Err = -101
		} else {
			profile_tmp.InsertBMObject()
		}
	}
	b.bk.Pr = tmp
	return nil
}

func (b *PHAuthProfilePush) Prepare(pr interface{}) error {
	req := pr.(auth.PHAuth)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHAuthProfilePush) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthProfilePush) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthProfilePush) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PHAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHAuthProfilePush) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PHAuth = b.BrickInstance().Pr.(auth.PHAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
