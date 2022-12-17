package service

func (s *Service) Test() error {

	return s.dao.CreateTestDemo()
}
