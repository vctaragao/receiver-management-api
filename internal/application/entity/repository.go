package entity

type Repository interface {
	AddReceiver(r *Receiver) (*Receiver, error)
	UpdateReceiver(r *Receiver) (*Receiver, error)
	GetReceiverWithPix(receiverId uint) (*Receiver, *Pix, error)

	FindReceiversBy(searchParams string, page int) ([]Receiver, error)
	FindReceivers(page int) ([]Receiver, error)

	AddPix(receiverId uint, p *Pix) (*Pix, error)
	UpdatePix(p *Pix) (*Pix, error)
}
