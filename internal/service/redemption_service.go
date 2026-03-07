package service

import (
    "errors"
    "govwallet-redemption/internal/repository"
)

type RedemptionService struct {
    StaffRepo      *repository.StaffRepository
    RedemptionRepo *repository.RedemptionRepository
}

func NewRedemptionService(sr *repository.StaffRepository, rr *repository.RedemptionRepository) *RedemptionService {
    return &RedemptionService{
        StaffRepo:      sr,
        RedemptionRepo: rr,
    }
}

func (rs *RedemptionService) Redeem(staffPassID string) error {
    team, err := rs.StaffRepo.GetTeamByStaffPassID(staffPassID)
    if err != nil {
        return err
    }
    if rs.RedemptionRepo.HasRedeemed(team) {
        return errors.New("team already redeemed")
    }
    return rs.RedemptionRepo.AddRedemption(team)
}
