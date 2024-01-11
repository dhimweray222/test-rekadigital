package config

import (
	"os"
	"strconv"
)

var (
	//db
	host               = os.Getenv("DB_HOST")
	port               = os.Getenv("DB_PORT")
	user               = os.Getenv("DB_USERNAME")
	password           = os.Getenv("DB_PASSWORD")
	dbName             = os.Getenv("DB_NAME")
	minConns           = os.Getenv("DB_POOL_MIN")
	maxConns           = os.Getenv("DB_POOL_MAX")
	TimeOutDuration, _ = strconv.Atoi(os.Getenv("DB_CONNECTION_TIMEOUT"))

	//host
	ServerHost     = os.Getenv("SERVER_URI")
	ServerPort     = os.Getenv("SERVER_PORT")
	EndpointPrefix = os.Getenv("ENDPOINT_PREFIX")
	DefaultLimit   = os.Getenv("DEFAULT_LIMIT")
)
