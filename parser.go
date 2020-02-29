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

	"gopkg.in/yaml.v3"
)

// Example values are used to capture information about a spec example.
type Example struct {
	Name    string
	Type    string
	Levels  []string
	Example interface{}
}

// String returns the value as a JSON formatted string.
func (e *Example) String() string {
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(e); err != nil {
		return "ERROR unable to encode value"
	}

	return b.String()

}

// ParseArray recusively parses a YAML map finding examples.
func ParseArray(a []interface{}, lvl []string,
	exs []*Example, name, extype string) ([]*Example, error) {
	for _, v := range a {
		var err error
		switch vv := v.(type) {
		case []interface{}:
			exs, err = ParseArray(vv, lvl, exs, name, extype)
		case map[string]interface{}:
			exs, err = ParseMap(vv, lvl, exs, name, extype)
		}

		if err != nil {
			return nil, err
		}
	}

	return exs, nil
}

// ParseMap recusively parses a YAML map finding examples.
func ParseMap(m map[string]interface{}, lvl []string,
	exs []*Example, name, extype string) ([]*Example, error) {
	ex := &Example{
		Name:   name,
		Type:   extype,
		Levels: lvl,
	}

	for k, v := range m {
		switch k {
		case "name":
			if n, ok := v.(string); ok {
				name = n
				ex.Name = n
			}
		case "type":
			if t, ok := v.(string); ok {
				extype = t
				ex.Type = t
			}
		}
	}

	for k, v := range m {
		switch k {
		case "example":
			ex.Example = v
		}

		var err error
		switch vv := v.(type) {
		case []interface{}:
			exs, err = ParseArray(vv, append(lvl, k), exs, name, extype)
		case map[string]interface{}:
			exs, err = ParseMap(vv, append(lvl, k), exs, name, extype)
		}

		if err != nil {
			return nil, err
		}
	}

	if len(lvl) > 2 && lvl[len(lvl)-2] == "properties" {
		ex.Name = lvl[len(lvl)-1]
	}

	if ex.Example != nil {
		return append(exs, ex), nil
	}

	return exs, nil
}

// ParseExamples parses examples from OpenAPI specification data into a slice
// of Example values.
func ParseExamples(data string) ([]*Example, error) {
	m := make(map[string]interface{}, 0)
	err := yaml.NewDecoder(bytes.NewBufferString(data)).Decode(&m)
	if err != nil {
		return nil, fmt.Errorf("unable to parse yaml: %w", err)
	}

	return ParseMap(m, []string{}, []*Example{}, "", "")
}

// ParseTests parses integration testing data from OpenAPI YAML specification
// data into a slice of SpecTest values.
func ParseTests(data string) ([]*Example, error) {
	m := make(map[string]interface{}, 0)
	err := yaml.NewDecoder(bytes.NewBufferString(data)).Decode(&m)
	if err != nil {
		return nil, fmt.Errorf("unable to parse yaml: %w", err)
	}

	return ParseMap(m, []string{}, []*Example{}, "", "")
}
