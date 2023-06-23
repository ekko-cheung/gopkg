package sqlgen

const (
	whereTemplate = `
func Where{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}) (string, []interface{}) {
	sqlBuild := strings.Builder{}
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

import "strings"

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
`
	setTemplate = `
func Set{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}) (string, []interface{}) {
	sqlBuild := strings.Builder{}
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

	return trimSql(sqlBuild.String()), params
}
`
	columnTemplate = `
func {{.FuncName}}Columns() string {
	return "{{range $index, $field := .Fields}}{{if eq $index $.FieldLength}}{{$field.SqlName}}{{else}}{{$field.SqlName}},{{end}}{{end}}"
}
`
	insertTemplate = `
func Insert{{.FuncName}}({{.ParamName}} *{{.ParamFullName}}, params []interface{}) (string, string, []interface{}){
	columns := strings.Builder{}
	values := strings.Builder{}
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

	return c, v, params
}
`
)
