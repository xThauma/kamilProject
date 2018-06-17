package main

import (
  "net/http"
  "log"
  "encoding/json"
 // "github.com/gocql/gocql"
  "github.com/kamilProject/Messages"
  "github.com/kamilProject/Cassandra"
  "github.com/gorilla/mux"
)

type heartbeatResponse struct {
  Status string `json:"status"`
  Code int `json:"code"`
}

func main() {
	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	router.HandleFunc("/api", Messages.Get)
	router.HandleFunc("/api/message", Messages.Post)
	router.HandleFunc("/api/messages/{email}", Messages.GetByEmail)
	router.HandleFunc("/api/post/{magic_number}", Messages.GetByMagicNumber)
 	log.Fatal(http.ListenAndServe(":8080", router))
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}


