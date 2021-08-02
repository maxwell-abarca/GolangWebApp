package Models

import (
	"PaginaWeb/DB"
	"PaginaWeb/Entities"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func RetrieveAllUsers() []Entities.Usuario {
	var usuario []Entities.Usuario

	rows, err := DB.Conexion().Query("EXEC RET_ALL_USUARIO_PR")
	if err != nil {
		fmt.Println("Error" + err.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			var Cedula int
			var Nombre, PrimerApellido, SegundoApellido, CorreoElectronico, Contrasena, FechaNacimiento, FotoPerfil string

			rows.Scan(&Cedula, &Nombre, &PrimerApellido, &SegundoApellido, &FechaNacimiento, &CorreoElectronico, &Contrasena, &FotoPerfil)

			usuario = append(usuario, Entities.Usuario{Cedula, Nombre, PrimerApellido, SegundoApellido, FechaNacimiento, CorreoElectronico, Contrasena, FotoPerfil})
		}
	}
	return usuario
}
func RetrieveUser(usuario Entities.Usuario) Entities.Usuario {
	var user Entities.Usuario
	rows, err := DB.Conexion().Query("EXEC RET_USUARIO_ID_PR @P_CORREO_ELECTRONICO=?,@P_CONTRASENA=?", usuario.CorreoElectronico, usuario.Contrasena)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&user.Cedula, &user.Nombre, &user.PrimerApellido, &user.SegundoApellido, &user.FechaNacimiento, &user.CorreoElectronico, &user.Contrasena, &user.FotoPerfil)
		}
	}
	return user
}
func CreateUser(usuario Entities.Usuario) {

	rows, err := DB.Conexion().Query("EXEC CRE_USUARIO_PR @P_CEDULA=?,@P_NOMBRE=?,@P_APELLIDO1=?,@P_APELLIDO2=?,@P_FECHA_NACIMIENTO=?,@P_CORREO_ELECTRONICO=?,@P_CONTRASENA=?,@P_FOTO_PERFIL=?", usuario.Cedula, usuario.Nombre, usuario.PrimerApellido, usuario.SegundoApellido, usuario.FechaNacimiento, usuario.CorreoElectronico, usuario.Contrasena, usuario.FotoPerfil)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		defer rows.Close()
	}

}
func DeleteUserById(usuario Entities.Usuario) {
	rows, err := DB.Conexion().Query("EXEC DEL_USUARIO_PR @P_CEDULA=?", usuario.Cedula)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		defer rows.Close()

	}
}
