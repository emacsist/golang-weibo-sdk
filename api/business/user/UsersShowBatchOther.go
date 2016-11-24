package user

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// UsersShowBatchOtherString : 批量获取其他用户的基本信息
// http://open.weibo.com/wiki/C/2/users/show_batch/other

/*
{"users":[{"id":5225532117,"idstr":"5225532117","class":1,"screen_name":"emacsist","name":"emacsist","province":"100","city":"1000","location":"其他","description":"2015, 加油.","url":"http://emacsist.github.io","profile_image_url":"http://tva1.sinaimg.cn/crop.0.0.179.179.50/005HDNC5gw1ejz5rhva7cj3050050glj.jpg","profile_url":"emacsist","domain":"emacsist","weihao":"","gender":"m","followers_count":16,"friends_count":62,"pagefriends_count":0,"statuses_count":24,"favourites_count":1,"created_at":"Mon Jul 21 14:11:38 +0800 2014","following":false,"allow_all_act_msg":false,"geo_enabled":true,"verified":false,"verified_type":-1,"remark":"","status":{"created_at":"Wed Oct 19 18:58:13 +0800 2016","id":4032367087971584,"mid":"4032367087971584","idstr":"4032367087971584","text":"广发信用卡 哈哈","textLength":15,"source_allowclick":0,"source_type":1,"source":"<a href=\"http://app.weibo.com/t/feed/6vtZb0\" rel=\"nofollow\">微博 weibo.com</a>","appid":780,"favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","pic_ids":[],"geo":null,"reposts_count":0,"comments_count":0,"attitudes_count":0,"isLongText":false,"mlevel":0,"visible":{"type":0,"list_id":0},"biz_feature":0,"hasActionTypeCard":0,"darwin_tags":[],"hot_weibo_tags":[],"text_tag_tips":[],"userType":0,"positive_recom_flag":0,"gif_ids":"","is_show_bulletin":2},"ptype":0,"allow_all_comment":true,"avatar_large":"http://tva1.sinaimg.cn/crop.0.0.179.179.180/005HDNC5gw1ejz5rhva7cj3050050glj.jpg","avatar_hd":"http://tva1.sinaimg.cn/crop.0.0.179.179.1024/005HDNC5gw1ejz5rhva7cj3050050glj.jpg","verified_reason":"","verified_trade":"","verified_reason_url":"","verified_source":"","verified_source_url":"","follow_me":false,"online_status":0,"bi_followers_count":2,"lang":"zh-cn","star":0,"mbtype":0,"mbrank":0,"block_word":0,"block_app":0,"credit_score":80,"user_ability":0,"urank":9}],"total_number":1,"states":[{"5225532117":1}]}
*/
func UsersShowBatchOtherString(param apiParam.UsersShowBatchOtherParam, accessToken string) (string, *apiResp.ErrorResp) {
	URL := apiHelper.BuildGetQuery(&param, apiURL.UsersShowBatchOtherURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// UsersShowBatchOtherObject : 以对象的方式返回
func UsersShowBatchOtherObject(param apiParam.UsersShowBatchOtherParam, accessToken string) (*apiResp.UsersShowBatchOtherResp, *apiResp.ErrorResp) {
	body, error := UsersShowBatchOtherString(param, accessToken)

	if error != nil {
		return nil, error
	}

	var resp apiResp.UsersShowBatchOtherResp
	jsonError := json.Unmarshal([]byte(body), &resp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "UsersShowBatchOtherObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &resp, nil
}
