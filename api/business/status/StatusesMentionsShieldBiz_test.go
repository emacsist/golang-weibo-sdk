package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesMentionsShieldBizString(t *testing.T) {
	param := apiParam.StatusesMentionsShieldBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 12341234

	body, err := StatusesMentionsShieldBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if err != nil {
		t.Errorf("TestStatusesMentionsShieldBizString error : %v\n", err)
		return
	}

	status := resp.StatusesMentionsShieldBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.Result)
	}
}

func TestStatusesMentionsShieldBizObject(t *testing.T) {
	param := apiParam.StatusesMentionsShieldBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 12341234

	resp, error := StatusesMentionsShieldBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesMentionsShieldBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.Result)
	}
}
