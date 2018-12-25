package max

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhPanelConf struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Ym        string            `json:"ym" bson:"ym"`
	Mkt       string            `json:"mkt" bson:"mkt"`
	PanelName string            `json:"panel_name" bson:"panel_name"`
	JarPath   string            `json:"jar_path" bson:"jar_path"`
	Clazz     string            `json:"clazz" bson:"clazz"`
	Conf      map[string]interface{} `json:"conf" bson:"conf"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhPanelConf) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhPanelConf) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhPanelConf) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhPanelConf) QueryId() string {
	return bd.Id
}

func (bd *PhPanelConf) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhPanelConf) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhPanelConf) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhPanelConf) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhPanelConf) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhPanelConf) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhPanelConf) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhPanelConf) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
