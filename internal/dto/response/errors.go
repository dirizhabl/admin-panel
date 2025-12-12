package response

type Err struct {
	Err string `json:"error"`
}

type Errs struct {
	Errs []string `json:"errors"`
}
