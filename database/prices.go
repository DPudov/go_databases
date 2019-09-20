package database

const (
	QueryInsertPrice = "insert into prices(price_value, haircut_id) values ($1, $2);"
)
