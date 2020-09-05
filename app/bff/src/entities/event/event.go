package event

type Event struct {
	Func func() error
}
