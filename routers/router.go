package routers

import (
	"db_lab01/models"
	routing "github.com/qiangxue/fasthttp-routing"
	"log"
)

func InitRouter(env *models.Env) *routing.Router {
	router := routing.New()
	for _, route := range routes {
		handler := route.Handler
		for _, middleware := range route.Wrappers {
			for _, h := range handler {
				h = middleware(env, h(env))
			}
		}

		if len(handler) > 0 {
			if route.Method == GET {
				router.Get(route.Path, handler[0](env))
			} else if route.Method == POST {
				router.Post(route.Path, handler[0](env))
			} else if route.Method == BOTH {
				router.Get(route.Path, handler[0](env)).Post(handler[1](env))
			}
		} else {
			log.Println("Empty handlers list")
		}

		log.Println(route.Name)
	}

	return router
}
