package output

type TransferOutput struct {
	Id         string `json:"id"`
	Successful bool   `json:"successful"`
	Reason     string `json:"reason,omitempty"`
	Timestamp  string `json:"timestamp"`
}

func NewTransferOutput(id string, successful bool, timestamp string, reason string) *TransferOutput {
	return &TransferOutput{
		Id:         id,
		Reason:     reason,
		Successful: successful,
		Timestamp:  timestamp,
	}
}
