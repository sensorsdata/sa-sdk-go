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

package main

import (
	"fmt"
	sdk "github.com/sensorsdata/sa-sdk-go"
)

func main() {
	c, err := sdk.InitLoggingConsumer("./log.data", false)
	if err != nil {
		fmt.Println(err)
		return
	}

	sa := sdk.InitSensorsAnalytics(c, "default", false)
	defer sa.Close()

	distinctId := "ABCDEF123456"
	event := "ViewProduct"
	properties := map[string]interface{}{
		"$ip":            "2.2.2.2",
		"ProductId":      "123456",
		"ProductCatalog": "Laptop Computer",
		"IsAddedToFav":   true,
	}

	err = sa.Track(distinctId, event, properties, true)
	if err != nil {
		fmt.Println("track failed", err)
		return
	}

	// 1. uint8
	err = sa.Track(distinctId, "TestUint8", map[string]interface{}{
		"uint8_value": uint8(8),
	}, true)
	fmt.Println("TestUint8:", err)

	// 2. uint16
	err = sa.Track(distinctId, "TestUint16", map[string]interface{}{
		"uint16_value": uint16(16),
	}, true)
	fmt.Println("TestUint16:", err)

	// 3. uint32
	err = sa.Track(distinctId, "TestUint32", map[string]interface{}{
		"uint32_value": uint32(32),
	}, true)
	fmt.Println("TestUint32:", err)

	// 4. uint64
	err = sa.Track(distinctId, "TestUint64", map[string]interface{}{
		"uint64_value": uint64(64),
	}, true)
	fmt.Println("TestUint64:", err)

	// 5. int8
	err = sa.Track(distinctId, "TestInt8", map[string]interface{}{
		"int8_value": int8(-8),
	}, true)
	fmt.Println("TestInt8:", err)

	// 6. int16
	err = sa.Track(distinctId, "TestInt16", map[string]interface{}{
		"int16_value": int16(-16),
	}, true)
	fmt.Println("TestInt16:", err)

	// 7. uint（平台相关，32 或 64 位）
	err = sa.Track(distinctId, "TestUint", map[string]interface{}{
		"uint_value": uint(128),
	}, true)
	fmt.Println("TestUint:", err)

	// 8. float32
	err = sa.Track(distinctId, "TestFloat32", map[string]interface{}{
		"float32_value": float32(3.14),
	}, true)
	fmt.Println("TestFloat32:", err)

	fmt.Println("track done")
}
