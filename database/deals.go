package database

const (
	QueryInsertDeal = "insert into deals(customer_id, haircut_id, employee_id, price_id) values($1, $2, $3, $4);"
)
