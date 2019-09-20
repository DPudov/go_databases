package routers

import (
	"db_lab01/models"
	routing "github.com/qiangxue/fasthttp-routing"
	"log"
)

type Wrapper func(*models.Env, routing.Handler) MyHandler

func RecoverMiddleware(env *models.Env, next routing.Handler) MyHandler {
	return func(env *models.Env) routing.Handler {
		return func(ctx *routing.Context) error {
			defer func() {
				if err := recover(); err != nil {
					log.Fatal(err)
				}
			}()
			err := next(ctx)
			return err
		}
	}
}
