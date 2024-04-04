# GoURL

Simple Go package to manage URLs. Inspired by JavaScript's URL class.

## Usage

```go
url, _ := gourl.New("https://example.com/api/v1")
// error handling skipped now

fmt.Println(url.Hostname()) // example.com

url.Hostname("hello.com")
fmt.Println(url.Href()) // https://hello.com/api/v1
```

## Methods

Without a parameter every function is a getter. When a value is passed it will be a setter, but also returns the new value.

|    Hash    |
|:----------:|
|   `Host`   | 
| `Hostname` | 
| `Password` |
| `Pathname` |
|   `Port`   |
| `Protocol` |
|  `Search`  |
| `Username` | 
