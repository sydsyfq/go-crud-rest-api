package handler

import (
	"database/sql"
	"time"

	"github.com/alphaloan/vehicle/datastore"
	"github.com/google/uuid"
)

type LoanCustomer struct {
	CustomerID		string	`json:"customer_id"`
	IDCardNumber	string	`json:"id_card_number"`
	FullName		string	`json:"full_name"`
	BirthDate		string	`json:"birth_date"`
	PhoneNumber		string	`json:"phone_number"`
	Email			string	`json:"email"`
	MonthlyIncome	string	`json:"monthly_income"`
	AddressStreet	string	`json:"address_street"`
	AddressCity		string	`json:"address_city"`
}

type LoanSubmission struct {
    SubmissionID            string `json:"submission_id"`
    VehicleType             string `json:"vehicle_type"`
    VehicleBrand            string `json:"vehicle_brand"`
    VehicleModel            string `json:"vehicle_model"`
    VehicleLicenseNumber    string `json:"vehicle_license_number"`
    VehicleOdometer         int    `json:"vehicle_odometer"`
    ManufacturingYear       int    `json:"manufacturing_year"`
    ProposedLoanAmount      int    `json:"proposed_loan_amount"`
    ProposedLoanTenureMonth int    `json:"proposed_loan_tenure_month"`
    IsCommercialVehicle     bool   `json:"is_commercial_vehicle"`
}
 
type LoanSubmitRequest struct {
    Customer     LoanCustomer   `json:"customer"`
    ProposedLoan LoanSubmission `json:"proposed_loan"`
}
 
type LoanSubmitResponse struct {
    CustomerID   *string `json:"customer_id"`
    SubmissionID *string `json:"submission_id"`
}

func converLoanCustomer(loanCustomer *LoanCustomer) *datastore.LoanCustomerRow {
	if loanCustomer == nil {
		return nil
	}

	parsedEmail := ""

	if loanCustomer.Email != nil {
		parsedEmail = *loanCustomer.Email
	}

	return &datastore.LoanCustomerRow {
		CustomerID:	uuid.New().String(),
		IDCardNumber:	loanCustomer.IDCardNumber,
		FullName:		loanCustomer.FullName,
		BirthDate:		loanCustomer.BirthDate,
		PhoneNumber:	loanCustomer.PhoneNumber,
		Email:	sql.NullString {
			String: parsedEmail,
			Valid:	loanCustomer.Email != nil && *loanCustomer.Email != "",
		},
		MonthlyIncome:	loanCustomer.MonthlyIncome,
		AddressStreet:	loanCustomer.AddressStreet,
		AddressCity:	loanCustomer.AddressCity,
	}
}