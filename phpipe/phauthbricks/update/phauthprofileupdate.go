package authupdate

import (
	"github.com/PharbersDeveloper/max-go/phmodel/profile"
	//"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	//"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type PHAuthProfileUpdateBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthProfileUpdateBrick) Exec() error {
	tmp := profile.PhProfile{}
	tmp.UpdateBMObject(*b.bk.Req)
	b.bk.Pr = tmp
	return nil
}

func (b *PHAuthProfileUpdateBrick) Prepare(pr interface{}) error {
	req := pr.(request.Request)
	//b.bk.Req = &req
	b.BrickInstance().Req = &req
	return nil
}

func (b *PHAuthProfileUpdateBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthProfileUpdateBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthProfileUpdateBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(profile.PhProfile)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHAuthProfileUpdateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval profile.PhProfile = b.BrickInstance().Pr.(profile.PhProfile)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
