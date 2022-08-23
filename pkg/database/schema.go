package database

const (
	SmsStatusPending = "PENDING"
)

type Sms struct {
	Pk       string `dynamodbav:"pk" json:"pk"` // Audience
	Sk       string `dynamodbav:"sk" json:"sk"` // Unique message identifier
	To       int    `dynamodbav:"to" json:"to"`
	Message  string `dynamodbav:"message" json:"message"`
	Status   string `dynamodbav:"status" json:"status"` // PENDING, SENT, FAILED
	StatusAt int64  `dynamodbav:"status_at" json:"status_at"`
}
