package Entities

type Usuario struct {
	Cedula            int    `json:"Cedula,string,omitempty"`
	Nombre            string `json:"Nombre"`
	PrimerApellido    string `json:"PrimerApellido"`
	SegundoApellido   string `json:"SegundoApellido"`
	FechaNacimiento   string `json:"FechaNacimiento,timestamp"`
	CorreoElectronico string `json:"CorreoElectronico"`
	Contrasena        string `json:"Contrasena"`
	FotoPerfil        string `json:"FotoPerfil"`
}
