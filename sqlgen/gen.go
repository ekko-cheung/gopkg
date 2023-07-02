/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	FieldLength   int
	TableName     string
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

type Table interface {
	TableName() string
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

	tableName := ""
	if t, ok := s.(Table); ok {
		tableName = t.TableName()
	}

	name := typ.Name()
	path := typ.PkgPath()
	sp := strings.Split(path, "/")

	return Data{
		Fields:        fields,
		FuncName:      name,
		ParamFullName: fmt.Sprintf("%s.%s", sp[len(sp)-1], name),
		ParamName:     string(name[0] + 32),
		FieldLength:   len(fields) - 1,
		TableName:     tableName,
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

func GenSelectColumns(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	parse, err := template.New("").Parse(columnTemplate)
	if err != nil {
		log.Fatalln("parse whereTemplate fail: ", err)
	}
	for i := range m {
		where := parseStruct(m[i])
		parse.Execute(f, where)
	}
	f.Close()
}

func GenInsert(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	parse, err := template.New("").Parse(insertTemplate)
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
	}
	if file != nil {
		file.Close()
	}
}

func GenAll(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	wt, _ := template.New("").Parse(whereTemplate)
	st, _ := template.New("").Parse(setTemplate)
	ct, _ := template.New("").Parse(columnTemplate)
	it, _ := template.New("").Parse(insertTemplate)
	for i := range m {
		data := parseStruct(m[i])
		wt.Execute(f, data)
		st.Execute(f, data)
		ct.Execute(f, data)
		it.Execute(f, data)
	}
	f.Close()
}
