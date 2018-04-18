package test

import (
	"fmt"
	"time"
	"testing"

	sdk "github.com/sensorsdata/sa-sdk-go"
)

const (
	FILE_NAME = "./cctest.log"
)

func TestLoggingConsumer(t *testing.T) {
	go MockServerRun()

	c, err := sdk.InitLoggingConsumer(FILE_NAME, false)
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
		t.Fatal("logging consumer track failed", err)
		return
	}

	time.Sleep(time.Millisecond)

	today := time.Now().Format("2006-01-02")
	logfile := fmt.Sprintf("%s.%s", FILE_NAME, today)
	estr, _ := judgeFile(logfile)
	if estr != "" {
		t.Fatal("logging consumer track failed", estr)
		return
	}

	t.Log("logging consumer ok")
}
