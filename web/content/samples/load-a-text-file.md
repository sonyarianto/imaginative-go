# Load a text file

## Scenario

It will show how to load a text file and display it on console.

Suppose a text file `my-text-file.txt`.

```
Dear Go designers,
I love using Go because it's easy to use and easy to deploy.

I hope Go will become very popular language in the future.

Best regards,
Sony AK
sony@sony-ak.com
```

Here is the code to load and display a text file content.

Save as `load-a-text-file.go`
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    fileContent, err := ioutil.ReadFile("my-text-file.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(fileContent))
}

```

Run with `go run load-a-text-file.go`