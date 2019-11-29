package max

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhCalcConf struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Ym            string                 `json:"ym" bson:"ym"`
	Mkt           string                 `json:"mkt" bson:"mkt"`
	CompanyId     string                 `json:"company_id" bson:"company_id"`
	PanelName     string                 `json:"panel_name" bson:"panel_name"`
	MaxName       string                 `json:"max_name" bson:"max_name"`
	MaxSearchName string                 `json:"max_search_name" bson:"max_search_name"`
	JarPath       string                 `json:"jar_path" bson:"jar_path"`
	Clazz         string                 `json:"clazz" bson:"clazz"`
	Conf          map[string]interface{} `json:"conf" bson:"conf"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhCalcConf) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhCalcConf) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhCalcConf) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhCalcConf) QueryId() string {
	return bd.Id
}

func (bd *PhCalcConf) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhCalcConf) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhCalcConf) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhCalcConf) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhCalcConf) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhCalcConf) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhCalcConf) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhCalcConf) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
