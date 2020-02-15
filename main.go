package main

//import "fmt"
import "net/http"
import "html/template"
import "path"

func main() {
http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // not yet implemented
var filepath = path.Join("index.html")
var tmpl, err = template.ParseFiles(filepath)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

//var data = map[string]interface{}{
//    "title": "Learning Golang Web",
//    "name":  "Batman",
//}

err = tmpl.Execute(w, "no data needed")
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
}
    })

  //  fmt.Println("server started at localhost:9000")
  //  http.ListenAndServe(":9000", nil)

//	http.HandleFunc("/", index)
	//http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}


