package samplecheckforward

import (
	"github.com/PharbersDeveloper/max-go/phmodel/config"
	"github.com/PharbersDeveloper/max-go/phmodel/samplecheck"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHSampleCheckBodyForwardBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

var maxForward maxconfig.PhForwardConfig

func (b *PHSampleCheckBodyForwardBrick) Exec() error {
	return nil
}

func (b *PHSampleCheckBodyForwardBrick) Prepare(pr interface{}) error {
	req := pr.(samplecheck.SampleCheckBody)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHSampleCheckBodyForwardBrick) Done(pkg string, idx int64, e error) error {
	maxForward.GenerateConfig()
	host := maxForward.HostA
	port := maxForward.PortA
	bmrouter.ForWardNextBrick(host, port, pkg, idx, b)
	return nil
}

func (b *PHSampleCheckBodyForwardBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHSampleCheckBodyForwardBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(samplecheck.SampleCheckBody)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHSampleCheckBodyForwardBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval samplecheck.SampleCheckBody = b.BrickInstance().Pr.(samplecheck.SampleCheckBody)
		jsonapi.ToJsonAPI(&reval, w)
	}
}