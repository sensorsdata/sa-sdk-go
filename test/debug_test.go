package test

import (
	"fmt"
	"testing"

	sdk "github.com/sensorsdata/sa-sdk-go"
)

func TestDebugConsumer(t *testing.T) {
	go MockServerRun()

	c, err := sdk.InitDebugConsumer("http://localhost:8106/sa", false, 1000)
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
		t.Fatal("debug consumer track failed", err)
		return
	}

	t.Log("debug consumer ok")
}
