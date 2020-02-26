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

// Request values are used to describe test API requests.
type Request struct {
	URL    string `json:"url"`
	Method string `json:"method"`
	Body   string `json:"body,omitempty"`
}

// Response values are used to describe test API responses.
type Response struct {
	Code int64  `json:"code"`
	Body string `json:"body,omitempty"`
}

// SpecTest values are used to specify an individual API endpoint test.
type SpecTest struct {
	Req *Request  `json:"req"`
	Res *Response `json:"res"`
}
