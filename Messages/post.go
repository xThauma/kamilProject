package Messages

import (
	"net/http"
	"encoding/json"
  	"github.com/kamilProject/Cassandra"
	"fmt"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	message, errs := FormToMessage(r)
	var created bool = false
	if len(errs) == 0 {
		if err := Cassandra.Session.Query(`
		INSERT INTO messages (email, content, title, magic_number) VALUES (?, ?, ?, ?) USING TTL 300`,
		message.Email, message.Content, message.Title, message.Magic_number).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}
	}

	if created {
		fmt.Println("created")
	} else {
		fmt.Println("errors", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
