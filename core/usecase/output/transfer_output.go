package output

type TransferOutput struct {
	Id         string `json:"id"`
	Successful bool   `json:"successful"`
	Timestamp  string `json:"timestamp"`
}

func NewTransferOutput(id string, successful bool, timestamp string) *TransferOutput {
	return &TransferOutput{
		Id:         id,
		Successful: successful,
		Timestamp:  timestamp,
	}
}
