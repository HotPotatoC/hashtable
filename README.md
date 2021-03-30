# Hashtable in Go

Very simple, idiomatic and thread-safe implementation of Hashtables for Golang using [Seperate chaining](https://en.wikipedia.org/wiki/Hash_table#Separate_chaining).

# Installation

```sh
❯ go get -u github.com/HotPotatoC/hashtable
```

# Usage

Set a value

```go
ht := hashtable.New()

ht.Set("user", "John")
```

Get a value

```go
ht.Get("user") // John
```

Remove a value

```go
ht.Remove("user") // 1
```

Iterate through the table using `Iter()`

```go
for entry := range ht.Iter() {
	fmt.Printf("key: %s | value: %v\n", entry.Key, entry.Value)
}
```

See more examples [here](https://github.com/HotPotatoC/hashtable/tree/master/examples)

# Methods

```go
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
```

# Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

# License

[MIT](https://choosealicense.com/licenses/mit/)

# Support

<a href="https://www.buymeacoffee.com/hotpotato" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
