package enums

type Type int

const (
	Deposit Type = iota
	Withdraw
	Bet
	Transfer
	UnknownType
)
