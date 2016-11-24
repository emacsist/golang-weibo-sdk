package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsMentionsBizString : 获取@当前登录用户的评论
// http://open.weibo.com/wiki/C/2/comments/mentions/biz
func CommentsMentionsBizString(param apiParam.CommentsMentionsBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsMentionsBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsMentionsBizObject : 以对象的方式返回
func CommentsMentionsBizObject(param apiParam.CommentsMentionsBizParam, accessToken string) (*apiResp.CommentsMentionsBizResp, *apiResp.ErrorResp) {
	body, error := CommentsMentionsBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsMentionsBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsMentionsBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
