package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsMentionsOtherString : 获取某个用户收到的评论列表
// http://open.weibo.com/wiki/C/2/comments/to_me/other
func CommentsMentionsOtherString(param apiParam.CommentsMentionsOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsMentionsOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsMentionsOtherObject : 以对象的方式返回
func CommentsMentionsOtherObject(param apiParam.CommentsMentionsOtherParam, accessToken string) (*apiResp.CommentsMentionsOtherResp, *apiResp.ErrorResp) {
	body, error := CommentsMentionsOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsMentionsOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsMentionsOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
