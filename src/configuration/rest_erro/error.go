package manager_error

type Error interface {
	Error() string
	GetCauses() string
}
