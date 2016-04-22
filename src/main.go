package main

import (
	"fmt"
	"taobao"
)

const (
	APPKEY      string = "23349763"
	APPSECRET   string = "531b017cbac1b5555ac2b1ddcc869c89"
	REST_URL    string = "http://gw.api.taobao.com/router/rest"
	FORMAT      string = "json"
	V           string = "2.0"
	SIGN_METHOD string = "md5"
)

func main() {
	auth := taobao.NewAuth(APPKEY, APPSECRET, REST_URL)
	var reqParams map[string]string
	reqParams = make(map[string]string)
	reqParams["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"

	result := auth.Execute("taobao.tbk.item.get", reqParams)
	fmt.Println("\nresult ==> \n", result)

	reqParams["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"
	reqParams["num_iids"] = "6535538417,39442448794,6956495372,,45587889166"
	result2 := auth.Execute("taobao.tbk.item.info.get", reqParams)

	fmt.Println("\nresult2 ==> \n", result2)
}
