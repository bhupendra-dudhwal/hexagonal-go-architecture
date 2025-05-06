package response

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"project_structure/internal/core/constants"
	"project_structure/internal/core/ports"
)

type response struct{}

func NewResponsepObj() ports.IResponse {
	return &response{}
}

func (r *response) JSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set(constants.HeaderContentType, constants.ContentTypeJSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (r *response) XML(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set(constants.HeaderContentType, constants.ContentTypeXML)
	w.WriteHeader(statusCode)
	xml.NewEncoder(w).Encode(data)
}

func (r *response) JSONCompressed(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(constants.HeaderContentEncoding, constants.ContentEncodingGZIP)
	w.Header().Set(constants.HeaderContentType, constants.ContentTypeJSON)
	w.WriteHeader(statusCode)

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	json.NewEncoder(gz).Encode(data)
	gz.Close()

	w.Write(buf.Bytes())
}

func (r *response) XMLCompressed(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(constants.HeaderContentEncoding, constants.ContentEncodingGZIP)
	w.Header().Set(constants.HeaderContentType, constants.ContentTypeXML)
	w.WriteHeader(statusCode)

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	xml.NewEncoder(gz).Encode(data)
	gz.Close()

	w.Write(buf.Bytes())
}
