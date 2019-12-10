package saurontypes

type Entry struct {
	Key   string
	Value interface{}
}

type StreamEvent struct {
	ID string
	Values map[string]interface{}
}

type Event struct {
	Source    string
	Type      string
	FlowID    string
	Timestamp string
	PusherID  string
	Project   string
	Details   string
}

func (e Event) ConvertToEntry() []Entry {
	entries := make([]Entry, 0)

	entries = append(entries,
		Entry{
			Key:   "source",
			Value: e.Source,
		},
		Entry{
			Key:   "type",
			Value: e.Type,
		},
		Entry{
			Key:   "flowID",
			Value: e.FlowID,
		},
		Entry{
			Key:   "timestamp",
			Value: e.Timestamp,
		},
		Entry{
			Key:   "pusherID",
			Value: e.PusherID,
		},
		Entry{
			Key:   "project",
			Value: e.Project,
		},
		Entry{
			Key:   "details",
			Value: e.Details,
		},
	)

	return entries
}
