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
	"errors"
	"fmt"
	"sync"
	"time"
)

var snowFlake, _ = NewSnowFlake(0, 0)

func GenId() int64 {
	id, err := snowFlake.NextId()
	if err != nil {
		return time.Now().UnixNano()
	}

	return id
}

type SnowFlake struct {
	mu sync.Mutex
	twepoch int64
	workerIdBits     int64
	datacenterIdBits int64
	sequenceBits     int64
	maxWorkerId     int64
	maxDatacenterId int64
	maxSequence     int64
	workerIdShift     int64
	datacenterIdShift int64
	timestampShift    int64
	datacenterId int64
	workerId int64
	sequence int64
	lastTimestamp int64
}

func (s *SnowFlake) timeGen() int64 {
	return time.Now().UnixMilli()
}

func (s *SnowFlake) tilNextMills() int64 {
	timeStampMill := s.timeGen()
	for timeStampMill <= s.lastTimestamp {
		timeStampMill = s.timeGen()
	}
	return timeStampMill
}
func (s *SnowFlake) NextId() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	nowTimestamp := s.timeGen()
	if nowTimestamp < s.lastTimestamp {
		return -1, fmt.Errorf("clock moved backwards, Refusing to generate id for %d milliseconds", s.lastTimestamp-nowTimestamp)
	}
	if nowTimestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & s.maxSequence
		if s.sequence == 0 {
			nowTimestamp = s.tilNextMills()
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = nowTimestamp
	return (nowTimestamp-s.twepoch)<<s.timestampShift |
			s.datacenterId<<s.datacenterIdShift |
			s.workerId<<s.workerIdShift |
			s.sequence,
		nil
}

func NewSnowFlake(workerId int64, datacenterId int64) (*SnowFlake, error) {
	mySnow := new(SnowFlake)
	mySnow.twepoch = time.Now().Unix()
	if workerId < 0 || datacenterId < 0 {
		return nil, errors.New("workerId or datacenterId must not lower than 0 ")
	}

	mySnow.workerIdBits = 5
	mySnow.datacenterIdBits = 5
	mySnow.sequenceBits = 12

	mySnow.maxWorkerId = -1 ^ (-1 << mySnow.workerIdBits)
	mySnow.maxDatacenterId = -1 ^ (-1 << mySnow.datacenterIdBits)
	mySnow.maxSequence = -1 ^ (-1 << mySnow.sequenceBits)

	if workerId >= mySnow.maxWorkerId || datacenterId >= mySnow.maxDatacenterId {
		return nil, errors.New("workerId or datacenterId must not higher than max value ")
	}
	mySnow.workerIdShift = mySnow.sequenceBits
	mySnow.datacenterIdShift = mySnow.sequenceBits + mySnow.workerIdBits
	mySnow.timestampShift = mySnow.sequenceBits + mySnow.workerIdBits + mySnow.datacenterIdBits

	mySnow.lastTimestamp = -1
	mySnow.workerId = workerId
	mySnow.datacenterId = datacenterId

	return mySnow, nil
}
