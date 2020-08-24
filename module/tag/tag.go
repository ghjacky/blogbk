package tag

import "blogbk/model"

func Get(t *model.STag) error {
	return t.Get()
}

func Fetch(ts *model.STags) error {
	return ts.FetchList()
}
