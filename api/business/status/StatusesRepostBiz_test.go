package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesRepostBizString(t *testing.T) {
	param := apiParam.StatusesRepostBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234213
	param.Status = "hello"

	body, err := StatusesRepostBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("StatusesRepostBizString error : %v\n", err)
		return
	}

	status := resp.StatusesRepostBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestStatusesRepostBizObject(t *testing.T) {
	param := apiParam.StatusesRepostBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234213
	param.Status = "hello"

	resp, error := StatusesRepostBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesRepostBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.ID)

	}
}
