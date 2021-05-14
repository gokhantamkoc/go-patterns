package iota_usage

type Operation uint

const (
	Create Operation = 1 << iota
	Update
)

func GetValueOfOperation(op Operation) uint {
	return uint(op)
}
