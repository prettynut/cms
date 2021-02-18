package cms

type NewsStruct struct {
	Name   string `json:"name"`
	Start  string `json:"start"`
	End    string `json:"end"`
	Status string `json:"status"`
}

type NewsListResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// type NewsStruct struct {
// 	Name   string `json:"name"`
// 	Start  string `json:"start"`
// 	End    string `json:"end"`
// 	Status string `json:"status"`
// }
