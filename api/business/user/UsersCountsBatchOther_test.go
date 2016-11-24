package user

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestUsersCountsBatchOtherString(t *testing.T) {
	param := apiParam.UsersCountsBatchOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uids = "5225532117"

	body, err := UsersCountsBatchOtherString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("UsersCountsBatchOtherString error : %v\n", err)
		return
	}

	com := []resp.UsersCountsBatchOtherResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {

		t.Errorf("get body mid => %v\n", com[0].ID)
	}
}

func TestUsersCountsBatchOtherObject(t *testing.T) {
	param := apiParam.UsersCountsBatchOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uids = "5225532117"

	resp, error := UsersCountsBatchOtherObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if error != nil {
		t.Errorf("TestUsersCountsBatchOtherObject error : %+v\n", error)
	} else {
		t.Errorf("get body mid => %v\n", (*resp)[0].ID)
	}
}
