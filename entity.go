package main

type SuccessResponse struct {
	Result string `json:"result"`
}

type ResultResponse struct {
	Result any `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "not found"
}
