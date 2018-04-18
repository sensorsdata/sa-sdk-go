package main
import (
	"fmt"
	sdk "github.com/sensorsdata/sa-sdk-go"
)

func main() {
	c, err := sdk.InitDebugConsumer("http://localhost:8106/sa?project=production", false, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	sa := sdk.InitSensorsAnalytics(c, "default", false)

	distinctId := "ABCDEF123456777"
	event := "ViewInfo"
	properties := map[string]interface{}{
		"$ip": "2.2.2.2",
		"ProductId": "123456",
		"ProductCatalog": "Laptop Computer",
		"IsAddedToFav": true,
	}

	err = sa.Track(distinctId, event, properties, false)
	if err != nil {
		fmt.Println("track failed", err)
		return
	}

	fmt.Println("track done")
}
