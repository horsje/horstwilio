package main

import "net/http"
import "log"
import "fmt"

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

const content = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Hello Monkey</Say>
</Response>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprint(w, content)
}
