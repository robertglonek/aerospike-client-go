// Copyright 2013-2014 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package atomic

import (
	"sync"
)

// AtomicBool implements a synchronized boolean value
type AtomicBool struct {
	val   bool
	mutex sync.RWMutex
}

// NewAtomicBool generates a new AtomicBoolean instance.
func NewAtomicBool(value bool) *AtomicBool {
	return &AtomicBool{
		val: value,
	}
}

// Get atomically retrieves the boolean value.
func (ab *AtomicBool) Get() bool {
	ab.mutex.RLock()
	res := ab.val
	ab.mutex.RUnlock()
	return res
}

// Set atomically sets the boolean value.
func (ab *AtomicBool) Set(newVal bool) {
	ab.mutex.Lock()
	ab.val = newVal
	ab.mutex.Unlock()
}

// GetAndToggle atomically retrieves the current boolean value first, and then toggles it.
func (ab *AtomicBool) GetAndToggle() bool {
	ab.mutex.Lock()
	val := ab.val
	ab.val = !ab.val
	ab.mutex.Unlock()
	return val
}

// ToggleAndGet atomically toggles the boolean value first, and then retrieves it.
func (ab *AtomicBool) ToggleAndGet() bool {
	ab.mutex.Lock()
	ab.val = !ab.val
	val := ab.val
	ab.mutex.Unlock()
	return val
}

// CompareAndSet atomically sets the boolean value to updated value, if the current value is as expected.
func (ab *AtomicBool) CompareAndSet(expect bool, update bool) bool {
	res := false
	ab.mutex.Lock()
	if ab.val == expect {
		ab.val = update
		res = true
	}
	ab.mutex.Unlock()
	return res
}
