package user

import (
	"encoding/json"
	"testing"

	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestUsersShowBizString(t *testing.T) {

	body, err := UsersShowBizString("2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("UsersShowBizString error : %v\n", err)
		return
	}

	com := resp.UsersShowBizResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {

		t.Errorf("get body mid => %v\n", com.City)
	}
}

func TestUsersShowBizObject(t *testing.T) {

	resp, error := UsersShowBizObject("2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if error != nil {
		t.Errorf("TesUsersShowBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body mid => %v\n", resp.City)
	}
}
