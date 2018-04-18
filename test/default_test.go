package test

import (
	"fmt"
	"testing"

	sdk "github.com/sensorsdata/sa-sdk-go"
)

func TestDefaultConsumer(t *testing.T) {
	go MockServerRun()

	c, err := sdk.InitDefaultConsumer("http://localhost:8106/sa", 1000)
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
		t.Fatal("default consumer track failed", err)
		return
	}

	t.Log("Default consumer ok")
}

