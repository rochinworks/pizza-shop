package handler

type handler struct {
	Repo Repository
}

// NewHanlder take a Repository interface
func NewHandler(r Repository) *handler {
	return &handler{Repo: r}
}
