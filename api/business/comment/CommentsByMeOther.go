package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsByMeOtherString : 获取某个用户发出的评论列表
// http://open.weibo.com/wiki/C/2/comments/by_me/other
func CommentsByMeOtherString(param apiParam.CommentsByMeOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsByMeOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsByMeOtherObject : 以对象的方式返回
func CommentsByMeOtherObject(param apiParam.CommentsByMeOtherParam, accessToken string) (*apiResp.CommentsByMeOtherResp, *apiResp.ErrorResp) {
	body, error := CommentsByMeOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsByMeOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsByMeOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
