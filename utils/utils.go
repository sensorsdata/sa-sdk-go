package utils
import (
	"os"
	"fmt"
	"time"
	"bytes"
	"errors"
	"net/url"
	"net/http"
	"compress/gzip"
	"encoding/base64"
)

func DoRequest(url, args string, to time.Duration) error {
	var resp *http.Response

	data := bytes.NewBufferString(args)

	req, _ := http.NewRequest("POST", url , data)

	client := &http.Client{Timeout: to}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Request failed to url: %s\n", url)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Http status is %d\n", resp.StatusCode)
		return errors.New("Bad http status")
	}
	return nil
}

func gzipData(data string) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write([]byte(data))
	if err != nil {
		zw.Close()
		return nil, err
	}
	zw.Close()

	return buf.Bytes(), nil
}

func encodeData(data string) (string, error) {
	gdata, err := gzipData(data)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(gdata)
	return encoded, nil
}

func GeneratePostDataList(data string) (string, error) {
	edata, err := encodeData(data)
	if err != nil {
		return "", err
	}

	v := url.Values{}
	v.Add("data_list", edata)
	v.Add("gzip", "1")

	uedata := v.Encode()

	return uedata, nil
}

func GeneratePostData(data string) (string, error) {
	edata, err := encodeData(data)
	if err != nil {
		return "", err
	}

	v := url.Values{}
	v.Add("data", edata)
	v.Add("gzip", "1")

	uedata := v.Encode()

	return uedata, nil
}

func NowMs() int64 {
	return time.Now().UnixNano() / 1000000
}

func DeepCopy(value map[string]interface{}) map[string]interface{} {
	ncopy := deepCopy(value)
	if nmap, ok := ncopy.(map[string]interface{}); ok {
		return nmap
	}

	return nil
}

func deepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = deepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = deepCopy(v)
		}

		return newSlice
	}

	return value
}
