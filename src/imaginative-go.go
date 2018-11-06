package main

import (
    "io"
    "log"
    "net/http"
    "html/template"
    "database/sql"
    "context"
    "net"
    "strings"
    "io/ioutil"
    _ "github.com/go-sql-driver/mysql"
    "github.com/mongodb/mongo-go-driver/mongo"
)

func defaultHome(w http.ResponseWriter, r *http.Request) {
    // because of / match everything route that not defined include / itself, so we have to check it
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    var templates = template.Must(template.ParseFiles("templates/editorial/index_imaginative_go.html"))
    templates.ExecuteTemplate(w, "index_imaginative_go.html", nil)
}

func original(w http.ResponseWriter, r *http.Request) {
    var templates = template.Must(template.ParseFiles("templates/editorial/index.html"))
    templates.ExecuteTemplate(w, "index.html", nil)
}

func seeCode(w http.ResponseWriter, r *http.Request) {
    // get the fn parameter (to define starting function name)
    fns, fnOK := r.URL.Query()["fn"]
    
    if !fnOK || len(fns[0]) < 1 {
        io.WriteString(w, "fn parameter is missing!")
        return
    }

    // define the variables
    start := "func " + fns[0]
    end := "// end of " + fns[0]

    // read the source code
    rawSourceCode, err := ioutil.ReadFile("imaginative-go.go")
    if err != nil {
        log.Fatal(err)
    }

    sourceCode := string(rawSourceCode)

    // start searching for function start
    startIndex := strings.Index(sourceCode, start)
    if startIndex > -1 {
        // start searching for function end -- TODO help us with regex please
        endIndex := strings.Index(sourceCode, end)
        if endIndex > -1 {
            io.WriteString(w, sourceCode[startIndex:endIndex])    
        } else {
            io.WriteString(w, "function " + start + " ending not found!")
            return    
        }
    } else {
        io.WriteString(w, "function " + start + " not found!")
        return
    } 
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello World!")
}
// end of helloWorld

func displayImaginativeGoSource(w http.ResponseWriter, r *http.Request) {
    b, err := ioutil.ReadFile("imaginative-go.go")
    if err != nil {
        log.Fatal(err)
    }

    log.Println(string(b))
}

func about(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{"OK": "PrasBox", "Nama": "Elan"}

    //var templates = template.Must(template.ParseFiles("about.html", "templates/ww.html", "templates/hours.html"))
    var templates = template.Must(template.ParseFiles("about.html"))
    //templates.ExecuteTemplate(w, "indexPage", data)
    templates.ExecuteTemplate(w, "about.html", data)
}

func mysqlSelectMultipleRows(w http.ResponseWriter, r *http.Request) {
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
        InternalId int `json:"internal_id"`
        Name string `json:"name"`
        Slug string `json:"slug"`
        ShortDescription string `json:"short_description"`
    }

    type Data struct {
        Category []Category
    }

    // prepare the database connection
    db, err := sql.Open("mysql", mysqlUsername + ":" + mysqlPassword + "@" + mysqlProtocol + "(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDatabaseName)
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

func mongodbSelectRows(w http.ResponseWriter, r *http.Request) {
    var templates = template.Must(template.ParseFiles("mongodb_select_rows.html"))
    
    // client, _ := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
    // db := client.Database("go_db")
    // collection := db.Collection("content_category")

    client, err := mongo.NewClient("mongodb://@localhost:27017")
    if err != nil { log.Fatal(err) }
    err = client.Connect(context.TODO())
    if err != nil { log.Fatal(err) }

    collection := client.Database("go_db").Collection("content_category")

    collection.InsertOne(context.Background(), map[string]string{"name": "Erlang", "slug": "erlang", "short_description": "Resource for learning Erlang language"})
    if err != nil { log.Fatal(err) }

    //collection.InsertOne(nil, map[string]string{"name": "C", "slug": "c", "short_description": "Resource for learning C language"})

    templates.ExecuteTemplate(w, "mongodb_select_rows.html", nil)
}

func genericPage(w http.ResponseWriter, r *http.Request) {
    var templates = template.Must(template.ParseFiles("templates/editorial/generic_imaginative_go.html"))
    templates.ExecuteTemplate(w, "generic_imaginative_go.html", nil)
}

func elementsPage(w http.ResponseWriter, r *http.Request) {
    var templates = template.Must(template.ParseFiles("templates/editorial/elements_imaginative_go.html"))
    templates.ExecuteTemplate(w, "elements_imaginative_go.html", nil)
}

func getQueryHandler(w http.ResponseWriter, r *http.Request) {
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

func mysqlSelectMultiRowsHandler(w http.ResponseWriter, r *http.Request) {
    type Category struct {
        name string
        slug string
        short_description string
    }

    db, err := sql.Open("mysql", "root:mysqlpass99!@tcp(" + localIpString + ":3306)/go_db")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
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
    mux := http.NewServeMux()
    
    // registers the handler for the given pattern. If a handler already exists for pattern, Handle panics
    mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
    mux.Handle("/assets-phantom/", http.StripPrefix("/assets-phantom/", http.FileServer(http.Dir("templates/phantom/assets"))))
    mux.Handle("/images-phantom/", http.StripPrefix("/images-phantom/", http.FileServer(http.Dir("templates/phantom/images"))))
    mux.Handle("/assets-editorial/", http.StripPrefix("/assets-editorial/", http.FileServer(http.Dir("templates/editorial/assets"))))
    mux.Handle("/images-editorial/", http.StripPrefix("/images-editorial/", http.FileServer(http.Dir("templates/editorial/images"))))
    
    // registers the handler function for the given pattern
    mux.HandleFunc("/", defaultHome)
    mux.HandleFunc("/generic-page", genericPage)
    mux.HandleFunc("/elements-page", elementsPage)
    mux.HandleFunc("/see-code", seeCode)
    mux.HandleFunc("/hello-world", helloWorld)
    mux.HandleFunc("/display-imaginative-go-source", displayImaginativeGoSource)
    mux.HandleFunc("/mysql-select-multiple-rows", mysqlSelectMultipleRows)
    mux.HandleFunc("/mongo-select-rows", mongodbSelectRows)
    mux.HandleFunc("/get-query", getQueryHandler)
    mux.HandleFunc("/mysql-select-multi-rows", mysqlSelectMultiRowsHandler)
    
    log.Fatal(http.ListenAndServe(":9899", mux))
}