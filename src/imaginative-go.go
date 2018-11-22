package main

// If you add an external package here, make sure it also added on
// docker/golang/Dockerfile so next time if you recreate all containers
// it will be installed.
import (
	"bytes"
	"context"
	"database/sql"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

// Prepare struct for syntax highlighter.
type ChromaRenderer struct {
	html  *blackfriday.HTMLRenderer
	theme string
}

// RenderNode is called with the node being traversed.
func (r *ChromaRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	switch node.Type {
	// We only care about the pre tag.
	case blackfriday.CodeBlock:
		// Set up a lexer.
		var lexer chroma.Lexer

		// Read the language from the annotation.
		lang := string(node.CodeBlockData.Info)
		if lang != "" {
			lexer = lexers.Get(lang)
		} else {
			// Analyze when no language annotation is given.
			lexer = lexers.Analyse(string(node.Literal))
		}

		// If no annotation was found and couldn't be analyzed, fallback.
		if lexer == nil {
			lexer = lexers.Fallback
		}

		// Set a syntax highlighting theme
		style := styles.Get(r.theme)
		if style == nil {
			style = styles.Fallback
		}

		// Apply highlighting with Chroma.
		iterator, err := lexer.Tokenise(nil, string(node.Literal))
		if err != nil {
			panic(err)
		}

		// An HTML formatter for the tokenized results.
		formatter := html.New()

		// Write out the highlighted code to the io.Writer.
		err = formatter.Format(w, style, iterator)
		if err != nil {
			panic(err)
		}

		// Move on to the next node.
		return blackfriday.GoToNext
	}

	// Didn't match the CodeBlock type, render it as is.
	return r.html.RenderNode(w, node, entering)
}

func (r *ChromaRenderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}
func (r *ChromaRenderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {}

func NewChromaRenderer(theme string) *ChromaRenderer {
	return &ChromaRenderer{
		html:  blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{}),
		theme: theme,
	}
}

// Handle / path.
func HomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var templates = template.Must(template.New("").ParseFiles("templates/_base.html", "templates/index.html"))

	// Execute template.
	templates.ExecuteTemplate(w, "_base.html", nil)
}

// Handle /content path.
func ContentHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content, err := ioutil.ReadFile("data/sample_hello_world.md")
	if err != nil {
		log.Fatal(err)
	}

	cr := NewChromaRenderer("perldoc")
	output := blackfriday.Run(content, blackfriday.WithRenderer(cr))
	output2 := string(output)

	var templates = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/_base.html", "templates/read_sample.html"))

	// Execute templates
	templates.ExecuteTemplate(w, "_base.html", output2)
}

// Handle /see-code path
func SeeCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get the fn parameter (to define starting function name)
	fns, fnOK := r.URL.Query()["fn"]

	// Check the fn parameter
	if !fnOK || len(fns[0]) < 1 {
		io.WriteString(w, "fn parameter is missing!")
		return
	}

	// Start marker
	start := "func " + fns[0]
	// End marker
	end := "// End of " + fns[0]

	// Read the source code (imaginative-go.go)
	sourceCode, err := ioutil.ReadFile("imaginative-go.go")
	if err != nil {
		log.Fatal(err)
	}

	dataSourceCode := string(sourceCode)

	// Start searching for function start  -- TODO help us with regex please
	startIndex := strings.Index(dataSourceCode, start)
	endIndex := strings.Index(dataSourceCode, end)
	if startIndex > -1 {
		// Function name start marker found

		// Start searching for function end -- TODO help us with regex please
		endIndex = strings.Index(dataSourceCode, end)
		if endIndex > -1 {
			// Function name (one block) found

			// We got the source code string on imaginative-go.go
			dataSourceCode = dataSourceCode[startIndex:endIndex]
		} else {
			// Function end marker not found
			io.WriteString(w, "function "+start+" ending not found!")
			return
		}
	} else {
		// Function start marker not found
		//io.WriteString(w, "function "+start+" not found!")
		//return
		dataSourceCode = ""
		endIndex = 0
	}

	// Start doing syntax highlight on it
	lexer := lexers.Get("go")
	iterator, _ := lexer.Tokenise(nil, dataSourceCode)
	style := styles.Get("github")

	// Do this if you want line number, formatter := html.New(html.WithLineNumbers())
	formatter := html.New()

	var buffDataSourceCode bytes.Buffer

	formatter.Format(&buffDataSourceCode, style, iterator)

	niceSourceCode := buffDataSourceCode.String()
	niceSourceCode = strings.Replace(niceSourceCode, `<pre style="background-color:#fff">`, `<pre style="background-color:#fff;width:100%;"><code>`, -1)
	niceSourceCode = strings.Replace(niceSourceCode, "</pre>", "</code></pre>", -1)

	// Read the source code (src/examples/[fn].go) (stand alone code version)
	saSourceCode, err := ioutil.ReadFile("examples/" + fns[0] + ".go")
	if err != nil {
		log.Fatal(err)
	}

	dataSaSourceCode := string(saSourceCode)

	// Start doing syntax highlight on it
	lexer = lexers.Get("go")
	iterator, _ = lexer.Tokenise(nil, dataSaSourceCode)
	style = styles.Get("github")

	// Do this if you want line number, formatter = html.New(html.WithLineNumbers())
	formatter = html.New()

	var buffDataSaSourceCode bytes.Buffer

	formatter.Format(&buffDataSaSourceCode, style, iterator)

	niceSaSourceCode := buffDataSaSourceCode.String()
	niceSaSourceCode = strings.Replace(niceSaSourceCode, `<pre style="background-color:#fff">`, `<pre style="background-color:#fff;width:100%;"><code>`, -1)
	niceSaSourceCode = strings.Replace(niceSaSourceCode, "</pre>", "</code></pre>", -1)

	// Execute template
	//templates.ExecuteTemplate(w, "sample_imaginative_go.html", map[string]interface{}{"sourceCode": niceSourceCode, "standAloneSourceCode": niceSaSourceCode, "id": fns[0]})
}

