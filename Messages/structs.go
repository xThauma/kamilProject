package Messages

type Message struct {
	Email 		string 		`json:"email"`
	Title  		string 		`json:"title"`
	Content     	string 		`json:"content"`
	Magic_number    int 		`json:"magic_number"`
}

type GetMessageResponse struct {
	Message Message `json:"message"`
}

type AllMessagesResponse struct {
	Messages []Message `json:"messages_data"`
}


type ErrorResponse struct {
	Errors []string `json:"errors"`
}

