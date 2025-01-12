package log

// Tagger is used to return a [Logger] with a particular tag.
type Tagger interface {
	Tag(string) Logger
}
