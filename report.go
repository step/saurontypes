package saurontypes

type Report struct {
	Job     string `json:"job"`
	Results string `json:"result"`
}

type Results struct {
	Results string `json:"result.json"`
}

type TestResult struct {
	Total   int          `json:"total"`
	Passed  []TestReport `json:"passed"`
	Failed  []TestReport `json:"failed"`
	Pending []TestReport `json:"pending"`
}

type TestReport struct {
	Suite string `json:"suite"`
	Title string `json:"title"`
}

type DBReport struct {
	Job     string
	Result  TestResult
	FlowID  string
	Project string
	Pusher  string
	Time    string
}