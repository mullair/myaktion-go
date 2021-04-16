package handler

import (
	"io"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

/*func Test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	io.WriteString(w, vars["name"])
}*/
