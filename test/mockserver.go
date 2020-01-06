/*
 * Created by dengshiwei on 2020/01/06.
 * Copyright 2015－2020 Sensors Data Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func init() {
	http.Handle("/sa", &mockHandler{})
	http.Handle("/debug", &mockHandler{})
}

type mockHandler struct{}

func (h *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("method invalid")
		http.Error(w, "method invalid", http.StatusBadRequest)
		return
	}

	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("body invalid")
		http.Error(w, "body invalid", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//log.Println("Mock server: raw data is", string(result))
	m, err := url.ParseQuery(string(result))
	if err != nil {
		fmt.Println("url query invalid")
		http.Error(w, "url query invalid", http.StatusBadRequest)
		return
	}
	dataFlag := false
	dataListFlag := false

	rawdata := m.Get("data")
	if rawdata != "" {
		dataFlag = true
	} else {
		rawdata = m.Get("data_list")
		if rawdata != "" {
			dataListFlag = true
		}
	}

	if rawdata == "" {
		fmt.Println("get raw data failed")
		http.Error(w, "no data or data_list error", http.StatusBadRequest)
		return
	}

	data, err := base64.StdEncoding.DecodeString(rawdata)
	if err != nil {
		fmt.Println("base64 decode failed", err)
		http.Error(w, "decode base64 failed", http.StatusBadRequest)
		return
	}

	buf := bytes.NewBuffer(data)

	zr, err := gzip.NewReader(buf)
	if err != nil {
		fmt.Println("gzip new reader faild", err)
		http.Error(w, "ungzip failed", http.StatusBadRequest)
		return
	}
	defer zr.Close()

	ungzips, _ := ioutil.ReadAll(zr)
	//log.Printf("Mock server: data(%d) is %s\n", len(ungzips), string(ungzips))
	estr := judge(ungzips, dataFlag, dataListFlag)
	if estr != "" {
		fmt.Println("judge failed", estr)
		http.Error(w, estr, http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusOK)
}

func MockServerRun() {
	http.ListenAndServe(":8106", nil)
}
