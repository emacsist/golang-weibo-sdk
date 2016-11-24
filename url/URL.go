package url

// CUrlPreFix : 商业api的前缀
const CUrlPreFix = "https://c.api.weibo.com"

// ParamNameAccessToken 参数accessToken的名字
const ParamNameAccessToken string = "access_token"

/* ========================微博内容数据（收费） START=====================================*/

//SearchStatusesLimitedURL : 搜索含某关键词的微博
const SearchStatusesLimitedURL string = CUrlPreFix + "/2/search/statuses/limited.json"

// StatusesRepostTimelineAllURL : 返回一条微博的全部转发微博列表 的URL
const StatusesRepostTimelineAllURL string = CUrlPreFix + "/2/statuses/repost_timeline/all.json"

// StatusesMentionsOtherURL : 获取@某人的微博
const StatusesMentionsOtherURL string = CUrlPreFix + "/2/statuses/mentions/other.json"

// StatusesUserTimelineBatchURL : 批量获取用户个人微博列表
const StatusesUserTimelineBatchURL string = CUrlPreFix + "/2/statuses/user_timeline_batch.json"

// PlaceUserTimelineOtherURL :获取某个用户的位置动态
const PlaceUserTimelineOtherURL string = CUrlPreFix + "/2/place/user_timeline/other.json"

// CommentsShowAllURL : 返回一条微博的全部评论列表
const CommentsShowAllURL string = CUrlPreFix + "/2/comments/show/all.json"

// CommentsByMeOtherURL : 获取某个用户发出的评论列表
const CommentsByMeOtherURL string = CUrlPreFix + "/2/comments/by_me/other.json"

// CommentsToMeOtherURL : 获取某个用户收到的评论列表
const CommentsToMeOtherURL string = CUrlPreFix + "/2/comments/to_me/other.json"

// CommentsTimelineOtherURL : 获取某个用户发出和收到的评论列表
const CommentsTimelineOtherURL string = CUrlPreFix + "/2/comments/timeline/other.json"

// CommentsMentionsOtherURL : 获取@某人的评论
const CommentsMentionsOtherURL string = CUrlPreFix + "/2/comments/mentions/other.json"

/* ========================微博内容数据（收费） END=====================================*/

/* ========================微博用户数据 START=====================================*/

// UsersShowBatchOtherURL : 批量获取其他用户的基本信息
const UsersShowBatchOtherURL string = CUrlPreFix + "/2/users/show_batch/other.json"

// TagsTagsBatchOtherURL : 批量获取用户标签
const TagsTagsBatchOtherURL string = CUrlPreFix + "/2/tags/tags_batch/other.json"

// UsersCountsBatchOtherURL : 批量获取用户的粉丝数、关注数、微博数
const UsersCountsBatchOtherURL string = CUrlPreFix + "/2/users/counts_batch/other.json"

/* ========================微博用户数据 END=====================================*/

/* ========================通用API START=====================================*/

/* ========================微博内容读取接口 START=====================================*/

// StatusesPublicTimelineBizURL : 获取最新的公共微博
const StatusesPublicTimelineBizURL string = CUrlPreFix + "/2/statuses/public_timeline/biz.json"

// StatusesFriendsTimelineBizURL :获取当前登录用户及其所关注用户的最新微博
const StatusesFriendsTimelineBizURL string = CUrlPreFix + "/2/statuses/friends_timeline/biz.json"

// StatusesUserTimelineBizURL : 获取当前登录用户发布的微博
const StatusesUserTimelineBizURL string = CUrlPreFix + "/2/statuses/user_timeline/biz.json"

// StatusesRepostTimelineBizURL : 获取当前登录用户发布的一条原创微博的最新转发微博
const StatusesRepostTimelineBizURL string = CUrlPreFix + "/2/statuses/repost_timeline/biz.json"

// StatusesMentionsBizURL : 获取@当前登录用户的微博
const StatusesMentionsBizURL string = CUrlPreFix + "/2/statuses/mentions/biz.json"

/* ========================微博内容读取接口 END=====================================*/

/* ========================微博信息读取接口 START=====================================*/

// StatusesShowBatchBizURL : 根据微博ID批量获取微博信息
const StatusesShowBatchBizURL string = CUrlPreFix + "/2/statuses/show_batch/biz.json"

// StatusesCountBizURL : 批量获取指定微博的转发数评论数喜欢数
const StatusesCountBizURL string = CUrlPreFix + "/2/statuses/count/biz.json"

/* ========================微博信息读取接口 END=====================================*/

/* ========================微博写入接口 START=====================================*/

// StatusesRepostBizURL : 转发一条微博信息
const StatusesRepostBizURL string = CUrlPreFix + "/2/statuses/repost/biz.json"

// StatusesDestroyBizURL ： 删除一条微博信息
const StatusesDestroyBizURL string = CUrlPreFix + "/2/statuses/destroy/biz.json"

// StatusesUpdateBizURL : 发布一条微博信息
const StatusesUpdateBizURL string = CUrlPreFix + "/2/statuses/update/biz.json"

// StatusesUploadBizURL : 上传图片并发布一条微博
const StatusesUploadBizURL string = CUrlPreFix + "/2/statuses/upload/biz.json"

