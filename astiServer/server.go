package astiServer

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Noofbiz/wormWars1/simulation"
)

var tmpl = template.Must(template.ParseGlob("astiServer/html/*"))

// StartServer Starts the server to serve http to the astillectron gui
func StartServer() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/sim", simHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./astiServer/static/"))))

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

	pop, _ := strconv.Atoi(r.Form.Get("pop"))
	initI, _ := strconv.Atoi(r.Form.Get("initI"))
	s2i, _ := strconv.ParseFloat(r.Form.Get("S2I"), 64)
	i2r, _ := strconv.ParseFloat(r.Form.Get("I2R"), 64)
	s2r, _ := strconv.ParseFloat(r.Form.Get("S2R"), 64)

	simulation.Simulate(pop, initI, s2i, i2r, s2r)

	err := tmpl.ExecuteTemplate(w, "sim", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
