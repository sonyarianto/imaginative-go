# Template

## Scenario

It will show how to use template in Go for web application.

### Sample 1

Display a HTML template for your web.

Install the required package first.

```
go get -u github.com/julienschmidt/httprouter
```

Prepare the HTML template.

Save as `index.html`

<script src="https://gist.github.com/sonyarianto/e29805ac7ea95cc28c44e07c75dc226a.js"></script>

Prepare the Go code.

Save as `template-html.go`

<script src="https://gist.github.com/sonyarianto/a4683d8e97a7d6bef7ccd0d7116b79db.js"></script>

Run with `go run template-html.go`