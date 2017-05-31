package main

import (
	"log"

	"github.com/coseyo/taobao-sdk-golang/taobao"
)

func main() {
	var reqParams map[string]interface{}
	reqParams = make(map[string]interface{})
	reqParams["platform"] = 1
	reqParams["q"] = "游戏机"
	reqParams["page_size"] = 2
	reqParams["page_no"] = 1
	reqParams["pid"] = "mm_xxxx"
	sdk := taobao.NewSDK("tttt", "eeee", "http://gw.api.taobao.com/router/rest")
	respMap, err := sdk.Execute("taobao.tbk.item.coupon.get", reqParams)
	log.Println("err", err)
	log.Println("respMap", respMap)

}
