/**
 * Copyright © 2020 David B. Haifley. All rights reserved.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *   http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package spect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request values are used to describe test API requests.
type Request struct {
	URL    string `json:"url"`
	Method string `json:"method"`
	Body   string `json:"body,omitempty"`
}

// NewRequest returns a pointer to a new request value.
func NewRequest(url, method, body string) *Request {
	return &Request{
		URL:    url,
		Method: method,
		Body:   body,
	}
}

// String returns the value as a JSON format string.
func (r *Request) String() string {
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(r); err != nil {
		return "ERROR unable to encode value"
	}

	return b.String()
}

// Equal tests for equality between two values.
func (r *Request) Equal(b *Request) bool {
	if b == nil {
		return false
	}

	switch {
	case r.URL != b.URL:
		return false
	case r.Method != b.Method:
		return false
	case r.Body != b.Body:
		return false
	}

	return true
}

// Response values are used to describe test API responses.
type Response struct {
	Code int    `json:"code"`
	Body string `json:"body,omitempty"`
}

// NewResponse returns a pointer to a new response value.
func NewResponse(code int, body string) *Response {
	return &Response{
		Code: code,
		Body: body,
	}
}

// String returns the value as a JSON format string.
func (r *Response) String() string {
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(r); err != nil {
		return "ERROR unable to encode value"
	}

	return b.String()
}

// Equal tests for equality between two values.
func (r *Response) Equal(b *Response) bool {
	if b == nil {
		return false
	}

	switch {
	case r.Code != b.Code:
		return false
	case r.Body != b.Body:
		return false
	}

	return true
}

// SpecTest values are used to specify an individual API endpoint test.
type SpecTest struct {
	cli *http.Client
	Req *Request  `json:"req,omitempty"`
	Exp *Response `json:"exp,omitempty"`
	Res *Response `json:"res,omitempty"`
}

// NewSpecTest returns a pointer to a new spec test value.
func NewSpecTest(cli *http.Client, req *Request, exp *Response) *SpecTest {
	return &SpecTest{
		cli: cli,
		Req: req,
		Exp: exp,
	}
}

// String returns the value as a JSON format string.
func (st *SpecTest) String() string {
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(st); err != nil {
		return "ERROR unable to encode value"
	}

	return b.String()
}

// Equal tests for equality between two values.
func (st *SpecTest) Equal(b *SpecTest) bool {
	if b == nil {
		return false
	}

	switch {
	case st.Req == nil && b.Req != nil:
		return false
	case st.Req != nil && !st.Req.Equal(b.Req):
		return false
	case st.Exp == nil && b.Exp != nil:
		return false
	case st.Exp != nil && !st.Exp.Equal(b.Exp):
		return false
	case st.Res == nil && b.Res != nil:
		return false
	case st.Res != nil && !st.Res.Equal(b.Res):
		return false
	}

	return true
}

// Run executes the test and returns a bool indicating pass or fail result.
func (st *SpecTest) Run() (bool, error) {
	if st.cli == nil {
		return false, fmt.Errorf("invalid test http client")
	}

	req, err := http.NewRequest(st.Req.Method, st.Req.URL,
		bytes.NewBufferString(st.Req.Body))
	if err != nil {
		return false, fmt.Errorf("invalid request: %w", err)
	}

	res, err := st.cli.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed: %w", err)
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("unable to read response body: %w", err)
	}

	st.Res = &Response{
		Code: res.StatusCode,
		Body: string(b),
	}

	return st.Exp.Equal(st.Res), nil
}
