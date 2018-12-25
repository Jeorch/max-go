package maxjobsend

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/bmxmpp"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/hashicorp/go-uuid"
	"io"
	"net/http"
	"strings"
)

type PHMaxJobSendBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHMaxJobSendBrick) Exec() error {
	var maxjob max.Phmaxjob = b.bk.Pr.(max.Phmaxjob)

	paction := maxjob2phaction(maxjob)
	msg, err := jsonapi.ToJsonString(&paction)
	println(msg)
	err = bmxmpp.Forward("driver@localhost", msg)
	//err = bmxmpp.Forward("test@localhost", msg)
	return err
}

func (b *PHMaxJobSendBrick) Prepare(pr interface{}) error {
	req := pr.(max.Phmaxjob)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHMaxJobSendBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHMaxJobSendBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHMaxJobSendBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.Phmaxjob)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHMaxJobSendBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.Phmaxjob = b.BrickInstance().Pr.(max.Phmaxjob)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func maxjob2phaction(maxjob max.Phmaxjob) max.PhAction {

	pactionId, _ := uuid.GenerateUUID()
	xmppConfId, _ := uuid.GenerateUUID()
	paction := max.PhAction{
		Id:pactionId,
	}
	paction.XmppConf.Id = xmppConfId
	paction.XmppConf.XmppReport = maxjob.UserID + "@localhost"

	paction.UserId = maxjob.UserID
	paction.CompanyId = maxjob.CompanyID
	paction.JobId = maxjob.JobID
	//TODO：配置化
	paction.MaxPath = "hdfs:///workData/Max"
	paction.PanelPath = "hdfs:///workData/Panel"
	paction.ProdLst = generateCompanyProd(maxjob)

	switch maxjob.Call {
	case "ymCalc":
		paction.CalcYmConf = generateCalcYmConf(maxjob)
	case "panel":
		paction.PanelConf = generatePanelConf(maxjob)
	case "calc":
		paction.CalcConf = generateCalcConf(maxjob)
	}

	//if paction.CalcYmConf.Id == "" {
	//	var tmp interface{}
	//	paction.CalcYmConf = null
	//}

	return paction

}

func generateCompanyProd(maxjob max.Phmaxjob) string {
	prodLst := ""
	companyProd := max.PhCompanyProd{}
	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = maxjob.CompanyID
	req := request.Request{}
	req.Res = "PhCompanyProd"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	err := companyProd.FindOne(c.(request.Request))
	if err != nil {
		return prodLst
	}

	prodLst = companyProd.ProdLst[0].(string)
	for _, v := range companyProd.ProdLst[1:] {
		prodLst = prodLst + "#" + v.(string)
	}

	return prodLst

}

func generateCalcYmConf(maxjob max.Phmaxjob) max.PhCalcYmConf {
	calcYmConf := max.PhCalcYmConf{}
	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = maxjob.CompanyID
	req := request.Request{}
	req.Res = "PhCalcYmConf"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("conditions", condi)
	err := calcYmConf.FindOne(c.(request.Request))
	if err != nil {
		return max.PhCalcYmConf{}
	}

	confMap := make(map[string]string)
	confMap["cpa_file"] = "hdfs:///workData/Client/" + maxjob.Cpa
	confMap["gycx_file"] = "hdfs:///workData/Client/" + maxjob.Gycx
	confMap["not_arrival_hosp_file"] = "hdfs:///workData/Client/" + maxjob.NotArrivalHospFile
	calcYmConf.Conf = confMap
	return calcYmConf

}

func generatePanelConf(maxjob max.Phmaxjob) []max.PhPanelConf {
	req := request.Request{}
	req.Res = "PhPanelConf"

	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = maxjob.CompanyID

	incond := request.Incond{}
	incond.Ky = "ym"
	yms := []string{}
	yms = strings.Split(maxjob.Yms, "#")
	incond.Vy = yms

	var condi1 []interface{}
	var condi2 []interface{}
	condi1 = append(condi1, eq)
	condi2 = append(condi2, incond)

	req.SetConnect("Eqcond", condi1)
	req.SetConnect("Incond", condi2)

	var reval []max.PhPanelConf
	var rlst []max.PhPanelConf
	err := bmmodel.FindMutil(req, &reval)
	if err != nil {
		return []max.PhPanelConf{}
	}

	for _, v := range reval[:] {
		tmpName, _ := uuid.GenerateUUID()
		v.Conf["cpa_file"] = "hdfs:///workData/Client/" + maxjob.Cpa
		v.Conf["gycx_file"] = "hdfs:///workData/Client/" + maxjob.Gycx
		v.Conf["not_arrival_hosp_file"] = "hdfs:///workData/Client/" + maxjob.NotArrivalHospFile
		v.ResetIdWithId_()
		v.PanelName = tmpName
		rlst = append(rlst, v)
	}

	return rlst
}

func generateCalcConf(maxjob max.Phmaxjob) []max.PhCalcConf {
	req := request.Request{}
	req.Res = "PhCalcConf"

	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = maxjob.CompanyID

	incond := request.Incond{}
	incond.Ky = "ym"
	incond.Vy = maxjob.Yms

	var condi1 []interface{}
	var condi2 []interface{}
	condi1 = append(condi1, eq)
	condi2 = append(condi2, incond)

	req.SetConnect("Eqcond", condi1)
	req.SetConnect("Incond", condi2)

	var reval []max.PhCalcConf
	var rlst []max.PhCalcConf
	err := bmmodel.FindMutil(req, &reval)
	if err != nil {
		return []max.PhCalcConf{}
	}

	for _, v := range reval {
		tmpName1, _ := uuid.GenerateUUID()
		tmpName2, _ := uuid.GenerateUUID()
		v.Conf["cpa_file"] = "hdfs:///workData/Client/" + maxjob.Cpa
		v.Conf["gycx_file"] = "hdfs:///workData/Client/" + maxjob.Gycx
		v.Conf["not_arrival_hosp_file"] = "hdfs:///workData/Client/" + maxjob.NotArrivalHospFile
		v.ResetIdWithId_()
		v.MaxName = tmpName1
		v.MaxSearchName = tmpName2
		//TODO:从panel结果中读取panelName.
		//v.PanelName
		rlst = append(rlst, v)
	}

	return rlst
}
