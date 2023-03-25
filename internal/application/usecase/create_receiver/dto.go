package create_receiver

type InputDto struct {
	CorporateName string
	Cpf           string
	Cnpj          string
	Email         string
	Pix_type      string
	Pix_key       string
}

type OutputDto struct {
	Id uint
	InputDto
}
