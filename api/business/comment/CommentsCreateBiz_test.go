package comment

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestCommentsCreateBizString(t *testing.T) {
	param := apiParam.CommentsCreateBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234234234
	param.Comment = "hello"

	body, err := CommentsCreateBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("CommentsCreateBizString error : %v\n", err)
		return
	}

	status := resp.CommentsCreateBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestCommentsCreateBizObject(t *testing.T) {
	param := apiParam.CommentsCreateBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Id = 1234234234
	param.Comment = "hello"

	resp, error := CommentsCreateBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestCommentsCreateBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.ID)

	}
}
