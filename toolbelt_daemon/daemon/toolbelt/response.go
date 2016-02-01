package toolbelt

import "encoding/json"

type PodStateResponse struct {
	State  	string	`json:"state"`
	Message string 	`json:"message"`
}

func GetPodStateResponseJson(state string, message string) string {
	json, _ := json.Marshal(&PodStateResponse{
		State: state,
		Message: message,
	})

	return string(json)
}