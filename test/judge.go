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
	"bufio"
	"encoding/json"
	"github.com/sensorsdata/sa-sdk-go/structs"
	"io"
	"os"
)

type EDL []structs.EventData

func judge(data []byte, dataFlag, dataListFlag bool) string {
	if dataFlag {
		return judgeData(data)
	}
	if dataListFlag {
		return judgeDataList(data)
	}
	return "error"
}

func judgeData(data []byte) string {
	ed := structs.EventData{}
	err := json.Unmarshal(data, &ed)
	if err != nil {
		return "unmarshal data failed"
	}

	return demoCompare(ed)
}

func judgeDataList(data []byte) string {
	edl := EDL{}
	err := json.Unmarshal(data, &edl)
	if err != nil {
		return "unmarshal data failed"
	}

	return demoListCompare(edl)
}

func judgeFile(name string) (string, int) {
	defer cleanFile(name)

	file, err := os.OpenFile(name, os.O_RDONLY, 0)
	if err != nil {
		return "open logging file failed", -1
	}
	defer file.Close()
	rb := bufio.NewReader(file)
	count := 0
	for {
		line, _, err := rb.ReadLine()
		if err == io.EOF {
			break
		}
		count++

		estr := judgeData(line)
		if estr != "" {
			return estr, count
		}
	}
	return "", count
}

func cleanFile(name string) {
	os.Remove(name)
}
