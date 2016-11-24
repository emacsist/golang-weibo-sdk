package tag

import (
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestTagsBizString(t *testing.T) {
	param := apiParam.TagsBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1

	body, err := TagsBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("TagsBizString error : %v\n", err)
		return
	}

	com := resp.TagsBizResp{Body: body}
	t.Errorf("get body mid => %v\n", com.GetTags())
}

func TestTagsBizObject(t *testing.T) {
	param := apiParam.TagsBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1

	resp, error := TagsBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE")

	if error != nil {
		t.Errorf("TestTagsBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body mid => %v\n", resp.GetTags())
	}
}
