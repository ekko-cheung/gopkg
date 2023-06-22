package sqlgen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

type Field struct {
	Name    string
	Type    string
	SqlName string
}

type Data struct {
	Fields        []Field
	FuncName      string
	ParamFullName string
	ParamName     string
}

func convert2Snake(s string) string {
	b := make([]byte, 0)
	b = append(b, s[0]+32)
	for i := 1; i < len(s); i++ {
		r := s[i]
		if r >= 65 && r <= 90 {
			b = append(b, '_')
			r = r + 32
		}
		b = append(b, r)
	}

	return string(b)
}

func parseStruct(s interface{}) Data {
	typ := reflect.TypeOf(s)
	numField := typ.NumField()
	fields := make([]Field, 0, numField)
	for i := 0; i < numField; i++ {
		field := typ.Field(i)
		if field.IsExported() {
			fields = append(fields, Field{
				Name:    field.Name,
				Type:    field.Type.String(),
				SqlName: convert2Snake(field.Name),
			})
		}
	}

	name := typ.Name()
	path := typ.PkgPath()
	sp := strings.Split(path, "/")
	return Data{
		Fields:        fields,
		FuncName:      name,
		ParamFullName: fmt.Sprintf("%s.%s", sp[len(sp)-1], name),
		ParamName:     string(name[0] + 32),
	}
}

func GenWhere(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	parse, err := template.New("").Parse(whereTemplate)
	if err != nil {
		log.Fatalln("parse whereTemplate fail: ", err)
	}
	for i := range m {
		where := parseStruct(m[i])
		parse.Execute(f, where)
	}
	f.Close()
}

func GenSet(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	parse, err := template.New("").Parse(setTemplate)
	if err != nil {
		log.Fatalln("parse whereTemplate fail: ", err)
	}
	for i := range m {
		where := parseStruct(m[i])
		parse.Execute(f, where)
	}
	f.Close()
}

func preGen(pkgName, path string) {
	os.MkdirAll(path, 0777)

	trimSqlFile := filepath.Join(path, "trimSql.go")
	file, err := os.OpenFile(trimSqlFile, os.O_APPEND, 0777)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.OpenFile(trimSqlFile, os.O_CREATE, 0777)
		parse, err := template.New("trimSqlFile").Parse(trimSqlTemplate)
		if err != nil {
			log.Fatalln("parse trimSqlTemplate fail: ", err)
		}
		parse.Execute(file, pkgName)
		file.Close()
	}
}
