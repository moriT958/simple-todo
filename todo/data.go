package todo

type todo struct {
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

type Data struct {
	Todos []todo `json:"todos"`
}
