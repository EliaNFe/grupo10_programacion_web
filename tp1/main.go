package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"

	fmt.Printf("servidor escuchando en http://localhost%s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //HOME DE LA PAGINA
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "<h1>404 - Página no encontrada</h1>")
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "index.html")
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}

}
