package create_receiver

type InputDto struct {
	CorporateName string
	Cpf           string
	Cnpj          string
	Email         string
	PixType       string
	PixKey        string
}

type OutputDto struct {
	Id uint
	InputDto
}
