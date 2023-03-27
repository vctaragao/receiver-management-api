package entity

type Repository interface {
	AddReceiver(r *Receiver) (*Receiver, error)
	UpdateReceiver(r *Receiver, corporateName, cpfCnpj, email string) (*Receiver, error)
	GetReceiverWithPix(receiverId uint) (*Receiver, *Pix, error)

	FindReceiversBy(searchParams string, page int) ([]Receiver, int, error)
	FindReceivers(page int) ([]Receiver, int, error)

	AddPix(receiverId uint, p *Pix) (*Pix, error)
	UpdatePix(p *Pix) (*Pix, error)
}
