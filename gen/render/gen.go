package render

import (
	"bytes"
	"go/format"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/hjhsamuel/zerostack/pkg/file"
)

func CreateGoTemplate(content, path string, params map[string]any) error {
	name := filepath.Base(path)
	if err := file.MkdirIfNotExist(filepath.Dir(path)); err != nil {
		return err
	}
	f, err := file.CreateFileIfNotExist(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buffer := new(bytes.Buffer)
	t := template.Must(template.New(name).
		Funcs(template.FuncMap{
			"FirstUpper": FirstUpper,
		}).Parse(content))
	err = t.Execute(buffer, params)
	if err != nil {
		return err
	}

	code := FormatGoCode(buffer.String())
	_, err = f.WriteString(code)
	return err
}

func OverwriteGoTemplate(content, path string, params map[string]any) error {
	name := filepath.Base(path)
	if err := file.MkdirIfNotExist(filepath.Dir(path)); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	buffer := new(bytes.Buffer)
	t := template.Must(template.New(name).
		Funcs(template.FuncMap{
			"FirstUpper": FirstUpper,
		}).Parse(content))
	err = t.Execute(buffer, params)
	if err != nil {
		return err
	}

	code := FormatGoCode(buffer.String())
	_, err = f.WriteString(code)
	return err
}

func FormatGoCode(content string) string {
	out, err := format.Source([]byte(content))
	if err != nil {
		return content
	}
	return string(out)
}

func FirstUpper(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
