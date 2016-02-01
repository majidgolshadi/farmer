package toolbelt

import (
	"os"
	"time"
	"os/exec"

	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/api/request"
	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/db"
)

func Create(req request.CreateRequest) {
	projectCodeDir := os.Getenv("TOOLBELT_TEMP_DIR") + "/" + time.Now().Format("20060102150405")

	cmd := exec.Command("git", "clone", req.RepoUrl, projectCodeDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		db.Set(req.PodName, GetPodStateResponseJson("error", err.Error()))
		return
	}

	cmd = exec.Command("git", "checkout", req.Pathspec)
	cmd.Dir = projectCodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		db.Set(req.PodName, GetPodStateResponseJson("error", err.Error()))
		return
	}

	cmd = exec.Command("toolbelt", "pod", "deploy", req.PodName)
	if req.Domain != "" {
		cmd.Args = append(cmd.Args, "-d ", req.Domain)
	}
	if req.Env != nil {
		cmd.Env = req.Env
	}
	cmd.Dir = projectCodeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	message := "done"
	status := "success"

	if err := cmd.Run(); err != nil {
		status = "error"
		message = err.Error()
	}

	exec.Command("rm -rf", projectCodeDir).Run()
	db.Set(req.PodName, GetPodStateResponseJson(status, message))
}

func State(pod string) (string, error) {
	return db.Get(pod)
}
