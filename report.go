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

type DBTestReport struct {
	Job     string
	Result  TestResult
	FlowID  string
	Project string
	Pusher  string
	Time    string
	SHA     string
}

type LintReport struct {
	RuleID   string `json:"ruleId"`
	Severity int    `json:"severity"`
	Message  string `json:"message"`
	Line     int    `json:"line"`
	Col      int    `json:"column"`
}

type LintResult struct {
	FileName            string `json:"filePath"`
	Messages            []LintReport
	ErrorCount          int    `json:"errorCount"`
	WarningCount        int    `json:"warningCount"`
	FixableErrorCount   int    `json:"fixableErrorCount"`
	FixableWarningCount int    `json:"fixableWarningCount"`
}

type DBLintReport struct {
	Job     string
	Result  []LintResult
	FlowID  string
	Project string
	Pusher  string
	Time    string
	SHA     string
}
