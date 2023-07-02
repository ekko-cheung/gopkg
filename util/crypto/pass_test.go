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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncoderPass(t *testing.T) {
	EncoderPass("abcd123456")
}

func TestComparePass(t *testing.T) {
	a := assert.New(t)

	rowPass := "abcd123456"
	hashPass := "$2a$10$zw9iZSvaO45ao2zgm3GOt.ricHfQLs/7xpCJt0blR5clVYnXcyT4W"
	errHashPass := "$2a$10$zw9iZSvaO45ao2zgm3GOt.ricHfQLs/7xpCJt0blR5clVYnXcyT4C"
	a.True(ComparePass(hashPass, rowPass))
	a.False(ComparePass(hashPass, errHashPass))
}
