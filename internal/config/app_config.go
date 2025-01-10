package config

import "github.com/jackc/pgx/v5/pgxpool"

type AppConfiguration struct {
	Conn *pgxpool.Pool
}

var AppConfig *AppConfiguration

func NewAppConfiguration(conn *pgxpool.Pool) {
	AppConfig = &AppConfiguration{
		Conn: conn,
	}
}
