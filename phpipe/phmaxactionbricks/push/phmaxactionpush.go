package maxactionpush
//
//import (
//	"fmt"
//	"github.com/Jeorch/max-go/phmodel/max"
//	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
//	"github.com/alfredyang1986/blackmirror/bmconfighandle"
//	"github.com/alfredyang1986/blackmirror/bmerror"
//	"github.com/alfredyang1986/blackmirror/bmexcelhandle"
//	"github.com/alfredyang1986/blackmirror/bmpipe"
//	"github.com/alfredyang1986/blackmirror/bmrouter"
//	"github.com/alfredyang1986/blackmirror/jsonapi"
//	"github.com/colinmarc/hdfs"
//	"github.com/hashicorp/go-uuid"
//	"io"
//	"net/http"
//	"os"
//)
//
//type PhMaxActionPushBrick struct {
//	bk *bmpipe.BMBrick
//}
//
///*------------------------------------------------
// * brick interface
// *------------------------------------------------*/
//
//func (b *PhMaxActionPushBrick) Exec() error {
//	var tmp max.PhAction = b.bk.Pr.(max.PhAction)
//	var err error
//	var cpaDesName string
//	var notArrivalHospDesName string
//	var gycDesName string
//
//	calcYmConf := tmp.CalcYmConf
//	cpa := calcYmConf.Conf["cpa_file"].(string)
//	gycx := calcYmConf.Conf["gycx_file"].(string)
//
//	cpaDesName, notArrivalHospDesName, gycDesName, err = pushCpaGyc(cpa, gycx)
//
//	calcYmConf.Conf["cpa_file"] = cpaDesName
//	calcYmConf.Conf["gycx_file"] = gycDesName
//	calcYmConf.Conf["not_arrival_hosp_file"] = notArrivalHospDesName
//	tmp.CalcYmConf = calcYmConf
//
//	b.BrickInstance().Pr = tmp
//	return err
//}
//
//func (b *PhMaxActionPushBrick) Prepare(pr interface{}) error {
//	req := pr.(max.PhAction)
//
//	b.BrickInstance().Pr = req
//	return nil
//}
//
//func (b *PhMaxActionPushBrick) Done(pkg string, idx int64, e error) error {
//	tmp, _ := bmpkg.GetPkgLen(pkg)
//	if int(idx) < tmp-1 {
//		bmrouter.NextBrickRemote(pkg, idx+1, b)
//	}
//	return nil
//}
//
//func (b *PhMaxActionPushBrick) BrickInstance() *bmpipe.BMBrick {
//	if b.bk == nil {
//		b.bk = &bmpipe.BMBrick{}
//	}
//	return b.bk
//}
//
//func (b *PhMaxActionPushBrick) ResultTo(w io.Writer) error {
//	pr := b.BrickInstance().Pr
//	tmp := pr.(max.PhAction)
//	err := jsonapi.ToJsonAPI(&tmp, w)
//	return err
//}
//
//func (b *PhMaxActionPushBrick) Return(w http.ResponseWriter) {
//	ec := b.BrickInstance().Err
//	if ec != 0 {
//		bmerror.ErrInstance().ErrorReval(ec, w)
//	} else {
//		var reval max.PhAction = b.BrickInstance().Pr.(max.PhAction)
//		jsonapi.ToJsonAPI(&reval, w)
//	}
//}
//
//func pushCpaGyc(cpa string, gyc string) (string, string, string, error) {
//	var err error
//	var cpaCsv string
//	var gycCsv string
//	var cpaDesName string
//	var gycDesName string
//	var notArrivalHospCsv string
//	var notArrivalHospDesName string
//	if cpa != "" {
//		cpaCsv, notArrivalHospCsv, err = cpa2csv(cpa)
//		cpaDesName, err = push2hdfs(cpaCsv)
//		notArrivalHospDesName, err = push2hdfs(notArrivalHospCsv)
//	}
//	if gyc != "" {
//		gycCsv, err = gyc2csv(gyc)
//		gycDesName, err = push2hdfs(gycCsv)
//	}
//	return cpaDesName, notArrivalHospDesName, gycDesName, err
//}
//
//func cpa2csv(cpaFile string) (string, string, error) {
//	var err error
//	var cpa string
//	var notArrivalHosp string
//	var bmRouter bmconfig.BMRouterConfig
//	bmRouter.GenerateConfig()
//	localCpa := bmRouter.TmpDir + cpaFile
//	cpa, err = bmexcelhandle.GenerateCSVFromXLSXFile(localCpa, 0)
//	notArrivalHosp, err = bmexcelhandle.GenerateCSVFromXLSXFile(localCpa, 1)
//	os.Remove(localCpa)
//	return cpa, notArrivalHosp, err
//}
//
//func gyc2csv(gycFile string) (string, error) {
//	var err error
//	var gyc string
//	var bmRouter bmconfig.BMRouterConfig
//	bmRouter.GenerateConfig()
//	localGyc := bmRouter.TmpDir + gycFile
//	gyc, err = bmexcelhandle.GenerateCSVFromXLSXFile(localGyc, 0)
//	os.Remove(localGyc)
//	return gyc, err
//}
//
//func push2hdfs(localFile string) (string, error) {
//	localDir := localFile
//	fileDesName, _ := uuid.GenerateUUID()
//	fmt.Println(fileDesName)
//	//TODO:hdfs 配置化
//	fileDesPath := "/workData/Client/" + fileDesName
//	client, _ := hdfs.New("192.168.100.137:9000")
//	err := client.CopyToRemote(localDir, fileDesPath)
//	os.Remove(localDir)
//	return fileDesName, err
//}
