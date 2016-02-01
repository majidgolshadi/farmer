package request

import (
	"errors"
	"regexp"
)

type CreateRequest struct {
	PodName  string		`json:"pod_name" binding:"required"`
	RepoUrl  string 	`json:"repo_url" binding:"required"`
	Pathspec string 	`json:"pathspec" binding:"required"`
	Domain	 string		`json:"domain"`
	Env		 []string 	`json:"env"`
}

func (request *CreateRequest) Validate() error {
	if ok, _ := regexp.MatchString("((git|ssh|http(s)?)|(git@[\\w\\.]+))(:(//)?)([\\w\\.@\\:/\\-~]+)(\\.git)(/)?", request.RepoUrl); !ok {
		return errors.New("Invalid git repository URL [" + request.RepoUrl + "]")
	}

	return nil
}
