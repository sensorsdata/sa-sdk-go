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
	c, err := sdk.InitConcurrentLoggingConsumer("./log.data", false)
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
	itemType := "num"
	itemId := "12312"

	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"uint_value": uint(1),
	})
	fmt.Println("uint_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"uint8_value": uint8(8),
	})
	fmt.Println("uint8_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"uint16_value": uint16(16),
	})
	fmt.Println("uint16_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"uint32_value": uint32(32),
	})
	fmt.Println("uint32_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"uint64_value": uint64(64),
	})
	fmt.Println("uint64_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"int_value": int(1),
	})
	fmt.Println("int_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"int8_value": int8(-8),
	})
	fmt.Println("int8_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"int16_value": int16(-16),
	})
	fmt.Println("int16_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"int32_value": int32(-32),
	})

	fmt.Println("int32_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"int64_value": int64(-64),
	})
	fmt.Println("int64_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"bool_value": true,
	})
	fmt.Println("bool_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"float32_value": float32(3.14),
	})
	fmt.Println("float32_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"float64_value": float64(3.141592653589793),
	})
	fmt.Println("float64_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"string_value": "hello",
	})
	fmt.Println("string_value err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"string_array": []string{"a", "b", "c"},
	})
	fmt.Println("string_array err:", err)
	err = sa.ItemSet(itemType, itemId, map[string]interface{}{
		"map_value": map[string]int{"a": 1},
	})

	fmt.Println("string_too_long err:", err)

	fmt.Println("track done")
}
