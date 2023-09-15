package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Something went wrong!!")

	}
	e := json.Unmarshal([]byte(body), x)

	if e != nil {
		fmt.Println("Something went wrong!!")

	}

}
