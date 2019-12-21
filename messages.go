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

// SauronConfig is the top level config for sauron
// This will come from the config file for sauron
// This will contain different config for different assignments
type SauronConfig struct {
	Assignments []Assignment `json:"assignments" mapstructure:"assignments"`
}

// Assignment is the config for a particular assignment
// including the name, description, prefix and the associated Tasks
type Assignment struct {
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	Prefix      string `json:"prefix" mapstructure:"prefix"`
	Tasks       []Task `json:"tasks" mapstructure:"tasks"`
}

// AngmarMessage is a struct that encapsulates the message that Angmar
// listens to on a queue for.
type AngmarMessage struct {
	URL     string `json:"url" mapstructure:"url"`
	Stream  string `json:"stream" mapstructure:"stream"`
	SHA     string `json:"sha" mapstructure:"sha"`
	FlowID  string `json:"flowID" mapstructure:"flowID"`
	Pusher  string `json:"pusher" mapstructure:"pusher"`
	Project string `json:"project" mapstructure:"project"`
	Tasks   []Task `json:"tasks" mapstructure:"tasks"`
}

// String returns a stringified version of AngmarMessage, but doesn't
// stringify the list of Tasks.
func (m AngmarMessage) String() string {
	var builder strings.Builder
	builder.WriteString("Flow ID: " + m.FlowID + "\n")
	builder.WriteString("Stream: " + m.Stream + "\n")
	builder.WriteString("URL: " + m.URL + "\n")
	builder.WriteString("Project: " + m.Project + "\n")
	builder.WriteString("SHA: " + m.SHA + "\n")
	builder.WriteString("Pusher: " + m.Pusher + "\n")
	return builder.String()
}

// UrukMessage is a struct that encapsulates the message that Uruk
// listens to on a queue for
type UrukMessage struct {
	FlowID       string
	Pusher       string
	Project      string
	Stream       string
	ImageName    string
	RepoLocation string
	DataPath     string
	Job          string
	SHA          string
}

// String returns a stringified version of UrukMessage
func (m UrukMessage) String() string {
	var builder strings.Builder
	builder.WriteString("Flow ID: " + m.FlowID + "\n")
	builder.WriteString("Pusher: " + m.Pusher + "\n")
	builder.WriteString("Stream: " + m.Stream + "\n")
	builder.WriteString("Image: " + m.ImageName + "\n")
	builder.WriteString("Project: " + m.Project + "\n")
	builder.WriteString("Repo Location: " + m.RepoLocation + "\n")
	builder.WriteString("SHA: " + m.SHA + "\n")
	return builder.String()
}

// ConvertAngmarToUrukMessages converts an Angmar message to a map of queue names and
// their respective UrukMessage
func ConvertAngmarToUrukMessages(angmarMessage AngmarMessage, repoLocation string) map[string]UrukMessage {
	urukMessages := make(map[string]UrukMessage)
	for _, task := range angmarMessage.Tasks {
		urukMessage := UrukMessage{
			FlowID:       angmarMessage.FlowID,
			Stream:       angmarMessage.Stream,
			Pusher:       angmarMessage.Pusher,
			Project:      angmarMessage.Project,
			SHA:          angmarMessage.SHA,
			ImageName:    task.ImageName,
			RepoLocation: repoLocation,
			DataPath:     task.Data,
			Job:          task.Name,
		}
		urukMessages[task.Queue] = urukMessage
	}
	return urukMessages
}
