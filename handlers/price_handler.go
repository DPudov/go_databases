package handlers

import (
	"db_lab01/database"
	"db_lab01/models"
	"encoding/json"
	"github.com/lib/pq"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func unmarshalPrice(c *routing.Context) *models.Price {
	price := &models.Price{}
	err := json.Unmarshal(c.PostBody(), price)
	if err != nil {
		log.Println("Error while unmarshalling price")
		return nil
	}
	return price
}

func PriceCreateHandler(env *models.Env) routing.Handler {
	return func(c *routing.Context) error {
		price := unmarshalPrice(c)

		_, err := env.DB.Exec(database.QueryInsertPrice, price.Value,
			price.HaircutID)
		if err != nil {

			if err, ok := err.(*pq.Error); ok {
				switch err.Code.Name() {
				case "foreign_key_violation":
					mess := &models.Err{}
					mess.Message = "Key violation for price"
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
