package attitude

import "testing"
import apiParam "github.com/emacsist/golang-weibo-sdk/param"
import apiHelper "github.com/emacsist/golang-weibo-sdk/helper"

func TestAttitudesShowBizObject(t *testing.T) {
	param := apiParam.AttitudesShowBizParam{}

	apiHelper.SetDefaultValues(&param)

	param.Page = 1
	param.Count = 2

	resp, error := AttitudesShowBizObject(param, "2.00ZrQ6BDnYTUdC7164d3f05bY5kDBE sdfsdf")

	if error != nil {
		t.Errorf("TestAttitudesShowBizObject error : %+v\n", error)
	} else {
		t.Errorf("get body => %v\n", resp.TotalNumber)

		t.Errorf("get body mid => %v\n", resp.Attitudes[0].ID)
	}
}
