package authothers

import (
	"crypto/md5"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/bmrouter/bmoauth"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/Jeorch/max-go/bmmodel/auth"
	"io"
	"net/http"
)

type PHAuthGenerateToken struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthGenerateToken) Exec() error {

	tmp := b.BrickInstance().Pr
	bmah := tmp.(auth.PHAuth)
	h := md5.New()
	io.WriteString(h, bmah.Id)

	token := fmt.Sprintf("%x", h.Sum(nil))

	bmah.Token = token
	b.BrickInstance().Pr = bmah

	err := bmoauth.PushToken(token)

	return err
}

func (b *PHAuthGenerateToken) Prepare(pr interface{}) error {
	req := pr.(auth.PHAuth)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHAuthGenerateToken) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthGenerateToken) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthGenerateToken) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PHAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHAuthGenerateToken) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PHAuth = b.BrickInstance().Pr.(auth.PHAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
