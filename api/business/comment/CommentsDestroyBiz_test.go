package comment

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestCommentsDestroyBizString(t *testing.T) {
	param := apiParam.CommentsDestroyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Cid = 1234234

	body, err := CommentsDestroyBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("CommentsDestroyBizString error : %v\n", err)
		return
	}

	status := resp.CommentsDestroyBizResp{}
	error := json.Unmarshal([]byte(body), &status)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", status.ID)
	}
}

func TestCommentsDestroyBizObject(t *testing.T) {
	param := apiParam.CommentsDestroyBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Cid = 1234234

	resp, error := CommentsDestroyBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestCommentsDestroyBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.ID)

	}
}
