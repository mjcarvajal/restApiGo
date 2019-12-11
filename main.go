package main


/*
	Importa los paquetes o modulos necesarios para la ejecución
*/
import (
	"encoding/json" // Manejo de codificaciòn y decodificación JSON
	"log" // Manejo de Errores 
	"net/http" //  Escribir el servidor peticiones, funcionalidades y respuesta 
	"github.com/gorilla/mux"
)

type product struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Brand string `json:"brand,omitempty"`
	DateProduct *DateProduct `json:"dateproduct,omitempty"`
}

type DateProduct struct {
	Lote string `json:"lote,omitempty"`
	Expiration string `json:"expiration,omitempty"`
}

var products []product


/*
	Funcion para traer todos los productos
	Entradas: 
	w variable con la que se responde al navegador
	req Variable con la información enviada por el navegador
*/
func getProdcutsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(products)

}

/*
	Función para traer el producto por el identificador
	Entradas: 
	w variable con la que se responde al navegador
	req Variable con la información enviada por el navegador
*/
func getProductbyIdEndpoint(w http.ResponseWriter, req *http.Request){

}
/*
	Función para crear un nuevo producto 
	Entradas: 
	w variable con la que se responde al navegador
	req Variable con la información enviada por el navegador
*/
func createProductEndpoint(w http.ResponseWriter, req *http.Request){

}

/*
	Función para eliminar un producto por el id
	Entradas: 
	w variable con la que se responde al navegador
	req Variable que contiene la información enviada por el navegador
*/

func deleteProductByIdEndpoint(w http.ResponseWriter, req *http.Request){

}


/*
	Función principal
*/
func main (){
	// Declaración y asignacion del router
	router:= mux.NewRouter()

	// Llenado del arreglo de productos

	products= append(products, product{Id: "1", Name: "Papas", Brand: "Margarita", DateProduct: &DateProduct{Lote: "10/02/2019", Expiration: "20/02/2020"}})
	products= append(products, product{Id: "2", Name: "Chitos", Brand: "Chirricos", DateProduct: &DateProduct{Lote: "10/02/2019", Expiration: "20/02/2020"}})

	// Endopoins 
	router.HandleFunc("/products", getProdcutsEndpoint).Methods("GET")
	router.HandleFunc("/products/{id}", getProductbyIdEndpoint).Methods("GET")
	router.HandleFunc("/products/{id}", createProductEndpoint).Methods("POST")
	router.HandleFunc("/products/{id}", deleteProductByIdEndpoint).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":3010", router))



}
