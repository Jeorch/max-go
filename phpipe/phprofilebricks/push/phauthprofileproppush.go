package profilepush

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
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type PHAuthProfileRSPush struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthProfileRSPush) Exec() error {
	var tmp auth.PhAuth = b.bk.Pr.(auth.PhAuth)
	profile_tmp := tmp.Profile

	eq := request.Eqcond{}
	eq.Ky = "profile_id"
	eq.Vy = profile_tmp.Id
	req := request.Request{}
	req.Res = "PhProfileProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("Eqcond", condi)
	fmt.Println(c)

	var qr profile.PhProfileProp
	err := qr.FindOne(c.(request.Request))
	if err != nil && err.Error() == "not found" {
		//panic(err)
		qr.Id_ = bson.NewObjectId()
		qr.Id = qr.Id_.Hex()
		qr.ProfileID = profile_tmp.Id
		qr.CompanyID = profile_tmp.Company.Id
		qr.InsertBMObject()
	}
	fmt.Println(qr)

	b.bk.Pr = tmp
	return nil
}

func (b *PHAuthProfileRSPush) Prepare(pr interface{}) error {
	req := pr.(auth.PhAuth)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHAuthProfileRSPush) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthProfileRSPush) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthProfileRSPush) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PhAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHAuthProfileRSPush) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PhAuth = b.BrickInstance().Pr.(auth.PhAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
