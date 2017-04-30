package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

/*
Data structure to show basic dynamic functionality. The struct
will be filled with data accordingly, and then handed to the
templating agent, which then inserts the data into the HTML
that is going to be sent to the user.
*/
type pageData struct {
	Title string
}

//Function to initialize the server
func init() {

/*
This section takes any command line parameters that may have
been given to the application on server start and checks them for
valid matches.
*/
	progArgs := os.Args[1:]
	for _, progArg := range progArgs {
	    /*
	    In this basic version of the server implementation only
	    a loggin parameter has been implemented.
	    */
		if (strings.EqualFold(progArg, "--log")||strings.EqualFold(progArg, "-l")) {
			shariffLogging(true)
		}
	}

/*
This prints the location the server has been started from on the
host machine into the log.
*/
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started from:", dir)

/*
This takes all *.gohtml files from the folder templates and
parses them for their Golang templating tags. The specialized
file-ending is intended to differentiate between conventional
HTML files and files that use Golang templating tags and as
such would not be valid for a browser.
*/
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

/*
The main function calls all the handle methods of the http
package for their respective \ac{URI}s.
*/
func main() {
    //The root path is treated as the index page
	http.HandleFunc("/", idx)
	/*
	The pub path provides all publically available resources
	necesary for the page to load. This includes pictures and
	cascading stylesheets.
	*/
	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir("pub/"))))
	/*
	This handler catches the request for the favicon icon and
	returns a not found response.
	*/
	http.Handle("/favicon.ico", http.NotFoundHandler())
	/*
	Here the shariff backend is called when the frontend is
	loaded by the browser and requests the additional
	information from the server.
	*/
	http.HandleFunc("/shariff/", shariff)
	/*
	Here the server is started and instructed to listen to the
	standard HTTP Port of 80. This can be configured as
	wished.
	*/
	http.ListenAndServe(":80", nil)
}

//Page-Handler that responds to with the Index page
func idx(w http.ResponseWriter, req *http.Request) {

    /*
    Here the Page data that can be inserted dynamically into the
    page is initialized
    */
	data := pageData{
		Title: "Go Title",
	}

    /*
    By executing the template and handing the method the
    ResponseWriter object, information which template should be
    used and the pageData object, the user is sent the parsed HTML
    Response. If an error occurs that will be logged and the user
    is sent an error response.
    */
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}