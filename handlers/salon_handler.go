package handlers

import (
	"db_lab01/database"
	"db_lab01/models"
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func unmarshalSalon(c *routing.Context) *models.Salon {
	salon := &models.Salon{}
	err := json.Unmarshal(c.PostBody(), salon)
	if err != nil {
		log.Println("Error while unmarshalling salon")
		return nil
	}
	return salon
}

func SalonCreateHandler(env *models.Env) routing.Handler {
	return func(c *routing.Context) error {
		salon := unmarshalSalon(c)
		log.Println(salon)
		_, err := env.DB.Exec(database.QueryInsertSalon, salon.Address)
		if err != nil {
			c.Error(err.Error(), fasthttp.StatusInternalServerError)
			return err
		}
		return nil
	}
}
