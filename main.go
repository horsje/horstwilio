package horstwilio

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", handler)
}

const content = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Hello Gopher, how are you?</Say>
	<Record timeout="5" />
</Response>`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := appengine.NewContext(r)
	for k, v := range r.PostForm {
		log.Infof(c, "%s: %v", k, v)
	}

	fmt.Fprint(w, content)
}
