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

const (
	whereTemplate = `
func Where{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}) (string, []interface{}) {
	sqlBuild := strings.Builder{}
	sqlBuild.Grow(50)
	{{range $field := .Fields}}
	{{- if eq $field.Type "string" -}}
	if {{$.ParamName}}.{{$field.Name}} != "" {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- end -}}
	{{- end -}}

	return trimSql(sqlBuild.String()), params
}
`
	trimSqlTemplate = `
package {{.}}

import (
	"fmt"
	"strings"
)

func trimSql(s string) string {
	if s == "" {
		return s
	}
	if strings.HasPrefix(s, "AND") {
		return s[3:]
	} else if strings.HasPrefix(s, "OR") {
		return s[2:]
	} else if strings.HasPrefix(s, ",") {
		return s[1:]
	} else if strings.HasSuffix(s, ",") {
		return s[:len(s)-1]
	}

	return s
}

func pgPlaceholder(i int) string {
	return fmt.Sprintf("$%d,", i)
}

func pgPlaceholder2(i int) string {
	return fmt.Sprintf("$%d ", i)
}

func pgPlaceholder3(i int) string {
	return fmt.Sprintf("$%d", i)
}
`
	setTemplate = `
func Set{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}) (string, []interface{}) {
	sqlBuild := strings.Builder{}
	sqlBuild.Grow(50)
	{{range $field := .Fields}}
	{{- if eq $field.Type "string" -}}
	if {{$.ParamName}}.{{$field.Name}} != "" {
		sqlBuild.WriteString(", {{$field.SqlName}} = ?")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ?")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ? ")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- end -}}
	{{- end -}}

	return "UPDATE {{.TableName}} SET " + trimSql(sqlBuild.String()), params
}
`
	columnTemplate = `
func {{.FuncName}}Columns() string {
	return "{{range $index, $field := .Fields}}{{if eq $index $.FieldLength}}{{$field.SqlName}}{{else}}{{$field.SqlName}},{{end}}{{end}}"
}
`
	insertTemplate = `
func Insert{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}) (string, []interface{}){
	columns := strings.Builder{}
	columns.Grow(50)
	columns.WriteString("INSERT INTO {{.TableName}}")
	values := strings.Builder{}
	values.Grow(50)
	columns.WriteString("(")
	values.WriteString("(")
{{range $field := .Fields}}

{{- if eq $field.Type "string"}}
	if {{$.ParamName}}.{{$field.Name}} != "" {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString("?,")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "int64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString("?,")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "int32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString("?,")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "int"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString("?,")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "uint64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString("?,")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "uint32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString("?,")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "uint"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString("?,")
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- end -}}
{{end}}
	c := trimSql(columns.String()) + ")"
	v := trimSql(values.String()) + ")"

	return c + "VALUES" + v, params
}
`
)
