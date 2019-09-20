package handlers

import (
	"db_lab01/database"
	"db_lab01/models"
	"encoding/json"
	"github.com/lib/pq"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func unmarshalDeal(c *routing.Context) *models.Deal {
	deal := &models.Deal{}
	err := json.Unmarshal(c.PostBody(), deal)
	if err != nil {
		log.Println("Error while unmarshalling deal")
		return nil
	}
	return deal
}

func DealCreateHandler(env *models.Env) routing.Handler {
	return func(c *routing.Context) error {
		deal := unmarshalDeal(c)

		_, err := env.DB.Exec(database.QueryInsertDeal, deal.CustomerID,
			deal.HaircutID, deal.EmployeeID, deal.PriceID)
		if err != nil {

			if err, ok := err.(*pq.Error); ok {
				switch err.Code.Name() {
				case "foreign_key_violation":
					mess := &models.Err{}
					mess.Message = "Key violation for deal"
					userOutput, _ := json.Marshal(mess)
					c.SetContentType(EncodingApplicationJSON)
					c.SetStatusCode(fasthttp.StatusNotFound)
					c.SetBody(userOutput)
					return nil
				}

			}

			c.Error(err.Error(), fasthttp.StatusInternalServerError)
			return err
		}
		return nil
	}
}
