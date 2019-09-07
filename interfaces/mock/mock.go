package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(string) string {
	return r.Contents
}
