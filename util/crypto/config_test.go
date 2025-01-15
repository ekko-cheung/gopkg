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

package crypto

import "testing"

type child struct {
	Value string `crypto:"-"`
}

type config struct {
	Value string
	Child child
}

func TestDecodeConfig(t *testing.T) {
	key := "abcdefgh"
	c := &config{
		Value: "a//xM5D6Q3XwT14eXPsa/A==",
		Child: child{
			Value: "a//xM5D6Q3XwT14eXPsa/A==",
		},
	}
	err := DecodeConfig(c, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", c)
}

func TestEncodeConfig(t *testing.T) {
	key := "abcdefgh"
	c := &config{
		Value: "helloworld",
		Child: child{
			Value: "helloworld",
		},
	}
	err := EncodeConfig(c, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", c)
}
