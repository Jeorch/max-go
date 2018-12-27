package maxactionsend

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/bmxmpp"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PhMaxActionSendBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhMaxActionSendBrick) Exec() error {
	var maxjob max.PhAction = b.bk.Pr.(max.PhAction)
	msg, err := jsonapi.ToJsonString(&maxjob)
	println(msg)
	err = bmxmpp.Forward("driver@localhost", msg)
	//err = bmxmpp.Forward("test@localhost", msg)
	return err
}

func (b *PhMaxActionSendBrick) Prepare(pr interface{}) error {
	req := pr.(max.PhAction)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhMaxActionSendBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhMaxActionSendBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhMaxActionSendBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PhAction)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhMaxActionSendBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PhAction = b.BrickInstance().Pr.(max.PhAction)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
