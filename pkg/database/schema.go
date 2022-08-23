package database

const (
	SmsStatusPending = "Pending"
	SmsStatusSuccess = "Success"
	SmsStatusFailed  = "Failed"
)

type Sms struct {
	Pk       string `dynamodbav:"pk" json:"pk"` // Status
	Sk       string `dynamodbav:"sk" json:"sk"` // Unique message identifier
	To       int    `dynamodbav:"to" json:"to"`
	Message  string `dynamodbav:"message" json:"message"`
	StatusAt int64  `dynamodbav:"status_at" json:"status_at"`
}
