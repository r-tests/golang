package main

import (
	json2 "encoding/json"
	"model"
	"net/http"
)

func ApiHandler(writer http.ResponseWriter, request *http.Request) {
	payment := model.Payment{Sender: "Peter", Receiver: "Rubeen", Amount: 3000.1}
	json, _ := json2.Marshal(payment)
	writer.Write(json)
}
