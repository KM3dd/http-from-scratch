package types

type Request struct {
	Operation string
	Route     []string // route contain all
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
