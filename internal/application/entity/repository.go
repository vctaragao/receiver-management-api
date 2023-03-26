package entity

type Repository interface {
	AddReceiver(r *Receiver) (*Receiver, error)
	AddPix(receiverId uint, p *Pix) (*Pix, error)
	GetReceiverWithPix(receiverId uint) (*Receiver, error)
}
