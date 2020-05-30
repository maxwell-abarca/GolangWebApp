package Models

import (
	"PaginaWeb/DB"
	"PaginaWeb/Entities"
	"fmt"
)

func RetrieveAllRopa() []Entities.Ropa {
	var ropa []Entities.Ropa
	rows, err := DB.Conexion().Query("EXEC RET_ALL_ROPA_PR")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var Id int
			var Tipo string
			var Precio float64
			var Cantidad int
			var Talla rune
			var Propietario int
			rows.Scan(&Id, &Tipo, &Precio, &Cantidad, &Talla)
			ropa = append(ropa, Entities.Ropa{Id, Tipo, Precio, Cantidad, Talla, Propietario})
		}
	}
	return ropa
}
func CreateCLothing(ropa Entities.Ropa) {

	rows, err := DB.Conexion().Query("EXEC PROCEDURE CRE_ROPA_PR @P_TIPO=?,@P_PRECIO=?,@P_CANTIDAD=?,@P_TALLA=?,@P_PROPIETARIO=?", ropa.Tipo, ropa.Precio, ropa.Cantidad, ropa.Talla, ropa.Propietario)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer rows.Close()

	}

}
