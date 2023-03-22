package entity

type Repository interface {
	AddReceiver(r *Receiver) error
}
