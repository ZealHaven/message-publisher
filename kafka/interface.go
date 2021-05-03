package kafka

type (
	IKafka interface {
		Connect() error
		Close() error
		Publish(topic string, content []byte) error
	}

	Kafka struct {}
)