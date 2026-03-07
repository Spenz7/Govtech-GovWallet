package repository

import (
    "encoding/json"
    "errors"
    "os"
    "time"
    "govwallet-redemption/internal/model"
)

type RedemptionRepository struct {
    FilePath   string
    Redemptions []model.Redemption
}

func NewRedemptionRepository(filePath string) (*RedemptionRepository, error) {
    repo := &RedemptionRepository{FilePath: filePath}
    f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    decoder := json.NewDecoder(f)
    err = decoder.Decode(&repo.Redemptions)
    if err != nil && err.Error() != "EOF" {
        return nil, err
    }
    return repo, nil
}

func (rr *RedemptionRepository) HasRedeemed(team string) bool {
    for _, r := range rr.Redemptions {
        if r.TeamName == team {
            return true
        }
    }
    return false
}

func (rr *RedemptionRepository) AddRedemption(team string) error {
    if rr.HasRedeemed(team) {
        return errors.New("team already redeemed")
    }
    redemption := model.Redemption{
        TeamName:   team,
        RedeemedAt: time.Now().UnixMilli(),
    }
    rr.Redemptions = append(rr.Redemptions, redemption)
    f, err := os.OpenFile(rr.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        return err
    }
    defer f.Close()
    encoder := json.NewEncoder(f)
    return encoder.Encode(rr.Redemptions)
}
