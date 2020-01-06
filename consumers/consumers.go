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
	"time"

	"github.com/sensorsdata/sa-sdk-go/structs"
	"github.com/sensorsdata/sa-sdk-go/utils"
)

type Consumer interface {
	Send(data structs.EventData) error
	Flush() error
	Close() error
	ItemSend(item structs.Item) error
}

func send(url string, data string, to time.Duration, list bool) error {
	pdata := ""

	if list {
		rdata, err := utils.GeneratePostDataList(data)
		if err != nil {
			return err
		}
		pdata = rdata
	} else {
		rdata, err := utils.GeneratePostData(data)
		if err != nil {
			return err
		}
		pdata = rdata
	}

	err := utils.DoRequest(url, pdata, to)
	if err != nil {
		return err
	}

	return nil
}
