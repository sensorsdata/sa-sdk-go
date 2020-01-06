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
