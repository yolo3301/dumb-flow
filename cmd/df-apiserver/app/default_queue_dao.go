package app

import (
	"fmt"

	model "github.com/yolo3301/dumb-flow/pkg/df-model"

	"github.com/Shopify/sarama"

	// "crypto/tls"
	// "crypto/x509"
	"encoding/json"
	"errors"
	"flag"
	// "io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	notificationTopic = "notification"
)


// DefaultQueueDAO - queue dao
type DefaultQueueDAO struct {
	producer sarama.SyncProducer
	/*
		usage:
		partition, offset, err := s.DataCollector.SendMessage(&sarama.ProducerMessage{
				Topic: "important",
				Value: sarama.StringEncoder(r.URL.RawQuery),
			})
	*/

	consumer sarama.Consumer
}

// NewDefaultQueueDAO init-connection
func NewDefaultQueueDAO() (*DefaultQueueDAO, error) {
	// prepare broker list - used both in producer and consumer
	brokers := flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
	if *brokers == "" {
		return nil, fmt.Errorf("Cannot get Kafka brokers from KAFKA_PEERS")
	}
	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	// prepare configs for producer
	// config := sarama.NewConfig()
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.Retry.Max = 10 // Retry up to 10 times to produce the message
	// config.Producer.Return.Successes = true

	// temp - use default config
	producer, err := sarama.NewSyncProducer(brokerList, nil)
	if err != nil {
		log.Fatal("Failed to initate producer")
	}

	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		log.Fatal("Failed to initate consumer")
	}

	return &DefaultQueueDAO{producer: producer, consumer: consumer}, nil

	// TODO: need to use Offset_manager to mark offset
}

// EnqueueNotification - enqueue notice
func (dao DefaultQueueDAO) EnqueueNotification(notifications []model.Notification) ([]model.Notification, error) {
	var noticeArr []model.Notification
	for _, notice := range notifications {
		// convert to JSON
		body, err := json.Marshal(notice)
		if err != nil {
			// failed
			continue
		}

		// assume only one partition for one topic, keep order
		partition, _, err := dao.producer.SendMessage(&sarama.ProducerMessage{
			Topic: notificationTopic,
			Value: sarama.StringEncoder(string(body)),
		})

		if partition != 0 || err != nil {
			// log error
			continue
		}

		// send message success
		noticeArr = append(noticeArr, notice)
	}

	return noticeArr, errors.New(string(len(notifications) - len(noticeArr)))

}

// DequeueNotification - dequeue notice
func (dao DefaultQueueDAO) DequeueNotification(count int) ([]model.Notification, error) {
	var noticeArr []model.Notification
	partitionList, err := dao.consumer.Partitions(notificationTopic)
	if err != nil {
		return noticeArr, err
	}

	partitionConsumer, err := dao.consumer.ConsumePartition(notificationTopic, partitionList[0], sarama.OffsetNewest)
	if err != nil {
		return noticeArr, err
	}

	for i := 0; i < count; i++ {
		message := <-partitionConsumer.Messages()
		var notice model.Notification
		err := json.Unmarshal(message.Value, &notice)
		if err != nil {
			break
		}
		noticeArr = append(noticeArr, notice)
	}

	diff := count - len(noticeArr)
	return noticeArr, errors.New(string(diff))
}

// EnqueueEvents -- enqueue events
func (dao DefaultQueueDAO) EnqueueEvents(topic string, events []model.Event) ([]model.Event, error) {
	var eventArr []model.Event
	for _, event := range events {
		body, err := json.Marshal(event)
		if err != nil {
			// failed
			continue
		}

		partition, _, err := dao.producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(string(body)),
		})

		if partition != 0 || err != nil {
			// log error
			continue
		}

		eventArr = append(eventArr, event)
	}

	return eventArr, errors.New(string(len(events) - len(eventArr)))
}

// DequeueEvents - dequeue events
func (dao DefaultQueueDAO) DequeueEvents(topic string, count int) ([]model.Event, error) {
	var eventArr []model.Event
	partitionList, err := dao.consumer.Partitions(topic)
	if err != nil {
		return eventArr, err
	}

	partitionConsumer, err := dao.consumer.ConsumePartition(topic, partitionList[0], sarama.OffsetNewest)
	if err != nil {
		return eventArr, err
	}

	for i := 0; i < count; i++ {
		message := <-partitionConsumer.Messages()
		var event model.Event
		err := json.Unmarshal(message.Value, &event)
		if err != nil {
			break
		}
		eventArr = append(eventArr, event)
	}

	diff := count - len(eventArr)
	return eventArr, errors.New(string(diff))
}

// Close - avoid memory leak
func (dao DefaultQueueDAO) Close() {
	dao.consumer.Close()
	dao.producer.Close()
}

// SanityCheck - check for QueueDao
func (dao DefaultQueueDAO) SanityCheck() (string, error) {
	// create topic: testqueue

	// producer - enqueue

	// consumer - dequeue

	return "", fmt.Errorf("Not implemented")
}
