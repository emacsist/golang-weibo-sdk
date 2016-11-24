package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsToMeOtherString : 获取某个用户收到的评论列表
// http://open.weibo.com/wiki/C/2/comments/to_me/other
func CommentsToMeOtherString(param apiParam.CommentsToMeOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsToMeOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsToMeOtherObject : 以对象的方式返回
func CommentsToMeOtherObject(param apiParam.CommentsToMeOtherParam, accessToken string) (*apiResp.CommentsToMeOtherResp, *apiResp.ErrorResp) {
	body, error := CommentsToMeOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsToMeOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsToMeOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
