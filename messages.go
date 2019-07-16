package saurontypes

import "strings"

// Task is a struct that encapsulates the mapping between the
// queue to place a task on and the name of the image to be run
type Task struct {
	Queue     string `json:"queue"`
	ImageName string `json:"imageName"`
}

// AngmarMessage is a struct that encapsulates the message that Angmar
// listens to on a queue for.
type AngmarMessage struct {
	Url     string `json:"url"`
	SHA     string `json:"sha"`
	Pusher  string `json:"pusher"`
	Project string `json:"project"`
	Tasks   []Task `json:"tasks"`
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
		}
		urukMessages[task.Queue] = urukMessage
	}
	return urukMessages
}
