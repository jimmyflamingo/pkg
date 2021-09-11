package tbdiags

type Diagnostic interface {
	Severity() Severity
	Description() Description
	Source() Source
}

type Severity rune

//go:generate go run golang.org/x/tools/cmd/stringer -type=Severity

const (
	Error   Severity = 'E'
	Warning Severity = 'W'
)

type Description struct {
	Address string
	Summary string
	Detail  string
}

type Source struct {
	Subject *SourceRange
	Context *SourceRange
}
