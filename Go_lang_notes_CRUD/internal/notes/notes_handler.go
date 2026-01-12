package notes

type Handler struct {
	repo *Repo
}

func NewHandler(*repo) *Handler {
	return &Handler{repo: repo}
}