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
    import sdk "github.com/sensorsdata/sa-sdk-go"

    // Gets the url of Sensors Analytics in the home page.
    SA_SERVER_URL = 'YOUR_SERVER_URL'

    // Initialized the Sensors Analytics SDK with Default Consumer
    consumer = sdk.InitDefaultConsumer(SA_SERVER_URL)
    sa = sdk.InitSensorsAnalytics(consumer)

    properties := map[string]interface{}{
         "price": 12,
         "name": "apple",
         "somedata": []string{"a", "b"},
    }

    // Track the event 'ServerStart'
    sa.track("ABCDEFG1234567", "ServerStart", properties, false)

    sa.Close()
```

## More Examples
([Examples](_examples)) 

## To Learn More
See our [full manual](http://www.sensorsdata.cn/manual/golang_sdk.html)

