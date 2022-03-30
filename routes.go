package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PostClassifiedsCtrl(w http.ResponseWriter, r *http.Request) {

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	var postInfo PostInformation

	json.Unmarshal(body, &postInfo)

	e.Encode(PostClassifieds(postInfo))
}

func GetClassifiedsCtrl(w http.ResponseWriter, r *http.Request) {

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	e.Encode(GetClassifieds())
}

// func DelateClassifiedsCtrl(w http.ResponseWriter, r *http.Request) {

// 	e := json.newEncoder()

// 	w.Header().Set("Content-Type", "applicaton/json")
	
// 	body, _ := ioutil.ReadAll(r.Body)

// 	json.Unmarshal(body, &)

// 	e.Encode(DeleteClassified())

// }