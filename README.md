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

See more examples [here](https://github.com/HotPotatoC/hashtable/examples)

# Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

# License

[MIT](https://choosealicense.com/licenses/mit/)

# Support

<a href="https://www.buymeacoffee.com/hotpotato" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
