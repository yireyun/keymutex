package keymutex

import (
	"sync"
)

// KeyMutex ...
type KeyMutex struct {
	locks  []sync.Mutex
	count  uint
	handle HashHandle
}

// New and initialize a new keymutex, ready for use.
// It require the number of mutexs(prime number is better) that you need and
// a hash handle(ELFHash, SDBMHash ...).
func New(count uint) *KeyMutex {
	var this KeyMutex
	this.count = count
	this.handle = ELFHash
	this.locks = make([]sync.Mutex, count, count)
	return &this
}

func NewByHash(count uint, handle HashHandle) *KeyMutex {
	var this KeyMutex
	this.count = count
	this.handle = handle
	this.locks = make([]sync.Mutex, count, count)
	return &this
}

// Count the number of mutexs
func (km *KeyMutex) Count() uint {
	return km.count
}

func (km *KeyMutex) LockID(idx uint) {
	km.locks[idx%km.count].Lock()
}
func (km *KeyMutex) UnlockID(idx uint) {
	km.locks[idx%km.count].Unlock()
}

func (km *KeyMutex) LockID64(idx uint64) {
	km.locks[idx%km.count].Lock()
}
func (km *KeyMutex) UnlockID64(idx uint64) {
	km.locks[idx%km.count].Unlock()
}

// Lock the key
func (km *KeyMutex) Lock(key string) {
	km.LockID(km.handle(key))
}

// Unlock the key
func (km *KeyMutex) Unlock(key string) {
	km.UnlockID(km.handle(key))
}
