package enums

type Type int

const (
	deposit Type = iota
	withdraw
	bet
	transfer
	unknown
)
