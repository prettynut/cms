package cms

import routing "github.com/qiangxue/fasthttp-routing"

func RoutePath(r *routing.Router) {
	news := r.Group("/news")
	news.Get("/", cms.GetNewsList)

}
