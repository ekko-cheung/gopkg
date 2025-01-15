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

package stream

type Filter[T any] func(T) bool

// Compare a and b
//
// if a > b return number > 0
//
// if a == b return 0
//
// if a < b return number < 0
type Compare[T any] func(a T, b T) int

type ForEach[T any] func(T)

type Match[T any] func(T) bool

type Stream[T any] interface {
	// Filter returns a stream consisting of the elements of this stream that match the given predicate
	Filter(Filter[T]) Stream[T]
	// ForEach performs an action for each element of this stream
	ForEach(ForEach[T])
	// Reverse returns a stream consisting of the elements of this stream, reverse the order of elements
	Reverse() Stream[T]
	// Count returns the count of elements in this stream
	Count() int
	// ToSlice return slice containing the elements of this stream
	ToSlice() []T
	// Limit returns a stream consisting of the elements of this stream, truncated to be no longer than maxSize in length
	Limit(int) Stream[T]
	// Min returns the minimum element of this stream according to the provided Comparator
	Min(Compare[T]) T
	// Max returns the maximum element of this stream according to the provided Comparator
	Max(Compare[T]) T
	// Skip returns a stream consisting of the remaining elements of this stream after discarding the first n elements of the stream
	// If this stream contains fewer than do nothing
	Skip(int) Stream[T]
	// Sorted Returns a stream consisting of the elements of this stream, sorted according to the given method
	Sorted(Compare[T]) Stream[T]
	// Peek returns a stream consisting of the elements of this stream,
	// additionally performing the provided action on each element as elements are consumed from the resulting stream
	Peek(ForEach[T]) Stream[T]
	// AnyMatch returns whether any elements of this stream match the provided predicate.
	// May not evaluate the predicate on all elements if not necessary for determining the result.
	// If the stream is empty then false is returned and the predicate is not evaluated
	AnyMatch(Match[T]) bool
	// AllMatch returns whether all elements of this stream match the provided predicate.
	// May not evaluate the predicate on all elements if not necessary for determining the result.
	// If the stream is empty then true is returned and the predicate is not evaluated
	AllMatch(Match[T]) bool
	// NoneMatch returns whether no elements of this stream match the provided predicate.
	// May not evaluate the predicate on all elements if not necessary for determining the result.
	// If the stream is empty then true is returned and the predicate is not evaluated
	NoneMatch(Match[T]) bool
	// FindAny returns an element that matches the given method, or any empty value if the stream is empty or not match
	FindAny(Match[T]) T
	// FindFirst returns an element describing the first element of this stream, or an empty value if the stream is empty
	FindFirst() T
	// FindLast returns an element describing the last element of this stream, or ant empty value if the stream is empty
	FindLast() T
}

func Map[T, R any](source []T, f func(T) R) []R {
	arr := make([]R, len(source))
	for i := range source {
		arr[i] = f(source[i])
	}

	return arr
}
