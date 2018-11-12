package main

import (
    "bytes"
	"fmt"
    "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"strings"
)

func main() {
	someSourceCode := `package main

import ("github.com/alecthomas/chroma")

func main() {
	log.Println("OK")
}	
`
	lexer := lexers.Get("go")
	iterator, _ := lexer.Tokenise(nil, someSourceCode)
	style := styles.Get("github")
	formatter := html.New(html.WithLineNumbers())

	var buff bytes.Buffer
	
	formatter.Format(&buff, style, iterator)

	niceSourceCode := buff.String()
	niceSourceCode = strings.Replace(niceSourceCode, `<pre style="background-color:#fff">`, `<pre style="background-color:#fff"><code>`, -1)
	niceSourceCode = strings.Replace(niceSourceCode, "</pre>", "</code></pre>", -1)

	fmt.Println(niceSourceCode)

	//fmt.Printf("[%s]", buff.String())
}