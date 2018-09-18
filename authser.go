package main

import (
	"github.com/Jeorch/max-go/phmodel/company"
	"github.com/Jeorch/max-go/phmodel/maxjob"
	"github.com/Jeorch/max-go/phmodel/profile"
	"github.com/Jeorch/max-go/phpipe/phauthbricks/others"
	"github.com/Jeorch/max-go/phpipe/phcompanybricks/push"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/delete"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/forword"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/generate"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/push"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/send"
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
	fac.RegisterModel("eqcond", &request.EQCond{})
	fac.RegisterModel("upcond", &request.UPCond{})
	fac.RegisterModel("fmcond", &request.FMUCond{})
	fac.RegisterModel("BMErrorNode", &bmerror.BMErrorNode{})

	fac.RegisterModel("PHAuth", &auth.PHAuth{})
	fac.RegisterModel("PHAuthProp", &auth.PHAuthProp{})
	fac.RegisterModel("PHCompany", &company.PHCompany{})
	fac.RegisterModel("PHProfile", &profile.PHProfile{})
	fac.RegisterModel("PHProfileProp", &profile.PHProfileProp{})

	fac.RegisterModel("phmaxjob", &maxjob.PHMaxJob{})

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
	 * maxjob bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHMaxJobForwardBrick", &maxjobforword.PHMaxJobForwardBrick{})
	fac.RegisterModel("PHMaxJobGenerateBrick", &maxjobgenerate.PHMaxJobGenerateBrick{})
	fac.RegisterModel("PHMaxJobDeleteBrick", &maxjobdelete.PHMaxJobDeleteBrick{})
	fac.RegisterModel("PHMaxJobPushBrick", &maxjobpush.PHMaxJobPushBrick{})
	fac.RegisterModel("PHMaxJobSendBrick", &maxjobsend.PHMaxJobSendBrick{})

	/*------------------------------------------------
	 * other bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthGenerateToken", &authothers.PHAuthGenerateToken{})

	r := bmrouter.BindRouter()
	http.ListenAndServe(":8081", r)
}
