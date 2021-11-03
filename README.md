# vue-in-go

## 介绍
在 在 Golang 应用中部署 Vue

## Reference

- [Embedding Vue.js Apps in Go](https://hackandsla.sh/posts/2021-06-18-embed-vuejs-in-go/)
- [Embedding Vue.js Apps in Go - Code Examples](https://gitlab.com/hackandsla.sh/blog-examples/-/tree/master/2021-06-18-vue-with-go)
- [//go:embed 入门](https://taoshu.in/go/how-to-use-go-embed.html)
- [Customizing Go Binaries with Build Tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags)


## Note 

`json.Marshal` : To convert a struct into json string
`json.Unmarshal` : To convert a json string into struct

```go
package main

import (
	"encoding/json"
	"log"
)

type employee struct {
	Name string `json:name`
	Age  int    `json:age`
}

func main() {
	var e1 employee
	eJSON := `{"name":"John","age": 21}`
	err := json.Unmarshal([]byte(eJSON), &e1)
	if err != nil {
		log.Fatalf("Error")
	}
}
```