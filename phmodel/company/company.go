package company

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PHCompany struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CompanyName string `json:"companyname" bson:"company_name"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PHCompany) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PHCompany) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PHCompany) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PHCompany) QueryId() string {
	return bd.Id
}

func (bd *PHCompany) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PHCompany) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PHCompany) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PHCompany) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PHCompany) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PHCompany) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PHCompany) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * profile interface
 *------------------------------------------------*/

func (bd PHCompany) IsCompanyRegisted() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic("dial db error")
	}
	defer session.Close()

	c := session.DB("test").C("PHCompany")
	n, err := c.Find(bson.M{"company_name": bd.CompanyName}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd PHCompany) Valid() bool {
	return bd.CompanyName != ""
}
