# Hello World

## Scenario

It will show how to do "hello, world" in Go. There is web and non-web sample for this.

### Sample 1 (on console)

Display 'hello, world' on console.

Save as `hello-world.go`

<script src="https://gist.github.com/sonyarianto/84e14db12bcdef21b6fdcda500f808e8.js"></script>

Run with `go run hello-world.go`

### Sample 2 (on web using net/http)

Display 'hello, world' on web browser. Using standard mux.

Save as `hello-world-on-web.go`

<script src="https://gist.github.com/sonyarianto/490999458905896135d16c1ddad92a7b.js"></script>

Run with `go run hello-world-on-web.go` and access on your browser at `http://localhost:3000`

### Sample 3 (on web using net/http)

Display 'hello, world' on web browser. Using standard mux. Render with Content Type `text/html` so it aware of HTML tags.

Save as `hello-world-on-web-2.go`

<script src="https://gist.github.com/sonyarianto/92a3d645af2370762cd749b2093da7d3.js"></script>

Run with `go run hello-world-on-web-2.go` and access on your browser at `http://localhost:3000`