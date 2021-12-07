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

package utility

import (
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
)

func MillisAsMicros(ms float64) common.MapStr {
	m := common.MapStr{}
	m["us"] = int(ms * 1000)
	return m
}

func TimeAsMicros(t time.Time) common.MapStr {
	if t.IsZero() {
		return nil
	}

	m := common.MapStr{}
	m["us"] = t.UnixNano() / 1000
	return m
}