package comment

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestCommentsShowBatchBizString(t *testing.T) {
	param := apiParam.CommentsShowBatchBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Cids = "1234123,123412"

	body, err := CommentsShowBatchBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("CommentsShowBatchBizString error : %v\n", err)
		return
	}

	com := []resp.CommentsShowBatchBizResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", com[0].ID)
	}
}

func TestCommentsShowBatchBizObject(t *testing.T) {
	param := apiParam.CommentsShowBatchBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Cids = "1234123,123412"

	resp, error := CommentsShowBatchBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestCommentsShowBatchBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", (*resp)[0].ID)
	}
}
