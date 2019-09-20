package routers

import (
	"db_lab01/handlers"
	"db_lab01/models"
	routing "github.com/qiangxue/fasthttp-routing"
)

const POST = 1
const GET = 2
const BOTH = 3

type MyHandler func(env *models.Env) routing.Handler

type Route struct {
	Name     string
	Path     string
	Method   int
	Wrappers []Wrapper
	//Handler  []routing.Handler
	Handler []MyHandler
}

var routes = []Route{
	{
		"CustomerCreate",
		"/api/customer/create",
		POST,
		[]Wrapper{RecoverMiddleware},
		[]MyHandler{handlers.CustomerCreateHandler},
	},
	{
		"DealCreate",
		"/api/deal/create",
		POST,
		[]Wrapper{RecoverMiddleware},
		[]MyHandler{handlers.DealCreateHandler},
	},
	{
		"EmployeeCreate",
		"/api/employee/create",
		POST,
		[]Wrapper{RecoverMiddleware},
		[]MyHandler{handlers.EmployeeCreateHandler},
	},
	{
		"HaircutCreate",
		"/api/haircut/create",
		POST,
		[]Wrapper{RecoverMiddleware},
		[]MyHandler{handlers.HaircutCreateHandler},
	},
	{
		"PriceCreate",
		"/api/price/create",
		POST,
		[]Wrapper{RecoverMiddleware},
		[]MyHandler{handlers.PriceCreateHandler},
	},
	{
		"SalonCreate",
		"/api/salon/create",
		POST,
		[]Wrapper{RecoverMiddleware},
		[]MyHandler{handlers.SalonCreateHandler},
	},
}
