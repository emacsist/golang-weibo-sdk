package friend

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestFriendshipsFollowersBizString(t *testing.T) {
	param := apiParam.FriendshipsFollowersBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Type = 23

	body, err := FriendshipsFollowersBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("FriendshipsFollowersBizString error : %v\n", err)
		return
	}

	com := resp.FriendshipsFollowersBizResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", com.TotalNumber)

		t.Errorf("get body mid => %v\n", com.Users[0].Idstr)
	}
}

func TestFriendshipsFollowersBizObject(t *testing.T) {
	param := apiParam.FriendshipsFollowersBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Type = 23

	resp, error := FriendshipsFollowersBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestFriendshipsFollowersBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Users[0].Idstr)
	}
}
