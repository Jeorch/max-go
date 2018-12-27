package export

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/colinmarc/hdfs"
	"github.com/hashicorp/go-uuid"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ExportMaxResult struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	JobID      string `json:"job_id" bson:"job_id"`
	CompanyID  string `json:"company_id" bson:"company_id"`
	Ym         string `json:"ym" bson:"ym"`
	Market     string `json:"market" bson:"market"`
	ResultPath string `json:"result_path" bson:"result_path"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *ExportMaxResult) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *ExportMaxResult) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *ExportMaxResult) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *ExportMaxResult) QueryId() string {
	return bd.Id
}

func (bd *ExportMaxResult) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *ExportMaxResult) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd ExportMaxResult) SetConnect(tag string, v interface{}) interface{} {
	return bd
}

func (bd ExportMaxResult) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *ExportMaxResult) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *ExportMaxResult) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *ExportMaxResult) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

func (bd *ExportMaxResult) Save2Local() {
	uuidTmp, _ := uuid.GenerateUUID()
	resultPath := bd.ResultPath
	originPath := resultPath[7:]
	destName := uuidTmp + "-" + bd.Ym + "-" + bd.Market + ".csv"
	//TODO:resource 文件配置化
	destPath := "/go/src/github.com/Jeorch/max-go/resource/" + destName
	//destPath := "resource/" + destName
	copyHdfsResultFile2Local(originPath, destPath)
	bd.ResultPath = destName
}

func copyHdfsResultFile2Local(originPath string, destName string) error {

	client, _ := hdfs.New("192.168.100.137:9000")

	defer client.Close()

	p, err := client.ReadDir(originPath)

	if err != nil {
		fmt.Println(err.Error())
	}
	for _,f := range p {
		if strings.HasSuffix(f.Name(), ".csv") {
			err = client.CopyToLocal(originPath + "/" + f.Name(), destName)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
