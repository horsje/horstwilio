package horstwilio

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

const (
	welcomeMsg = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say>Hello Gopher, how are you?</Say>
	<Record timeout="5" />
</Response>`
	echoMsg = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Play>%s</Play>
</Response>`
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")

	rec := r.FormValue("RecordingUrl")
	if rec == "" {
		fmt.Fprint(w, welcomeMsg)
		return
	}

	fmt.Fprintf(w, echoMsg, rec)
}
