package max

import (
	"github.com/Jeorch/max-go/phmodel/xmpp"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type PhAction struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	JobId      string   `json:"job_id" bson:"job_id"`
	UserId     string   `json:"user_id" bson:"user_id"`
	CompanyId  string   `json:"company_id" bson:"company_id"`
	//Yms        []string `json:"yms" bson:"yms"`
	PanelPath  string   `json:"panel_path" bson:"panel_path"`
	MaxPath    string   `json:"max_path" bson:"max_path"`
	ProdLst    string   `json:"prod_lst" bson:"prod_lst"`
	//CreateTime int64    `json:"create_time" bson:"create_time"`

	XmppConf   xmpp.PhXmppConf `json:"xmppConf" jsonapi:"relationships"`
	CalcYmConf PhCalcYmConf    `json:"calcYmConf" jsonapi:"relationships"`
	PanelConf  []PhPanelConf   `json:"panelConf" jsonapi:"relationships"`
	CalcConf   []PhCalcConf    `json:"calcConf" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhAction) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhAction) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhAction) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhAction) QueryId() string {
	return bd.Id
}

func (bd *PhAction) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhAction) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhAction) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "xmppConf":
		bd.XmppConf = v.(xmpp.PhXmppConf)
	case "calcYmConf":
		bd.CalcYmConf = v.(PhCalcYmConf)
	case "panelConf":
		var rst []PhPanelConf
		for _, item := range v.([]interface{}) {
			tmp := item.(PhPanelConf)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.PanelConf = rst
	case "calcConf":
		var rst []PhCalcConf
		for _, item := range v.([]interface{}) {
			tmp := item.(PhCalcConf)
			if len(tmp.Id) > 0 {
				rst = append(rst, tmp)
			}
		}
		bd.CalcConf = rst
	}
	return bd
}

func (bd PhAction) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhAction) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhAction) CoverBMObject() error {
	return bmmodel.CoverOne(bd)
}

func (bd *PhAction) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhAction) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
