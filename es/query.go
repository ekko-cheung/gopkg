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

func Query(options ...Option) *reqBody {
	r := &reqBody{
		Query: &query{
			Bool: &boolean{},
		},
		Highlight: &highlight{Fields: make(map[string]interface{})},
	}

	for i := range options {
		options[i](r)
	}

	return r
}

type reqBody struct {
	From      int64             `json:"from,omitempty"`
	Size      int64             `json:"size,omitempty"`
	Query     *query            `json:"query,omitempty"`
	Source    []string          `json:"source,omitempty"`
	Sort      []map[string]sort `json:"sort,omitempty"`
	Highlight *highlight        `json:"highlight,omitempty"`
}

type query struct {
	Bool   *boolean                 `json:"bool,omitempty"`
	Filter []map[string]interface{} `json:"filter,omitempty"`
}

type boolean struct {
	Should  []interface{} `json:"should,omitempty"`
	Must    []interface{} `json:"must,omitempty"`
	MustNot []interface{} `json:"must_not,omitempty"`
}

type match struct {
	Match map[string]interface{} `json:"match,omitempty"`
}

type matchPhrase struct {
	MatchPhrase map[string]interface{} `json:"match_phrase,omitempty"`
}

type highlight struct {
	PreTags  []string               `json:"pre_tags,omitempty"`
	PostTags []string               `json:"post_tags,omitempty"`
	Fields   map[string]interface{} `json:"fields,omitempty"`
}

type sorter string

type sort struct {
	Order sorter `json:"order,omitempty"`
}

const (
	ASC  = sorter("asc")
	DESC = sorter("desc")
)
