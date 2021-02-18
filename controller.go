package cms

import (
	"encoding/json"
	"net/http"

	routing "github.com/qiangxue/fasthttp-routing"
)

func NewsListCtrl(c *routing.Context) error {

	list := GetNewsList()
	// js, _ := json.Marshal(data)

	var res NewsListResponse
	// json.Unmarshal(js, &res)
	res.Data = list
	res.Code = 1
	res.Message = "Success"
	js, _ := json.Marshal(res)
	c.Write(js)
	// c.SetBody(js)

	c.SetStatusCode(http.StatusOK)
	// c.Success("json", js)

	return nil
}
