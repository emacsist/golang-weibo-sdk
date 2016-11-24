package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsReplyBizString : 回复一条评论.
// http://open.weibo.com/wiki/C/2/comments/reply/biz
func CommentsReplyBizString(param apiParam.CommentsReplyBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.CommentsReplyBizURL, URLParam)
	return apiHelper.APIPost(apiURL.CommentsReplyBizURL, URLParam)
}

// CommentsReplyBizObject : 以对象的方式返回
func CommentsReplyBizObject(param apiParam.CommentsReplyBizParam, accessToken string) (*apiResp.CommentsReplyBizResp, *apiResp.ErrorResp) {
	body, error := CommentsReplyBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.CommentsReplyBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsReplyBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
