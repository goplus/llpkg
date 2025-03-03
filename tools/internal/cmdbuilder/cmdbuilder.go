package cmdbuilder

import (
	"fmt"
	"strings"
)

type CmdBuilderSerilizer func(k, v string) string

type Options func(*CmdBuilder)

type CmdBuilder struct {
	Serilazier CmdBuilderSerilizer
	cmds       map[string]string
}

func WithConanSerilazier() Options {
	return func(cb *CmdBuilder) {
		cb.Serilazier = func(k, v string) string {
			return fmt.Sprintf("--%s=%s", k, v)
		}
	}
}

func NewCmdBuilder(opts ...Options) *CmdBuilder {
	c := &CmdBuilder{
		cmds: make(map[string]string),
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *CmdBuilder) Set(k, v string) {
	c.cmds[k] = v
}

func (c *CmdBuilder) Args() []string {
	var cmds []string

	for k, v := range c.cmds {
		cmds = append(cmds, c.Serilazier(k, v))
	}

	return cmds
}

func (c *CmdBuilder) String() string {
	return strings.Join(c.Args(), " ")
}
