package comment

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestCommentsByMeOtherString(t *testing.T) {
	param := apiParam.CommentsByMeOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	body, err := CommentsByMeOtherString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("CommentsByMeOtherString error : %v\n", err)
		return
	}

	com := resp.CommentsByMeOtherResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", com.TotalNumber)

		t.Errorf("get body mid => %v\n", com.Comments[0].ID)
	}
}

func TestCommentsByMeOtherObject(t *testing.T) {
	param := apiParam.CommentsByMeOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	resp, error := CommentsByMeOtherObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestCommentsByMeOtherObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Comments[0].ID)
	}
}
