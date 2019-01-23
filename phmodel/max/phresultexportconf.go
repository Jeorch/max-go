package max

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhResultExportConf struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	MaxName    string `json:"max_name" bson:"max_name"`
	ExportName string `json:"export_name" bson:"export_name"`
	Clazz      string `json:"clazz" bson:"clazz"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhResultExportConf) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhResultExportConf) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhResultExportConf) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhResultExportConf) QueryId() string {
	return bd.Id
}

func (bd *PhResultExportConf) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhResultExportConf) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhResultExportConf) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhResultExportConf) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhResultExportConf) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhResultExportConf) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhResultExportConf) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhResultExportConf) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
