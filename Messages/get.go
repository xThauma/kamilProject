package Messages

import (
	"net/http"
	"encoding/json"
  	"github.com/kamilProject/Cassandra"
  	"github.com/kamilProject/checkmail"
	"github.com/gorilla/mux"
	//"fmt"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var messageList []Message
	m := map[string]interface{}{}

	query := "SELECT magic_number, email, content, title FROM messages"
	iterable := Cassandra.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		messageList = append(messageList, Message{
			Email: 		    m["email"].(string),
			Content: 	    m["content"].(string),
			Title:     	    m["title"].(string),
			Magic_number:       m["magic_number"].(int),
		})
		m = map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(AllMessagesResponse{Messages: messageList})
}

func GetByEmail(w http.ResponseWriter, r *http.Request) {
	var messageList []Message
	m := map[string]interface{}{}
	var errs []string
	vars := mux.Vars(r)
	emailMessage := vars["email"]
	err := checkmail.ValidateFormat(emailMessage)
	if err != nil {
		errs = append(errs, err.Error())
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	} else {
	query := "SELECT magic_number, email,content,title FROM messages where email = ? allow filtering"
	iterable := Cassandra.Session.Query(query, emailMessage).Iter()
	for iterable.MapScan(m) {
		messageList = append(messageList, Message{
			Email: 		    m["email"].(string),
			Content: 	    m["content"].(string),
			Title:     	    m["title"].(string),
			Magic_number:       m["magic_number"].(int),
		})
		m = map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(AllMessagesResponse{Messages: messageList})
	}
}

func GetByMagicNumber(w http.ResponseWriter, r *http.Request) {
	var messageList []Message
	m := map[string]interface{}{}
	vars := mux.Vars(r)
	magicNumberMessage := vars["magic_number"]
	query := "SELECT magic_number, email,content,title FROM messages where magic_number = ?"
	iterable := Cassandra.Session.Query(query, magicNumberMessage).Iter()
	for iterable.MapScan(m) {
		messageList = append(messageList, Message{
			Email: 		    m["email"].(string),
			Content: 	    m["content"].(string),
			Title:     	    m["title"].(string),
			Magic_number:       m["magic_number"].(int),
		})
		m = map[string]interface{}{}
	json.NewEncoder(w).Encode(AllMessagesResponse{Messages: messageList})
	queryToDel := "DELETE FROM messages where magic_number = ?"	
	Cassandra.Session.Query(queryToDel, magicNumberMessage).Exec()
	}
}

