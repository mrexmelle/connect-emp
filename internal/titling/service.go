package titling

import (
	"database/sql"
	"time"

	"github.com/mrexmelle/connect-emp/internal/config"
	"github.com/mrexmelle/connect-emp/internal/localerror"
)

type Service struct {
	ConfigService     *config.Service
	TitlingRepository Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
) *Service {
	return &Service{
		ConfigService:     cfg,
		TitlingRepository: r,
	}
}

func (s *Service) Create(req PostRequestDto) (*ViewEntity, error) {
	sd, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	var ed sql.NullTime
	if req.EndDate == "" {
		ed.Valid = false
	} else {
		ed.Time, err = time.Parse("2006-01-02", req.EndDate)
		ed.Valid = (err == nil)
	}

	result, err := s.TitlingRepository.Create(&Entity{
		Ehid:      req.Ehid,
		StartDate: sd,
		EndDate:   ed,
		Title:     req.Title,
	})
	if err != nil {
		return nil, err
	}
	return toViewEntity(result), nil
}

func (s *Service) RetrieveById(id int) (*ViewEntity, error) {
	result, err := s.TitlingRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return toViewEntity(result), nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id int) error {
	return s.TitlingRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id int) error {
	err := s.TitlingRepository.DeleteById(id)
	return err
}

func (s *Service) RetrieveByEhidOrderByStartDate(ehid string, orderDir string) ([]ViewEntity, error) {
	if orderDir != OrderAsc && orderDir != OrderDesc && orderDir != OrderNone {
		return []ViewEntity{}, localerror.ErrBadQueryParam
	}
	result, err := s.TitlingRepository.FindByEhidOrderByStartDate(ehid, orderDir)
	if err != nil {
		return []ViewEntity{}, err
	}
	return toViewEntitySlice(result), nil
}

func (s *Service) RetrieveCurrentByNodeId(nodeId string) ([]ViewEntity, error) {
	result, err := s.TitlingRepository.FindCurrentByNodeId(nodeId)
	if err != nil {
		return []ViewEntity{}, err
	}
	return toViewEntitySlice(result), nil
}

func (s *Service) RetrieveCurrentByEhid(ehid string) ([]ViewEntity, error) {
	result, err := s.TitlingRepository.FindCurrentByEhid(ehid)
	if err != nil {
		return []ViewEntity{}, err
	}
	return toViewEntitySlice(result), nil
}
