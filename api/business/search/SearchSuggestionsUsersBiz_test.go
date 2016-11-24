package search

import (
	"encoding/json"
	"testing"

	apiHelper "github.com/emacsist/golang-weibo-sdk/helper"
	apiParam "github.com/emacsist/golang-weibo-sdk/param"
	"github.com/emacsist/golang-weibo-sdk/resp"
)

func TestSearchSuggestionsUsersBizString(t *testing.T) {
	param := apiParam.SearchSuggestionsUsersBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Q = "Hello"
	param.Count = 2

	body, err := SearchSuggestionsUsersBizString(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBEsdfs")

	if err != nil {
		t.Errorf("SearchSuggestionsUsersBizString error : %v\n", err)
		return
	}

	com := []resp.SearchSuggestionsUsersBizResp{}
	error := json.Unmarshal([]byte(body), &com)
	if error != nil {
		t.Errorf("to json error : %v\n", error.Error())
	} else {
		t.Errorf("get body => %v\n", com[0].ScreenName)
	}
}

func TestSearchSuggestionsUsersBizObject(t *testing.T) {
	param := apiParam.SearchSuggestionsUsersBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Q = "Hello"
	param.Count = 2

	resp, error := SearchSuggestionsUsersBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestSearchSuggestionsUsersBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", (*resp)[0].ScreenName)
	}
}
