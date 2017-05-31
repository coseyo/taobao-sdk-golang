package taobao

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
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
	FORMAT      string = "json"
	V           string = "2.0"
	SIGN_METHOD string = "md5"
)

func NewSDK(APPKey, APPSecret, RestURL string) *Auth {
	return &Auth{appkey: APPKey, appsecret: APPSecret, requrl: RestURL}
}

func (this *Auth) invoke(method string, params map[string]string, methodType string) (map[string]interface{}, error) {
	params["method"] = method
	respMap, err := this.request(params, methodType)
	return respMap, err
}

func (this *Auth) request(params map[string]string, methodType string) (respMap map[string]interface{}, err error) {
	args := Params{
		format:      FORMAT,
		app_key:     this.appkey,
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

	resp, err := http.Post(this.requrl, "application/x-www-form-urlencoded", strings.NewReader(urlParams.Encode()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &respMap)
	return
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
	for _, k := range keys {
		basestring += k + paramsMap[k]
	}
	basestring += this.appsecret

	data := []byte(basestring)
	signedString := fmt.Sprintf("%x", md5.Sum(data))
	return strings.ToUpper(signedString)
}

func (this *Auth) Execute(apiname string, params map[string]interface{}) (respMap map[string]interface{}, err error) {
	var paramsMap map[string]string
	paramsMap = make(map[string]string)

	for k, v := range params {
		if s, ok := v.(string); ok {
			paramsMap[k] = s
		} else if _, ok := v.(int); ok {
			paramsMap[k] = strconv.Itoa(v.(int))
		} else {
			err = errors.New("格式错误，map 格式只支持string, int")
			return
		}
	}

	respMap, err = this.invoke(apiname, paramsMap, "POST")
	return
}
