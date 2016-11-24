package tag

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestTagsTagsBatchOtherString(t *testing.T) {
	param := apiParam.TagsTagsBatchOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uids = "5225532117"

	body, err := TagsTagsBatchOtherString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("TagsTagsBatchOtherString error : %v\n", err)
		return
	}

	com := []resp.TagsTagsBatchOtherResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {

		t.Errorf("get body mid => %v\n", com[0].ID)
	}
}

func TestTagsTagsBatchOtherObject(t *testing.T) {
	param := apiParam.TagsTagsBatchOtherParam{}

	apiHelper.SetDefaultValues(&param)

	param.Uids = "5225532117"

	resp, error := TagsTagsBatchOtherObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if error != nil {
		t.Errorf("TestTagsTagsBatchOtherObject error : %+v\n", error)
	} else {
		t.Errorf("get body mid => %v\n", (*resp)[0].ID)
	}
}
