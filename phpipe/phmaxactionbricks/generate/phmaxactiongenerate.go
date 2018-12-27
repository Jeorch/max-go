package maxactiongenerate

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/hashicorp/go-uuid"
	"io"
	"net/http"
)

type PhMaxActionGenerateBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhMaxActionGenerateBrick) Exec() error {

	jobId, _ := uuid.GenerateUUID()
	xmppConfId, _ := uuid.GenerateUUID()
	calcConfId, _ := uuid.GenerateUUID()

	maxJob := b.bk.Pr.(max.PhAction)
	maxJob.Id = jobId
	maxJob.JobId = jobId
	//maxJob.CreateTime = time.Now().UnixNano()

	maxJob.CalcYmConf.Id = calcConfId
	maxJob.XmppConf.Id = xmppConfId
	maxJob.XmppConf.XmppReport = maxJob.UserId

	maxJob, err := generateCalcYmConf(maxJob)

	b.BrickInstance().Pr = maxJob
	return err
}

func (b *PhMaxActionGenerateBrick) Prepare(pr interface{}) error {
	req := pr.(max.PhAction)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhMaxActionGenerateBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhMaxActionGenerateBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhMaxActionGenerateBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PhAction)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhMaxActionGenerateBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PhAction = b.BrickInstance().Pr.(max.PhAction)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func generateCalcYmConf(paction max.PhAction) (max.PhAction, error) {

	calcYmConf := paction.CalcYmConf

	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = paction.CompanyId
	req := request.Request{}
	req.Res = "PhCalcYmConf"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)

	err := calcYmConf.FindOne(c.(request.Request))
	paction.CalcYmConf = calcYmConf

	return paction, err
}
