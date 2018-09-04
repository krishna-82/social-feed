package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/krishna-82/social-feed/config"
	. "github.com/krishna-82/social-feed/dao"
	. "github.com/krishna-82/social-feed/models"
)

var config = Config{}
var dao = FeedDAO{}

// GET list of people User follows
/* It uses FindAll method of DAO to fetch list of people from database. 
	Get the lastest posts from these users and sent back as reponse to API
*/
func getNewFeeds(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	followTo, err := dao.GetFollowTo(params["user_id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var page int
	page = 1
	if params["page"] > 0 {
		page = params["page"]
	} 
	limit = page * 10
	feeds, err := dao.GetFeeds(followTo, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, feeds)
}

func Search(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var page int
	page = 1
	if params["page"] > 0 {
		page = params["page"]
	} 
	limit = page * 10
	searchResult, err := dao.Search(params["keyword"], limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, searchResult)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
// The code  creates a controller for each endpoint, then expose an HTTP server on port 3000.
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/{user_id}/{page}", getNewFeeds).Methods("GET")
	r.HandleFunc("/api/{keyword}", search).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
