package database

const (
	QueryInsertCustomer = "insert into customers(fullname, email, gender, address) values ($1, $2, $3, $4);"

)
