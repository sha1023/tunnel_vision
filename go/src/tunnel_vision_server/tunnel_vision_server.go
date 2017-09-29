package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "net/http"
    "os"
)

//TODO make this safer and more reliable:
var saveDir string = "/tmp/"

func get(w http.ResponseWriter, r *http.Request) {
    name := mux.Vars(r)["name"]
    text, err := ioutil.ReadFile(saveDir + name)
    if err != nil {
        fmt.Fprintf(w, "error retrieving " + name + ":\n" + err.Error() + "\n")
        return
    }
    fmt.Fprintf(w, string(text))
}

func save(w http.ResponseWriter, r *http.Request) {
    name := mux.Vars(r)["name"]
    body, err := ioutil.ReadAll(r.Body)
    if  err != nil {
        fmt.Println("error retrieving body for " + name + ":\n" + err.Error() + "\n")
    }
    err = ioutil.WriteFile(saveDir + name, body, 0644)
    if  err != nil {
        fmt.Println("error writing body for " + name + ":\n" + err.Error() + "\n")
    }
}

func endpoints(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "\t/: \t\t\t\t\t endpoints\n")
    fmt.Fprintf(w, "\t/tv/: \t\t\t tunnel vision\n")
    fmt.Fprintf(w, "\t/api/save/{name}: \t\t\t [POST] save {name}\n")
    fmt.Fprintf(w, "\t/api/get/{name}: \t\t\t get {name}\n")
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please specify path to index.html")
        os.Exit(1)
    }
    htmlPath := os.Args[1]
    r := mux.NewRouter()
    r.HandleFunc("/api/get/{name}", get)
    r.HandleFunc("/api/save/{name}", save).Methods("POST")
    r.HandleFunc("/", endpoints)
    r.PathPrefix("/tv/").Handler(http.StripPrefix("/tv/", http.FileServer(http.Dir(htmlPath))))
    http.Handle("/tv/", r)
    http.Handle("/", r)

    port := "1920"
    fmt.Print("Starting Server on port: " + port + "\n")
    err := http.ListenAndServe(":1920", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
}
