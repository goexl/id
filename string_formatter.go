package id

type stringFormatter interface {
	format(id *Id) string
}