// Handle /hello-world path
func SampleHelloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "hello, world")
} // End of SampleHelloWorld

// Handle /hello-world-2 path
func SampleHelloWorld2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "<h1>hello, world<h1>")
} // End of SampleHelloWorld2

func displayImaginativeGoSource(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b, err := ioutil.ReadFile("imaginative-go.go")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]interface{}{"OK": "PrasBox", "Nama": "Elan"}

	//var templates = template.Must(template.ParseFiles("about.html", "templates/ww.html", "templates/hours.html"))
	var templates = template.Must(template.ParseFiles("about.html"))
	//templates.ExecuteTemplate(w, "indexPage", data)
	templates.ExecuteTemplate(w, "about.html", data)
}

func mysqlSelectMultipleRows(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// prepare the function for template
	funcMap := template.FuncMap{
		// the name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}

	// prepare the template
	var templates = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/phantom/mysql_select_multiple_rows.html"))

	// prepare the structure
	type Category struct {
		InternalId       int    `json:"internal_id"`
		Name             string `json:"name"`
		Slug             string `json:"slug"`
		ShortDescription string `json:"short_description"`
	}

	type Data struct {
		Category []Category
	}

	// prepare the database connection
	db, err := sql.Open("mysql", mysqlUsername+":"+mysqlPassword+"@"+mysqlProtocol+"("+mysqlHost+":"+mysqlPort+")/"+mysqlDatabaseName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// prepare the SQL (using Query)
	rows, err := db.Query("SELECT id, `name`, slug, short_description FROM content_category ORDER BY `name` ASC")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	// prepare the data
	rowsData := make([]Category, 0)

	// start loop to selected table
	for rows.Next() {
		// prepare the variable
		var category Category

		// scan each row
		err := rows.Scan(&category.InternalId, &category.Name, &category.Slug, &category.ShortDescription)
		if err != nil {
			log.Fatal(err)
		}

		// save to variable
		rowsData = append(rowsData, category)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	templates.ExecuteTemplate(w, "mysql_select_multiple_rows.html", Data{Category: rowsData})
}

func mongodbSelectRows(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var templates = template.Must(template.ParseFiles("mongodb_select_rows.html"))

	// client, _ := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	// db := client.Database("go_db")
	// collection := db.Collection("content_category")

	client, err := mongo.NewClient("mongodb://@localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("go_db").Collection("content_category")

	collection.InsertOne(context.Background(), map[string]string{"name": "Erlang", "slug": "erlang", "short_description": "Resource for learning Erlang language"})
	if err != nil {
		log.Fatal(err)
	}

	//collection.InsertOne(nil, map[string]string{"name": "C", "slug": "c", "short_description": "Resource for learning C language"})

	templates.ExecuteTemplate(w, "mongodb_select_rows.html", nil)
}

func getQueryHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	// log.Println("Url Param 'key' is: " + string(key))
	io.WriteString(w, key)
}

func mysqlSelectMultiRowsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type Category struct {
		name              string
		slug              string
		short_description string
	}

	db, err := sql.Open("mysql", "root:mysqlpass99!@tcp("+localIpString+":3306)/go_db")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Query statement for reading data
	rows, err := db.Query("SELECT slug, name, short_description FROM content_category ORDER BY name ASC")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var category Category

		err := rows.Scan(&category.slug, &category.name, &category.short_description)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Some global vars
var localIpString string = GetOutboundIP().String()
var mysqlHost string = "mysql"
var mysqlUsername string = "root"
var mysqlPassword string = "mysqlpassword"
var mysqlProtocol string = "tcp"
var mysqlDatabaseName string = "go_db"
var mysqlPort string = "3306"

// Define function for template
var funcMap = template.FuncMap{
	"toHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
}

// Prepare all templates
//var templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))

func main() {
	mux := httprouter.New()

	// Serve static files
	mux.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	//mux.ServeFiles("/assets-phantom/*filepath", http.Dir("templates/phantom/assets/"))
	//mux.ServeFiles("/images-phantom/*filepath", http.Dir("templates/phantom/images/"))
	//mux.ServeFiles("/assets-editorial/*filepath", http.Dir("templates/editorial/assets/"))
	//mux.ServeFiles("/images-editorial/*filepath", http.Dir("templates/editorial/images/"))

	// Registers the handler function for the given pattern
	mux.GET("/", HomeHandler)
	mux.GET("/content", ContentHandler)
	mux.GET("/see-code/:slug", SeeCode)
	mux.GET("/result/hello-world", SampleHelloWorld)
	mux.GET("/result/hello-world-2", SampleHelloWorld2)
	mux.GET("/display-imaginative-go-source", displayImaginativeGoSource)
	mux.GET("/mysql-select-multiple-rows", mysqlSelectMultipleRows)
	mux.GET("/mongo-select-rows", mongodbSelectRows)
	mux.GET("/get-query", getQueryHandler)
	mux.GET("/mysql-select-multi-rows", mysqlSelectMultiRowsHandler)

	// Start listen and serve
	log.Fatal(http.ListenAndServe(":9899", mux))
}
