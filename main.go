package horstwilio

import "net/http"

import "fmt"

func init() {
	http.HandleFunc("/", handler)
}

const content = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Hello Gopher</Say>
</Response>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprint(w, content)
}
