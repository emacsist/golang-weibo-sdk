package status

import (
	"encoding/json"
	"os"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestStatusesUploadBizString(t *testing.T) {
	param := apiParam.StatusesUploadBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Status = "hello"
	param.Pic, _ = os.Open("/tmp/fluxlog.txt")

	body, err := StatusesUploadBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("StatusesUploadBizString error : %v\n", err)
		return
	}

	status := resp.StatusesUploadBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestStatusesUploadBizObject(t *testing.T) {
	param := apiParam.StatusesUploadBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Status = "hello"

	resp, error := StatusesUploadBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestStatusesUploadBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.ID)

	}
}
