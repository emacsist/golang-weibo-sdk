package comment

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// CommentsTimelineOtherString : 获取某个用户发出和收到的评论列表
// http://open.weibo.com/wiki/C/2/comments/timeline/other
func CommentsTimelineOtherString(param apiParam.CommentsTimelineOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.CommentsTimelineOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// CommentsTimelineOtherObject : 以对象的方式返回
func CommentsTimelineOtherObject(param apiParam.CommentsTimelineOtherParam, accessToken string) (*apiResp.CommentsTimelineOtherResp, *apiResp.ErrorResp) {
	body, error := CommentsTimelineOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.CommentsTimelineOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "CommentsTimelineOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
