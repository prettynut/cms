package cms

import (
	"encoding/json"

	routing "github.com/qiangxue/fasthttp-routing"
)

func NewsListCtrl(c *routing.Context) error {

	// var js Request
	// err := json.Unmarshal(c.Request.Body(), &js)
	// if err != nil {

	// }
	list := GetNewsList()
	js, _ := json.Marshal(data)

	var res ResponseStruct
	json.Unmarshal(js, &res)

	c.Write(js)
	// c.SetBody(js)

	c.SetStatusCode(res.HTTPStatus)
	// c.Success("json", js)

	return nil
}
