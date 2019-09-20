package database

const (
	QueryInsertEmployee = "insert into employees(fullname, role, salon_id, status, email, gender, isWorking) " +
		"values ($1, $2, $3, $4, $5, $6, $7);"
)
