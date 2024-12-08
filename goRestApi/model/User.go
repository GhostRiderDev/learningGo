package model

import "encoding/json"

type User struct {
	Name string `json:"name"`
	Age  json.Number    `json:"age"`
}
