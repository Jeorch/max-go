package xmpp

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhXmppConf struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	//XmppHost    string `json:"xmpp_host" bson:"xmpp_host"`
	//XmppPort    string `json:"xmpp_port" bson:"xmpp_port"`
	//XmppUser    string `json:"xmpp_user" bson:"xmpp_user"`
	//XmppPwd     string `json:"xmpp_pwd" bson:"xmpp_pwd"`
	//XmppListens string `json:"xmpp_listens" bson:"xmpp_listens"`
	XmppReport  string `json:"xmpp_report" bson:"xmpp_report"`
	//XmppPoolNum string `json:"xmpp_pool_num" bson:"xmpp_pool_num"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhXmppConf) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhXmppConf) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhXmppConf) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhXmppConf) QueryId() string {
	return bd.Id
}

func (bd *PhXmppConf) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhXmppConf) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhXmppConf) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd PhXmppConf) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhXmppConf) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhXmppConf) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhXmppConf) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhXmppConf) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
