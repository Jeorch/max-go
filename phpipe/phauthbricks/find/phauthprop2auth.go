package authfind

import (
	"fmt"
	"github.com/Jeorch/max-go/phmodel/auth"
	"github.com/Jeorch/max-go/phmodel/company"
	"github.com/Jeorch/max-go/phmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type PHAuthProp2AuthBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PHAuthProp2AuthBrick) Exec() error {
	prop := b.bk.Pr.(auth.PHAuthProp)
	reval, err := findAuth(prop)
	profile, err := findProfile(prop)
	profileProp, err := findProfileProp(profile)
	company, err := findProfileCompany(profileProp)
	profile.Company = company
	reval.Profile = profile
	b.bk.Pr = reval
	return err
}

func (b *PHAuthProp2AuthBrick) Prepare(pr interface{}) error {
	req := pr.(auth.PHAuthProp)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHAuthProp2AuthBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PHAuthProp2AuthBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHAuthProp2AuthBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(auth.PHAuth)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHAuthProp2AuthBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval auth.PHAuth = b.BrickInstance().Pr.(auth.PHAuth)
		jsonapi.ToJsonAPI(&reval, w)
	}
}

/*------------------------------------------------
 * brick inner function
 *------------------------------------------------*/

func findAuth(prop auth.PHAuthProp) (auth.PHAuth, error) {
	eq := request.EQCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.AuthID)
	req := request.Request{}
	req.Res = "PHAuth"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("eqcond", condi)
	fmt.Println(c)

	reval := auth.PHAuth{}
	err := reval.FindOne(c.(request.Request))

	return reval, err

}

func findProfile(prop auth.PHAuthProp) (profile.PHProfile, error) {
	eq := request.EQCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.ProfileID)
	req := request.Request{}
	req.Res = "PHProfile"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("eqcond", condi)
	fmt.Println(c)

	reval := profile.PHProfile{}
	err := reval.FindOne(c.(request.Request))
	reval.Password = ""

	return reval, err
}

func findProfileProp(phProfile profile.PHProfile) (profile.PHProfileProp, error) {
	eq := request.EQCond{}
	eq.Ky = "profile_id"
	eq.Vy = phProfile.Id
	req := request.Request{}
	req.Res = "PHProfileProp"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("eqcond", condi)
	fmt.Println(c)

	reval := profile.PHProfileProp{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}

func findProfileCompany(prop profile.PHProfileProp) (company.PHCompany, error) {
	eq := request.EQCond{}
	eq.Ky = "_id"
	eq.Vy = bson.ObjectIdHex(prop.CompanyID)
	req := request.Request{}
	req.Res = "PHCompany"
	var condi []interface{}
	condi = append(condi, eq)
	c := req.SetConnect("eqcond", condi)
	fmt.Println(c)

	reval := company.PHCompany{}
	err := reval.FindOne(c.(request.Request))

	return reval, err
}
