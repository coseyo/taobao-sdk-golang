package taobao

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

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Auth struct {
	appkey    string
	appsecret string
	requrl    string
}

type Params struct {
	app_key     string // appKey
	timestamp   string // 时间戳
	format      string // 相应格式 json
	v           string // api版本 2.0
	sign_method string // 签名的摘要算法md5
	// method      string // api接口名称
	// session     string // 是否需要授权
}

const (
	APPKEY      string = "2bb4qq63"
	APPSECRET   string = "53wb0r7cbac1bt555ac4b1ddcc569c89"
	REST_URL    string = "http://gw.api.taobao.com/router/rest"
	FORMAT      string = "json"
	V           string = "2.0"
	SIGN_METHOD string = "md5"
)

func NewSDK() *Auth {
	return &Auth{appkey: APPKEY, appsecret: APPSECRET, requrl: REST_URL}
}

func (this *Auth) invoke(method string, params map[string]string, methodType string) (string, error) {
	params["method"] = method
	resBody, err := this.request(params, methodType)
	return resBody, err
}

func (this *Auth) request(params map[string]string, methodType string) (string, error) {

	args := Params{
		// method:      params["method"],
		format:      FORMAT,
		app_key:     APPKEY,
		timestamp:   time.Now().Format("2006-01-02 15:04:05"),
		v:           V,
		sign_method: SIGN_METHOD,
	}

	urlParams := url.Values{}
	urlParams.Add("format", args.format)
	urlParams.Add("app_key", args.app_key)
	urlParams.Add("timestamp", args.timestamp)
	urlParams.Add("v", args.v)
	urlParams.Add("sign_method", args.sign_method)

	for k, v := range params {
		urlParams.Add(k, v)
	}
	urlParams.Add("sign", this.sign(args, params))

	baseUrl, err := url.Parse(this.requrl)
	if err != nil {
		panic(err)
	}
	baseUrl.RawQuery = urlParams.Encode()
	// fmt.Println(urlParams.Encode())
	fmt.Println("request ==> \n", baseUrl.String())
	// reqBody := []byte(baseUrl.String())

	resp, err := http.Post(this.requrl, "application/x-www-form-urlencoded", strings.NewReader(urlParams.Encode()))

	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	return body, nil
}

// Convert struct to map
func convert(params Params) map[string]string {
	var paramsMap map[string]string
	paramsMap = make(map[string]string)

	vt := reflect.TypeOf(params)
	vv := reflect.ValueOf(params)

	for i := 0; i < vt.NumField(); i++ {
		k := vt.Field(i).Name
		v := vv.FieldByName(k)
		paramsMap[k] = v.String()
	}
	return paramsMap
}

// Generate signature
func (this *Auth) sign(params Params, fields map[string]string) string {
	paramsMap := convert(params)

	for k, _ := range fields {
		paramsMap[k] = fields[k]
	}
	basestring := this.appsecret
	var keys []string
	for k := range paramsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		basestring += k + paramsMap[k]
	}
	basestring += this.appsecret
	fmt.Println(basestring)

	data := []byte(basestring)
	signedString := fmt.Sprintf("%x", md5.Sum(data))
	return strings.ToUpper(signedString)
}

func (this *Auth) Execute(apiname string, params map[string]interface{}) string {

	var paramsMap map[string]string
	paramsMap = make(map[string]string)

	for k, v := range params {
		if s, ok := v.(string); ok {
			paramsMap[k] = s
		} else if _, ok := v.(int); ok {
			paramsMap[k] = strconv.Itoa(v.(int))
		} else {
			panic("格式错误，map 格式只支持string, int")
		}
	}

	body, err := this.invoke(apiname, paramsMap, "POST")
	if err != nil {
		panic(err)
	}
	return body
}
