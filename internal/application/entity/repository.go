package entity

type Repository interface {
	AddReceiver(r *Receiver) (*Receiver, error)
	UpdateReceiver(r *Receiver) (*Receiver, error)
	GetReceiverWithPix(receiverId uint) (*Receiver, *Pix, error)

	AddPix(receiverId uint, p *Pix) (*Pix, error)
	UpdatePix(p *Pix) (*Pix, error)
}
