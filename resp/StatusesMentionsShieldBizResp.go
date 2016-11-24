package resp

// StatusesMentionsShieldBizResp ： 屏蔽某个@我的微博及后续由其转发引起的@提及 api结果对象
// 文档写着返回的是:
/*
{
    "result":true
}
*/
//实际返回的是 "result":"true"
type StatusesMentionsShieldBizResp struct {
	Result string `json:"result"`
}
