package port

type PublisherConnection interface {
	PublisherConfig(name string) error
	PublishMsg(msg []byte) error
	Close()
}
