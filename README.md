# Sensors Analytics

This is the official Golang SDK for Sensors Analytics.

## Easy Installation

You can get Sensors Analytics SDK using go.

```
	go get github.com/sensorsdata/sa-sdk-go
```
Or update sdk with
```
	go get -u github.com/sensorsdata/sa-sdk-go
	
```

Once the SDK is successfully installed, use the Sensors Analytics SDK likes:

```golang
package main

import (
	sdk "github.com/sensorsdata/sa-sdk-go"
	"log"
)

// Gets the url of Sensors Analytics in the home page.
const (
	SA_SERVER_URL = "YOUR_SERVER_URL"

	TIMEOUT = 2000 // milliseconds
)

func main() {
	// Initialized the Sensors Analytics SDK with Default Consumer
	consumer, err := sdk.InitDefaultConsumer(SA_SERVER_URL, TIMEOUT)
	if err != nil {
		log.Fatal("init error", err)
	}
	sa := sdk.InitSensorsAnalytics(consumer, "", false)

	properties := map[string]interface{}{
		"price":    12,
		"name":     "apple",
		"somedata": []string{"a", "b"},
	}

	// Track the event 'ServerStart'
	err = sa.Track("ABCDEFG1234567", "ServerStart", properties, false)
	if err != nil {
		log.Print("track error", err)
	}

	sa.Close()
}
```

## More Examples
([Examples](_examples)) 

## To Learn More
See our [full manual](http://www.sensorsdata.cn/manual/golang_sdk.html)<br>
或者加入神策官方 SDK QQ 讨论群：<br><br>
![ QQ 讨论群](https://github.com/sensorsdata/sa-sdk-android/blob/master/docs/qrCode.jpeg)

## License

Copyright 2015－2019 Sensors Data Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

**禁止一切基于神策数据开源 SDK 的所有商业活动！**
