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

const testSpecData = `
openapi: "3.0.2"
info:
  title: "Circonus Application Programming Interface"
  version: "0.6.1"
servers:
  - url: "https://api.circonus.com"
paths:
  /broker:
    get:
      tags:
        - Broker
      summary: Retrieves broker configuration details for the current account.
      responses:
        "200":
         content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Broker"
        "400":
          $ref: "#/components/responses/ErrorBadRequest"
        "401":
          $ref: "#/components/responses/ErrorUnauthorized"
        "403":
          $ref: "#/components/responses/ErrorForbidden"
        "404":
          $ref: "#/components/responses/ErrorNotFound"
        "5XX":
          $ref: "#/components/responses/ErrorInternal"
  /broker/{id}:
    get:
      tags:
        - Broker
      summary: Retrieves details for the specified broker by ID.
      description: >
        Retrieves the configuration details for a broker specified by the ID
        in the request path.
      parameters:
        - in: path
          name: id
          description: >
            The broker ID of the configuration details to retrieve.
          required: true
          schema:
            type: integer
            format: int64
            example: 1
      responses:
        "200":
          description: Broker configuration information.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Broker"
        "400":
          $ref: "#/components/responses/ErrorBadRequest"
        "401":
          $ref: "#/components/responses/ErrorUnauthorized"
        "403":
          $ref: "#/components/responses/ErrorForbidden"
        "404":
          $ref: "#/components/responses/ErrorNotFound"
        "5XX":
          $ref: "#/components/responses/ErrorInternal"
components:
  securitySchemes:
    App-Name:
      in: header
      name: X-Circonus-App-Name
      type: apiKey
    Auth-Token:
      in: header
      name: X-Circonus-Auth-Token
      type: apiKey
  schemas:
    Tag:
      type: string
      description: >
        An associated tag.
        A tag is just a string, with or without a colon, such as 'foo', 'bar',
        'datacenter:london', or 'os:linux'. The part of the string before the
        colon is considered the category the tag is in; Tag strings without a
        colon will place the string in the 'uncategorized' category. Circonus
        will lowercase the contents of the string before storing it.
      example: datacenter:primary
    Broker:
      type: object
      properties:
        _cid:
          type: string
          description: >
            The primary identifier of a broker configuration.
            A string containing a broker cid.
          example: /broker/1
        _details:
          type: array
          description: >
            An array of details on every broker that is grouped together.
          example:
            [
              {
                "cn": "us1.noit.circonus.net",
                "external_host": "example.circonus.net",
                "external_port": 8080,
                "ipaddress": "75.102.43.87",
                "minimum_version_required": 1367937537,
                "modules": ["cim", "circonuswindowsagent", "collectd", "dcm"],
                "port": null,
                "skew": "0.00257706642150879",
                "status": null,
                "version": 1370978917,
              },
            ]
          items:
            $ref: "#/components/schemas/BrokerDetails"
        _latitude:
          type: number
          format: double
          description: >
            The latitude of the broker.
            A floating point number indicating GPS location.
          example: "39.043"
        _longitude:
          type: number
          format: double
          description: >
            The longitude of the broker.
            A floating point number indicating GPS location.
          example: "-77.487"
        _name:
          type: string
          description: >
            The name of the broker.
            A string containing freeform text.
          example: Ashburn, VA, US
        _tags:
          type: array
          description: >
            The tags associated with this broker.
            An array of tags. The tags in the array are automatically sorted,
            de-duplicated and transformed into their lowercase canonical form.
          example: ["datacenter:primary", "foo:bar"]
          items:
            $ref: "#/components/schemas/Tag"
        _type:
          type: string
          description: >
            The type of broker, whether public or private.
            A string containing either 'circonus' (a cloud based broker run by
            Circonus) or 'enterprise' (an in-house enterprise broker).
          example: circonus
tags:
  - name: Broker
    description: >
      The Broker service provides access to configuration information about
      brokers servers.
`

func TestParseExamples(t *testing.T) {
	examples, err := ParseExamples(testSpecData)
	if err != nil {
		t.Fatal(err)
	}

	if len(examples) != 9 {
		t.Fatalf("Expected length: 9, got: %v", len(examples))
	}

	expA := `{"Name":"id","Type":"integer","Levels":["paths","/broker/{id}",` +
		`"get","parameters","schema"],"Example":1}` + "\n"
	expB := `{"Name":"_cid","Type":"string","Levels":["components",` +
		`"schemas","Broker","properties","_cid"],"Example":"/broker/1"}` + "\n"
	passA, passB := false, false
	for _, ex := range examples {
		switch {
		case ex.String() == expA:
			passA = true
		case ex.String() == expB:
			passB = true
		}
	}

	if !passA {
		t.Errorf("Expected example: %v, got: %v", expA, examples)
	}

	if !passB {
		t.Errorf("Expected example: %v, got: %v", expB, examples)
	}
}
