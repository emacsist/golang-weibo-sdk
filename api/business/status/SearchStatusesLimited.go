//Package api 与微博相关的包，这里是关键词搜索的包
package status

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
	apiURL "github.com/emacsist/golang-weibo-sdk/url"
)

// SearchStatusesLimited 调用搜索接口
// 返回的结果如下
/*
{
    "statuses": [
        {
            "created_at": "Wed Oct 19 15:10:18 +0800 2016",
            "id": 4032309726687993,
            "mid": "4032309726687993",
            "idstr": "4032309726687993",
            "text": "\u3010\u7ffb\u8bd1\u3011\u4f7f\u7528 sync.ErrGroup \u5b9e\u73b0\u5e76\u53d1\u641c\u7d22\u6587\u4ef6 - Go \u6280\u672f\u793e\u533a - golang http:\/\/t.cn\/RVS3DRm",
            "textLength": 86,
            "source_allowclick": 0,
            "source_type": 1,`json:"appid"`
            "source": "<a href=\"http:\/\/app.weibo.com\/t\/feed\/4PaMVH\" rel=\"nofollow\">JiaThis\u5206\u4eab\u6309\u94ae<\/a>",
            "appid": 10327,
            "favorited": false,
            "truncated": false,
            "in_reply_to_status_id": "",
            "in_reply_to_user_id": "",
            "in_reply_to_screen_name": "",
            "pic_ids": [],
            "geo": null,
            "user": {
                "id": 1889019865,
                "idstr": "1889019865",
                "class": 1,
                "screen_name": "ASTA\u8c22",
                "name": "ASTA\u8c22",
                "province": "31",
                "city": "4",
                "location": "\u4e0a\u6d77 \u5f90\u6c47\u533a",
                "description": "beego\u4f5c\u8005\uff0c\u300aGo Web\u7f16\u7a0b\u300b\u4f5c\u8005\u3002\u6210\u529f\u7684\u4e60\u60ef\u6bd4\u6210\u529f\u66f4\u91cd\u8981",
                "url": "http:\/\/github.com\/astaxie",
                "profile_image_url": "http:\/\/tva4.sinaimg.cn\/crop.92.10.266.266.50\/709827d9jw8ezx11051d9j20dc08wwer.jpg",
                "profile_url": "533452688",
                "domain": "kaibao001",
                "weihao": "533452688",
                "gender": "m",
                "followers_count": 12519,
                "friends_count": 339,
                "pagefriends_count": 2,
                "statuses_count": 3534,
                "favourites_count": 30,
                "created_at": "Fri Dec 10 22:01:45 +0800 2010",
                "following": false,
                "allow_all_act_msg": false,
                "geo_enabled": false,
                "verified": false,
                "verified_type": 220,
                "remark": "",
                "ptype": 0,
                "allow_all_comment": false,
                "avatar_large": "http:\/\/tva4.sinaimg.cn\/crop.92.10.266.266.180\/709827d9jw8ezx11051d9j20dc08wwer.jpg",
                "avatar_hd": "http:\/\/tva4.sinaimg.cn\/crop.92.10.266.266.1024\/709827d9jw8ezx11051d9j20dc08wwer.jpg",
                "verified_reason": "",
                "verified_trade": "",
                "verified_reason_url": "",
                "verified_source": "",
                "verified_source_url": "",
                "follow_me": false,
                "online_status": 0,
                "bi_followers_count": 257,
                "lang": "zh-cn",
                "star": 0,
                "mbtype": 0,
                "mbrank": 0,
                "block_word": 0,
                "block_app": 0,
                "level": 7,
                "type": 1,
                "ulevel": 0,
                "badge": {
                    "uc_domain": 0,
                    "enterprise": 0,
                    "anniversary": 0,
                    "taobao": 0,
                    "travel2013": 0,
                    "gongyi": 0,
                    "gongyi_level": 0,
                    "bind_taobao": 1,
                    "hongbao_2014": 1,
                    "suishoupai_2014": 0,
                    "dailv": 0,
                    "zongyiji": 0,
                    "vip_activity1": 0,
                    "unread_po
	TotalNumber int32 ol": 0,
                    "daiyan": 0,
                    "ali_1688": 0,
                    "vip_activity2": 0,
                    "suishoupai_2016": 0,
                    "fools_day_2016": 0,
                    "uefa_euro_2016": 0,
                    "super_star_2016": 0,
                    "unread_pool_ext": 0,
                    "self_media": 0,
                    "olympic_2016": 0,
                    "dzwbqlx_2016": 1,
                    "discount_2016": 0,
                    "wedding_2016": 0,
                    "shuang11_2016": 0
                },
                "badge_top": "",
                "has_ability_tag": 1,
                "extend": {
                    "privacy": {
                        "mobile": 0
                    },
                    "mbprivilege": "0000000000000000000000000000000000000000000000000000000000000000"
                },
                "credit_score": 80,
                "user_ability": 8,
                "urank": 33
            },
            "reposts_count": 0,
            "comments_count": 0,
            "attitudes_count": 1,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 0,
            "url_objects": [
                {
                    "url_ori": "http:\/\/t.cn\/RVS3DRm",
                    "info": {
                        "url_short": "http:\/\/t.cn\/RVS3DRm",
                        "url_long": "https:\/\/gocn.io\/question\/161",
                        "type": 0,
                        "result": true,
                        "title": "\u3010\u7ffb\u8bd1\u3011\u4f7f\u7528 sync.ErrGroup \u5b9e\u73b0\u5e76\u53d1\u641c\u7d22\u6587\u4ef6 - Go \u6280\u672f\u793e\u533a - golang ",
                        "description": "",
                        "last_modified": 1476861010,
                        "transcode": 0
                    },
                    "like_count": 0,
                    "isActionType": false,
                    "follower_count": 0,
                    "asso_like_count": 0,
                    "card_info_un_integrity": false,
                    "super_topic_status_count": 0,
                    "super_topic_photo_count": 0
                }
            ],
            "hasActionTypeCard": 0,
            "darwin_tags": [],
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "rid": "0_0_0_2633239956795357624",
            "userType": 0,
            "positive_recom_flag": 0,
            "gif_ids": "",
            "is_show_bulletin": 2,
            "status": 0,
            "relation": null,
            "category": 31,
            "base62_id": "Edywss3QR"
        }
    ],
    "total_number": 14836
}
*/
func SearchStatusesLimitedString(param apiParam.SearchStatusesLimitedParam, accessToken string) (string, *apiResp.ErrorResp) {
	// return BuildQuery(&param, url.SearchStatusesLimitedURL, accessToken)
	var URL = apiHelper.BuildGetQuery(&param, apiURL.SearchStatusesLimitedURL, accessToken)
	logrus.Infof("api invoke [%v]\n", URL)
	return apiHelper.APIGet(URL)
}

// SearchStatusesLimitedObject ： 以对象的方式返回
func SearchStatusesLimitedObject(param apiParam.SearchStatusesLimitedParam, accessToken string) (*apiResp.SearchStatusesLimitedResp, *apiResp.ErrorResp) {
	body, error := SearchStatusesLimitedString(param, accessToken)
	if error != nil {
		return nil, error
	}
	var statusesResp apiResp.SearchStatusesLimitedResp
	jsonError := json.Unmarshal([]byte(body), &statusesResp)

	if jsonError != nil {
		return nil, &apiResp.ErrorResp{Error: jsonError.Error(), ErrorCode: apiResp.JSONParseErrorCode, Request: "SearchStatusesLimitedObject" + "==>" + apiResp.JSONParseErrorCodeMsg}
	}
	return &statusesResp, error
}
