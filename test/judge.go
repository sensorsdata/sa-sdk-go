package test
import (
	"os"
	"io"
	"bufio"
	"encoding/json"
	"github.com/sensorsdata/sa-sdk-go/structs"
)

type EDL []structs.EventData

func judge(data []byte, dataFlag, dataListFlag bool) string {
	if dataFlag {
		return judgeData(data)
	}
	if dataListFlag {
		return judgeDataList(data)
	}
	return "error"
}

func judgeData(data []byte) string {
	ed := structs.EventData{}
	err := json.Unmarshal(data, &ed)
	if err != nil {
		return "unmarshal data failed"
	}

	return demoCompare(ed)
}

func judgeDataList(data []byte) string {
	edl := EDL{}
	err := json.Unmarshal(data, &edl)
	if err != nil {
		return "unmarshal data failed"
	}

	return demoListCompare(edl)
}

func judgeFile(name string) (string, int) {
	defer cleanFile(name)

	file, err := os.OpenFile(name, os.O_RDONLY, 0)
	if err != nil {
		return "open logging file failed", -1
	}
	defer file.Close()
	rb := bufio.NewReader(file)
	count := 0
	for {
		line, _, err := rb.ReadLine()
		if err == io.EOF {
			break
		}
		count++

		estr := judgeData(line)
		if estr != "" {
			return estr, count
		}
	}
	return "", count
}

func cleanFile(name string) {
	os.Remove(name)
}
