package hx

const (
	idPrefix = "hx"
)

func Id(id string) string {
	return idPrefix + "-" + id
}

func HashId(id string) string {
	return "#" + Id(id)
}
