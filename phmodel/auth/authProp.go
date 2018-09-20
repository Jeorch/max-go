package auth

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PHAuthProp struct {
	Id  string        `json:"Id"`
	Id_ bson.ObjectId `bson:"_id"`

	AuthID    string `json:"auth_id" bson:"auth_id"`
	ProfileID string `json:"profile_id" bson:"profile_id"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PHAuthProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PHAuthProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PHAuthProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PHAuthProp) QueryId() string {
	return bd.Id
}

func (bd *PHAuthProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PHAuthProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PHAuthProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PHAuthProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PHAuthProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PHAuthProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PHAuthProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
