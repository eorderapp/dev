package main
//docker run --name sqlserver -d mssql-articles
//docker inspect

//docker run -it --link sqlserver --rm mssql-articles  /opt/mssql-tools/bin/sqlcmd -S 172.17.0.2 -U sa -P "puieMonta140!"


import (
	"net/http"
	"log"
	"context"
	"github.com/gorilla/mux"
	"fmt"
	"database/sql"
)
import (
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
)

var db *sql.DB = nil


func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	article := getArticleById(vars["category"])
	fmt.Println("Article get", article)
	fmt.Fprintf(w, "Category: %v\n", article)
}

func main() {
	db = initSqlConnection()
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func initSqlConnection() *sql.DB {
	query := url.Values{}
	query.Add("database", fmt.Sprintf("%s", "TEST"))

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword("sa", "puieMonta140!"),
		Host:     fmt.Sprintf("%s:%d", "172.17.0.2", 1433),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	connectionString := u.String()
	fmt.Println(connectionString)
	db, err := sql.Open("mssql", connectionString)
	if err!= nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}
	return db

}

func getArticleById(id string) string {

	rows, err := db.Query("select * from articles where id='1'")
	if err != nil  {
		log.Fatal("error retrieving columns for id :",id)
	}
	defer rows.Close()
	var idd int
	var article string
	rows.Next()
	err = rows.Scan(&idd, &article)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	return article
}
