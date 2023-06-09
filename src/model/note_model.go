package model

type Note struct {
	NoteId  int    `form:"noteid" json:"noteid"`
	UserId  int    `form:"userid" json:"userid"`
	Owner   string `form:"owner" json:"owner"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

type CreatedNote struct {
	UserId  int    `form:"userid" json:"userid"`
	Owner   string `form:"owner" json:"owner"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Note
}

type SingleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Note
}

type CreatedResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    CreatedNote
}

type DeletedResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
