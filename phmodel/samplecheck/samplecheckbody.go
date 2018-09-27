package samplecheck

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type SampleCheckBody struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	JobID           string                 `json:"job_id" bson:"job_id"`
	UserID          string                 `json:"user_id" bson:"user_id"`
	CompanyID       string                 `json:"company_id" bson:"company_id"`
	Ym              string                 `json:"ym" bson:"ym"`
	Market          string                 `json:"market" bson:"market"`
	Hospital        map[string]interface{} `json:"hospital" bson:"hospital"`
	Product         map[string]interface{} `json:"product" bson:"product"`
	Sales           map[string]interface{} `json:"sales" bson:"sales"`
	NotFindHospital []interface{}          `json:"notfindhospital" bson:"notfindhospital"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *SampleCheckBody) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *SampleCheckBody) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *SampleCheckBody) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *SampleCheckBody) QueryId() string {
	return bd.Id
}

func (bd *SampleCheckBody) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *SampleCheckBody) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd SampleCheckBody) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd SampleCheckBody) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *SampleCheckBody) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *SampleCheckBody) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *SampleCheckBody) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
