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
	"encoding/json"
	"github.com/bytedance/sonic"
	"testing"
)

func TestQuery(t *testing.T) {
	body := Query(
		Page(10, 20),
		TermFilter("name", "mike"),
		Source("name", "age"),
		ShouldMatch("from", "NewYork"),
		MustMatch("grade", 1),
		MustNotMatch("sex", 0),
		RangeFilter("score", map[string]interface{}{"gt": 80, "lt": 50}),
		HighlightField("name"),
	)

	bytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestMap(t *testing.T) {
	body := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should":   []map[string]interface{}{{"match": map[string]interface{}{"from": "NewYork"}}},
				"must":     []map[string]interface{}{{"match": map[string]interface{}{"grade": 1}}},
				"must_not": []map[string]interface{}{{"match": map[string]interface{}{"sex": 0}}},
			},
			"filter": []map[string]interface{}{
				{"term": map[string]interface{}{"name": "mike"}},
				{"range": map[string]interface{}{"score": map[string]interface{}{"gt": 80, "lt": 50}}},
			},
		},
		"source": []string{"name", "age"},
		"from":   180,
		"size":   20,
		"sort":   map[string]interface{}{},
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func BenchmarkQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		body := Query(
			Page(10, 20),
			TermFilter("name", "mike"),
			Source("name", "age"),
			ShouldMatch("from", "NewYork"),
			MustMatch("grade", 1),
			MustNotMatch("sex", 0),
			RangeFilter("score", map[string]interface{}{"gt": 80, "lt": 50}),
		)

		_, _ = sonic.Marshal(body)
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		body := map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should":   []map[string]interface{}{{"match": map[string]interface{}{"from": "NewYork"}}},
					"must":     []map[string]interface{}{{"match": map[string]interface{}{"grade": 1}}},
					"must_not": []map[string]interface{}{{"match": map[string]interface{}{"sex": 0}}},
				},
				"filter": []map[string]interface{}{
					{"term": map[string]interface{}{"name": "mike"}},
					{"range": map[string]interface{}{"score": map[string]interface{}{"gt": 80, "lt": 50}}},
				},
			},
			"source": []string{"name", "age"},
			"from":   180,
			"size":   20,
			"sort":   map[string]interface{}{},
		}

		_, _ = sonic.Marshal(body)
	}
}
