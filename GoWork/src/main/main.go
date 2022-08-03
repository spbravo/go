// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Transaction_date string `json:"Transaction_Date"`
	Region_Id        string `json:"Region_Id"`
	Location_Id      string `json:"Location_Id"`
	Bar_Code_Id      string `json:"Bar_Code_Id"`
	QTY_In_Loc       int    `json:"QTY_In_Loc"`
	QTY_In_Transit   int    `json:"QTY_In_Transit"`
}

var Stock []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Stock)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Stock {
		if article.Bar_Code_Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	decoder := json.NewDecoder(strings.NewReader(string(reqBody)))
	var article Article
	for decoder.More() {
		err := decoder.Decode(&article)
		checkError(err)
		Stock = append(Stock, article)
	}

	json.NewEncoder(w).Encode(Stock)
}

/*
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Stock {
		if article.Id == id {
			Stock = append(Stock[:index], Stock[index+1:]...)
		}
	}

}
*/
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	//myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	/*Stock = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}*/
	handleRequests()
}
