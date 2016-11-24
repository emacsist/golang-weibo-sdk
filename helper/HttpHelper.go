package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	"github.com/emacsist/golang-weibo-sdk/url"
	"github.com/Sirupsen/logrus"
)

// APIGet ： API中get请求返回string的helper方法
func APIGet(URL string) (string, *apiResp.ErrorResp) {
	resp, err := http.Get(URL)
	if err != nil {
		// if error
		logrus.Warnf("get[%v] falied. reason : %v", URL, err.Error())
		return "", &apiResp.ErrorResp{Error: err.Error(), ErrorCode: apiResp.NetWorkErrorCode, Request: apiResp.NetWorkErrorCodeMsg}
	}

	return respToString(resp)
}

// APIPost ： API中post请求返回string的helper方法
func APIPost(URL string, paramBody string) (string, *apiResp.ErrorResp) {
	resp, err := http.Post(URL, url.ContentTypeURLEncoded, strings.NewReader(paramBody))
	if err != nil {
		// if error
		logrus.Warnf("post[url=%v, param=%v] falied. reason : %v\n", URL, paramBody, err.Error())
		return "", &apiResp.ErrorResp{Error: err.Error(), ErrorCode: apiResp.NetWorkErrorCode, Request: apiResp.NetWorkErrorCodeMsg}
	}
	return respToString(resp)
}

// APIPostMultiPart : 带上传文件的post请求
func APIPostMultiPart(req *http.Request) (string, *apiResp.ErrorResp) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// if error
		logrus.Warnf("post[url=%v, param=%v] falied. reason : %v\n", req.RequestURI, req.ParseForm(), err.Error())
		return "", &apiResp.ErrorResp{Error: err.Error(), ErrorCode: apiResp.NetWorkErrorCode, Request: apiResp.NetWorkErrorCodeMsg}
	}
	return respToString(resp)
}

func respToString(resp *http.Response) (string, *apiResp.ErrorResp) {
	defer resp.Body.Close()
	//OK
	body, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		logrus.Warnf("read searchStatusesLimited Response error : %v", error.Error())
		return "", &apiResp.ErrorResp{Error: error.Error(), ErrorCode: apiResp.IOErrorCode, Request: apiResp.IOErrorCodeMsg}
	}
	if resp.StatusCode != 200 {
		var apiRespError apiResp.ErrorResp
		jsonError := json.Unmarshal(body, &apiRespError)
		if jsonError == nil {
			//转换为json不出错的话，表示是微博的error_code
			return "", &apiRespError
		}
		return "", &apiResp.ErrorResp{Error: string(body), ErrorCode: resp.StatusCode, Request: resp.Request.URL.Host + resp.Request.URL.RequestURI()}
	}
	return string(body), nil
}
