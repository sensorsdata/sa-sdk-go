package main
import (
	"fmt"
	sdk "github.com/sensorsdata/sa-sdk-go"
)

func main() {
	c, err := sdk.InitDefaultConsumer("http://localhost:8106/sa", 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	sa := sdk.InitSensorsAnalytics(c, "default", false)
	defer sa.Close()

	distinctId := "12345"
	event := "ViewProduct"
	properties := map[string]interface{}{
		"price": 12,
		"name": "apple",
		"somedata": []string{"a", "b"},
	}

	err = sa.Track(distinctId, event, properties, true)
	if err != nil {
		fmt.Println("track failed", err)
		return
	}

	fmt.Println("track done")
}
