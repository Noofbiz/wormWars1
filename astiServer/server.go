package astiServer

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("astiServer/html/*"))

// StartServer Starts the server to serve http to the astillectron gui
func StartServer() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/sim", simHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("astiServer/static"))))

	go http.ListenAndServe(":4000", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func simHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	d := make(map[string]string)
	d["pop"] = r.Form.Get("pop")
	d["initI"] = r.Form.Get("initI")
	d["S2I"] = r.Form.Get("S2I")
	d["I2R"] = r.Form.Get("I2R")
	d["S2R"] = r.Form.Get("S2R")

	err := tmpl.ExecuteTemplate(w, "sim", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
