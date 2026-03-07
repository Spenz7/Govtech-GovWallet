package tests

import (
    "govwallet-redemption/internal/model"
    "govwallet-redemption/internal/repository"
    "govwallet-redemption/internal/service"
    "testing"
    "os"
)

func setup() *service.RedemptionService {
    // Create temporary CSV file
    csvFile := "data/test_staff.csv"
    os.WriteFile(csvFile, []byte("staff_pass_id,team_name,created_at\nS1,TeamA,0\nS2,TeamB,0"), 0644)
    staffRepo, _ := repository.NewStaffRepository(csvFile)

    jsonFile := "data/test_redemptions.json"
    os.WriteFile(jsonFile, []byte("[]"), 0644)
    redemptionRepo, _ := repository.NewRedemptionRepository(jsonFile)

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
