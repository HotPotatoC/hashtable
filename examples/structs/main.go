package main

import (
	"fmt"
	"time"

	"github.com/HotPotatoC/hashtable"
)

type post struct {
	Title    string
	Content  string
	PostedAt time.Time
}

func main() {
	ht := hashtable.New()

	ht.Set("post::1", &post{
		Title:    "Post Title",
		Content:  "Post Content",
		PostedAt: time.Now(),
	})

	result, _ := ht.Get("post::1")

	post := result.(*post)

	fmt.Println(post.Title)
	fmt.Println(post.Content)
	fmt.Println(post.PostedAt)
}
