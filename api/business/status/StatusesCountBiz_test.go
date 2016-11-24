package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesCountBizString(t *testing.T) {
	param := apiParam.StatusesCountBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Ids = "1231234"

	body, err := StatusesCountBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("StatusesCountBizString error : %v\n", err)
		return
	}

	com := []resp.StatusesCountBizResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {

		t.Errorf("get body mid => %v\n", com[0].ID)
	}
}

func TestStatusesCountBizObject(t *testing.T) {
	param := apiParam.StatusesCountBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Ids = "1231234"

	resp, error := StatusesCountBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if error != nil {
		t.Errorf("TestStatusesCountBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body mid => %v\n", (*resp)[0].ID)
	}
}
