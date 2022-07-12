package build

type Command struct {
	name    string
	script  []string
	flags   map[string]flagType
	flagsFn map[string][]string
	sub     map[string]*Command
}

func NewCommand(name string) *Command {
	return &Command{
		name:    name,
		script:  make([]string, 0),
		flags:   make(map[string]flagType),
		flagsFn: make(map[string][]string),
		sub:     make(map[string]*Command),
	}
}

func (c *Command) Serialize() {

}
