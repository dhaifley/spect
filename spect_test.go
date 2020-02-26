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
