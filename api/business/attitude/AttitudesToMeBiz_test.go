package attitude

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestAttitudesToMeBizString(t *testing.T) {
	param := apiParam.AttitudesToMeBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	body, err := AttitudesToMeBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("AttitudesToMeBizString error : %v\n", err)
		return
	}

	com := resp.AttitudesToMeBizResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", com.TotalNumber)

		t.Errorf("get body mid => %v\n", com.Attitudes[0].ID)
	}
}

func TestAttitudesToMeBizObject(t *testing.T) {
	param := apiParam.AttitudesToMeBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	resp, error := AttitudesToMeBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestAttitudesToMeBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Attitudes[0].ID)
	}
}
