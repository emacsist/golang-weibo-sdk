package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	apiResp "github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesShowBatchBizString(t *testing.T) {
	param := apiParam.StatusesShowBatchBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Ids = "4032721045198077"

	body, err := StatusesShowBatchBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("StatusesShowBatchBizString error : %v\n", err)
		return
	}

	status := apiResp.StatusesShowBatchBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.TotalNumber)

		t.Errorf("get body mid => %v\n", status.Statuses[0].ID)
	}
}

func TestStatusesShowBatchBizObject(t *testing.T) {
	param := apiParam.StatusesShowBatchBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Ids = "4032721045198077"

	resp, error := StatusesShowBatchBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesShowBatchBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Statuses[0].ID)
	}
}
