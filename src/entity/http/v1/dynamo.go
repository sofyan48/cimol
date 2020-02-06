package v1

// DynamoItem ..
type DynamoItem struct {
	ID              string
	Data            string
	History         string
	ReceiverAddress string
	StatusText      string
	Type            string
}

// Testing ...
type Testing struct {
	ID string `dynamo:"id"`
}
