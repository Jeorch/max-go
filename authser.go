package main

import (
	"github.com/Jeorch/max-go/phmodel/company"
	"github.com/Jeorch/max-go/phmodel/profile"
	"github.com/Jeorch/max-go/phpipe/phauthbricks/others"
	"github.com/Jeorch/max-go/phpipe/phcompanybricks/push"
	"github.com/Jeorch/max-go/phpipe/phprofilebricks/push"
	"net/http"

	"github.com/Jeorch/max-go/phmodel/auth"
	"github.com/Jeorch/max-go/phpipe/phauthbricks/find"
	"github.com/Jeorch/max-go/phpipe/phauthbricks/push"
	"github.com/Jeorch/max-go/phpipe/phauthbricks/update"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmrouter"
)

func main() {

	fac := bmsingleton.GetFactoryInstance()

	/*------------------------------------------------
	 * model object
	 *------------------------------------------------*/
	fac.RegisterModel("request", &request.Request{})
	fac.RegisterModel("eq_cond", &request.EQCond{})
	fac.RegisterModel("up_cond", &request.UPCond{})
	fac.RegisterModel("fm_cond", &request.FMUCond{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})

	fac.RegisterModel("PHAuth", &auth.PHAuth{})
	fac.RegisterModel("PHAuthProp", &auth.PHAuthProp{})
	fac.RegisterModel("PHCompany", &company.PHCompany{})
	fac.RegisterModel("PHProfile", &profile.PHProfile{})
	fac.RegisterModel("PHProfileProp", &profile.PHProfileProp{})

	/*------------------------------------------------
	 * auth find bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthFindProfileBrick", &authfind.PHAuthFindProfileBrick{})
	//fac.RegisterModel("PHProfile2ProfileProp", &authfind.PHProfile2ProfileProp{})
	//fac.RegisterModel("PHProfileProp2ProfileBrick", &authfind.PHProfileProp2ProfileBrick{})
	fac.RegisterModel("PHProfile2AuthProp", &authfind.PHProfile2AuthProp{})
	fac.RegisterModel("PHAuthProp2AuthBrick", &authfind.PHAuthProp2AuthBrick{})

	/*------------------------------------------------
	 * auth push bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthRSPushBrick", &authpush.PHAuthRSPushBrick{})
	fac.RegisterModel("PHAuthPushBrick", &authpush.PHAuthPushBrick{})

	/*------------------------------------------------
	 * auth update bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthProfileUpdateBrick", &authupdate.PHAuthProfileUpdateBrick{})

	/*------------------------------------------------
	 * company bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthCompanyPush", &companypush.PHAuthCompanyPush{})

	/*------------------------------------------------
	 * profile bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthProfilePush", &profilepush.PHAuthProfilePush{})
	fac.RegisterModel("PHAuthProfileRSPush", &profilepush.PHAuthProfileRSPush{})

	/*------------------------------------------------
	 * other bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthGenerateToken", &authothers.PHAuthGenerateToken{})

	r := bmrouter.BindRouter()
	http.ListenAndServe(":8080", r)
}
