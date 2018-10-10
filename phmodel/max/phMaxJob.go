package max

import (
	"errors"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2/bson"
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
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	jobID := bd.JobID
	call := bd.Call
	jobCall := jobID + call

	//defer client.Del(jobCall)

	r, err := client.Get(jobCall).Result()
	if r == "" {
		return nil
	} else if r != "" {
		return errors.New("exist job_id call job")
	} else {
		return err
	}
}

func (bd *PHMaxJob) PushJobIdCall() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	pipe := client.Pipeline()

	jobID := bd.JobID
	call := bd.Call
	jobCall := jobID + call

	pipe.Incr(jobCall)
	pipe.Expire(jobCall, 1*time.Minute)
	_, err := pipe.Exec()

	fmt.Println(jobCall)
	return err
}
