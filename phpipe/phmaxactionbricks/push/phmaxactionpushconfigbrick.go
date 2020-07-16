package maxactionpush

import (
	"github.com/PharbersDeveloper/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
	"strings"
)

type PhActionPushConfigBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhActionPushConfigBrick) Exec() error {
	var err error
	tmp := b.bk.Pr.(max.PhAction)
	err = pushCompanyProd(tmp)
	err = pushCalcYmConf(tmp)
	err = pushPanelConf(tmp)
	err = pushCalcConf(tmp)
	b.bk.Pr = tmp
	return err
}

func (b *PhActionPushConfigBrick) Prepare(pr interface{}) error {
	req := pr.(max.PhAction)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhActionPushConfigBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhActionPushConfigBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhActionPushConfigBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PhAction)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhActionPushConfigBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PhAction = b.BrickInstance().Pr.(max.PhAction)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

func pushCalcYmConf(paction max.PhAction) error {
	if paction.CalcYmConf.Clazz == "" {
		return nil
	}
	tmp := paction.CalcYmConf
	tmp.CompanyId = paction.CompanyId
	err := tmp.InsertBMObject()
	return err
}

func pushPanelConf(paction max.PhAction) error {
	if paction.PanelConf == nil {
		return nil
	}
	for _, v := range paction.PanelConf {
		if v.Id == "" {
			panic("no PanelConf id")
		}
		tmp := v
		tmp.CompanyId = paction.CompanyId
		err := tmp.InsertBMObject()
		return err
	}
	return nil
}

func pushCalcConf(paction max.PhAction) error {
	if paction.CalcConf == nil {
		return nil
	}
	for _, v := range paction.CalcConf {
		if v.Id == "" {
			panic("no CalcConf id")
		}
		tmp := v
		tmp.CompanyId = paction.CompanyId
		err := tmp.InsertBMObject()
		if err != nil {
			return err
		}
	}
	return nil
}

func pushCompanyProd(paction max.PhAction) error {
	//TODO: Check CompanyProd Exist
	var companyProd max.PhCompanyProd
	companyProd.Id_ = bson.NewObjectId()
	companyProd.ResetIdWithId_()
	companyProd.CompanyId = paction.CompanyId
	prodLstStr := strings.Split(paction.ProdLst, "#")
	var prodLst []interface{}
	for _, v := range prodLstStr {
		prodLst = append(prodLst, v)
	}
	companyProd.CompanyName = prodLstStr[0]
	companyProd.ProdLst = prodLst
	companyProd.InsertBMObject()
	return nil
}
