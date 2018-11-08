package settings

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

var (
	Listen           = ":8080"
	HmacSampleSecret = []byte("whatever")
	LogFile          = "./beauty.log"
	Domain           = "****.com"
	DefaultOrigin    = "http://www.xsfuture.com"
	Local            = map[string]string{}
)

func init() {
	bytes, err := ioutil.ReadFile("./conf_latest.json")
	if err != nil {
		log.Println("readfile: ", err.Error())
		return
	}
	if json.Unmarshal(bytes, &Local); err != nil {
		log.Println("unmarshal: ", err.Error())
		return
	}
}
