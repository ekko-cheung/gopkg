/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package es

import (
	"fmt"
	"github.com/bytedance/sonic"
	"io"
)

type Highlight interface {
	SetField(string, string)
}

type hits[T any] struct {
	Index     string              `json:"_index,omitempty"`
	Type      string              `json:"_type,omitempty"`
	Id        string              `json:"_id,omitempty"`
	Score     float64             `json:"_score"`
	Source    T                   `json:"_source,omitempty"`
	Highlight map[string][]string `json:"highlight,omitempty"`
}

type resp[T any] struct {
	Hits struct {
		Hits []hits[T] `json:"hits"`
	} `json:"hits"`
}

func GetHits[T Highlight](r io.Reader) []T {
	resp := new(resp[T])
	err := sonic.ConfigFastest.NewDecoder(r).Decode(resp)
	if err != nil {
		fmt.Println(err)
	}
	hits := resp.Hits.Hits
	if len(hits) == 0 {
		return []T{}
	}
	result := make([]T, len(hits))
	for i, hit := range hits {
		val := hit.Source
		for k, v := range hit.Highlight {
			if len(v) != 0 {
				val.SetField(k, v[0])
			}
		}
		result[i] = val
	}

	return result
}
