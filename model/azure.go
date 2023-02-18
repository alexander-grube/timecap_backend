package model

// {
// "subscriptionId":"SUBSCRIPTION_ID",
// "detailedMessage":
// 	{"text":
// 	"Bug #5110 (This is another Test)
// 	Area: AREA 2\r\n-
// 	Iteration: SPRINT\r\n-
// 	State: New\r\n-
// 	Assigned to: ASSIGNEE <EMAIL@ADRESS.COM>\r\n-
// 	Severity: 1 - Critical\r\n"},
// 	"resource":
// 		{"id":5110,
// 		"url":"https://dev.azure.com/ORGANIZATION/PROJECT/_workitems/edit/5110"},
type Azure struct {
	SubscriptionID  string `json:"subscriptionId"`
	DetailedMessage struct {
		Text AzureText `json:"text"`
	} `json:"detailedMessage"`
	Resource struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"resource"`
}

type AzureText struct {
	Type       string
	ID         string
	Title      string
	Area       string
	Iteration  string
	State      string
	AssignedTo string
	Severity   string
}
