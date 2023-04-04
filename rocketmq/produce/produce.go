package produce

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/beslow/goblog/config"
	log "github.com/sirupsen/logrus"
)

var instance rocketmq.Producer

func init() {
	if os.Getenv("GO_TEST") != "" {
		return
	}

	var err error

	addr := config.Rocketmq.NamesrvAddr + ":9876"

	instance, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{addr})),
		producer.WithRetry(2),
		producer.WithGroupName("goblog"),
	)

	if err != nil {
		log.Errorf("new producer error: %s", err.Error())
		os.Exit(1)
	}

	err = instance.Start()
	if err != nil {
		log.Errorf("start producer error: %s", err.Error())
		os.Exit(1)
	}
}

func Do(topic, body string) {
	if os.Getenv("GO_TEST") != "" {
		return
	}

	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(body),
	}

	res, err := instance.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
}

func Close() {
	if os.Getenv("GO_TEST") != "" {
		return
	}

	err := instance.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
