package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type Event struct {
	Type    string  `json:"type"` //事件类型
	Repo    Repo    `json:"repo"` //操作的仓库
	Payload Payload `json:"payload"`
}

type Payload struct {
	Ref_type string `json:"ref_type"` //操作的对象
	Action   string `json:"action"`   //如何操作
	Size     int    `json:"size"`
}

type Repo struct {
	Name string `json:"name"`
}

var events []Event

func FetchTheRecentActivity(username string) error {
	r, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		return err
	}

	return Display()
}

func Display() error {
	yellow := color.New(color.FgYellow).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()

	fmt.Println(green("Output:"))
	fmt.Print("\n")
	for _, event := range events {
		var output string
		switch event.Type {
		case "CreateEvent":
			output = fmt.Sprintf("- Created %s in %s", event.Payload.Ref_type, event.Repo.Name)
		case "DeleteEvent":
			output = fmt.Sprintf("- Deleted %s in %s", event.Payload.Ref_type, event.Repo.Name)
		case "PushEvent":
			output = fmt.Sprintf("- Pushed %d commit(s) to %s", event.Payload.Size, event.Repo.Name)
		case "PublicEvent":
			output = fmt.Sprintf("- Opened %s", event.Repo.Name)
		case "IssuesEvent":
			output = fmt.Sprintf("- %s a issue in %s", event.Payload.Action, event.Repo.Name)
		case "ForkEvent,GollumEvent":
			output = fmt.Sprintf("- %s in %s", event.Type, event.Repo.Name)
		default:
			output = fmt.Sprintf("- %s : %s in %s", event.Type, event.Payload.Action, event.Repo.Name)
		}
		fmt.Println(yellow(output))
		fmt.Print("---------------------------------------------------------------------\n")
	}
	return nil
}
