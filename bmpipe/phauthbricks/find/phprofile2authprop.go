package authfind

import (
	"fmt"
	"github.com/Jeorch/max-go/bmmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/Jeorch/max-go/bmmodel/auth"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"net/http"
	"io"
)

type PHProfile2AuthProp struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHProfile2AuthProp) Exec() error {
	var tmp profile.PHProfile = b.bk.Pr.(profile.PHProfile)
	eq := request.EQCond{}
	eq.Ky = "profile_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "PHAuthProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var reval auth.PHAuthProp
	err := reval.FindOne(c.(request.Request))
	b.bk.Pr = reval
	return err
}

func (b *PHProfile2AuthProp) Prepare(pr interface{}) error {
	req := pr.(profile.PHProfile)
	b.BrickInstance().Pr = req
	//b.bk.Pr = req
	return nil
}

func (b *PHProfile2AuthProp) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHProfile2AuthProp) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHProfile2AuthProp) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PHAuthProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHProfile2AuthProp) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PHAuth = b.BrickInstance().Pr.(auth.PHAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
