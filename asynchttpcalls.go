package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func PrintRes(url string, resp chan []byte) {
	respBody, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		body, err := io.ReadAll(respBody.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			resp <- body
		}
	}
}
func GetApi(w http.ResponseWriter, r *http.Request) {
	respArr := []string{}
	urlObj := struct {
		URL []string
	}{}
	input, err := io.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.Unmarshal(input, &urlObj)
	urls := urlObj.URL
	if len(urls) == 0 {
		json.NewEncoder(w).Encode("Provide url")
		return
	} else {
		resp := make(chan []byte)
		for _, url := range urls {
			go PrintRes(url, resp)
		}
		for range urls {
			val := <-resp
			respArr = append(respArr, string(val))
		}
	}
	json.NewEncoder(w).Encode(respArr)
	return
}
func main() {
	server := mux.NewRouter().StrictSlash(true)
	server.HandleFunc("/getapi", GetApi).Methods("GET")
	http.ListenAndServe(":8080", server)
}
