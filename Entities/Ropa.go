package Entities

type Ropa struct {
	Id          int     `json:"Id,string,omitempty"`
	Imagen      string  `json:"Imagen"`
	Tipo        string  `json:"Tipo"`
	Precio      float64 `json:"Precio,string"`
	Cantidad    int     `json:"Cantidad,string"`
	Talla       string  `json:"Talla"`
	Propietario int     `json:"Propietario,string"`
}

func GetRopa(id int, imagen string, tipo string, precio float64, talla string, propietario int) *Ropa {
	r := new(Ropa)
	r.Id = id
	r.Imagen = imagen
	r.Tipo = tipo
	r.Precio = precio
	r.Talla = talla
	r.Propietario = propietario
	return r
}
