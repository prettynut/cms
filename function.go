package cms

func GetNewsList() []NewsStruct {

	var list []NewsStruct
	news := NewsStruct{
		Name:   "TestNews1",
		Start:  "2021-01-01",
		End:    "2021-12-31",
		Status: "Active",
	}
	list = append(list, news)

	news = NewsStruct{
		Name:   "TestNews2",
		Start:  "2019-01-01",
		End:    "2020-12-31",
		Status: "Active",
	}
	list = append(list, news)

	news = NewsStruct{
		Name:   "TestNews3",
		Start:  "2022-01-01",
		End:    "2022-12-31",
		Status: "Inactive",
	}
	list = append(list, news)
	return list
}
