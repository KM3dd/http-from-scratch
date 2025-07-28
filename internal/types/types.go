package types

type Request struct {
	Operation string   // GET / POST
	Route     []string // route contain all
	Headers   map[string]string
	Body      string
}

type Response struct {
	Code            int
	Message         string
	Headers         map[string]string
	Body            string
	ContentType     string
	ContentLength   int
	ContentEncoding []string
}
