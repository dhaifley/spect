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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestString(t *testing.T) {
	a := &Request{
		URL:    "test.com",
		Method: "GET",
		Body:   "test",
	}

	exp := `{"url":"test.com","method":"GET","body":"test"}` + "\n"
	if a.String() != exp {
		t.Errorf("Expected string: %v, got: %v", exp, a.String())
	}
}

func TestRequestEqual(t *testing.T) {
	a := NewRequest("test.com", "GET", "test")
	b := NewRequest("test.com", "GET", "test")
	if !a.Equal(b) {
		t.Error("Expected equal: true, got: false")
	}

	b.Body = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b.Method = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b.URL = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b = nil
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}
}

func TestResponseString(t *testing.T) {
	a := &Response{
		Code: 200,
		Body: "test",
	}

	exp := `{"code":200,"body":"test"}` + "\n"
	if a.String() != exp {
		t.Errorf("Expected string: %v, got: %v", exp, a.String())
	}
}

func TestResponseEqual(t *testing.T) {
	a := NewResponse(200, "test")
	b := NewResponse(200, "test")
	if !a.Equal(b) {
		t.Error("Expected equal: true, got: false")
	}

	b.Body = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b.Code = 500
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b = nil
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}
}

func TestSpecTestString(t *testing.T) {
	a := &SpecTest{
		Req: &Request{
			URL:    "test.com",
			Method: "GET",
			Body:   "test",
		},
		Res: &Response{
			Code: 200,
			Body: "test",
		},
	}

	exp := `{"req":{"url":"test.com","method":"GET","body":"test"},"res":` +
		`{"code":200,"body":"test"}}` + "\n"
	if a.String() != exp {
		t.Errorf("Expected string: %v, got: %v", exp, a.String())
	}
}

func TestSpecTestEqual(t *testing.T) {
	a := NewSpecTest(nil,
		NewRequest("test.com", "GET", "test"),
		NewResponse(200, "test"))

	b := NewSpecTest(nil,
		NewRequest("test.com", "GET", "test"),
		NewResponse(200, "test"))

	if !a.Equal(b) {
		t.Error("Expected equal: true, got: false")
	}

	b.Res = NewResponse(500, "error")
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	a.Res = NewResponse(200, "test")
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b.Exp.Body = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	a.Exp = nil
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b.Req.Body = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	a.Req = nil
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b = nil
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}
}

func TestSpecTestRun(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintln(w, "test")
		}))
	defer ts.Close()

	st := NewSpecTest(&http.Client{},
		NewRequest(ts.URL, "GET", ""),
		NewResponse(200, "test\n"))

	pass, err := st.Run()
	if err != nil {
		t.Error(err)
	}

	if !pass {
		t.Errorf("Expected response: %v, got: %v", st.Exp, st.Res)
	}
}
