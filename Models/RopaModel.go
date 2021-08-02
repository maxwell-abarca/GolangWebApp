package Models

import (
	"PaginaWeb/DB"
	"PaginaWeb/Entities"
	"fmt"
)

func RetrieveAllClothing(cedula int) []*Entities.Ropa {
	var ropa []*Entities.Ropa

	rows, err := DB.Conexion().Query("EXEC RET_ALL_ROPA_PR @P_CEDULA=?", cedula)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var Id int
			var Imagen string
			var Tipo string
			var Precio float64
			var Talla string
			var Propietario int
			rows.Scan(&Id, &Imagen, &Tipo, &Precio, &Talla, &Propietario)
			ropa = append(ropa, Entities.GetRopa(Id, Imagen, Tipo, Precio, Talla, Propietario))
		}
	}
	return ropa
}
func CreateClothing(ropa Entities.Ropa) {
	rows, err := DB.Conexion().Query("EXEC CRE_ROPA_PR @P_IMAGEN=?,@P_TIPO=?,@P_PRECIO=?,@P_CANTIDAD=?,@P_TALLA=?,@P_PROPIETARIO=?", ropa.Imagen, ropa.Tipo, ropa.Precio, ropa.Cantidad, ropa.Talla, ropa.Propietario)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer rows.Close()
	}
}
