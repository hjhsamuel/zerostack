package render

import (
	"bytes"
	"encoding/json"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

func CreateGoTemplate(content, path string, info *entities.GenInfo) error {
	name := filepath.Base(path)
	if err := file.MkdirIfNotExist(filepath.Dir(path)); err != nil {
		return err
	}
	f, err := file.CreateFileIfNotExist(path)
	if err != nil {
		return err
	}
	defer f.Close()

	code, err := GetRenderedContent(name, content, info)
	if err != nil {
		return err
	}
	_, err = f.WriteString(code)
	return err
}

func OverwriteGoTemplate(content, path string, info *entities.GenInfo) error {
	name := filepath.Base(path)
	if err := file.MkdirIfNotExist(filepath.Dir(path)); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	code, err := GetRenderedContent(name, content, info)
	if err != nil {
		return err
	}
	_, err = f.WriteString(code)
	return err
}

func GetRenderedContent(name, content string, info *entities.GenInfo) (string, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	var params map[string]any
	if err = json.Unmarshal(body, &params); err != nil {
		return "", err
	}

	return GetRenderedContentByParams(name, content, params)
}

func GetRenderedContentByParams(name, content string, params map[string]any) (string, error) {
	buffer := new(bytes.Buffer)
	t := template.Must(template.New(name).
		Funcs(template.FuncMap{
			"FirstUpper": FirstUpper,
			"ToDocPath":  ToDocPath,
		}).Parse(content))
	err := t.Execute(buffer, params)
	if err != nil {
		return "", err
	}
	code := FormatGoCode(buffer.String())
	return code, nil
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

func ToDocPath(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '/'
	})
	if len(parts) == 0 {
		return ""
	}
	for index, part := range parts {
		if strings.HasPrefix(part, ":") {
			parts[index] = "{" + part[1:] + "}"
		}
	}
	return "/" + strings.Join(parts, "/")
}
