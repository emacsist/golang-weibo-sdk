package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsCreateBizString : 评论 一条微博.
// http://open.weibo.com/wiki/C/2/comments/create/biz
func CommentsCreateBizString(param apiParam.CommentsCreateBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URLParam := apiHelper.BuildPostQuery(&param, accessToken)
	logrus.Infof("api invoke [url=%v, param=%v]\n", apiURL.CommentsCreateBizURL, URLParam)
	return apiHelper.APIPost(apiURL.CommentsCreateBizURL, URLParam)
}

// CommentsCreateBizObject : 以对象的方式返回
func CommentsCreateBizObject(param apiParam.CommentsCreateBizParam, accessToken string) (*apiResp.CommentsCreateBizResp, *apiResp.ErrorResp) {
	body, error := CommentsCreateBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.CommentsCreateBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsCreateBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
