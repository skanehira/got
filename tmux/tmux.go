package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Tmux struct {
	Name     string
	Attached bool
}

type Session struct {
	SessionName string
	WindowName  string
	HostName    string
	Status      string
	Created     string
}

func New() *Tmux {
	tmux := &Tmux{}
	var attached bool
	name := currentSessionName()

	if name != "" {
		attached = true
	}

	tmux.Name = name
	tmux.Attached = attached

	return tmux
}

func currentSessionName() string {
	env := os.Getenv("TMUX")
	if env == "" {
		return ""
	}

	result := strings.Split(env, ",")
	return result[len(result)-1]
}

func (t *Tmux) SessionList() []*Session {
	format := "#{session_name} #{window_name} #h #{?session_attached,attached,unattached} #{session_created}"

	output, err := exec.Command("tmux", "ls", "-F", format).Output()

	if err != nil {
		return make([]*Session, 0)
	}

	return t.parseOutput(string(output))
}

func (t *Tmux) NewSession() error {
	cmd := exec.Command("tmux")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (t *Tmux) AttachSession(name string) error {
	cmd := exec.Command("tmux", "attach", "-t", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (t *Tmux) KillSession(name string) error {
	cmd := exec.Command("tmux", "kill-session", "-t", name)
	return cmd.Run()
}

func (t *Tmux) parseOutput(output string) []*Session {
	var sessions []*Session

	list := strings.Split(string(output), "\n")
	for i, l := range list {
		if len(list)-1 == i {
			break
		}

		cs := strings.Split(l, " ")
		sessions = append(sessions, &Session{
			SessionName: cs[0],
			WindowName:  cs[1],
			HostName:    cs[2],
			Status:      cs[3],
			Created:     t.parseDate(cs[4]),
		})
	}

	return sessions
}

func (t *Tmux) parseDate(date string) string {
	unixtime, err := strconv.ParseInt(date, 10, 64)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return time.Unix(unixtime, 0).Format("2006/01/02 15:04:05")
}
