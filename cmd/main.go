package main

import (
    "fmt"
    "os"
    "govwallet-redemption/internal/repository"
    "govwallet-redemption/internal/service"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <staff_pass_id>")
        return
    }
    staffPassID := os.Args[1]

    staffRepo, err := repository.NewStaffRepository("data/staff_mapping.csv")
    if err != nil {
        fmt.Println("Error loading staff data:", err)
        return
    }

    redemptionRepo, err := repository.NewRedemptionRepository("data/redemptions.json")
    if err != nil {
        fmt.Println("Error loading redemption data:", err)
        return
    }

    svc := service.NewRedemptionService(staffRepo, redemptionRepo)
    err = svc.Redeem(staffPassID)
    if err != nil {
        fmt.Println("Redemption failed:", err)
    } else {
        fmt.Println("Redemption successful for team!")
    }
}
