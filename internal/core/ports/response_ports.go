package ports

import "net/http"

type IResponse interface {
	JSON(w http.ResponseWriter, data any, statusCode int)
	XML(w http.ResponseWriter, data any, statusCode int)
	JSONCompressed(w http.ResponseWriter, data interface{}, statusCode int)
	XMLCompressed(w http.ResponseWriter, data interface{}, statusCode int)
}
