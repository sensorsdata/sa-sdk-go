package consumers

import (
	"encoding/json"
	"time"

	"github.com/sensorsdata/sa-sdk-go/structs"
)

const (
	BATCH_DEFAULT_MAX = 50
)

type BatchConsumer struct {
	Url     string
	Max     int
	buffer  []structs.EventData
	Timeout time.Duration
}

func InitBatchConsumer(url string, max, timeout int) (*BatchConsumer, error) {
	if max > BATCH_DEFAULT_MAX {
		max = BATCH_DEFAULT_MAX
	}

	c := &BatchConsumer{Url: url, Max: max, Timeout: time.Duration(timeout) * time.Millisecond}
	c.buffer = make([]structs.EventData, 0, max)

	return c, nil
}

func (c *BatchConsumer) Send(data structs.EventData) error {
	c.buffer = append(c.buffer, data)
	if len(c.buffer) >= c.Max {
		return c.Flush()
	}
	return nil
}

func (c *BatchConsumer) Flush() error {
	if len(c.buffer) == 0 {
		return nil
	}
	jdata, err := json.Marshal(c.buffer)
	if err != nil {
		return err
	}

	err = send(c.Url, string(jdata), c.Timeout, true)

	c.buffer = c.buffer[:0]

	return err
}

func (c *BatchConsumer) Close() error {
	return c.Flush()
}
