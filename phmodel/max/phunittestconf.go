package max

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhUnitTestConf struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Ym            string            `json:"ym" bson:"ym"`
	Mkt           string            `json:"mkt" bson:"mkt"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhUnitTestConf) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhUnitTestConf) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhUnitTestConf) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhUnitTestConf) QueryId() string {
	return bd.Id
}

func (bd *PhUnitTestConf) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhUnitTestConf) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhUnitTestConf) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhUnitTestConf) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhUnitTestConf) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhUnitTestConf) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhUnitTestConf) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhUnitTestConf) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
