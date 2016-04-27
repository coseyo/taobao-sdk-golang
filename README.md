# taobao-sdk-golang

### 注意
请替换自己的appkey、appsecret

### 用法示例见src/main.go文件
```
/**
 * 淘宝SDK，go语言实现版
 * 用法示例 :
 */ /*

	var reqParams map[string]interface{}
	reqParams = make(map[string]interface{})
	reqParams["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url"

	sdk := taobao.NewSDK()
	result := sdk.Execute("taobao.tbk.item.get", reqParams)
	fmt.Println("\nresult ==> \n", result)

*/
```

### 运行
> make build && make run
