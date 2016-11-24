package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesMentionsBizString(t *testing.T) {
	param := apiParam.StatusesMentionsBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	body, err := StatusesMentionsBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if err != nil {
		t.Errorf("TestStatusesMentionsBizString error : %v\n", err)
		return
	}

	status := resp.StatusesMentionsBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.TotalNumber)

		t.Errorf("get body mid => %v\n", status)
	}
}

func TestStatusesMentionsBizObject(t *testing.T) {
	param := apiParam.StatusesMentionsBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	resp, error := StatusesMentionsBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesMentionsBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Statuses[0].ID)
	}
}
