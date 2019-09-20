package handlers

import (
	"db_lab01/database"
	"db_lab01/models"
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func unmarshalCustomer(c *routing.Context) *models.Customer {
	customer := &models.Customer{}
	err := json.Unmarshal(c.PostBody(), customer)
	if err != nil {
		log.Println("Error while unmarshalling customer")
		return nil
	}
	return customer
}

func CustomerCreateHandler(env *models.Env) routing.Handler {
	return func(c *routing.Context) error {
		customer := unmarshalCustomer(c)

		_, err := env.DB.Exec(database.QueryInsertCustomer, customer.FullName,
			customer.Email, customer.Gender, customer.Address)
		if err != nil {
			c.Error(err.Error(), fasthttp.StatusInternalServerError)
			return err
		}
		return nil
	}
}
