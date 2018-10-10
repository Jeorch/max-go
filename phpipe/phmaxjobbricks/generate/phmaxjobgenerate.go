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
	jobid, _ := uuid.GenerateUUID()
	mj := max.PHMaxJob{
		Id:        jobid,
		UserID:    "jeorch",
		CompanyID: "5afa53bded925c05c6f69c54",
		JobID:     jobid,
		Date:      time.Now().String(),
		Call:      "JobGenerate",
	}

	b.BrickInstance().Pr = mj
	return nil
}

func (b *PHMaxJobGenerateBrick) Prepare(pr interface{}) error {
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
	tmp := pr.(max.PHMaxJob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobGenerateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PHMaxJob = b.BrickInstance().Pr.(max.PHMaxJob)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
