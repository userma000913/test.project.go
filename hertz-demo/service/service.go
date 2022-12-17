package service

import (
	"hertz_demo/conf"
	"hertz_demo/dao"
	"hertz_demo/manager"
)

type Service struct {
	c *conf.AppConfig

	dao *dao.Dao
	mgr *manager.Manager
}

func New(c *conf.AppConfig) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c),
		mgr: manager.New(c),
	}
}

func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}

	if s.mgr != nil {
		s.mgr.Close()
	}
}
