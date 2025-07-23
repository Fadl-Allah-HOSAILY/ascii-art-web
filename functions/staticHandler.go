package functions

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w,  "Method not allowed !", http.StatusMethodNotAllowed,)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/static") {
		ErrorHandler(w,  "Statut not found !", http.StatusNotFound,)
		return
	} else {
		infos, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			ErrorHandler(w,  "Statut not found !", http.StatusNotFound,)
			return
		} else if infos.IsDir() {
			ErrorHandler(w,  "Access Forbidden !", http.StatusForbidden,)
			return
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}
