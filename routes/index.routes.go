package routes

import "net/http"

// HomeHandler function to handler home router
func HomeHandler(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("Hello World2!"))
}
