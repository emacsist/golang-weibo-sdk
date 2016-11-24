package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestSearchStatusesLimitedString(t *testing.T) {
	param := apiParam.SearchStatusesLimitedParam{}

	apiHelper.SetDefaultValues(&param)

	param.Q = "Golang"
	param.Page = 1
	param.Count = 2

	body, err := SearchStatusesLimitedString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdf")

	if err != nil {
		t.Errorf("SearchStatusesLimitedString error : %v\n", err)
		return
	}

	status := resp.SearchStatusesLimitedResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.TotalNumber)

		t.Errorf("get body mid => %v\n", status.Statuses[0].ID)
	}
}

func TestSearchStatusesLimitedObject(t *testing.T) {
	param := apiParam.SearchStatusesLimitedParam{}

	apiHelper.SetDefaultValues(&param)

	param.Q = "Golang"
	param.Page = 1
	param.Count = 2

	resp, error := SearchStatusesLimitedObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("SearchStatusesLimitedObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Statuses[0].ID)
	}
}
