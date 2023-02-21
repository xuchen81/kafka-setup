package mai

func Test() {
	AsyncProducer, err := sarama.NewAsyncProducer([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, config)

}

func main() {
	Test()
	SyncProducer.SendMessage(&sarama.ProducerMessage{})
}
}
