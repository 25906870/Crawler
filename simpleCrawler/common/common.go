package common

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
