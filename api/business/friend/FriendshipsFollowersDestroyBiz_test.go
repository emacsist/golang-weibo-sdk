package friend

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestFriendshipsFollowersDestroyBizString(t *testing.T) {
	param := apiParam.FriendshipsFollowersDestroyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uid = 1234123

	body, err := FriendshipsFollowersDestroyBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("FriendshipsFollowersDestroyBizString error : %v\n", err)
		return
	}

	status := resp.FriendshipsFollowersDestroyBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestFriendshipsFollowersDestroyBizObject(t *testing.T) {
	param := apiParam.FriendshipsFollowersDestroyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uid = 1234123

	resp, error := FriendshipsFollowersDestroyBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestFriendshipsFollowersDestroyBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.ID)

	}
}
