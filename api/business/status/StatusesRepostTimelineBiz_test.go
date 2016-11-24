package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesRepostTimelineBizString(t *testing.T) {
	param := apiParam.StatusesRepostTimelineBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 4032721045198077
	param.Page = 1
	param.Count = 2

	body, err := StatusesRepostTimelineBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("StatusesRepostTimelineBizString error : %v\n", err)
		return
	}

	status := apiResp.StatusesRepostTimelineBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.TotalNumber)

		t.Errorf("get body mid => %v\n", status.Reposts[0].ID)
	}
}

func TestStatusesRepostTimelineBizObject(t *testing.T) {
	param := apiParam.StatusesRepostTimelineBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 4032721045198077
	param.Page = 1
	param.Count = 2

	resp, error := StatusesRepostTimelineBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesRepostTimelineBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Reposts[0].ID)
	}
}
