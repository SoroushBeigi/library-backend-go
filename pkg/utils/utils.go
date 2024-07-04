package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(request *http.Request, x interface{}){
	if body, error := io.ReadAll(request.Body); error==nil{
		if err := json.Unmarshal([]byte(body), x); err!=nil{
			return 
		}
	}

}