package browser

import (
	"os"
	"os/exec"
	"testing"
)

func mockedExecCommand(command string, args...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestLaunchBrowser(t *testing.T) {
	execCommand = mockedExecCommand
	defer func(){ execCommand = exec.Command }()

	LaunchBrowser("http://example.org")
}
