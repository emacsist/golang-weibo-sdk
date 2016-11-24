package status

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesUploadURLTextBizString(t *testing.T) {
	param := apiParam.StatusesUploadURLTextBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Status = "hello"

	body, err := StatusesUploadURLTextBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("StatusesUploadURLTextBizString error : %v\n", err)
		return
	}

	status := resp.StatusesUploadURLTextBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestStatusesUploadURLTextBizObject(t *testing.T) {
	param := apiParam.StatusesUploadURLTextBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Status = "hello"

	resp, error := StatusesUploadURLTextBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesUploadURLTextBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.ID)

	}
}
