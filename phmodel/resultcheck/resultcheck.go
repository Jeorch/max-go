package resultcheck

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type ResultCheck struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	JobID      string                 `json:"job_id" bson:"job_id"`
	UserID     string                 `json:"user_id" bson:"user_id"`
	CompanyID  string                 `json:"company_id" bson:"company_id"`
	Ym         string                 `json:"ym" bson:"ym"`
	Market     string                 `json:"market" bson:"market"`
	Indicators map[string]interface{} `json:"indicators" bson:"indicators"`
	Trend      []interface{}          `json:"trend" bson:"trend"`
	Region     []interface{}          `json:"region" bson:"region"`
	Mirror     map[string]interface{} `json:"mirror" bson:"mirror"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *ResultCheck) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *ResultCheck) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *ResultCheck) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *ResultCheck) QueryId() string {
	return bd.Id
}

func (bd *ResultCheck) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *ResultCheck) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd ResultCheck) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd ResultCheck) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *ResultCheck) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *ResultCheck) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *ResultCheck) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
