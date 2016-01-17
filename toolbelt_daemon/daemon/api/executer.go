package api

import (
	"os"
	"time"
	"os/exec"

	"github.com/ravaj-group/toolbelt/daemon/api/request"
	"strings"
)

func Execute(req request.CmdRequest) (int, string) {
	if err := req.Validate(); err != nil {
		return 400, err.Error()
	}

	projectCodeDir := os.Getenv("TOOLBELT_TEMP_DIR") + "/" + time.Now().Format("20060102150405")

	cmd := exec.Command("git", "clone", req.RepoUrl, projectCodeDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return 500, err.Error()
	}

	cmd = exec.Command("git", "checkout", req.Pathspec)
	cmd.Dir = projectCodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return 500, err.Error()
	}

	args := strings.Split(req.Cmd, " ")
	cmd = exec.Command("toolbelt", args...)
	cmd.Env = req.Env
	cmd.Dir = projectCodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	message := ""
	statusCode := 201

	if err := cmd.Run(); err != nil {
		statusCode = 500
		message = err.Error()
	}

	exec.Command("rm -rf", projectCodeDir).Run()
	return statusCode, message
}
