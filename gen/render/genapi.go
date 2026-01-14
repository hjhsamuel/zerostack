package render

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/example"
	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
	"github.com/hjhsamuel/zerostack/pkg/parser/api"
)

//go:embed api.gotmpl
var apiTpl string

const (
	apiFilePath = "internal/api/api.go"
	baseApiDir  = "api"
)

func CreateApiFile(base *entities.BaseInfo, path string) error {
	var absPath string
	if filepath.IsAbs(path) {
		absPath = path
	} else {
		absPath = filepath.Join(base.Home, path)
	}

	baseApiDirPath := filepath.Join(base.Home, base.Service, baseApiDir)

	var isInit bool
	if !file.Exists(absPath) {
		isInit = true
		// 如果 api 文件/文件夹不存在，则重定向到默认的 api 文件夹
		absPath = baseApiDirPath
	}

	if !file.Exists(baseApiDirPath) {
		if err := file.MkdirIfNotExist(baseApiDirPath); err != nil {
			return err
		}
	}
	if isInit {
		err := file.WriteFile(filepath.Join(baseApiDirPath, "example.api"), example.ApiExampleFile)
		if err != nil {
			return err
		}
	}
	// 扫描解析 api 文件
	afs, err := prepareApi(absPath)
	if err != nil {
		return err
	}

}

func prepareApi(path string) ([]*entities.ApiFile, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	res := make([]*entities.ApiFile, 0)
	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			out, err := prepareApi(filepath.Join(path, entry.Name()))
			if err != nil {
				return nil, err
			}
			res = append(res, out...)
		}
	} else {
		p, err := api.NewParser(path)
		if err != nil {
			return nil, err
		}
		astDef, err := p.Parse()
		if err != nil {
			return nil, err
		}
		res = append(res, &entities.ApiFile{
			Path: path,
			Api:  astDef,
		})
	}
	return res, nil
}
