package infra

type Pagination struct {
	CurrentPage int
	PerPage     int
}

func (p *Pagination) Offset() int {
	return (p.CurrentPage - 1) * p.PerPage
}

func (p *Pagination) Limit() int {
	return p.PerPage
}
