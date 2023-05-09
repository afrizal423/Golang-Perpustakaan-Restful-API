package buku

type BukuService struct {
	Repository IBukuRepository
}

func NewBukuService(Repository IBukuRepository) *BukuService {
	return &BukuService{
		Repository,
	}
}
