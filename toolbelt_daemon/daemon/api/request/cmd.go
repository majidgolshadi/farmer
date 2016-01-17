package request

import (
	"errors"
	"regexp"
)

type CmdRequest struct {
	Cmd      string 	`json:"cmd" binding:"required"`
	RepoUrl  string 	`json:"repo_url" binding:"required"`
	Pathspec string 	`json:"pathspec" binding:"required"`
	Env		 []string 	`json:"env"`
}

func (request *CmdRequest) Validate() error {
	if ok, _ := regexp.MatchString("((git|ssh|http(s)?)|(git@[\\w\\.]+))(:(//)?)([\\w\\.@\\:/\\-~]+)(\\.git)(/)?", request.RepoUrl); !ok {
		return errors.New("Invalid git repository URL [" + request.RepoUrl + "]")
	}

	return nil
}
