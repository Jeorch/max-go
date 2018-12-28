package main

import (
	"github.com/Jeorch/max-go/phmodel/company"
	"github.com/Jeorch/max-go/phmodel/export"
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/Jeorch/max-go/phmodel/profile"
	"github.com/Jeorch/max-go/phmodel/resultcheck"
	"github.com/Jeorch/max-go/phmodel/samplecheck"
	"github.com/Jeorch/max-go/phmodel/xmpp"
	"github.com/Jeorch/max-go/phpipe/phauthbricks/others"
	"github.com/Jeorch/max-go/phpipe/phcompanybricks/push"
	"github.com/Jeorch/max-go/phpipe/phexportbricks/forward"
	"github.com/Jeorch/max-go/phpipe/phmaxactionbricks/find"
	"github.com/Jeorch/max-go/phpipe/phmaxactionbricks/generate"
	"github.com/Jeorch/max-go/phpipe/phmaxactionbricks/send"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/delete"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/generate"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/push"
	"github.com/Jeorch/max-go/phpipe/phmaxjobbricks/send"
	"github.com/Jeorch/max-go/phpipe/phprofilebricks/push"
	"github.com/Jeorch/max-go/phpipe/phresultcheck/forward"
	"github.com/Jeorch/max-go/phpipe/phsamplecheck/forward"
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"net/http"
	"sync"

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
	fac.RegisterModel("Request", &request.Request{})
	fac.RegisterModel("Eqcond", &request.Eqcond{})
	fac.RegisterModel("Upcond", &request.Upcond{})
	fac.RegisterModel("Fmcond", &request.Fmcond{})
	fac.RegisterModel("BmErrorNode", &bmerror.BmErrorNode{})

	fac.RegisterModel("PhAuth", &auth.PhAuth{})
	fac.RegisterModel("PhAuthProp", &auth.PhAuthProp{})
	fac.RegisterModel("PhCompany", &company.PhCompany{})
	fac.RegisterModel("PhProfile", &profile.PhProfile{})
	fac.RegisterModel("PhProfileProp", &profile.PhProfileProp{})

	fac.RegisterModel("Phmaxjob", &max.Phmaxjob{})
	fac.RegisterModel("PhAction", &max.PhAction{})
	fac.RegisterModel("PhCalcYmConf", &max.PhCalcYmConf{})
	fac.RegisterModel("PhPanelConf", &max.PhPanelConf{})
	fac.RegisterModel("PhCalcConf", &max.PhCalcConf{})
	fac.RegisterModel("PhXmppConf", &xmpp.PhXmppConf{})

	fac.RegisterModel("SampleCheckSelecter", &samplecheck.SampleCheckSelecter{})
	fac.RegisterModel("SampleCheckBody", &samplecheck.SampleCheckBody{})
	fac.RegisterModel("ResultCheck", &resultcheck.ResultCheck{})
	fac.RegisterModel("ExportMaxResult", &export.ExportMaxResult{})

	/*------------------------------------------------
	 * auth find bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthFindProfileBrick", &authfind.PHAuthFindProfileBrick{})
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
	 * max bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHMaxJobGenerateBrick", &maxjobgenerate.PHMaxJobGenerateBrick{})
	fac.RegisterModel("PHMaxJobDeleteBrick", &maxjobdelete.PHMaxJobDeleteBrick{})
	fac.RegisterModel("PHMaxJobPushBrick", &maxjobpush.PHMaxJobPushBrick{})
	fac.RegisterModel("PHMaxJobPushPanelBrick", &maxjobpush.PHMaxJobPushPanelBrick{})
	fac.RegisterModel("PHMaxJobSendBrick", &maxjobsend.PHMaxJobSendBrick{})

	/*------------------------------------------------
	 * max action bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PhMaxActionGenerateBrick", &maxactiongenerate.PhMaxActionGenerateBrick{})
	fac.RegisterModel("PhMaxActionCalcYmBrick", &maxactionfind.PhMaxActionCalcYmBrick{})
	fac.RegisterModel("PhMaxActionPanelBrick", &maxactionfind.PhMaxActionPanelBrick{})
	fac.RegisterModel("PhMaxActionCalcBrick", &maxactionfind.PhMaxActionCalcBrick{})

	//fac.RegisterModel("PhMaxActionPushBrick", &maxactionpush.PhMaxActionPushBrick{})
	//fac.RegisterModel("PhMaxActionPushPanelBrick", &maxactionpush.PhMaxActionPushPanelBrick{})

	fac.RegisterModel("PhMaxActionSendBrick", &maxactionsend.PhMaxActionSendBrick{})

	/*------------------------------------------------
	 * sample check bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PhSelecterForMarketBrick", &samplecheckforward.PhSelecterForMarketBrick{})
	fac.RegisterModel("PHSampleCheckSelecterForwardBrick", &samplecheckforward.PHSampleCheckSelecterForwardBrick{})
	fac.RegisterModel("PHSampleCheckBodyForwardBrick", &samplecheckforward.PHSampleCheckBodyForwardBrick{})

	/*------------------------------------------------
	 * result check bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHResultCheckForwardBrick", &resultcheckforward.PHResultCheckForwardBrick{})

	/*------------------------------------------------
	 * export bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHExportMaxResultForwardBrick", &exportforward.PHExportMaxResultForwardBrick{})

	/*------------------------------------------------
	 * other bricks object
	 *------------------------------------------------*/
	fac.RegisterModel("PHAuthGenerateToken", &authothers.PHAuthGenerateToken{})

	r := bmrouter.BindRouter()

	var once sync.Once
	var bmRouter bmconfig.BMRouterConfig
	once.Do(bmRouter.GenerateConfig)

	http.ListenAndServe(":"+bmRouter.Port, r)
}
