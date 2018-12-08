# Template

## Scenario

It will show how to use template in Go for web application.

Install the required package first.

```
go get -u github.com/julienschmidt/httprouter
```

### Sample 1 (simple one)

Display HTML template for your web.

Prepare the HTML template.

Save as `index.html`

<script src="https://gist.github.com/sonyarianto/e29805ac7ea95cc28c44e07c75dc226a.js"></script>

Prepare the Go code.

Save as `template-html.go`

<script src="https://gist.github.com/sonyarianto/a4683d8e97a7d6bef7ccd0d7116b79db.js"></script>

Run with `go run template-html.go` and check on web browser `http://localhost:3000`.

### Sample 2 (passing simple variable)

Display HTML template that contains simple data from handler.

Prepare the HTML template.

Save as `index2.html`

<script src="https://gist.github.com/sonyarianto/529d9a39dfa27ba740057459d573ffae.js"></script>

Prepare the Go code.

Save as `template-html-2.go`

<script src="https://gist.github.com/sonyarianto/6667a480a776b2d0e31d18560521f609.js"></script>

Run with `go run template-html-2.go` and check on web browser `http://localhost:3000`.

### Sample 3 (passing multiple variable)

Display HTML template that contains multiple variable from handler.

Prepare the HTML template.

Save as `index3.html`

<script src="https://gist.github.com/sonyarianto/b4948c947aa2ed95c623be7a9245caa5.js"></script>

Prepare the Go code. In this example will use map as data structure.

Save as `template-html-3.go`

<script src="https://gist.github.com/sonyarianto/0b78b340106595407694d3894c2e9bc3.js"></script>

Run with `go run template-html-3.go` and check on web browser `http://localhost:3000`.

Another example is using struct like below.

<script src="https://gist.github.com/sonyarianto/a91eb9f9b7a8784fd071a748823b918d.js"></script>

Run with `go run template-html-4.go` and check on web browser `http://localhost:3000`.

### Sample 4 (loop in template)

Display HTML template that contains array of data and loop it on template.

Prepare the HTML template.

Save as `index4.html`

<script src="https://gist.github.com/sonyarianto/82c4c59ea4555a21b19e87a517ffcd7c.js"></script>

Prepare the Go code.

Save as `template-html-5.go`

<script src="https://gist.github.com/sonyarianto/0c07470ddbdf816eac627ea34128be75.js"></script>

Run with `go run template-html-5.go` and check on web browser `http://localhost:3000`.