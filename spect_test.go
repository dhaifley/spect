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
	a := &Request{
		URL:    "test.com",
		Method: "GET",
		Body:   "test",
	}

	b := &Request{
		URL:    "test.com",
		Method: "GET",
		Body:   "test",
	}

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
	a := &Response{
		Code: 200,
		Body: "test",
	}

	b := &Response{
		Code: 200,
		Body: "test",
	}

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

	b := &SpecTest{
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

	if !a.Equal(b) {
		t.Error("Expected equal: true, got: false")
	}

	b.Res.Body = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b.Req.Body = "error"
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}

	b = nil
	if a.Equal(b) {
		t.Error("Expected equal: false, got: true")
	}
}
