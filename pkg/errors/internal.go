package errors

type Internal struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"error"`
}

func (e Internal) Error() string {
	return e.Err.Error()
}
