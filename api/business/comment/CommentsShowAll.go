package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsShowAllString : 返回一条微博的全部评论列表.
// http://open.weibo.com/wiki/C/2/comments/show/all
func CommentsShowAllString(param apiParam.CommentsShowAllParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsShowAllURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsShowAllObject : 以对象的方式返回
func CommentsShowAllObject(param apiParam.CommentsShowAllParam, accessToken string) (*apiResp.CommentsShowAllResp, *apiResp.ErrorResp) {
	body, error := CommentsShowAllString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsShowAllResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsShowAllObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
