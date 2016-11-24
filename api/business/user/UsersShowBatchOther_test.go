package user

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestUsersShowBatchOtherString(t *testing.T) {
	param := apiParam.UsersShowBatchOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uids = "5225532117"

	body, err := UsersShowBatchOtherString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("UsersShowBatchOtherString error : %v\n", err)
		return
	}

	com := resp.UsersShowBatchOtherResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", com.TotalNumber)

		t.Errorf("get body mid => %v\n", com.Users[0].ID)
	}
}

func TestUsersShowBatchOtherObject(t *testing.T) {
	param := apiParam.UsersShowBatchOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uids = "5225532117"

	resp, error := UsersShowBatchOtherObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestUsersShowBatchOtherObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Users[0].ID)
	}
}
