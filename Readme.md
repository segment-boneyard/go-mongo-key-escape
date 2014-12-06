# go-mongo-key-escape

[Escape MongoDB keys](http://docs.mongodb.org/manual/faq/developers/#faq-dollar-sign-escaping) for characters . and $.

## Example

```go
package main

import "github.com/segmentio/go-mongo-key-escape"

func main() {
    mongo.Escape("event.thing")
    // event\uFF0Ething
    mongo.Unescape("event\uFF0Ething")
    // event.thing

    mongo.Escape("event$thing")
    // event\uFF04thing
    mongo.Unescape("event\uFF04thing")
    // event$thing
}
```

## License

MIT
