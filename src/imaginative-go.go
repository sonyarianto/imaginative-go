package main

// If you add an external package here, make sure it also added on
// docker/golang/Dockerfile so next time if you recreate all containers
// it will be installed
import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/mongo"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
    "https://github.com/alecthomas/chroma"
)

// This is handle / path
func defaultHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Prepare the template for home page
    var templates = template.Must(template.ParseFiles("templates/editorial/index_imaginative_go.html"))

	// Execute template
    templates.ExecuteTemplate(w, "index_imaginative_go.html", nil)
}

// This is handle /see-code path
func seeCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
    end := "// end of " + fns[0]

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
		// Function name found

        // Start searching for function end -- TODO help us with regex please
		endIndex := strings.Index(dataSourceCode, end)
		if endIndex > -1 {
		      // Do nothing when found
        } else {
			// Function end marker not found
            io.WriteString(w, "function "+start+" ending not found!")
			return
		}
	} else {
		// Function start marker not found
        io.WriteString(w, "function "+start+" not found!")
		return
	}

    // Read the source code (src/examples/[fn].go)
    saSourceCode, err := ioutil.ReadFile("examples/" + fns[0] + ".go")
    if err != nil {
        log.Fatal(err)
    }

    dataSaSourceCode := string(saSourceCode)

	// Prepare templates
	var templates = template.Must(template.ParseGlob("templates/editorial/*.html"))
	
    // Execute template
    templates.ExecuteTemplate(w, fns[0] + ".html", map[string]interface{}{"sourceCode": dataSourceCode[startIndex:endIndex], "standAloneSourceCode": dataSaSourceCode, "id": fns[0]})
}

func helloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "hello, world")
}

// end of helloWorld

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

func genericPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var templates = template.Must(template.ParseFiles("templates/editorial/generic_imaginative_go.html"))
	templates.ExecuteTemplate(w, "generic_imaginative_go.html", nil)
}

func elementsPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var templates = template.Must(template.ParseFiles("templates/editorial/elements_imaginative_go.html"))
	templates.ExecuteTemplate(w, "elements_imaginative_go.html", nil)
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

var localIpString string = GetOutboundIP().String()
var mysqlHost string = "mysql"
var mysqlUsername string = "root"
var mysqlPassword string = "mysqlpassword"
var mysqlProtocol string = "tcp"
var mysqlDatabaseName string = "go_db"
var mysqlPort string = "3306"

func main() {
	// allocates and returns a new ServeMux
	mux := httprouter.New()

	mux.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	mux.ServeFiles("/assets-phantom/*filepath", http.Dir("templates/phantom/assets/"))
	mux.ServeFiles("/images-phantom/*filepath", http.Dir("templates/phantom/images/"))
	mux.ServeFiles("/assets-editorial/*filepath", http.Dir("templates/editorial/assets/"))
	mux.ServeFiles("/images-editorial/*filepath", http.Dir("templates/editorial/images/"))

	// registers the handler function for the given pattern
	mux.GET("/", defaultHome)
	mux.GET("/generic-page", genericPage)
	mux.GET("/elements-page", elementsPage)
	mux.GET("/see-code", seeCode)
	mux.GET("/hello-world", helloWorld)
	mux.GET("/display-imaginative-go-source", displayImaginativeGoSource)
	mux.GET("/mysql-select-multiple-rows", mysqlSelectMultipleRows)
	mux.GET("/mongo-select-rows", mongodbSelectRows)
	mux.GET("/get-query", getQueryHandler)
	mux.GET("/mysql-select-multi-rows", mysqlSelectMultiRowsHandler)

	log.Fatal(http.ListenAndServe(":9899", mux))
}
