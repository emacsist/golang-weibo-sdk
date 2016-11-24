package search

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// SearchSuggestionsUsersBizString : 搜用户搜索建议
// http://open.weibo.com/wiki/C/2/search/suggestions/users/biz
func SearchSuggestionsUsersBizString(param apiParam.SearchSuggestionsUsersBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.SearchSuggestionsUsersBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// SearchSuggestionsUsersBizObject : 以对象的方式返回
func SearchSuggestionsUsersBizObject(param apiParam.SearchSuggestionsUsersBizParam, accessToken string) (*[]apiResp.SearchSuggestionsUsersBizResp, *apiResp.ErrorResp) {
	body, error := SearchSuggestionsUsersBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp []apiResp.SearchSuggestionsUsersBizResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "SearchSuggestionsUsersBizObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
