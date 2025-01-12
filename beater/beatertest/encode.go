// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package beatertest

import (
	"bytes"
	"encoding/json"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/esleg/eslegclient"
)

// EncodeEventDocs encodes the events in es using EncodeEventDoc.
//
// The results are suitable for passing to approvaltest.ApproveEventDocs.
func EncodeEventDocs(es ...beat.Event) [][]byte {
	out := make([][]byte, len(es))
	for i, e := range es {
		out[i] = EncodeEventDoc(e)
	}
	return out
}

// EncodeEventDoc encodes e as a JSON document, using the libbeat
// Elasticsearch output.
//
// The result is suitable for passing to approvaltest.ApproveEventDocs.
func EncodeEventDoc(e beat.Event) []byte {
	var buf bytes.Buffer
	encoder := eslegclient.NewJSONEncoder(&buf, false)
	if err := encoder.AddRaw(&e); err != nil {
		panic(err)
	}

	var indented bytes.Buffer
	if err := json.Indent(&indented, buf.Bytes(), "", "  "); err != nil {
		panic(err)
	}
	return indented.Bytes()
}
