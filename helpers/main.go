package main

import (
	"bytes"
	"go/format"
	"os"
	"strings"
	"text/template"
)

type TagEntry struct {
	TagTitle string
	TagLower string
}

func generateCode(tags []string) error {
	componentTemplate := "// {{.TagTitle}} returns a HTML5 <{{.TagLower}}> element with no default style.\nfunc {{.TagTitle}}(opts ...Option) Component {\n\treturn Any(\"{{.TagLower}}\", opts...)\n}\n\n"
	tmpl, err := template.New("template").Parse(componentTemplate)
	if err != nil {
		return err
	}

	var outBuf bytes.Buffer
	outBuf.Write([]byte("package sydercomponents\n\n"))

	for _, tag := range tags {
		if tag == "" {
			continue
		}

		entry := TagEntry{
			TagTitle: strings.ToUpper(string(tag[0])) + tag[1:],
			TagLower: tag,
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, entry)
		if err != nil {
			return err
		}

		outBuf.Write(buf.Bytes())
	}

	fout, err := format.Source(outBuf.Bytes())
	if err != nil {
		return err
	}

	err = os.WriteFile("../html.go", fout, 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateTests(tags []string) error {
	componentTemplate := "//<{{.TagLower}}>\nfunc Test{{.TagTitle}}_Options(t *testing.T) {\n\tcomp := sc.{{.TagTitle}}(sc.WithID(\"id\"), sc.WithData(\"name\", \"john\"))\n\n\trequire.Equal(t, \"id\", comp.GetID())\n\tv, has := comp.GetData(\"name\")\n\trequire.Equal(t, \"john\", v)\n\trequire.True(t, has)\n}\n\nfunc Test{{.TagTitle}}_Output(t *testing.T) {\n\tid := \"id\"\n\tcomp := sc.{{.TagTitle}}(sc.WithID(id))\n\tassertElementWithIDExists(t, comp, \"{{.TagLower}}\", id)\n}\n\n"
	tmpl, err := template.New("template").Parse(componentTemplate)
	if err != nil {
		return err
	}

	var outBuf bytes.Buffer
	outBuf.Write([]byte("package sydercomponents_test\n\nimport (\n\t\"testing\"\n\n\tsc \"github.com/geneton-ltd/syder_components\"\n\t\"github.com/stretchr/testify/require\"\n)\n\n"))

	for _, tag := range tags {
		if tag == "" {
			continue
		}

		entry := TagEntry{
			TagTitle: strings.ToUpper(string(tag[0])) + tag[1:],
			TagLower: tag,
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, entry)
		if err != nil {
			return err
		}

		outBuf.Write(buf.Bytes())
	}

	fout, err := format.Source(outBuf.Bytes())
	if err != nil {
		return err
	}

	err = os.WriteFile("../html_test.go", fout, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	tagsContent, err := os.ReadFile("tag_list.txt")
	if err != nil {
		panic(err)
	}

	tags := strings.Split(string(tagsContent), "\n")

	err = generateCode(tags)
	if err != nil {
		panic(err)
	}

	err = generateTests(tags)
	if err != nil {
		panic(err)
	}
}
