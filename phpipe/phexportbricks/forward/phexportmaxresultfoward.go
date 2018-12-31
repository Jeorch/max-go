package exportforward

import (
	"github.com/Jeorch/max-go/phmodel/config"
	"github.com/Jeorch/max-go/phmodel/export"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHExportMaxResultForwardBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

var maxForward maxconfig.PhForwardConfig

func (b *PHExportMaxResultForwardBrick) Exec() error {
	return nil
}

func (b *PHExportMaxResultForwardBrick) Prepare(pr interface{}) error {
	req := pr.(export.ExportMaxResult)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHExportMaxResultForwardBrick) Done(pkg string, idx int64, e error) error {
	maxForward.GenerateConfig()
	host := maxForward.HostB
	port := maxForward.PortB
	bmrouter.ForWardNextBrick(host, port, pkg, idx, b)
	return nil
}

func (b *PHExportMaxResultForwardBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHExportMaxResultForwardBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(export.ExportMaxResult)

	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHExportMaxResultForwardBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval export.ExportMaxResult = b.BrickInstance().Pr.(export.ExportMaxResult)
		reval.Save2Local()
		jsonapi.ToJsonAPI(&reval, w)
	}
}
