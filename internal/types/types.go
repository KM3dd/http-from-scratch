package types

type Request struct {
	Operation string
	Route     string
	Headers   []string
	Body      string
}

type Response struct {
	Code          int
	Message       string
	Headers       string
	Body          string
	ContentType   string
	ContentLength int
}
