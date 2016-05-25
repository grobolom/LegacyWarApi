package main

import (
    "fmt"
    "net/http"

    // Third Party
    "github.com/julienschmidt/httprouter"
)

func main() {
    // Instantiate router
    r := httprouter.New()

    // set up a GET handler
    r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        fmt.Fprint(w, "Welcome!\n")
    })

    // httprouter does cool stuff I think
    http.ListenAndServe("localhost:3000", r)
}
