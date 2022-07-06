// cform.go
package cform

import (
    "html/template"
    "net/http"
    C "contactform/producer"
    "encoding/json"
    "fmt"
)

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("forms.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // Handle form data
	data, err := json.Marshal(details)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	C.SendMsg(data)
        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}
