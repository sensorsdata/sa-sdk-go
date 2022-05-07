/*
 * Created by dengshiwei on 2020/01/06.
 * Copyright 2015－2020 Sensors Data Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package consumers

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/sensorsdata/sa-sdk-go/structs"
)

const (
	BATCH_DEFAULT_MAX = 50
)

type BatchConsumer struct {
	Url        string
	Max        int
	DataBuffer []structs.EventData
	ItemBuffer []structs.Item
	Timeout    time.Duration
	lock       sync.Mutex
}

func InitBatchConsumer(url string, max, timeout int) (*BatchConsumer, error) {
	if max > BATCH_DEFAULT_MAX {
		max = BATCH_DEFAULT_MAX
	}

	c := &BatchConsumer{Url: url, Max: max, Timeout: time.Duration(timeout) * time.Millisecond}
	c.DataBuffer = make([]structs.EventData, 0, max)

	return c, nil
}

func (c *BatchConsumer) Send(data structs.EventData) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.DataBuffer = append(c.DataBuffer, data)
	if (len(c.DataBuffer) + len(c.ItemBuffer)) < c.Max {
		return nil
	}

	return c.Flush()
}

func (c *BatchConsumer) Flush() error {
	// 刷新 Event 数据
	if len(c.DataBuffer) != 0 {
		jdata, err := json.Marshal(c.DataBuffer)
		if err != nil {
			return err
		}

		err = send(c.Url, string(jdata), c.Timeout, true)
		c.DataBuffer = c.DataBuffer[:0]
		if err != nil {
			return err
		}
	}

	// 刷新 Item 数据
	if len(c.ItemBuffer) != 0 {
		itemData, err := json.Marshal(c.ItemBuffer)
		if err != nil {
			return err
		}

		err = send(c.Url, string(itemData), c.Timeout, true)

		c.ItemBuffer = c.ItemBuffer[:0]
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *BatchConsumer) Close() error {
	return c.Flush()
}

func (c *BatchConsumer) ItemSend(item structs.Item) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.ItemBuffer = append(c.ItemBuffer, item)
	if (len(c.DataBuffer) + len(c.ItemBuffer)) < c.Max {
		return nil
	}

	return c.Flush()
}
