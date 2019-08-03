package consumers

import (
	"encoding/json"
	"time"

	"github.com/sensorsdata/sa-sdk-go/structs"
)

type DefaultConsumer struct {
	Url     string
	Timeout time.Duration
}

func InitDefaultConsumer(url string, timeout int) (*DefaultConsumer, error) {
	return &DefaultConsumer{Url: url, Timeout: time.Duration(timeout) * time.Millisecond}, nil
}

func (c *DefaultConsumer) Send(data structs.EventData) error {
	jdata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return send(c.Url, string(jdata), c.Timeout, false)
}

func (c *DefaultConsumer) ItemSend(item structs.Item) error {
	itemData, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return send(c.Url, string(itemData), c.Timeout, false)
}

func (c *DefaultConsumer) Flush() error {
	return nil
}

func (c *DefaultConsumer) Close() error {
	return nil
}
