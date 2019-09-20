package handlers

import (
	"db_lab01/database"
	"db_lab01/models"
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func unmarshalHaircut(c *routing.Context) *models.Haircut {
	haircut := &models.Haircut{}
	err := json.Unmarshal(c.PostBody(), haircut)
	if err != nil {
		log.Println("Error while unmarshalling haircut")
		return nil
	}
	return haircut
}

func HaircutCreateHandler(env *models.Env) routing.Handler {
	return func(c *routing.Context) error {

		haircut := unmarshalHaircut(c)
		log.Println("received haircut", haircut)
		_, err := env.DB.Exec(database.QueryInsertHaircut, haircut.Name, haircut.Description)
		if err != nil {
			log.Println(err)
			c.Error(err.Error(), fasthttp.StatusInternalServerError)
			return err
		}
		return nil
	}
}
