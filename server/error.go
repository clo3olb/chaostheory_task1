package server

type Error interface {
	ResponseHeader() int
	ResponseBody() string
}

type ClientError struct {
	Status int   `json:"-"`
	Cause  error `json:"-"`
}

func (c ClientError) StatusCode() int {
	return c.Status
}

// client implements error interface
func (c ClientError) Error() string {
	return c.Cause.Error()
}

func NewClientError(status int, err error) ClientError {
	return ClientError{status, err}
}
