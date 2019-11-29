package maxjobdelete

import (
	"fmt"
	"github.com/PharbersDeveloper/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHMaxJobDeleteBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHMaxJobDeleteBrick) Exec() error {
	var tmp max.Phmaxjob = b.bk.Pr.(max.Phmaxjob)
	jobid := tmp.JobID
	//TODO:删除redis中的单次jobid
	fmt.Println(jobid)

	return nil
}

func (b *PHMaxJobDeleteBrick) Prepare(pr interface{}) error {
	req := pr.(max.Phmaxjob)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHMaxJobDeleteBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHMaxJobDeleteBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHMaxJobDeleteBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.Phmaxjob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobDeleteBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.Phmaxjob = b.BrickInstance().Pr.(max.Phmaxjob)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
