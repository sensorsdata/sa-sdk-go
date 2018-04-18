package test

import (
	"fmt"
	"testing"

	sdk "github.com/sensorsdata/sa-sdk-go"
)

func TestBatchConsumer(t *testing.T) {
	go MockServerRun()

	c, err := sdk.InitBatchConsumer("http://localhost:8106/sa", 3, 1000)
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
		t.Fatal("batch consumer track failed", err)
		return
	}
	err = sa.Track(distinctId, event, properties, true)
	if err != nil {
		t.Fatal("batch consumer track failed", err)
		return
	}
	err = sa.Track(distinctId, event, properties, true)
	if err != nil {
		t.Fatal("batch consumer track failed", err)
		return
	}

	t.Log("batch consumer ok")
}
