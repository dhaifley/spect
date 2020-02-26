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
)

// Request values are used to describe test API requests.
type Request struct {
	URL    string `json:"url"`
	Method string `json:"method"`
	Body   string `json:"body,omitempty"`
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
	Code int64  `json:"code"`
	Body string `json:"body,omitempty"`
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
	Req *Request  `json:"req"`
	Res *Response `json:"res"`
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
	case !st.Req.Equal(b.Req):
		return false
	case !st.Res.Equal(b.Res):
		return false
	}

	return true
}
