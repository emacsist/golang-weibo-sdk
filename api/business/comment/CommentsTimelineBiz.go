package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsTimelineBizString : 获取当前登录用户发出及收到的评论列表
// http://open.weibo.com/wiki/C/2/comments/timeline/biz
func CommentsTimelineBizString(param apiParam.CommentsTimelineBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsTimelineBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsTimelineBizObject : 以对象的方式返回
func CommentsTimelineBizObject(param apiParam.CommentsTimelineBizParam, accessToken string) (*apiResp.CommentsTimelineBizResp, *apiResp.ErrorResp) {
	body, error := CommentsTimelineBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsTimelineBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsTimelineBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
