package create_receiver

type InputDto struct {
	CorporateName string
	CpfCnpj       string
	Email         string
	PixType       string
	PixKey        string
}

type OutputDto struct {
	Id uint
	InputDto
}
