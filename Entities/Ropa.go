package Entities

type Ropa struct {
	Id          int     `json:"Id"`
	Tipo        string  `json:"Tipo"`
	Precio      float64 `json:"Precio"`
	Cantidad    int     `json:"Cantidad"`
	Talla       rune    `json:"Talla"`
	Propietario int     `json:"Propietario"`
}
