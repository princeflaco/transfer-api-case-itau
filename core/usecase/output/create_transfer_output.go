package output

type CreateTransferOutput struct {
	Id         string `json:"id"`
	Successful bool   `json:"successful"`
	Reason     string `json:"reason,omitempty"`
	Timestamp  string `json:"timestamp"`
}

func NewCreateTransferOutput(id string, successful bool, timestamp string, reason string) *CreateTransferOutput {
	return &CreateTransferOutput{
		Id:         id,
		Reason:     reason,
		Successful: successful,
		Timestamp:  timestamp,
	}
}
