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

package test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/sensorsdata/sa-sdk-go"
	"github.com/sensorsdata/sa-sdk-go/utils"
)

const (
	CCFILE_NAME = "./test.log"
	NUM_WRITERS = 3
	NUM_RECORDS = 100
)

func ccwriter() {
	c, err := sdk.InitConcurrentLoggingConsumer(CCFILE_NAME, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	sa := sdk.InitSensorsAnalytics(c, "default", false)
	defer sa.Close()

	distinctId := DemoDistinctId
	event := DemoEventString
	properties := utils.DeepCopy(DemoProperties)
	properties["$time"] = DemoTime

	for i := 0; i < NUM_RECORDS; i++ {
		err = sa.Track(distinctId, event, properties, true)
		if err != nil {
			fmt.Println("concurrentlogging consumer track failed", err)
			return
		}
	}
}

func TestConcurrentLoggingConsumer(t *testing.T) {
	go MockServerRun()

	for i := 0; i < NUM_WRITERS; i++ {
		go ccwriter()
	}

	//10ms is enough
	time.Sleep(time.Millisecond * 10)

	today := time.Now().Format("2006-01-02")
	logfile := fmt.Sprintf("%s.%s", CCFILE_NAME, today)
	estr, count := judgeFile(logfile)
	if estr != "" {
		t.Fatal("concurrentlogging consumer track failed", estr)
		return
	}

	if count != NUM_WRITERS*NUM_RECORDS {
		t.Fatal("concurrentlogging consumer track failed, count not match", count)
		return
	}

	fmt.Println("concurrent records count match")
	t.Log("concurrentlogging consumer ok")
}
