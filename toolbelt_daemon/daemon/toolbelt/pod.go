package toolbelt

import (
	"os"
	"time"
	"os/exec"

	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/api/request"
	"github.com/ravaj-group/farmer/toolbelt_daemon/daemon/db"
)


func Create(req request.CreateRequest) {
	codeDir := os.Getenv("TOOLBELT_TEMP_DIR") + "/" + time.Now().Format("20060102150405")

	set(req.PodName, PodStateJson(PreDeploy, "Pod instantiation"))
	if err := GitCloneBranch(req.RepoUrl, req.Pathspec, codeDir); err != nil {
		set(req.PodName, PodStateJson(Error, err.Error()))
		return
	}

	set(req.PodName, PodStateJson(Deploying, "Deploing " + req.PodName))

	cmd := exec.Command("toolbelt", "pod", "deploy", req.PodName)
	cmd.Dir    = codeDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if req.Domain != "" {
		cmd.Args = append(cmd.Args, "-d ", req.Domain)
	}

	if req.Env != nil {
		cmd.Env = req.Env
	}

	if err := cmd.Run(); err != nil {
		set(req.PodName, PodStateJson(Error, err.Error()))
	}

	exec.Command("rm -rf", codeDir).Run()
	set(req.PodName, PodStateJson(Success, "Done"))
}

func State(pod string) (string, error) {
	return get(pod)
}

func set(podName string, value string) error {
	return db.Set("/deploy/" + podName, value)
}

func get(podName string) (string, error) {
	return db.Get("/deploy/" + podName)
}
