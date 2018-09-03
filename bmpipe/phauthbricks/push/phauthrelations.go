package authpush

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/Jeorch/max-go/bmmodel/auth"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type PHAuthRSPushBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthRSPushBrick) Exec() error {
	var tmp auth.PHAuth = b.bk.Pr.(auth.PHAuth)
	eq := request.EQCond{}
	eq.Ky = "auth_id"
	eq.Vy = tmp.Id
	req := request.Request{}
	req.Res = "PHAuthProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	fmt.Println(c)

	var qr auth.PHAuthProp
	err := qr.FindOne(c.(request.Request))
	if err != nil && err.Error() == "not found" {
		//panic(err)
		qr.Id_ = bson.NewObjectId()
		qr.Id = qr.Id_.Hex()
		qr.AuthID = tmp.Id
		qr.ProfileID = tmp.Profile.Id
		qr.InsertBMObject()
	}
	fmt.Println(qr)
	return nil
}

func (b *PHAuthRSPushBrick) Prepare(pr interface{}) error {
	req := pr.(auth.PHAuth)
	//b.bk.Pr = req
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHAuthRSPushBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthRSPushBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthRSPushBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PHAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHAuthRSPushBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PHAuth = b.BrickInstance().Pr.(auth.PHAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
