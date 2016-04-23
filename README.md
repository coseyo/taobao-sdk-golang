# taobao-sdk-golang

### 注意
请替换自己的appkey、appsecret

### 用法示例见src/main.go文件
```
/**
 * 淘宝SDK，go语言实现版
 * 用法示例 :
 */ /*

	var reqParams map[string]string
	reqParams = make(map[string]string)
	reqParams["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"
	reqParams["num_iids"] = "6535538417,39442448794,6956495372,,45587889166"

	sdk := taobao.NewSDK()
	result := sdk.Execute(Address, reqParams)

	fmt.Println("\nresult ==> \n", result)

*/
```