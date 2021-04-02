package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	if len(e) == 0 {
		return ""
	}

	es := e[field]

	if len(es) == 0 {
		return ""
	}

	return es[0]
}
