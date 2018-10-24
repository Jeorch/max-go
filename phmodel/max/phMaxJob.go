package max

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmexcelhandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/colinmarc/hdfs"
	"github.com/go-redis/redis"
	"github.com/hashicorp/go-uuid"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type PHMaxJob struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	UserID             string `json:"user_id" bson:"user_id"`
	CompanyID          string `json:"company_id" bson:"company_id"`
	JobID              string `json:"job_id" bson:"job_id"`
	Date               string `json:"date" bson:"date"`
	Call               string `json:"call" bson:"call"`
	Percentage         int 	  `json:"percentage" bson:"percentage"`
	Message            string `json:"message" bson:"message"`
	Cpa                string `json:"cpa" bson:"cpa"`
	Gycx               string `json:"gycx" bson:"gycx"`
	NotArrivalHospFile string `json:"not_arrival_hosp_file" bson:"not_arrival_hosp_file"`
	Yms                string `json:"yms" bson:"yms"`
	//Args               map[string]interface{} `json:"args" bson:"args"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PHMaxJob) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PHMaxJob) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PHMaxJob) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PHMaxJob) QueryId() string {
	return bd.Id
}

func (bd *PHMaxJob) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PHMaxJob) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PHMaxJob) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PHMaxJob) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PHMaxJob) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PHMaxJob) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PHMaxJob) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *PHMaxJob) CheckJobIdCall() error {

	//var err error

	var cpaDesName string
	var notArrivalHospDesName string
	var gycDesName string

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	jobID := bd.JobID
	call := bd.Call
	jobCall := jobID + call

	cpa, err := client.HGet(jobCall, "cpa").Result()

	if cpa != "" {
		notArrivalHospFile, _ := client.HGet(jobCall, "not_arrival_hosp_file").Result()
		gycx, _ := client.HGet(jobCall, "gycx").Result()

		bd.Cpa = cpa
		bd.NotArrivalHospFile = notArrivalHospFile
		bd.Gycx = gycx
	} else {

		cpaDesName, notArrivalHospDesName, gycDesName, err = pushCpaGyc(bd.Cpa, bd.Gycx)

		cpa_tmp, _ := client.HGet(jobCall, "cpa").Result()
		if cpa_tmp != "" {
			notArrivalHospFile_tmp, _ := client.HGet(jobCall, "not_arrival_hosp_file").Result()
			gycx_tmp, _ := client.HGet(jobCall, "gycx").Result()

			bd.Cpa = cpa_tmp
			bd.NotArrivalHospFile = notArrivalHospFile_tmp
			bd.Gycx = gycx_tmp
			return nil
		}

		bd.Cpa = cpaDesName
		bd.NotArrivalHospFile = notArrivalHospDesName
		bd.Gycx = gycDesName

		_, err = client.HSet(jobCall, "cpa", bd.Cpa).Result()
		_, err = client.HSet(jobCall, "not_arrival_hosp_file", bd.NotArrivalHospFile).Result()
		_, err = client.HSet(jobCall, "gycx", bd.Gycx).Result()

		client.Expire(jobCall, 60 * time.Second)
	}
	bd.Id = bd.JobID
	bd.Date = time.Now().String()
	return err

}

func pushCpaGyc(cpa string, gyc string) (string, string, string, error) {
	var err error
	var cpaCsv string
	var gycCsv string
	var cpaDesName string
	var gycDesName string
	var notArrivalHospCsv string
	var notArrivalHospDesName string
	if cpa != "" {
		cpaCsv, notArrivalHospCsv, err = cpa2csv(cpa)
		cpaDesName, err = push2hdfs(cpaCsv)
		notArrivalHospDesName, err = push2hdfs(notArrivalHospCsv)
	}
	if gyc != "" {
		gycCsv, err = gyc2csv(gyc)
		gycDesName, err = push2hdfs(gycCsv)
	}
	return cpaDesName, notArrivalHospDesName, gycDesName, err
}

func cpa2csv(cpaFile string) (string, string, error) {
	var err error
	var cpa string
	var notArrivalHosp string
	localCpa := "resource/" + cpaFile
	cpa, err = bmexcelhandle.GenerateCSVFromXLSXFile(localCpa, 0)
	notArrivalHosp, err = bmexcelhandle.GenerateCSVFromXLSXFile(localCpa, 1)
	os.Remove(localCpa)
	return cpa, notArrivalHosp, err
}

func gyc2csv(gycFile string) (string, error) {
	var err error
	var gyc string
	localGyc := "resource/" + gycFile
	gyc, err = bmexcelhandle.GenerateCSVFromXLSXFile(localGyc, 0)
	os.Remove(localGyc)
	return gyc, err
}

func push2hdfs(localFile string) (string, error) {
	localDir := localFile
	fileDesName, _ := uuid.GenerateUUID()
	fmt.Println(fileDesName)
	fileDesPath := "/workData/Client/" + fileDesName
	client, _ := hdfs.New("192.168.100.137:9000")
	err := client.CopyToRemote(localDir, fileDesPath)
	os.Remove(localDir)
	return fileDesName, err
}