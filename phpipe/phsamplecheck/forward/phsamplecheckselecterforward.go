package samplecheckforward

import (
	"github.com/Jeorch/max-go/phmodel/samplecheck"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHSampleCheckSelecterForwardBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHSampleCheckSelecterForwardBrick) Exec() error {
	return nil
}

func (b *PHSampleCheckSelecterForwardBrick) Prepare(pr interface{}) error {
	req := pr.(samplecheck.SampleCheckSelecter)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHSampleCheckSelecterForwardBrick) Done(pkg string, idx int64, e error) error {
	//TODO：forward配置化
	host := "max-client"
	port := "9000"
	bmrouter.ForWardNextBrick(host, port, pkg, idx, b)
	return nil
}

func (b *PHSampleCheckSelecterForwardBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHSampleCheckSelecterForwardBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(samplecheck.SampleCheckSelecter)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHSampleCheckSelecterForwardBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval samplecheck.SampleCheckSelecter = b.BrickInstance().Pr.(samplecheck.SampleCheckSelecter)
		jsonapi.ToJsonAPI(&reval, w)
	}
}