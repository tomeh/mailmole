package contracts

type Listener interface {
	Publish(string)
}
