package maxjobgenerate

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/hashicorp/go-uuid"
	"io"
	"net/http"
	"time"
)

type PHMaxJobGenerateBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHMaxJobGenerateBrick) Exec() error {

	//TODO:是否需要把单次jobid存入redis中
	//TODO:等做完nhwa, 改成动态获取userid和companyid
	jobId, _ := uuid.GenerateUUID()

	maxJob := b.bk.Pr.(max.Phmaxjob)
	maxJob.Id = jobId
	maxJob.JobID = jobId
	maxJob.Date = time.Now().String()

	b.BrickInstance().Pr = maxJob
	return nil
}

func (b *PHMaxJobGenerateBrick) Prepare(pr interface{}) error {
	req := pr.(max.Phmaxjob)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHMaxJobGenerateBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHMaxJobGenerateBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHMaxJobGenerateBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.Phmaxjob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobGenerateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.Phmaxjob = b.BrickInstance().Pr.(max.Phmaxjob)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
