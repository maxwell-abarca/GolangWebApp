package Models

import (
	"PaginaWeb/DB"
	"PaginaWeb/Entities"
	"fmt"
)

func RetrieveException(id int) Entities.Exception {
	var exception Entities.Exception
	rows, err := DB.Conexion().Query("EXEC RET_EXCEPTION_ID @P_ID=?", id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var id int
			var message string
			rows.Scan(&id, &message)
			exception.Id = id
			exception.Message = message
		}
	}
	return exception
}
