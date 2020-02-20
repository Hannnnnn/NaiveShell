package command

const (
	INPUT_REDIRECTION = 0
	OUTPUT_REDIRECTION = 1
)

const (
	COMMAND = iota
	REDIRECTION_COMMAND

)
type RedirectionType int
// Ordinary command
type Command struct {
	Name string
	Args []string

}

type ListCommand struct {

}

type RedirectionCommand struct {
	Direction RedirectionType
	Command interface{}

}