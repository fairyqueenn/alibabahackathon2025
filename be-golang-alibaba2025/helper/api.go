package helper

type (
	BasicResponseStruct struct {
		Status int `json:"code"`
	}

	APIResponseStruct struct {
		BasicResponseStruct
		Data any `json:"data"`
	}

	APIResponseErrorStruct struct {
		BasicResponseStruct
		Error string `json:"error"`
	}
)

func BasicAPIResponse(code int) BasicResponseStruct {
	var r BasicResponseStruct
	r.Status = code
	return r
}

func APIResponse(code int, data any) APIResponseStruct {
	var r APIResponseStruct
	r.Status = code
	r.Data = data
	return r
}

func APIResponseError(code int, err string) APIResponseErrorStruct {
	var r APIResponseErrorStruct
	r.Status = code
	r.Error = err
	return r
}
