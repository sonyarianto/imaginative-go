package main

// If you add an external package here, make sure it also added on
// docker/golang/Dockerfile so next time if you recreate all containers
// it will be installed.
import (
	"context"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type Tag struct {
	Tag string `bson:"tag" json:"tag"`
}

type Content struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Title            string             `bson:"title" json:"title"`
	Slug             string             `bson:"slug" json:"slug"`
	ShortDescription string             `bson:"short_description" json:"short_description"`
	ContentFile      string             `bson:"content_file" json:"content_file"`
	Tags             []Tag              `bson:"tags" json:"tags"`
}

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

func MongoDBConnect() *mongo.Database {
	// Prepare database.
	client, err := mongo.NewClient(os.Getenv("IGO_MONGODB_URI"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database.
	err = client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Select a database.
	db := client.Database(os.Getenv("IGO_MONGODB_DATABASE"))

	return db
}

// Handle / path.
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Do the connection and select database.
	db := MongoDBConnect()

	// Do the query to a collection on database.
	c, err := db.Collection("sample_content").Find(nil, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close(nil)

	var content []Content

	// Start looping on the query result.
	for c.Next(context.TODO()) {
		eachContent := Content{}

		err := c.Decode(&eachContent)
		if err != nil {
			log.Fatal(err)
		}

		content = append(content, eachContent)
	}

	// Prepare data structure for data passed to template.
	type TemplateData struct {
		Content []Content
	}

	templateData := TemplateData{Content: content}

	// Parse templates.
	var templates = template.Must(template.New("").ParseFiles("templates/_base.html", "templates/index.html"))

	// Execute template.
	templates.ExecuteTemplate(w, "_base.html", templateData)
}

func ReadContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the parameter.
	slug := ps.ByName("slug")

	// Do the connection and select database.
	db := MongoDBConnect()

	result := Content{}

	// Do the query to a collection on database.
	db.Collection("sample_content").FindOne(nil, bson.D{{"slug", slug}}).Decode(&result)

	// Get content file (in markdown format).
	fileContent, err := ioutil.ReadFile("data/content/" + result.ContentFile)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare renderer.
	cr := NewChromaRenderer("perldoc")
	content := string(blackfriday.Run(fileContent, blackfriday.WithRenderer(cr)))

	// Prepare data structure for data passed to template.
	type TemplateData struct {
		Content template.HTML
	}

	templateData := TemplateData{Content: template.HTML(content)}

	// Parse templates.
	var templates = template.Must(template.New("").ParseFiles("templates/_base.html", "templates/read-content.html"))

	// Execute template.
	templates.ExecuteTemplate(w, "_base.html", templateData)
}

// func displayImaginativeGoSource(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	b, err := ioutil.ReadFile("imaginative-go.go")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println(string(b))
// }

// func mysqlSelectMultipleRows(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	// prepare the function for template
// 	funcMap := template.FuncMap{
// 		// the name "inc" is what the function will be called in the template text.
// 		"inc": func(i int) int {
// 			return i + 1
// 		},
// 	}

// 	// prepare the template
// 	var templates = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/phantom/mysql_select_multiple_rows.html"))

// 	// prepare the structure
// 	type Category struct {
// 		InternalId       int    `json:"internal_id"`
// 		Name             string `json:"name"`
// 		Slug             string `json:"slug"`
// 		ShortDescription string `json:"short_description"`
// 	}

// 	type Data struct {
// 		Category []Category
// 	}

// 	// prepare the database connection
// 	db, err := sql.Open("mysql", mysqlUsername+":"+mysqlPassword+"@"+mysqlProtocol+"("+mysqlHost+":"+mysqlPort+")/"+mysqlDatabaseName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	// prepare the SQL (using Query)
// 	rows, err := db.Query("SELECT id, `name`, slug, short_description FROM content_category ORDER BY `name` ASC")
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}
// 	defer rows.Close()

// 	// prepare the data
// 	rowsData := make([]Category, 0)

// 	// start loop to selected table
// 	for rows.Next() {
// 		// prepare the variable
// 		var category Category

// 		// scan each row
// 		err := rows.Scan(&category.InternalId, &category.Name, &category.Slug, &category.ShortDescription)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// save to variable
// 		rowsData = append(rowsData, category)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	templates.ExecuteTemplate(w, "mysql_select_multiple_rows.html", Data{Category: rowsData})
// }

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

func main() {
	mux := httprouter.New()

	// Serve static files
	mux.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	// Registers the handler function for the given pattern
	mux.GET("/", Home)
	mux.GET("/content/:slug", ReadContent)

	// Start listen and serve
	log.Fatal(http.ListenAndServe(":9899", mux))
}
