package enums

type Status int

const (
	Accepted Status = iota
	Rejected
	Pending
	UnknownStatus
)
