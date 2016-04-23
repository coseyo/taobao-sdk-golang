package main

import (
	"fmt"
	"taobao"
)

func main() {

	var reqParams map[string]interface{}
	reqParams = make(map[string]interface{})
	reqParams["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"

	sdk := taobao.NewSDK()
	result := sdk.Execute("taobao.tbk.item.get", reqParams)
	fmt.Println("\nresult ==> \n", result)

	var reqParams2 map[string]interface{}
	reqParams2 = make(map[string]interface{})
	reqParams2["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"
	reqParams2["num_iids"] = "6535538417,39442448794,6956495372,,45587889166"

	result2 := sdk.Execute("taobao.tbk.item.info.get", reqParams2)

	fmt.Println("\nresult2 ==> \n", result2)

	var reqParams3 map[string]interface{}
	reqParams3 = make(map[string]interface{})
	reqParams3["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"
	reqParams3["relate_type"] = 1
	reqParams3["num_iid"] = 6535538417

	result3 := sdk.Execute("taobao.tbk.item.recommend.get", reqParams3)

	fmt.Println("\nresult3 ==> \n", result3)
}
