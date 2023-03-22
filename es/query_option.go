package es

type Option func(body *reqBody)

func ShouldMatch(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Bool.Should = append(body.Query.Bool.Should, match{Match: map[string]interface{}{field: val}})
	}
}

func MustMatch(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Bool.Must = append(body.Query.Bool.Must, match{Match: map[string]interface{}{field: val}})
	}
}

func MustNotMatch(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Bool.MustNot = append(body.Query.Bool.MustNot, match{Match: map[string]interface{}{field: val}})
	}
}

func ShouldMatchPhrase(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Bool.Should = append(body.Query.Bool.Should, matchPhrase{MatchPhrase: map[string]interface{}{field: val}})
	}
}

func MustMatchPhrase(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Bool.Must = append(body.Query.Bool.Must, matchPhrase{MatchPhrase: map[string]interface{}{field: val}})
	}
}

func MustNotMatchPhrase(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Bool.MustNot = append(body.Query.Bool.MustNot, matchPhrase{MatchPhrase: map[string]interface{}{field: val}})
	}
}

func Page(from int64, size int64) Option {
	return func(body *reqBody) {
		body.From = (from - 1) * size
		body.Size = size
	}
}

func Source(fields ...string) Option {
	return func(body *reqBody) {
		body.Source = append(body.Source, fields...)
	}
}

func TermFilter(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Filter = append(body.Query.Filter, map[string]interface{}{"term": map[string]interface{}{field: val}})
	}
}

func RangeFilter(field string, val interface{}) Option {
	return func(body *reqBody) {
		body.Query.Filter = append(body.Query.Filter, map[string]interface{}{"range": map[string]interface{}{field: val}})
	}
}

func Sort(field string, s sorter) Option {
	return func(body *reqBody) {
		body.Sort = append(body.Sort, map[string]sort{field: {Order: s}})
	}
}

func HighlightField(fields ...string) Option {
	return func(body *reqBody) {
		for i := range fields {
			body.Highlight.Fields[fields[i]] = struct{}{}
		}
	}
}

func PreTags(tags ...string) Option {
	return func(body *reqBody) {
		body.Highlight.PreTags = append(body.Highlight.PreTags, tags...)
	}
}

func PostTags(tags ...string) Option {
	return func(body *reqBody) {
		body.Highlight.PostTags = append(body.Highlight.PostTags, tags...)
	}
}
