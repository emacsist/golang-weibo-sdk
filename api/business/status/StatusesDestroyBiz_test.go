package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesDestroyBizString(t *testing.T) {
	param := apiParam.StatusesDestroyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234124

	body, err := StatusesDestroyBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if err != nil {
		t.Errorf("TestStatusesDestroyBizString error : %v\n", err)
		return
	}

	status := resp.StatusesDestroyBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestStatusesDestroyBizObject(t *testing.T) {
	param := apiParam.StatusesDestroyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234123

	resp, error := StatusesDestroyBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesDestroyBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.Idstr)
	}
}
