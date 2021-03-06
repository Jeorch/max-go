package authfind

import (
	"fmt"
	"github.com/PharbersDeveloper/max-go/phmodel/auth"
	"github.com/PharbersDeveloper/max-go/phmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHProfile2AuthProp struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHProfile2AuthProp) Exec() error {
	var tmp profile.PhProfile = b.bk.Pr.(profile.PhProfile)
	eq := request.Eqcond{}
	eq.Ky = "profile_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "PhAuthProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("Eqcond", condi)
	fmt.Println(c)

	var reval auth.PhAuthProp
	err := reval.FindOne(c.(request.Request))
	b.bk.Pr = reval
	return err
}

func (b *PHProfile2AuthProp) Prepare(pr interface{}) error {
	req := pr.(profile.PhProfile)
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
	tmp := pr.(auth.PhAuthProp)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHProfile2AuthProp) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PhAuth = b.BrickInstance().Pr.(auth.PhAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
