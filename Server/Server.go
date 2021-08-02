package Server

import (
	"PaginaWeb/Entities"
	"PaginaWeb/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

var usuario Entities.Usuario
var ropa Entities.Ropa

func RetrieveUsers(w http.ResponseWriter, request *http.Request) {

	json.NewEncoder(w).Encode(Models.RetrieveAllUsers())
}

func RetrieveClothes(w http.ResponseWriter, request *http.Request) {
	cedula, err := strconv.Atoi(request.URL.RawQuery)
	if err != nil {
		err.Error()
	}
	json.NewEncoder(w).Encode(Models.RetrieveAllClothing(cedula))
}

func CreateUsers(w http.ResponseWriter, request *http.Request) {

	bodyBytes, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	} else {
		json.Unmarshal(bodyBytes, &usuario)
		Models.CreateUser(usuario)
		json.NewEncoder(w).Encode(Models.RetrieveException(1))
	}
}
func CreateClothes(w http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.Unmarshal(bodyBytes, &ropa)
		Models.CreateClothing(ropa)
		json.NewEncoder(w).Encode(Models.RetrieveException(2))
	}
}
func RetrieveUser(w http.ResponseWriter, request *http.Request) {
	query:=request.URL.Query()
	usuario.CorreoElectronico=query.Get("CorreoElectronico")
	usuario.Contrasena=query.Get("Contrasena")
	json.NewEncoder(w).Encode(Models.RetrieveUser(usuario))
}

func DeleteUsers(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&usuario)
	if err != nil {
		panic(err)
	}
	Models.DeleteUserById(usuario)
}

func InitServer() {
	router := mux.NewRouter()
	router.HandleFunc("/usuario", CreateUsers).Methods("POST")
	router.HandleFunc("/usuario", RetrieveUser).Methods("GET")
	router.HandleFunc("/usuario", DeleteUsers).Methods("DELETE")
	router.HandleFunc("/ropa", RetrieveClothes).Methods("GET")
	router.HandleFunc("/ropa", CreateClothes).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./src")))
	http.ListenAndServe(":3000", router)

}
