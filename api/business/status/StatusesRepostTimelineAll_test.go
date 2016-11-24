package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesRepostTimelineAllString(t *testing.T) {
	param := apiParam.StatusesRepostTimelineAllParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 4032721045198077
	param.Page = 1
	param.Count = 2

	body, err := StatusesRepostTimelineAllString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("StatusesRepostTimelineAllString error : %v\n", err)
		return
	}

	status := resp.StatusesRepostTimelineAllResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.TotalNumber)

		t.Errorf("get body mid => %v\n", status.Reposts[0].ID)
	}
}

func TestStatusesRepostTimelineAllObject(t *testing.T) {
	param := apiParam.StatusesRepostTimelineAllParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 4032721045198077
	param.Page = 1
	param.Count = 2

	resp, error := StatusesRepostTimelineAllObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesRepostTimelineAllObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Reposts[0].ID)
	}
}
