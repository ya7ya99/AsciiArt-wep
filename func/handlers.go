package functions

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Data struct holds the result, banner, text, and type of the ASCII art
type Data struct {
	Result string
	banner string
	text   string
	Type   string
}
var D Data

// ERRORS struct holds the error information
type ERRORS struct {
	PageTitle string
	Message   string
	ErrCde    int
}
var ERR ERRORS

// Welcom is the handler function for the welcome page
func Welcom(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/index.html")
	tmpl2, err2 := template.ParseFiles("../templates/errors.html")
	if err2 != nil {
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}
	if err != nil {
		ChooseErr(500, w)
		tmpl2.Execute(w, ERR)
		return
	}

	if r.URL.Path != "/" {
		ChooseErr(404, w)
		tmpl2.Execute(w, ERR)
		return
	}

	if r.Method != "GET" {
		ChooseErr(405, w)
		tmpl2.Execute(w, ERR)
		return
	}

	tmpl.Execute(w, nil)
}

// Last is the handler function for the ASCII art generation page
func Last(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/result.html")
	tmpl2, err2 := template.ParseFiles("../templates/errors.html")
	if err != nil {
		ChooseErr(500, w)
		tmpl2.Execute(w, ERR)
		return
	}
	if err2 != nil {
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/ascii-art" {
		ChooseErr(404, w)
		tmpl2.Execute(w, ERR)
		return
	}
	if r.Method != "POST" {
		ChooseErr(405, w)
		tmpl2.Execute(w, ERR)
		return
	}
	D.text = r.FormValue("ljomla")
	D.text = strings.ReplaceAll(D.text,"\r\n","\n")
	if len(D.text) > 250 {
		ChooseErr(777, w)
		tmpl2.Execute(w, ERR)
		return
	}
	D.banner = r.FormValue("banner")
	if D.text == "" || D.banner == "" {
		ChooseErr(400, w)
		tmpl2.Execute(w, ERR)
		return
	}

	D.Result = FS(D.banner, D.text)
	if D.Result == "ERORR" {
		ChooseErr(400, w)
		tmpl2.Execute(w, ERR)
		return
	}
	tmpl.Execute(w, D)
}

// ServeStyle serves the CSS files for the application
func ServeStyle(w http.ResponseWriter, r *http.Request) {
	tmpl2, err2 := template.ParseFiles("../templates/errors.html")
	if err2 != nil {
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}
	fs := http.StripPrefix("/styles/", http.FileServer(http.Dir("../styles")))
	if r.URL.Path == "/styles/" {
		ChooseErr(404, w)
		tmpl2.Execute(w, ERR)
		return
	}
	fs.ServeHTTP(w, r)
}

// chooseErr sets the error message and HTTP status code based on the provided error code
func ChooseErr(code int, w http.ResponseWriter) {
	if code == 404 {
		ERR.PageTitle = "Error 404"
		ERR.Message = "The page web doesn't exist\nError 404"
		ERR.ErrCde = code
		w.WriteHeader(code)
	} else if code == 405 {
		ERR.PageTitle = "Error 405"
		ERR.Message = "The method is not alloweded\nError 405"
		ERR.ErrCde = code
		w.WriteHeader(code)
	} else if code == 400 {
		ERR.PageTitle = "Error 400"
		ERR.Message = "Bad Request\nError 400"
		ERR.ErrCde = code
		w.WriteHeader(code)
	} else if code == 500 {
		ERR.PageTitle = "Error 500"
		ERR.Message = "Internal Server Error\nError 500"
		ERR.ErrCde = code
		w.WriteHeader(code)
	}else if code == 777 {
		ERR.PageTitle = "Error 400"
		ERR.Message = "Bad Request: Text have more than 250 characters\nError 400"
		ERR.ErrCde = 400
	}
}

// Download handles the download of the generated ASCII art
func Download(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Type := r.FormValue("type")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art."+Type)

	switch Type {
	case "html":
		D.Result = "<pre>" + D.Result + "</pre>"
		w.Header().Set("Content-Type", "text/html")
	case "txt":
		w.Header().Set("Content-Type", "text/plain")
	}

	w.Header().Set("Content-Length", strconv.Itoa(len( D.Result)))
	w.Write([]byte( D.Result))
}
