package maxjob

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
)

type PHMaxJob struct {
	Id        		string            	`json:"id"`
	Id_       		bson.ObjectId     	`bson:"_id"`

	UserID	  		string 				`json:"user_id" bson:"user_id"`
	CompanyID	  	string 				`json:"company_id" bson:"company_id"`
	Date	  		string 				`json:"date" bson:"date"`
	Call	  		string 				`json:"call" bson:"call"`
	Args	  		map[string]interface{} 				`json:"args" bson:"args"`

}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PHMaxJob) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PHMaxJob) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PHMaxJob) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PHMaxJob) QueryId() string {
	return bd.Id
}

func (bd *PHMaxJob) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PHMaxJob) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PHMaxJob) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PHMaxJob) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PHMaxJob) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PHMaxJob) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PHMaxJob) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}