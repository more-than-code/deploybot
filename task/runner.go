package task

import (
	"log"
	"os"
	"os/exec"

	"github.com/more-than-code/deploybot/model"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) DoTask(t model.Task, args []string) error {
	if t.Config.Script != "" {
		cmd := exec.Command("echo", t.Config.Script, "/var/opt/mypipe")
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, args...)
		output, err := cmd.Output()
		log.Println(string(output))
		if err != nil {
			return err
		}
	}

	return nil
}
