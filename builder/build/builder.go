package build

import (
	"runtime"
	"strings"
)

type Builder struct {
	OS         string
	PATH       string
	WORKDIR    string
	ExportName string
	ExportType string
}

func NewBuilder() *Builder {
	builder := &Builder{OS: runtime.GOOS}
	switch builder.OS {
	case "windows":
		builder.ExportType = BATCH
	case "linux":
		builder.ExportType = SHELL
	}
	return builder
}

func Build(args []string, flags map[string]any) {
	lexer := NewLexer()
	lexer.Read(args[0])
	stream := NewStream(lexer.CMD)
	root := NewCommand(strings.Split(args[0], ".")[0])
	ParserCmd(stream, root)
	builder := NewBuilder()
	if lexer.ENV["EXPORT"] != "" {
		builder.ExportName = lexer.ENV["EXPORT"]
	}
	if flags["ExportName"].(string) != "" {
		builder.ExportName = flags["ExportName"].(string)
	}
	if flags["ExportType"].(string) != "" {
		builder.ExportType = flags["ExportType"].(string)
	}
	if flags["OS"].(string) != "" {
		switch strings.ToLower(flags["OS"].(string)) {
		case "windows":
			builder.ExportType = BATCH
		case "linux":
			builder.ExportType = SHELL
		default:
			panic("Unsupport current os")
		}
	}
}
