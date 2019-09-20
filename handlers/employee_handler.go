package handlers

import (
	"db_lab01/database"
	"db_lab01/models"
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

func unmarshalEmployee(c *routing.Context) *models.Employee {
	employee := &models.Employee{}
	err := json.Unmarshal(c.PostBody(), employee)
	if err != nil {
		log.Println("Error while unmarshalling employee")
		return nil
	}
	return employee
}

func EmployeeCreateHandler(env *models.Env) routing.Handler {
	return func(c *routing.Context) error {
		employee := unmarshalEmployee(c)

		_, err := env.DB.Exec(database.QueryInsertEmployee, employee.FullName, employee.Role,
			employee.SalonID, employee.Email, employee.Gender, employee.IsWorking)
		if err != nil {
			c.Error(err.Error(), fasthttp.StatusInternalServerError)
			return err
		}
		return nil
	}
}
