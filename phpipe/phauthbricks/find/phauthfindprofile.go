package authfind

import (
	"fmt"
	"github.com/Jeorch/max-go/phmodel/auth"
	"github.com/Jeorch/max-go/phmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
	"reflect"
)

type PHAuthFindProfileBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthFindProfileBrick) Exec() error {
	var tmp profile.PHProfile
	err := tmp.FindOne(*b.bk.Req)
	b.bk.Pr = tmp
	return err
}

func (b *PHAuthFindProfileBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	fmt.Println(req)
	b.BrickInstance().Req = &req
	return nil
}

func (b *PHAuthFindProfileBrick) Done(pkg string, idx int64, e error) error {

	if e != nil && e.Error() == "not found" {
		b.bk.Err = -102
	}

	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthFindProfileBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthFindProfileBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	if reflect.ValueOf(pr).Type().Name() == "PHProfile" {
		tmp := pr.(profile.PHProfile)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	} else {
		tmp := pr.(auth.PHAuth)
		err := jsonapi.ToJsonAPI(&tmp, w)
		return err
	}
}

func (b *PHAuthFindProfileBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PHAuth = b.BrickInstance().Pr.(auth.PHAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
