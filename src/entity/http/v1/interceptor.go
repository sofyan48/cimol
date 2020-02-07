package v1

// StateFullKinesis ...
type StateFullKinesis struct {
	Data   *DynamoItem `json:"data"`
	Status string      `json:"queue"`
	Stack  string      `json:"stack"`
}
