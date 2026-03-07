package repository

import (
    "encoding/csv"
    "errors"
    "os"
    "strconv"
    "govwallet-redemption/internal/model"
)

type StaffRepository struct {
    StaffMap map[string]string
}

func NewStaffRepository(csvPath string) (*StaffRepository, error) {
    f, err := os.Open(csvPath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    r := csv.NewReader(f)
    rows, err := r.ReadAll()
    if err != nil {
        return nil, err
    }

    staffMap := make(map[string]string)
    for i, row := range rows {
        if i == 0 {
            continue // skip header
        }
        staffPass := row[0]
        team := row[1]
        staffMap[staffPass] = team
        // created_at is ignored for now
        _, _ = strconv.ParseInt(row[2], 10, 64)
    }

    return &StaffRepository{StaffMap: staffMap}, nil
}

func (sr *StaffRepository) GetTeamByStaffPassID(id string) (string, error) {
    team, ok := sr.StaffMap[id]
    if !ok {
        return "", errors.New("staff pass ID not found")
    }
    return team, nil
}
