package samplecheck

import (
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/bmredis"
	"gopkg.in/mgo.v2/bson"
)

type SampleCheckSelecter struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	JobID     string        `json:"job_id" bson:"job_id"`
	CompanyID string        `json:"company_id" bson:"company_id"`
	YmList    []interface{} `json:"ym_list" bson:"ym_list"`
	MktList   []interface{} `json:"mkt_list" bson:"mkt_list"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *SampleCheckSelecter) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *SampleCheckSelecter) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *SampleCheckSelecter) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *SampleCheckSelecter) QueryId() string {
	return bd.Id
}

func (bd *SampleCheckSelecter) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *SampleCheckSelecter) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd SampleCheckSelecter) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd SampleCheckSelecter) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *SampleCheckSelecter) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *SampleCheckSelecter) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *SampleCheckSelecter) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (scs *SampleCheckSelecter) GetAllYms() SampleCheckSelecter {
	client := bmredis.GetRedisClient()
	defer client.Close()
	ymLst, err := client.SMembers(scs.JobID + "ym").Result()
	if err != nil {
		panic("no yms found")
	}
	rst := make([]interface{}, 0)
	for _, v := range ymLst {
		rst = append(rst, v)
	}
	scs.YmList = rst
	return *scs

}

func (scs *SampleCheckSelecter) GetAllMkt() SampleCheckSelecter {
	req := request.Request{}
	req.Res = "PhCalcConf"
	eq := request.Eqcond{}
	eq.Ky = "company_id"
	eq.Vy = scs.CompanyID
	var condi1 []interface{}
	condi1 = append(condi1, eq)
	req = req.SetConnect("Eqcond", condi1).(request.Request)

	var reval []max.PhCalcConf
	err := bmmodel.FindMutil(req, &reval)
	if err != nil {
		panic("no PhCalcConf found")
	}
	rst := make([]interface{}, 0)
	for _, v := range reval {
		rst = append(rst, v.Mkt)
	}
	scs.MktList = rst
	return *scs
}
