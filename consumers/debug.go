package consumers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sensorsdata/sa-sdk-go/structs"
	"github.com/sensorsdata/sa-sdk-go/utils"
)

type DebugConsumer struct {
	Url       string
	WriteData bool
	Timeout   time.Duration
}

func InitDebugConsumer(surl string, writeData bool, timeout int) (*DebugConsumer, error) {
	u, err := url.Parse(surl)
	if err != nil {
		return nil, err
	}
	u.Path = "/debug"

	return &DebugConsumer{Url: u.String(), WriteData: writeData, Timeout: time.Duration(timeout) * time.Millisecond}, nil
}

func (c *DebugConsumer) send(url string, data string, to time.Duration, list bool) error {
	pdata := ""

	if list {
		rdata, err := utils.GeneratePostDataList(data)
		if err != nil {
			return err
		}
		pdata = rdata
	} else {
		rdata, err := utils.GeneratePostData(data)
		if err != nil {
			return err
		}
		pdata = rdata
	}

	res, status, err := doRequestDebug(url, pdata, to, c.WriteData)
	if err != nil && status == 0 {
		fmt.Fprintf(os.Stderr, "Send failed: %s\n", err)
		return err
	}

	if status == 200 {
		fmt.Fprintf(os.Stdout, "Valid message: %s\n", string(data))
		return nil
	} else {
		fmt.Fprintf(os.Stderr, "Invalid message: %s\n", string(data))
		fmt.Fprintf(os.Stderr, "Ret_code: %d\n", status)
		fmt.Fprintf(os.Stderr, "Ret_content: %s\n", res)
	}

	if status >= 300 {
		return errors.New("Bad http status")
	}

	return nil
}

func doRequestDebug(url, args string, to time.Duration, writeData bool) (string, int, error) {
	var resp *http.Response

	data := bytes.NewBufferString(args)

	req, _ := http.NewRequest("POST", url, data)

	if !writeData {
		req.Header.Add("Dry-Run", "true")
	}

	client := &http.Client{Timeout: to}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Request failed to url: %s\n", url)
		return "", 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Request failed http status is %d\n", resp.StatusCode)
		return "", resp.StatusCode, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request debug read response failed: %s\n", err)
		return "", resp.StatusCode, err
	}

	return string(body), resp.StatusCode, nil
}

func (c *DebugConsumer) Send(data structs.EventData) error {
	jdata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.send(c.Url, string(jdata), c.Timeout, false)
}

func (c *DebugConsumer) ItemSend(item structs.Item) error {
	itemData, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return c.send(c.Url, string(itemData), c.Timeout, false)
}

func (c *DebugConsumer) Flush() error {
	return nil
}

func (c *DebugConsumer) Close() error {
	return nil
}
