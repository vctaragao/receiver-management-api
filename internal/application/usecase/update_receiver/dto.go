package update_receiver

type InputDto struct {
	ReceiverId    uint
	CorporateName string
	CpfCnpj       string
	Email         string
	PixType       string
	PixKey        string
}

type OutputDto struct {
	ReceiverId    uint
	CorporateName string
	CpfCnpj       string
	Email         string
	Status        string
	PixType       string
	PixKey        string
}
