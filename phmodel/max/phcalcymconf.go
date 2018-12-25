package max

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhCalcYmConf struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CompanyId string      `json:"company_id" bson:"company_id"`
	JarPath   string      `json:"jar_path" bson:"jar_path"`
	Clazz     string      `json:"clazz" bson:"clazz"`
	Conf      interface{} `json:"conf"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhCalcYmConf) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhCalcYmConf) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhCalcYmConf) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhCalcYmConf) QueryId() string {
	return bd.Id
}

func (bd *PhCalcYmConf) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhCalcYmConf) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhCalcYmConf) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhCalcYmConf) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhCalcYmConf) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhCalcYmConf) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhCalcYmConf) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhCalcYmConf) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
