package main

import "fmt"

var  (
	host = "localhost"
	user = "muhammadsarvar"
	password = "onam2575"
	dbname = "tezparsel"
	port = 5432
)

var DB_CONFIG = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
						host, user, password, dbname, port,
						)