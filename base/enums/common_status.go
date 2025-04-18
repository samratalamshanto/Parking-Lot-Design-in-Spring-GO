package enums

type ECommonStatus struct {
	name        string
	code        string
	description string
}

func (e *ECommonStatus) Name() string {
	return e.name
}

func (e *ECommonStatus) Code() string {
	return e.code
}

func (e *ECommonStatus) Description() string {
	return e.description
}

var (
	SUCCESS    = ECommonStatus{name: "SUCCESS", code: "0000", description: "success"}
	FAILED     = ECommonStatus{name: "FAILED", code: "9999", description: "failed"}
	INVALID    = ECommonStatus{name: "INVALID", code: "1111", description: "invalid"}
	INPROGRESS = ECommonStatus{name: "INPROGRESS", code: "2222", description: "InProgress"}
)
