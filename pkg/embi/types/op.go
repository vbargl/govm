package types

type Op interface {
	Process() error
}
