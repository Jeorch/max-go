package company

import (
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PhCompany struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CompanyName string  `json:"companyname" bson:"company_name"`
	Process     float64 `json:"process" bson:"process"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhCompany) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhCompany) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhCompany) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhCompany) QueryId() string {
	return bd.Id
}

func (bd *PhCompany) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhCompany) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhCompany) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhCompany) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhCompany) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhCompany) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhCompany) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * profile interface
 *------------------------------------------------*/

func (bd PhCompany) IsCompanyRegisted() bool {
	var bmMongoConfig bmconfig.BMMongoConfig
	bmMongoConfig.GenerateConfig()
	session, err := mgo.Dial(bmMongoConfig.Host + ":" + bmMongoConfig.Port)
	if err != nil {
		return true
	}
	defer session.Close()

	c := session.DB("test").C("PhCompany")
	n, err := c.Find(bson.M{"company_name": bd.CompanyName}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd PhCompany) Valid() bool {
	return bd.CompanyName != ""
}
