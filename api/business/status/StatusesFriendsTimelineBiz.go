package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// StatusesFriendsTimelineBizString : 返回一条微博的全部转发微博列表.
// http://open.weibo.com/wiki/C/2/statuses/friends_timeline/biz
func StatusesFriendsTimelineBizString(param apiParam.StatusesFriendsTimelineBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.StatusesFriendsTimelineBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// StatusesFriendsTimelineBizObject : 以对象的方式返回
func StatusesFriendsTimelineBizObject(param apiParam.StatusesFriendsTimelineBizParam, accessToken string) (*apiResp.StatusesFriendsTimelineBizResp, *apiResp.ErrorResp) {
	body, error := StatusesFriendsTimelineBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var statusesResp apiResp.StatusesFriendsTimelineBizResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "StatusesFriendsTimelineBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, nil
}
