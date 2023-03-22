package es

func Query(options ...Option) *reqBody {
	r := &reqBody{
		Query: &query{
			Bool: &boolean{},
		},
		Highlight: &highlight{Fields: make(map[string]interface{})},
	}

	for i := range options {
		options[i](r)
	}

	return r
}

type reqBody struct {
	From      int64             `json:"from,omitempty"`
	Size      int64             `json:"size,omitempty"`
	Query     *query            `json:"query,omitempty"`
	Source    []string          `json:"source,omitempty"`
	Sort      []map[string]sort `json:"sort,omitempty"`
	Highlight *highlight        `json:"highlight,omitempty"`
}

type query struct {
	Bool   *boolean                 `json:"bool,omitempty"`
	Filter []map[string]interface{} `json:"filter,omitempty"`
}

type boolean struct {
	Should  []interface{} `json:"should,omitempty"`
	Must    []interface{} `json:"must,omitempty"`
	MustNot []interface{} `json:"must_not,omitempty"`
}

type match struct {
	Match map[string]interface{} `json:"match,omitempty"`
}

type matchPhrase struct {
	MatchPhrase map[string]interface{} `json:"match_phrase,omitempty"`
}

type highlight struct {
	PreTags  []string               `json:"pre_tags,omitempty"`
	PostTags []string               `json:"post_tags,omitempty"`
	Fields   map[string]interface{} `json:"fields,omitempty"`
}

type sorter string

type sort struct {
	Order sorter `json:"order,omitempty"`
}

const (
	ASC  = sorter("asc")
	DESC = sorter("desc")
)
