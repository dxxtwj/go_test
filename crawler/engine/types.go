package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) PaseResult
}

type PaseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) PaseResult {
	return PaseResult{}
}
