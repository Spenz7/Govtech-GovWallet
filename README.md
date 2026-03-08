# GovWallet Redemption

## Project Description

This project simulates the redemption of gifts for teams in a department, using staff pass IDs for verification.
It is implemented in **Golang** and consists of:

- A **CLI** to redeem gifts for a given staff pass ID.
- **Unit tests** to verify the redemption logic.
- Data stored in CSV (`staff_mapping.csv`) and JSON (`redemptions.json`) files.

---

## Project Structure

```
govwallet-redemption/
├─ cmd/
│  └─ main.go
├─ internal/
│  ├─ model/
│  ├─ repository/
│  └─ service/
├─ data/
│  ├─ staff_mapping.csv
│  └─ redemptions.json
├─ tests/
│  └─ redemption_service_test.go
├─ go.mod
└─ README.md
```

---

## Prerequisites

- Golang installed (version 1.20+ recommended)
- Git (optional, for cloning repo)

---

## How to Run CLI

1. Open terminal or PowerShell in the project root.

2. Make sure `data/staff_mapping.csv` exists and has sample staff IDs:

```csv
staff_pass_id,team_name,created_at
S1,TeamA,1677744000000
S2,TeamB,1677744000000
```

3. Run the CLI with a staff pass ID:

```bash
go run ./cmd/main.go <staff_pass_id>
```

Example:

```bash
go run ./cmd/main.go S1
```

Output if successful:

```text
Redemption successful for team TeamA!
```

Output if the team already redeemed:

```text
Redemption failed: team already redeemed
```

Output if staff pass ID is not found:

```text
Redemption failed: staff pass ID not found
```

---

## How to Run Unit Tests

From the project root, run:

```bash
go test ./...
```

This runs all automated unit tests in `tests/`.

Expected output if all tests pass:

```text
ok      govwallet-redemption/tests 0.652s
```

The tests cover:

-   Successful redemption

-   Duplicate redemption

-   Invalid staff pass ID

* * * * *

Assumptions / Notes
-------------------

-   Redemption data is stored in `redemptions.json` as:

```json

[
  {
    "team_name": "TeamA",
    "redeemed_at": 1677744000000
  }
]
```

-   Staff data is stored in CSV format with headers:

```text

staff_pass_id,team_name,created_at
```

-   `created_at` is in epoch milliseconds. For unit tests, it is set to `0` for simplicity.

-   Each team can only redeem once.

-   CLI and unit tests are self-contained --- tests use `test_staff.csv` and `test_redemptions.json` in `data/`.

* * * * *

Running Notes for Reviewers
---------------------------

-   Run `go test ./...` to verify automated unit tests.

-   Run `go run ./cmd/main.go <staff_pass_id>` to test CLI functionality.

-   No external dependencies are required other than Golang.

* * * * *

Author / Submission
-------------------

-   Implemented by: Spencer Ng

-   Submitted to: GovWallet / GDP Team
