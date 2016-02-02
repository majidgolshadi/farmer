package toolbelt

import "encoding/json"

type PodState struct {
	State  	string	`json:"state"`
	Message string 	`json:"message"`
}

func PodStateJson(state string, message string) string {
	json, _ := json.Marshal(&PodState{
		State: state,
		Message: message,
	})

	return string(json)
}