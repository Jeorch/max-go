package maxgenerate

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmcommon/bmsingleton/bmpkg"
	"github.com/Jeorch/max-go/phmodel/max"
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"io/ioutil"
	"net/http"
)

type PhActionReadBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

func (b *PhActionReadBrick) Exec() error {
	var bmRouter bmconfig.BMRouterConfig
	bmRouter.GenerateConfig()
	maxConfig := b.bk.Pr.(max.PhMaxConfig)
	configPath := bmRouter.TmpDir + "/" + maxConfig.ConfigFile
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	jsonstr := string(file)
	rst, _ := jsonapi.FromJsonAPI(jsonstr)
	paction := rst.(max.PhAction)
	b.bk.Pr = paction
	return err
}

func (b *PhActionReadBrick) Prepare(pr interface{}) error {
	req := pr.(max.PhMaxConfig)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PhActionReadBrick) Done(pkg string, idx int64, e error) error {
	tmp, _ := bmpkg.GetPkgLen(pkg)
	if int(idx) < tmp-1 {
		bmrouter.NextBrickRemote(pkg, idx+1, b)
	}
	return nil
}

func (b *PhActionReadBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PhActionReadBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(max.PhAction)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PhActionReadBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval max.PhAction = b.BrickInstance().Pr.(max.PhAction)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
