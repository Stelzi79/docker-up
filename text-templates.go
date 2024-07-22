package main

import (
	"bytes"
	"html/template"
	"strings"
)

const (
	colorsTest   = "{{.Colors.Yellow}}🏷️ Colors Test Template-String"
	debugEnabled = "{{.Colors.Yellow}}🐛 Debugging enabled"
	name         = "{{.Colors.Blue}}docker-up"
)

type tmps struct {
	colorsTest   string
	debugEnabled string
	name         string
}

var templates = tmps{
	colorsTest:   parseTmpl(colorsTest),
	debugEnabled: parseTmpl(debugEnabled),
	name:         parseTmpl(name),
}

var replacers = []map[string]string{
	{"\t": "  "}, // Replace tabs with two spaces
}

func parseTmpl(tmp string) string {
	t1 := template.New("t1")

	for k, v := range replacers[0] {
		tmp = strings.ReplaceAll(tmp, k, v)
	}
	tmp += "{{.Format.Reset}}"

	t1, err := t1.Parse(tmp)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	t1.Execute(&out, map[string]any{"Colors": Colors, "Format": Format})

	return out.String()
}
