package endpoints

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	if vars["config"] == "ip" {
		conn, _ := net.Dial("udp", "8.8.8.8:80")
		fmt.Fprintf(w, "System: %v Config: %v -> %v\n", vars["system"], vars["config"], strings.Split(conn.LocalAddr().String(), ":")[0])
	}

}
