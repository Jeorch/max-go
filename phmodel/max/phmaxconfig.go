package max

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhMaxConfig struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ConfigFile string `json:"configFile"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhMaxConfig) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhMaxConfig) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhMaxConfig) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhMaxConfig) QueryId() string {
	return bd.Id
}

func (bd *PhMaxConfig) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhMaxConfig) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhMaxConfig) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhMaxConfig) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhMaxConfig) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhMaxConfig) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhMaxConfig) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhMaxConfig) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