// StatusesUploadURLTextBizURL : 发布一条微博同时指定上传的图片或图片url
const StatusesUploadURLTextBizURL string = CUrlPreFix + "/2/statuses/upload_url_text/biz.json"

// StatusesFilterCreateBizURL : 屏蔽某条微博
const StatusesFilterCreateBizURL string = CUrlPreFix + "/2/statuses/filter/create/biz.json"

// StatusesMentionsShieldBizURL : 屏蔽某个@我的微博及后续由其转发引起的@提及
const StatusesMentionsShieldBizURL string = CUrlPreFix + "/2/statuses/mentions/shield/biz.json"

/* ========================微博写入接口 END=====================================*/

/* ========================评论内容读取接口 START=====================================*/

// CommentsShowBizURL : 获取当前登录用户发布的微博下的评论列表
const CommentsShowBizURL string = CUrlPreFix + "/2/comments/show/biz.json"

// CommentsByMeBizURL : 获取当前登录用户发出的评论列表
const CommentsByMeBizURL string = CUrlPreFix + "/2/comments/by_me/biz.json"

// CommentsToMeBizURL : 获取当前登录用户收到的评论列表
const CommentsToMeBizURL string = CUrlPreFix + "/2/comments/to_me/biz.json"

// CommentsTimelineBizURL : 获取当前登录用户发出及收到的评论列表
const CommentsTimelineBizURL string = CUrlPreFix + "/2/comments/timeline/biz.json"

// CommentsMentionsBizURL : 获取@当前登录用户的评论
const CommentsMentionsBizURL string = CUrlPreFix + "/2/comments/mentions/biz.json"

/* ========================评论内容读取接口 END=====================================*/

/* ========================评论信息读取接口 START=====================================*/

// CommentsShowBatchBizURL : 批量获取评论内容
const CommentsShowBatchBizURL string = CUrlPreFix + "/2/comments/show_batch/biz.json"

/* ========================评论信息读取接口 END=====================================*/

/* ========================评论写入接口 START=====================================*/

// CommentsCreateBizURL : 评论一条微博
const CommentsCreateBizURL string = CUrlPreFix + "/2/comments/create/biz.json"

// CommentsDestroyBizURL : 删除一条评论
const CommentsDestroyBizURL string = CUrlPreFix + "/2/comments/destroy/biz.json"

// CommentsReplyBizURL : 回复一条评论
const CommentsReplyBizURL string = CUrlPreFix + "/2/comments/reply/biz.json"

/* ========================评论写入接口 END=====================================*/

/* ========================用户读取接口 START=====================================*/

// UsersShowBizURL : 获取当前登录用户信息
const UsersShowBizURL string = CUrlPreFix + "/2/users/show/biz.json"

// TagsBizURL : 获取当前登录用户的标签
const TagsBizURL string = CUrlPreFix + "/2/tags/biz.json"

/* ========================用户读取接口 END=====================================*/

/* ========================赞列表读取接口 START=====================================*/

// AttitudesToMeBizURL : 获取当前登录用户收到赞的列表
const AttitudesToMeBizURL string = CUrlPreFix + "/2/attitudes/to_me/biz.json"

// AttitudesShowBizURL : 根据微博ID返回某条微博的赞列表
const AttitudesShowBizURL string = CUrlPreFix + "/2/attitudes/show/biz.json"

/* ========================赞列表读取接口 END=====================================*/

/* ========================关系读取接口 START=====================================*/

// FriendshipsFriendsBizURL : 获取当前登录用户的关注列表
const FriendshipsFriendsBizURL string = CUrlPreFix + "/2/friendships/friends/biz.json"

// FriendshipsFollowersBizURL : 获取当前登录用户的粉丝列表
const FriendshipsFollowersBizURL string = CUrlPreFix + "/2/friendships/followers/biz.json"

/* ========================关系读取接口 END=====================================*/

/* ========================关系写入接口 START=====================================*/

// FriendshipsCreateBizURL : 关注某用户
const FriendshipsCreateBizURL string = CUrlPreFix + "/2/friendships/create/biz.json"

// FriendshipsDestroyBizURL : 取消关注
const FriendshipsDestroyBizURL string = CUrlPreFix + "/2/friendships/destroy/biz.json"

// FriendshipsFollowersDestroyBizURL : 移除当前登录用户的粉丝
const FriendshipsFollowersDestroyBizURL string = CUrlPreFix + "/2/friendships/followers/destroy/biz.json"

// FriendshipsRemarkUpdateBizURL : 更新关注人备注
const FriendshipsRemarkUpdateBizURL string = CUrlPreFix + "/2/friendships/remark/update/biz.json"

/* ========================关系写入接口 END=====================================*/

/* ========================搜索用户接口 START=====================================*/

// SearchSuggestionsUsersBizURL : 搜用户搜索建议
const SearchSuggestionsUsersBizURL string = CUrlPreFix + "/2/search/suggestions/users/biz.json"

/* ========================搜索用户接口 END=====================================*/

/* ========================通用API END=====================================*/

// ContentTypeURLEncoded : 发送请求的类型
const ContentTypeURLEncoded string = "application/x-www-form-urlencoded"

//以 other 结尾的，都是收费的api
