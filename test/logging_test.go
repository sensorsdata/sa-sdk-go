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
)

const (
	FILE_NAME = "./cctest.log"
)

func TestLoggingConsumer(t *testing.T) {
	go MockServerRun()

	c, err := sdk.InitLoggingConsumer(FILE_NAME, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	sa := sdk.InitSensorsAnalytics(c, "default", false)
	defer sa.Close()

	distinctId := DemoDistinctId
	event := DemoEventString
	properties := DemoProperties
	properties["$time"] = DemoTime

	err = sa.Track(distinctId, event, properties, true)
	if err != nil {
		t.Fatal("logging consumer track failed", err)
		return
	}

	time.Sleep(time.Millisecond)

	today := time.Now().Format("2006-01-02")
	logfile := fmt.Sprintf("%s.%s", FILE_NAME, today)
	estr, _ := judgeFile(logfile)
	if estr != "" {
		t.Fatal("logging consumer track failed", estr)
		return
	}

	t.Log("logging consumer ok")
}
