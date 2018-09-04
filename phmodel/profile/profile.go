package profile

import (
	"github.com/Jeorch/max-go/phmodel/company"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type PHProfile struct {
	Id        		string            	`json:"id"`
	Id_       		bson.ObjectId     	`bson:"_id"`

	Username 		string 				`json:"username" bson:"username"`
	Password		string 				`json:"password" bson:"password"`
	Company	  		company.PHCompany 	`json:"company" jsonapi:"relationships"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PHProfile) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PHProfile) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PHProfile) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PHProfile) QueryId() string {
	return bd.Id
}

func (bd *PHProfile) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PHProfile) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PHProfile) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "company":
		bd.Company = v.(company.PHCompany)
	}
	return bd
}

func (bd PHProfile) QueryConnect(tag string) interface{} {
	switch tag {
	case "company":
		return bd.Company
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PHProfile) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PHProfile) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PHProfile) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * profile interface
 *------------------------------------------------*/

func (bd PHProfile) IsUserRegisted() bool {
	session, err := mgo.Dial("max-mongo:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("PHProfile")
	n, err := c.Find(bson.M{"username": bd.Username}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd PHProfile) Valid() bool {
	return bd.Username != ""
}
