package saurontypes

import "strings"

// AngmarMessage is a struct that encapsulates the message that Angmar
// listens to on a queue for.
type AngmarMessage struct {
	Url     string
	SHA     string
	Pusher  string
	Project string
	Tasks   []string
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
