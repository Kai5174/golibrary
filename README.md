# golibrary
存了一些开发项目时常见的函数, 封装好了以便以后直接拿来用.

## Usage
```go
import "github.com/Kai5174/golibrary"

func() main {
    content, err := golibrary.ReadFile("hello.txt")
    if err != nil {
        log.Panicln(err.Error())
    }
    fmt.Println(string(content))
}
```