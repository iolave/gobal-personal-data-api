package humans

type HumanId struct {
	Country      string `json: "country"`
	DocumentType string `json: "documentType"`
	DocumentId   string `json: "documentId"`
}

type Human struct {
	Id      HumanId `json: "id"`
	Name    string  `json: "name"`
	Gender  string  `json: "gender"`
	City    string  `json: "city"`
	Address string  `json: "address"`
}
