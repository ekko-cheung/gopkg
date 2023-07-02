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

package util

import (
	"context"
	"log"
	"testing"
	"time"
)

type subject1 struct{}

func (s subject1) Update(ctx context.Context, value string) {
	log.Println("subject1: ", value)
}

type subject2 struct{}

func (s subject2) Update(ctx context.Context, value string) {
	log.Println("subject2:", value)
}

func TestObserver(t *testing.T) {
	s := NewSimpleSubject[string]()
	s.Register(subject1{})
	s.Register(subject2{})

	t.Run("test_Notify", func(t *testing.T) {
		s.Notify(context.Background(), "notify content")
	})

	t.Run("test_AsyncNotify", func(t *testing.T) {
		s.AsyncNotify(context.Background(), "async notify content")
		time.Sleep(time.Second)
	})
}
