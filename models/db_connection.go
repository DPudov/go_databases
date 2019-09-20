package models

type Database struct {
	Host                   string
	Port                   int
	User                   string
	Password               string
	Name                   string
	Driver                 string
	MaxConnections         int
	MigrationsUpFileName   string
	MigrationsDownFileName string
}
