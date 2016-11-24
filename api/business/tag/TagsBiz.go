package tag

import (
	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// TagsBizString : 批量获取用户标签
// http://open.weibo.com/wiki/C/2/tags/biz
func TagsBizString(param apiParam.TagsBizParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.TagsBizURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// TagsBizObject : 以对象的方式返回
func TagsBizObject(param apiParam.TagsBizParam, accessToken string) (*apiResp.TagsBizResp, *apiResp.ErrorResp) {
	body, error := TagsBizString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.TagsBizResp
	resp.Body = body
	return &resp, nil
}
