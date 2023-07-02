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
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func GenPgWhere(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	parse, err := template.New("").Parse(pgWhereTemplate)
	if err != nil {
		log.Fatalln("parse whereTemplate fail: ", err)
	}
	for i := range m {
		where := parseStruct(m[i])
		parse.Execute(f, where)
	}
	f.Close()
}

func GenPgSet(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	parse, err := template.New("").Parse(pgSetTemplate)
	if err != nil {
		log.Fatalln("parse whereTemplate fail: ", err)
	}
	for i := range m {
		where := parseStruct(m[i])
		parse.Execute(f, where)
	}
	f.Close()
}

func GenPgInsert(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	parse, err := template.New("").Parse(pgInsertTemplate)
	if err != nil {
		log.Fatalln("parse whereTemplate fail: ", err)
	}
	for i := range m {
		where := parseStruct(m[i])
		parse.Execute(f, where)
	}
	f.Close()
}

func GenPgAll(pkgName, path, fileName string, m ...interface{}) {
	preGen(pkgName, path)

	f, err := os.OpenFile(filepath.Join(path, fileName+".go"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("open file fail: ", err)
	}

	wt, _ := template.New("").Parse(pgWhereTemplate)
	st, _ := template.New("").Parse(pgSetTemplate)
	ct, _ := template.New("").Parse(columnTemplate)
	it, _ := template.New("").Parse(pgInsertTemplate)
	for i := range m {
		data := parseStruct(m[i])
		wt.Execute(f, data)
		st.Execute(f, data)
		ct.Execute(f, data)
		it.Execute(f, data)
	}
	f.Close()
}
