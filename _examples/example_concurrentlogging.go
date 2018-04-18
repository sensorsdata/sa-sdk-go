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
		"$ip": "2.2.2.2",
		"ProductId": "123456",
		"ProductCatalog": "Laptop Computer",
		"IsAddedToFav": true,
	}

	err = sa.Track(distinctId, event, properties, true)
	if err != nil {
		fmt.Println("track failed", err)
		return
	}

	fmt.Println("track done")
}
