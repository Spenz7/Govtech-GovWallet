package tests

import (
    "govwallet-redemption/internal/repository"
    "govwallet-redemption/internal/service"
    "testing"
    "os"
)

func setup() *service.RedemptionService {
    // Ensure the data folder exists
    os.MkdirAll("data", os.ModePerm)

    // Create temporary CSV file
    csvFile := "data/test_staff.csv"
    os.WriteFile(csvFile, []byte("staff_pass_id,team_name,created_at\nS1,TeamA,0\nS2,TeamB,0"), 0644)

    staffRepo, err := repository.NewStaffRepository(csvFile)
    if err != nil {
        panic("failed to load staff repository: " + err.Error())
    }

    jsonFile := "data/test_redemptions.json"
    os.WriteFile(jsonFile, []byte("[]"), 0644)
    redemptionRepo, err := repository.NewRedemptionRepository(jsonFile)
    if err != nil {
        panic("failed to load redemption repository: " + err.Error())
    }

    return service.NewRedemptionService(staffRepo, redemptionRepo)
}

func TestRedeemSuccess(t *testing.T) {
    svc := setup()
    err := svc.Redeem("S1")
    if err != nil {
        t.Fatal("expected success, got error:", err)
    }
}

func TestRedeemDuplicate(t *testing.T) {
    svc := setup()
    _ = svc.Redeem("S1")
    err := svc.Redeem("S1")
    if err == nil {
        t.Fatal("expected error for duplicate redemption")
    }
}

func TestRedeemInvalidStaff(t *testing.T) {
    svc := setup()
    err := svc.Redeem("S999")
    if err == nil {
        t.Fatal("expected error for invalid staff ID")
    }
}
