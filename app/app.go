package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/skanehira/got/tmux"
)

type App struct {
	tmux *tmux.Tmux
}

func New() *App {
	return &App{
		tmux: tmux.New(),
	}
}

func (a *App) Run() {
	for {
		a.menu()
	}
}

func (a *App) newSession() {
	if err := a.tmux.NewSession(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func (a *App) menu() {
	type menu struct {
		Name string
		Do   func()
	}
	var menus []menu

	if !a.tmux.Attached {
		menus = append(menus, menu{"New Session", func() { a.newSession() }})
	}

	menus = append(menus, menu{"Session List", func() { a.sessionList() }})

	prompt := promptui.Select{
		Label: "Menu",
		Templates: &promptui.SelectTemplates{
			Label:    `{{ . | green }}`,
			Active:   `{{ .Name | red }}`,
			Inactive: ` {{ .Name | cyan }}`,
			Selected: `{{ .Name | yellow }}`,
		},
		Items: menus,
		Size:  20,
	}

	i, _, err := prompt.Run()

	if err != nil {
		if isEOF(err) || isInterrupt(err) {
			os.Exit(0)
		}
		fmt.Println(err)
		os.Exit(-1)
	}

	menus[i].Do()
}

func (a *App) sessionList() {
	sessions := a.tmux.SessionList()

	list := promptui.Select{
		Label: "Attaching Session: " + a.tmux.Name,
		Templates: &promptui.SelectTemplates{
			Label:    ` {{ . | green }}`,
			Active:   listTemplate("red"),
			Inactive: listTemplate("cyan"),
			Selected: listTemplate("yellow"),
		},
		Searcher: func(input string, index int) bool {
			session := sessions[index]
			name := strings.Replace(strings.ToLower(session.SessionName), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		},
		Items: sessions,
		Size:  20,
	}

	i, _, err := list.Run()

	if err != nil {
		if isEOF(err) {
			os.Exit(0)
		}

		if isInterrupt(err) {
			return
		}

		fmt.Println(err)
		os.Exit(-1)
	}

	if err := a.selectAction(sessions[i].SessionName); err != nil {
		if isEOF(err) {
			os.Exit(0)
		}

		if isInterrupt(err) {
			return
		}

		fmt.Println(err)
		os.Exit(-1)
	}
}

func (a *App) selectAction(name string) error {
	type action struct {
		Name string
		Do   func(name string) error
	}

	var actions []action

	if !a.tmux.Attached {
		actions = append(actions,
			action{
				Name: "Attach",
				Do: func(name string) error {
					return a.tmux.AttachSession(name)
				},
			})
	}

	actions = append(actions,
		action{
			Name: "Kill",
			Do: func(name string) error {
				return a.tmux.KillSession(name)
			},
		})

	prompt := promptui.Select{
		Label: "Action",
		Templates: &promptui.SelectTemplates{
			Label:    `{{ . | green }}`,
			Active:   `{{ .Name | red }}`,
			Inactive: ` {{ .Name | cyan }}`,
			Selected: `{{ .Name | yellow }}`,
		},
		Items: actions,
		Size:  20,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return err
	}

	return actions[i].Do(name)
}

func listTemplate(color string) string {
	template := `{{ printf "Name: %s" .SessionName | color }}	{{ printf "Window: %s" .WindowName | color }}	{{ printf "Host: %s" .HostName | color }}	{{ printf "Status: %s" .Status | color }}	{{ printf "Creaed: %s" .Created | color }}`

	return strings.Replace(template, "color", color, -1)
}

func isEOF(err error) bool {
	if err == promptui.ErrEOF {
		return true
	}

	return false
}

func isInterrupt(err error) bool {
	if err == promptui.ErrInterrupt {
		return true
	}

	return false
}
