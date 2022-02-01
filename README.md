# Reveal URL

This is an experiment out of curiosity to reveal the url from URL shorteners in Go.

[Same attempt in Node.js](https://github.com/wellingguzman/reveal-url)

## Example

```go
package main

import (
    "fmt"

    "wellingguzman.com/reveal_url"
)

func main() {
  var urls []string;
  var url string = "https://bit.ly/3ARwV0J"

  reveal_url.Reveal(url, &urls)

  if (urls == nil) {
    fmt.Printf("%s has no redirection\n", url)
  }

  for i, v := range urls {
    fmt.Printf("%v -> %v\n", i, v)
  }
}
```
