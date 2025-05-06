package constants

type Request string

const (
	REQUEST_KEY Request = "request_id"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderContentEncoding = "Content-Encoding"
	ContentTypeJSON       = "application/json"
	ContentTypeXML        = "application/xml"
	ContentEncodingGZIP   = "gzip"
)
