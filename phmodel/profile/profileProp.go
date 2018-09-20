package profile

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PHProfileProp struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	ProfileID string `json:"profileid" bson:"profile_id"`
	CompanyID string `json:"companyid" bson:"company_id"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PHProfileProp) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PHProfileProp) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PHProfileProp) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PHProfileProp) QueryId() string {
	return bd.Id
}

func (bd *PHProfileProp) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PHProfileProp) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PHProfileProp) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PHProfileProp) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PHProfileProp) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PHProfileProp) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PHProfileProp) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
