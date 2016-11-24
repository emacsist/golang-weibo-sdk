package comment

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestCommentsReplyBizString(t *testing.T) {
	param := apiParam.CommentsReplyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234234234
	param.Comment = "hello"

	body, err := CommentsReplyBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("CommentsReplyBizString error : %v\n", err)
		return
	}

	status := resp.CommentsReplyBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestCommentsReplyBizObject(t *testing.T) {
	param := apiParam.CommentsReplyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234234234
	param.Comment = "hello"

	resp, error := CommentsReplyBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestCommentsReplyBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.ID)

	}
}
