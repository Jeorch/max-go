package max

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhCompanyProd struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	CompanyId   string `json:"company_id" bson:"company_id"`
	CompanyName string `json:"company_name" bson:"company_name"`
	ProdLst     []interface{} `json:"prod_lst" bson:"prod_lst"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhCompanyProd) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhCompanyProd) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhCompanyProd) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhCompanyProd) QueryId() string {
	return bd.Id
}

func (bd *PhCompanyProd) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhCompanyProd) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhCompanyProd) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhCompanyProd) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhCompanyProd) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhCompanyProd) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhCompanyProd) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhCompanyProd) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
