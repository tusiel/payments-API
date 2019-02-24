package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Payment is the overall payment model
type Payment struct {
	Type           string              `json:"type,omitempty" bson:"type,omitempty"`
	ID             *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Version        *uint16             `json:"version,omitempty" bson:"version,omitempty"`
	OrganisationID string              `json:"organisation_id,omitempty" bson:"organisation_id,omitempty"`
	Attributes     Attributes          `json:"attributes,omitempty" bson:"attributes,omitempty"`
}

// Attributes is the Attributes model
type Attributes struct {
	Amount               string             `json:"amount,omitempty" bson:"amount,omitempty"`
	BeneficiaryParty     BeneficiaryParty   `json:"beneficiary_party,omitempty" bson:"beneficiary_party,omitempty"`
	ChargesInformation   ChargesInformation `json:"charges_information,omitempty" bson:"charges_information,omitempty"`
	Currency             string             `json:"currency,omitempty" bson:"currency,omitempty"`
	DebtorParty          DebtorParty        `json:"debtor_party,omitempty" bson:"debtor_party,omitempty"`
	EndToEndReference    string             `json:"end_to_end_reference,omitempty" bson:"end_to_end_reference,omitempty"`
	FX                   FX                 `json:"fx,omitempty" bson:"fx,omitempty"`
	NumericReference     string             `json:"numeric_reference,omitempty" bson:"numeric_reference,omitempty"`
	PaymentID            string             `json:"payment_id,omitempty" bson:"payment_id,omitempty"`
	PaymentPurpose       string             `json:"payment_purpose,omitempty" bson:"payment_purpose,omitempty"`
	PaymentScheme        string             `json:"payment_scheme,omitempty" bson:"payment_scheme,omitempty"`
	PaymentType          string             `json:"payment_type,omitempty" bson:"payment_type,omitempty"`
	ProcessingDate       string             `json:"processing_date,omitempty" bson:"processing_date,omitempty"`
	Reference            string             `json:"reference,omitempty" bson:"reference,omitempty"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type,omitempty" bson:"scheme_payment_sub_type,omitempty"`
	SchemePaymentType    string             `json:"scheme_payment_type,omitempty" bson:"scheme_payment_type,omitempty"`
	SponsorParty         SponsorParty       `json:"sponsor_party,omitempty" bson:"sponsor_party,omitempty"`
}

// BeneficiaryParty is the BeneficiaryParty model
type BeneficiaryParty struct {
	AccountName       string  `json:"account_name,omitempty" bson:"account_name,omitempty"`
	AccountNumber     string  `json:"account_number,omitempty" bson:"account_number,omitempty"`
	AccountNumberCode string  `json:"account_number_code,omitempty" bson:"account_number_code,omitempty"`
	AccountType       *uint16 `json:"account_type,omitempty" bson:"account_type,omitempty"`
	Address           string  `json:"address,omitempty" bson:"address,omitempty"`
	BankID            string  `json:"bank_id,omitempty" bson:"bank_id,omitempty"`
	BankIDCode        string  `json:"bank_id_code,omitempty" bson:"bank_id_code,omitempty"`
	Name              string  `json:"name,omitempty" bson:"name,omitempty"`
}

// ChargesInformation is the ChargesInformation model
type ChargesInformation struct {
	BearerCode              string         `json:"bearer_code,omitempty" bson:"bearer_code,omitempty"`
	SenderCharges           []SenderCharge `json:"sender_charges,omitempty" bson:"sender_charges,omitempty"`
	ReceiverChargesAmount   string         `json:"receiver_charges_amount,omitempty" bson:"receiver_charges_amount,omitempty"`
	ReceiverChargesCurrency string         `json:"receiver_charges_currency,omitempty" bson:"receiver_charges_currency,omitempty"`
}

// SenderCharge is the SenderCharge model
type SenderCharge struct {
	Amount   string `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency string `json:"currency,omitempty" bson:"currency,omitempty"`
}

// DebtorParty is the DebtorParty model
type DebtorParty struct {
	AccountName       string `json:"account_name,omitempty" bson:"account_name,omitempty"`
	AccountNumber     string `json:"account_number,omitempty" bson:"account_number,omitempty"`
	AccountNumberCode string `json:"account_number_code,omitempty" bson:"account_number_code,omitempty"`
	Address           string `json:"address,omitempty" bson:"address,omitempty"`
	BankID            string `json:"bank_id,omitempty" bson:"bank_id,omitempty"`
	BankIDCode        string `json:"bank_id_code,omitempty" bson:"bank_id_code,omitempty"`
	Name              string `json:"name,omitempty" bson:"name,omitempty"`
}

// FX is the FX model
type FX struct {
	ContractReference string `json:"contract_reference,omitempty" bson:"contract_reference,omitempty"`
	ExchangeRate      string `json:"exchange_rate,omitempty" bson:"exchange_rate,omitempty"`
	OriginalAmount    string `json:"original_amount,omitempty" bson:"original_amount,omitempty"`
	OriginalCurrency  string `json:"original_currency,omitempty" bson:"original_currency,omitempty"`
}

// SponsorParty is the SponsorParty model
type SponsorParty struct {
	AccountNumber string `json:"account_number,omitempty" bson:"end_to_end_reference,omitempty"`
	BankID        string `json:"bank_id,omitempty" bson:"bank_id,omitempty"`
	BankIDCode    string `json:"bank_id_code,omitempty" bson:"bank_id_code,omitempty"`
}
