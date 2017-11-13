package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const serverPort = "8080"

func AddressesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(getAddress(vars["cep"])))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("goCep RESTFull API \n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/addresses/{cep}", AddressesHandler)
	http.ListenAndServe(":"+serverPort, r)
}

func getAddress(cep string) string {
	var client http.Client
	resp, err := client.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := string(body)
	return string(result)
}
