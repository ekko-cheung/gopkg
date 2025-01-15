//go:build go1.18

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

package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitMap := NewBitMap()
	assertions := assert.New(t)
	assertions.EqualValues(bitMap.SetBit("a", 0, 1), 0)
	assertions.EqualValues(bitMap.SetBit("a", 1, 1), 0)
	assertions.EqualValues(bitMap.GetBit("a", 0), 1)
	assertions.EqualValues(bitMap.GetBit("a", 1), 1)
	assertions.EqualValues(bitMap.BitCount("a"), 2)
}
