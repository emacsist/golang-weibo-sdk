package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesMentionsOtherString(t *testing.T) {
	param := apiParam.StatusesMentionsOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	body, err := StatusesMentionsOtherString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if err != nil {
		t.Errorf("TestStatusesMentionsOtherString error : %v\n", err)
		return
	}

	status := resp.StatusesMentionsOtherResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.TotalNumber)

		t.Errorf("get body mid => %v\n", status.Statuses[0].ID)
	}
}

func TestStatusesMentionsOtherObject(t *testing.T) {
	param := apiParam.StatusesMentionsOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	resp, error := StatusesMentionsOtherObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesMentionsOtherObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Statuses[0].ID)
	}
}
