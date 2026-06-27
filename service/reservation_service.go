package service

import (
	"errors"
	"spotsync-api/dto"
	"spotsync-api/repository"
)

type ReservationService interface {
	ReserveSpot(userID uint, req dto.CreateReservationRequest) (*dto.APIResponse, error)
	GetMyReservations(userID uint) ([]dto.MyReservationResponse, error)
	CancelReservation(userID uint, role string, reservationID uint) error
	GetAllReservations() (interface{}, error)
}

type reservationService struct {
	resRepo repository.ReservationRepository
}

func NewReservationService(repo repository.ReservationRepository) ReservationService {
	return &reservationService{resRepo: repo}
}

func (s *reservationService) ReserveSpot(userID uint, req dto.CreateReservationRequest) (*dto.APIResponse, error) {
	res, err := s.resRepo.CreateAtomicReservation(userID, req.ZoneID, req.LicensePlate)
	if err != nil {
		if err.Error() == "zone_full" {
			return nil, errors.New("parking zone is full")
		}
		return nil, err
	}
	return &dto.APIResponse{
		Success: true,
		Data:    res,
	}, nil
}

func (s *reservationService) GetMyReservations(userID uint) ([]dto.MyReservationResponse, error) {
	data, err := s.resRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	res := make([]dto.MyReservationResponse, len(data))
	for i, item := range data {
		res[i] = dto.MyReservationResponse{
			ID:           item.ID,
			LicensePlate: item.LicensePlate,
			Status:       item.Status,
			CreatedAt:    item.CreatedAt,
			Zone: dto.ZoneBrief{
				ID:   item.Zone.ID,
				Name: item.Zone.Name,
				Type: item.Zone.Type,
			},
		}
	}
	return res, nil
}

func (s *reservationService) CancelReservation(userID uint, role string, reservationID uint) error {
	res, err := s.resRepo.FindByID(reservationID)
	if err != nil {
		return errors.New("not_found")
	}
	if role != "admin" && res.UserID != userID {
		return errors.New("forbidden")
	}
	return s.resRepo.UpdateStatus(reservationID, "cancelled")
}

func (s *reservationService) GetAllReservations() (interface{}, error) {
	return s.resRepo.FindAllWithPreload()
}