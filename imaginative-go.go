package main

// If you add an external package here, make sure it also added on
// docker/golang/Dockerfile so next time if you recreate all containers
// it will be installed
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
	"net/http"
	"os"
)

// Tag is for tag on Content struct.
type Tag struct {
	Tag string `bson:"tag" json:"tag"`
}

// Content is struct for content on this web project.
type Content struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Title            string             `bson:"title" json:"title"`
	Slug             string             `bson:"slug" json:"slug"`
	ShortDescription string             `bson:"short_description" json:"short_description"`
	ContentFile      string             `bson:"content_file" json:"content_file"`
	Tags             []Tag              `bson:"tags" json:"tags"`
}

// ChromaRenderer is struct for syntax highlighter.
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

// RenderHeader is used for render header.
func (r *ChromaRenderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}

// RenderFooter is used for render footer.
func (r *ChromaRenderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {}

// NewChromaRenderer is used for renderer.
func NewChromaRenderer(theme string) *ChromaRenderer {
	return &ChromaRenderer{
		html:  blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{}),
		theme: theme,
	}
}

// MongoDBConnect is used to connect to MongoDB.
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

// GetAllContent is to get all active content for the website.
func GetAllContent() []Content {
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

	return content
}

// Home is handler for / path.
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	content := GetAllContent()

	// Prepare data structure for data passed to template.
	type TemplateData struct {
		Content []Content
		Env     string
	}

	templateData := TemplateData{Content: content, Env: os.Getenv("IGO_ENV")}

	// Parse templates.
	var templates = template.Must(template.New("").ParseFiles("web/templates/_base.html", "web/templates/index.html"))

	// Execute template.
	templates.ExecuteTemplate(w, "_base.html", templateData)
}

// ReadContent is handler for reading content on this web.
func ReadContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the parameter.
	slug := ps.ByName("slug")

	// Do the connection and select database.
	db := MongoDBConnect()

	result := Content{}

	// Do the query to a collection on database.
	if err := db.Collection("sample_content").FindOne(nil, bson.D{{"slug", slug}}).Decode(&result); err != nil {
		http.NotFound(w, r)
		return
	}

	// Get content file (in markdown format).
	fileContent, err := ioutil.ReadFile("web/content/samples/" + result.ContentFile)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare renderer.
	cr := NewChromaRenderer("paraiso-light")
	content := string(blackfriday.Run(fileContent, blackfriday.WithRenderer(cr)))

	// Prepare data structure for data passed to template.
	type TemplateData struct {
		Content template.HTML
		Slug    string
		Env     string
	}

	templateData := TemplateData{Content: template.HTML(content), Slug: slug, Env: os.Getenv("IGO_ENV")}

	// Parse templates.
	var templates = template.Must(template.New("").ParseFiles("web/templates/_base.html", "web/templates/read-content.html"))

	// Execute template.
	templates.ExecuteTemplate(w, "_base.html", templateData)
}

func main() {
	mux := httprouter.New()

	// Serve static files.
	mux.ServeFiles("/assets/*filepath", http.Dir("web/assets/"))

	// Registers the handler function for the given pattern.
	mux.GET("/", Home)
	mux.GET("/content/:slug", ReadContent)

	// Start listen and serve.
	log.Fatal(http.ListenAndServe(":9899", mux))
}
