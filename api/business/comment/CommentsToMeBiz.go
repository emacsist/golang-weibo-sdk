package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsToMeBizString : 获取某个用户发出的评论列表
// http://open.weibo.com/wiki/C/2/comments/to_me/biz
func CommentsToMeBizString(param apiParam.CommentsToMeBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsToMeBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsToMeBizObject : 以对象的方式返回
func CommentsToMeBizObject(param apiParam.CommentsToMeBizParam, accessToken string) (*apiResp.CommentsToMeBizResp, *apiResp.ErrorResp) {
	body, error := CommentsToMeBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsToMeBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsToMeBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
