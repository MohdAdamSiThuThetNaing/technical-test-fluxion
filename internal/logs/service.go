package logs

type Service struct {
	Repository *Repository
}

func NewService() *Service {

	return &Service{
		Repository: NewRepository(),
	}
}

func (s *Service) GetLogs() ([]Log, error) {
	return s.Repository.FindAll()
}