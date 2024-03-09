package formmethod

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// defaultInputName is the default name of the input field
const defaultInputName = "_method"

// Middleware is a middleware that changes the request method based on the value of a form field
var Middleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch method := strings.ToUpper(r.PostFormValue(defaultInputName)); method {
			case http.MethodPut, http.MethodPatch, http.MethodDelete:
				r.Method = method
			}
		}
		next.ServeHTTP(w, r)
	})
}

// TemplateField returns a hidden input field with the method value
func TemplateField(method string) template.HTML {
	return template.HTML(
		fmt.Sprintf(
			`<input type="hidden" name="%s" value="%s">`,
			defaultInputName,
			method,
		),
	)
}
