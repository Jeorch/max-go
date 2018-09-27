package maxjobforword

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

type PHMaxJobForwardBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHMaxJobForwardBrick) Exec() error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PHMaxJob)
	msg, err := jsonapi.ToJsonString(&tmp)
	err = bmxmpp.Forward("cui@localhost", msg)
	return err
}

func (b *PHMaxJobForwardBrick) Prepare(pr interface{}) error {
	req := pr.(max.PHMaxJob)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHMaxJobForwardBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHMaxJobForwardBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHMaxJobForwardBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PHMaxJob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobForwardBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PHMaxJob = b.BrickInstance().Pr.(max.PHMaxJob)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
