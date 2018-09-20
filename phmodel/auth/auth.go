package auth

import (
	"github.com/Jeorch/max-go/phmodel/profile"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PHAuth struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Profile profile.PHProfile `json:"profile" jsonapi:"relationships"`

	Token string `json:"token"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PHAuth) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PHAuth) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PHAuth) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PHAuth) QueryId() string {
	return bd.Id
}

func (bd *PHAuth) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PHAuth) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PHAuth) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "profile":
		bd.Profile = v.(profile.PHProfile)
	}
	return bd
}

func (bd PHAuth) QueryConnect(tag string) interface{} {
	switch tag {
	case "profile":
		return bd.Profile
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PHAuth) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PHAuth) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PHAuth) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
