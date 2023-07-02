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
	pgInsertTemplate = `
func Insert{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}) (string, []interface{}){
	columns := strings.Builder{}
	columns.Grow(50)
	columns.WriteString("INSERT INTO {{.TableName}}")
	values := strings.Builder{}
	values.Grow(50)
	columns.WriteString("(")
	values.WriteString("(")
	i := 1
{{range $field := .Fields}}

{{- if eq $field.Type "string"}}
	if {{$.ParamName}}.{{$field.Name}} != "" {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString(pgPlaceholder(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "int64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString(pgPlaceholder(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "int32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString(pgPlaceholder(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "int"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString(pgPlaceholder(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "uint64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString(pgPlaceholder(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "uint32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString(pgPlaceholder(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- else if eq $field.Type "uint"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		columns.WriteString("{{$field.SqlName}},")
		values.WriteString(pgPlaceholder(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- end -}}
{{end}}
	c := trimSql(columns.String()) + ")"
	v := trimSql(values.String()) + ")"

	return c + "VALUES" + v, params
}
`
	pgWhereTemplate = `
func Where{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}, i int) (string, []interface{}, int) {
	sqlBuild := strings.Builder{}
	sqlBuild.Grow(50)
	{{range $field := .Fields}}
	{{- if eq $field.Type "string" -}}
	if {{$.ParamName}}.{{$field.Name}} != "" {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder2(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder2(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder2(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder2(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder2(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder2(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString("AND {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder2(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- end -}}
	{{- end -}}

	return trimSql(sqlBuild.String()), params, i
}
`
	pgSetTemplate = `
func Set{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}, i int) (string, []interface{}, int) {
	sqlBuild := strings.Builder{}
	sqlBuild.Grow(50)
	{{range $field := .Fields}}
	{{- if eq $field.Type "string" -}}
	if {{$.ParamName}}.{{$field.Name}} != "" {
		sqlBuild.WriteString(", {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder3(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder3(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder3(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "int"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder3(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint64"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder3(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint32"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder3(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{else if eq $field.Type "uint"}}
	if {{$.ParamName}}.{{$field.Name}} != 0 {
		sqlBuild.WriteString(", {{$field.SqlName}} = ")
		sqlBuild.WriteString(pgPlaceholder3(i))
		i++
		params = append(params, {{$.ParamName}}.{{$field.Name}})
	}
	{{- end -}}
	{{- end -}}

	return "UPDATE {{.TableName}} SET" + trimSql(sqlBuild.String()), params, i
}
`
)
