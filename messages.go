package saurontypes

import (
	"strings"
)

// Task is a struct that encapsulates the mapping between the
// queue to place a task on and the name of the image to be run
type Task struct {
	Queue     string `json:"queue" mapstructure:"queue"`
	ImageName string `json:"image" mapstructure:"image"`
	Name      string `json:"name" mapstructure:"name"`
	Data      string `json:"data" mapstructure:"data"`
}

type SauronConfig struct {
	Assignments []Assignment `json:"assignments" mapstructure:"assignments"`
}

type Assignment struct {
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	Prefix      string `json:"prefix" mapstructure:"prefix"`
	Tasks       []Task `json:"tasks" mapstructure:"tasks"`
}

// AngmarMessage is a struct that encapsulates the message that Angmar
// listens to on a queue for.
type AngmarMessage struct {
	Url     string `json:"url" mapstructure:"url"`
	SHA     string `json:"sha" mapstructure:"sha"`
	Pusher  string `json:"pusher" mapstructure:"pusher"`
	Project string `json:"project" mapstructure:"project"`
	Tasks   []Task `json:"tasks" mapstructure:"tasks"`
}

// String returns a stringified version of AngmarMessage, but doesn't
// stringify the list of Tasks.
func (m AngmarMessage) String() string {
	var builder strings.Builder
	builder.WriteString("URL: " + m.Url + "\n")
	builder.WriteString("Project: " + m.Project + "\n")
	builder.WriteString("SHA: " + m.SHA + "\n")
	builder.WriteString("Pusher: " + m.Pusher + "\n")
	return builder.String()
}

// UrukMessage is a struct that encapsulates the message that Uruk
// listens to on a queue for
type UrukMessage struct {
	ImageName    string
	RepoLocation string
	DataPath     string
}

// String returns a stringified version of UrukMessage
func (m UrukMessage) String() string {
	var builder strings.Builder
	builder.WriteString("Image: " + m.ImageName + "\n")
	builder.WriteString("Repo Location: " + m.RepoLocation + "\n")
	return builder.String()
}

// ConvertAngmarToUrukMessages converts an Angmar message to a map of queue names and
// their respective UrukMessage
func ConvertAngmarToUrukMessages(angmarMessage AngmarMessage, repoLocation string) map[string]UrukMessage {
	urukMessages := make(map[string]UrukMessage)
	for _, task := range angmarMessage.Tasks {
		urukMessage := UrukMessage{
			ImageName:    task.ImageName,
			RepoLocation: repoLocation,
			DataPath:     task.Data,
		}
		urukMessages[task.Queue] = urukMessage
	}
	return urukMessages
}
