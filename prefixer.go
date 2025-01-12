package log

// Prefix is used to return a [Logger] with a particular prefix.
type Prefixer interface {
	Prefix(...string) Logger
}
