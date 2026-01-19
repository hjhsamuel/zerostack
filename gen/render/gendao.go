package render

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/hjhsamuel/zerostack/gen/entities"
	"github.com/hjhsamuel/zerostack/pkg/file"
)

//go:embed dao/dao.gotmpl
var daoTpl string

//go:embed dao/mysql.gotmpl
var mysqlTpl string

const (
	daoFilePath   = "internal/dao/dao.go"
	mysqlFilePath = "internal/dao/mysql.go"
)

const (
	DatabaseMysql = "mysql"
)

func CreateDaoFile(base *entities.BaseInfo, database string) error {
	// create dao file
	absDaoPath := filepath.Join(base.SrvHome, daoFilePath)
	if err := createDaoFile(absDaoPath); err != nil {
		return err
	}
	// create database file
	switch database {
	case DatabaseMysql:
		absMysqlPath := filepath.Join(base.SrvHome, mysqlFilePath)
		return createMysqlFile(absMysqlPath)
	default:
		return nil
	}
}

func createDaoFile(path string) error {
	if file.Exists(path) {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(daoTpl)
	return err
}

func createMysqlFile(path string) error {
	if file.Exists(path) {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(mysqlTpl)
	return err
}
