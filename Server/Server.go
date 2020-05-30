package Server

import (
	"PaginaWeb/Entities"
	"PaginaWeb/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var usuario Entities.Usuario
var ropa Entities.Ropa

func RetrieveUsers(w http.ResponseWriter, request *http.Request) {
	json.NewEncoder(w).Encode(Models.RetrieveAllUsers())
}
func RetrieveClothes(w http.ResponseWriter, request *http.Request) {
	json.NewEncoder(w).Encode(Models.RetrieveAllRopa())
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
		Models.CreateCLothing(ropa)
		json.NewEncoder(w).Encode(Models.RetrieveException(2))
	}
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
	router.HandleFunc("/registro_usuario", CreateUsers).Methods("POST")
	router.HandleFunc("/usuarios", RetrieveUsers).Methods("GET")
	router.HandleFunc("/usuarios", DeleteUsers).Methods("DELETE")
	router.HandleFunc("/ropa", RetrieveClothes).Methods("GET")
	router.HandleFunc("/registro_ropa", CreateClothes).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./src")))
	http.ListenAndServe(":3000", router)

}
