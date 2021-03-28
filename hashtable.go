package hashtable

import (
	"sync"

	"github.com/cespare/xxhash/v2"
)

const (
	minLoadFactor = 0.25
	maxLoadFactor = 0.75
	// DefaultSize is the default capacity of the table (16)
	DefaultSize = 16
)

// HashTable data structure
type HashTable interface {
	// Set inserts a new key-value pair item into the hash table
	Set(k string, v interface{})
	// Get returns the value of the given key
	// and a boolean which returns false if the
	// lookup result is nil otherwise true
	Get(k string) (interface{}, bool)
	// Remove deletes an item by the given key
	// and returns the deleted count
	Remove(k string) int
	// Iter returns an iterator for the hashtable
	Iter() <-chan *Entry
	// Exist returns true if an item with the given key exists in the table
	// otherwise returns false
	Exist(k string) bool
	// Len represents the size of the hash table
	Len() int
}

type hashTable struct {
	buckets []*Bucket
	nSize   int
	mtx     sync.RWMutex
}

// Bucket represents the hash table bucket
type Bucket struct {
	Head *Entry
}

// Entry represents an entry inside the bucket
type Entry struct {
	Key   string
	Value interface{}
	Next  *Entry
}

// New creates a new hashtable
func New() HashTable {
	return &hashTable{
		buckets: make([]*Bucket, DefaultSize),
		nSize:   0,
	}
}

// NewWithSize creates a new hashtable with the given size
func NewWithSize(size uint) HashTable {
	return &hashTable{
		buckets: make([]*Bucket, size),
		nSize:   0,
	}
}

// newHashTable creates a new hashtable
func newHashTable(size uint) *hashTable {
	return &hashTable{
		buckets: make([]*Bucket, size),
		nSize:   0,
	}
}

func (ht *hashTable) Set(k string, v interface{}) {
	ht.mtx.Lock()
	defer ht.mtx.Unlock()

	initialSize := ht.nSize
	ht.insert(k, v)
	newSize := ht.nSize - initialSize
	if newSize > 0 {
		ht.verifyLoadFactor()
	}
}

func (ht *hashTable) Get(k string) (interface{}, bool) {
	ht.mtx.RLock()
	defer ht.mtx.RUnlock()
	result := ht.lookup(k)
	if result == nil {
		return nil, false
	}
	return result.Value, true
}

func (ht *hashTable) Remove(k string) int {
	ht.mtx.Lock()
	defer ht.mtx.Unlock()
	initialSize := ht.nSize
	ht.delete(k)
	newSize := initialSize - ht.nSize
	if newSize > 0 {
		ht.verifyLoadFactor()
	}
	return newSize
}

func (ht *hashTable) Iter() <-chan *Entry {
	ch := make(chan *Entry)
	go func() {
		ht.mtx.RLock()
		ht.iterate(ch)
		ht.mtx.RUnlock()
	}()
	return ch
}

func (ht *hashTable) Exist(k string) bool {
	ht.mtx.RLock()
	defer ht.mtx.RUnlock()
	return ht.lookup(k) != nil
}

func (ht *hashTable) Len() int {
	ht.mtx.RLock()
	defer ht.mtx.RUnlock()
	return ht.nSize
}

// insert stores a new entry into an empty bucket
// or chains a new entry into an existing entry with the same key (internal-use only)
func (ht *hashTable) insert(k string, v interface{}) {
	index := ht.hashkey(k, len(ht.buckets))
	entry := ht.newEntry(k, v)
	if ht.buckets[index] == nil {
		ht.buckets[index] = &Bucket{}
		entry.Next = ht.buckets[index].Head
		ht.buckets[index].Head = entry
		ht.nSize++
		return
	}

	for iterator := ht.buckets[index].Head; iterator != nil; iterator = iterator.Next {
		if iterator.Next == nil {
			entry.Next = ht.buckets[index].Head
			ht.buckets[index].Head = entry
			break
		}

		if iterator.Next.Key == k {
			iterator.Next.Value = v
			break
		}
	}

	ht.nSize++
}

// delete removes an entry (internal-use only)
func (ht *hashTable) delete(k string) {
	index := ht.hashkey(k, len(ht.buckets))

	if ht.buckets[index] == nil || ht.buckets[index].Head == nil {
		return
	}

	if ht.buckets[index].Head.Key == k {
		ht.buckets[index].Head = ht.buckets[index].Head.Next
		ht.nSize--
		return
	}

	iterator := ht.buckets[index].Head
	for iterator.Next != nil {
		if iterator.Next.Key == k {
			iterator.Next = iterator.Next.Next
			ht.nSize--
			return
		}
		iterator = iterator.Next
	}
}

// lookup iterates through the table then returns
// an entry with the matching key (internal-use only)
func (ht *hashTable) lookup(k string) *Entry {
	index := ht.hashkey(k, len(ht.buckets))
	if ht.buckets[index] == nil {
		return nil
	}

	iterator := ht.buckets[index].Head
	for iterator != nil {
		if iterator.Key == k {
			return iterator
		}
		iterator = iterator.Next
	}
	return nil
}

// iterate iterates through every available buckets then
// sends an entry into the channel (internal-use only)
func (ht *hashTable) iterate(ch chan<- *Entry) {
	for _, bucket := range ht.buckets {
		if bucket != nil {
			for entry := bucket.Head; entry != nil; entry = entry.Next {
				ch <- entry
			}
		}
	}
	close(ch)
}

// loadFactor calculates and returns the table current load factor (internal-use only)
func (ht *hashTable) loadFactor() float32 {
	return float32(ht.nSize) / float32(len(ht.buckets))
}

// verifyLoadFactor checks wether the current load factor has
// exceeded the maximum threshold (0.75) or minimal threshold (0.25)
// then updates the table capacity (internal-use only)
func (ht *hashTable) verifyLoadFactor() {
	if ht.nSize == 0 {
		return
	}

	lf := ht.loadFactor()
	if lf > maxLoadFactor {
		ht.updateCapacity(uint(ht.nSize * 2))
	} else if lf < minLoadFactor {
		ht.updateCapacity(uint(len(ht.buckets) / 2))
	}
}

// updateCapacity modifies the table size (internal-use only)
func (ht *hashTable) updateCapacity(size uint) {
	newTable := newHashTable(size)
	for _, bucket := range ht.buckets {
		for bucket != nil && bucket.Head != nil {
			newTable.insert(bucket.Head.Key, bucket.Head.Value)
			bucket.Head = bucket.Head.Next
		}
	}
	ht.buckets = newTable.buckets
}

// newEntry creates a new key-value node (internal-use only)
func (ht *hashTable) newEntry(key string, value interface{}) *Entry {
	return &Entry{
		Key:   key,
		Value: value,
		Next:  nil,
	}
}

// hashkey creates the index for the hashtable bucket array (internal-use only)
func (ht *hashTable) hashkey(key string, size int) uint64 {
	return xxhash.Sum64([]byte(key)) % uint64(size)
}
