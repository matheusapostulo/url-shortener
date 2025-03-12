package port

type Dependencies interface {
	BuildDependencies() (any, error)
}
