package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func Conexion() *sql.DB {
	var server = "LAPTOP-HDT2V47T"
	var user = "Maxwell"
	var password = "/Maximun751999"
	var port = 1433
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)

	db, err := sql.Open("mssql", connString)

	if err != nil {
		panic(err.Error())
	}

	return db
}
