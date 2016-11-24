package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsDestroyBizString : 评论 一条微博.
// http://open.weibo.com/wiki/C/2/comments/destroy/biz
func CommentsDestroyBizString(param apiParam.CommentsDestroyBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.CommentsDestroyBizURL, URLParam)
	return apiHelper.APIPost(apiURL.CommentsDestroyBizURL, URLParam)
}

// CommentsDestroyBizObject : 以对象的方式返回
func CommentsDestroyBizObject(param apiParam.CommentsDestroyBizParam, accessToken string) (*apiResp.CommentsDestroyBizResp, *apiResp.ErrorResp) {
	body, error := CommentsDestroyBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.CommentsDestroyBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsDestroyBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
