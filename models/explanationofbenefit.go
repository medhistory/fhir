// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ExplanationOfBenefit struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                                               `bson:"status,omitempty" json:"status,omitempty"`
	Type                  *CodeableConcept                                     `bson:"type,omitempty" json:"type,omitempty"`
	SubType               *CodeableConcept                                     `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                   string                                               `bson:"use,omitempty" json:"use,omitempty"`
	Patient               *Reference                                           `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod        *Period                                              `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created               *FHIRDateTime                                        `bson:"created,omitempty" json:"created,omitempty"`
	Enterer               *Reference                                           `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer               *Reference                                           `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider              *Reference                                           `bson:"provider,omitempty" json:"provider,omitempty"`
	Priority              *CodeableConcept                                     `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserveRequested *CodeableConcept                                     `bson:"fundsReserveRequested,omitempty" json:"fundsReserveRequested,omitempty"`
	FundsReserve          *CodeableConcept                                     `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Related               []ExplanationOfBenefitRelatedClaimComponent          `bson:"related,omitempty" json:"related,omitempty"`
	Prescription          *Reference                                           `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription  *Reference                                           `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                 *ExplanationOfBenefitPayeeComponent                  `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral              *Reference                                           `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility              *Reference                                           `bson:"facility,omitempty" json:"facility,omitempty"`
	Claim                 *Reference                                           `bson:"claim,omitempty" json:"claim,omitempty"`
	ClaimResponse         *Reference                                           `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	Outcome               string                                               `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition           string                                               `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PreAuthRef            []string                                             `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	PreAuthRefPeriod      []Period                                             `bson:"preAuthRefPeriod,omitempty" json:"preAuthRefPeriod,omitempty"`
	CareTeam              []ExplanationOfBenefitCareTeamComponent              `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	SupportingInfo        []ExplanationOfBenefitSupportingInformationComponent `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Diagnosis             []ExplanationOfBenefitDiagnosisComponent             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure             []ExplanationOfBenefitProcedureComponent             `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Precedence            *uint32                                              `bson:"precedence,omitempty" json:"precedence,omitempty"`
	Insurance             []ExplanationOfBenefitInsuranceComponent             `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident              *ExplanationOfBenefitAccidentComponent               `bson:"accident,omitempty" json:"accident,omitempty"`
	Item                  []ExplanationOfBenefitItemComponent                  `bson:"item,omitempty" json:"item,omitempty"`
	AddItem               []ExplanationOfBenefitAddedItemComponent             `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Adjudication          []ExplanationOfBenefitAdjudicationComponent          `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Total                 []ExplanationOfBenefitTotalComponent                 `bson:"total,omitempty" json:"total,omitempty"`
	Payment               *ExplanationOfBenefitPaymentComponent                `bson:"payment,omitempty" json:"payment,omitempty"`
	FormCode              *CodeableConcept                                     `bson:"formCode,omitempty" json:"formCode,omitempty"`
	Form                  *Attachment                                          `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote           []ExplanationOfBenefitNoteComponent                  `bson:"processNote,omitempty" json:"processNote,omitempty"`
	BenefitPeriod         *Period                                              `bson:"benefitPeriod,omitempty" json:"benefitPeriod,omitempty"`
	BenefitBalance        []ExplanationOfBenefitBenefitBalanceComponent        `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ExplanationOfBenefit"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ExplanationOfBenefit), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ExplanationOfBenefit) GetBSON() (interface{}, error) {
	x.ResourceType = "ExplanationOfBenefit"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "explanationOfBenefit" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type explanationOfBenefit ExplanationOfBenefit

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ExplanationOfBenefit) UnmarshalJSON(data []byte) (err error) {
	x2 := explanationOfBenefit{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ExplanationOfBenefit(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ExplanationOfBenefit) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ExplanationOfBenefit"
	} else if x.ResourceType != "ExplanationOfBenefit" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ExplanationOfBenefit, instead received %s", x.ResourceType))
	}
	return nil
}

type ExplanationOfBenefitRelatedClaimComponent struct {
	BackboneElement `bson:",inline"`
	Claim           *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship    *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference       *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ExplanationOfBenefitPayeeComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Party           *Reference       `bson:"party,omitempty" json:"party,omitempty"`
}

type ExplanationOfBenefitCareTeamComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider        *Reference       `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible     *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role            *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification   *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}

type ExplanationOfBenefitSupportingInformationComponent struct {
	BackboneElement `bson:",inline"`
	Sequence        *uint32          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	TimingDate      *FHIRDateTime    `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod    *Period          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueBoolean    *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueString     string           `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity   *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueAttachment *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueReference  *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Reason          *Coding          `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ExplanationOfBenefitDiagnosisComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept  `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference        `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `bson:"onAdmission,omitempty" json:"onAdmission,omitempty"`
	PackageCode              *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}

type ExplanationOfBenefitProcedureComponent struct {
	BackboneElement          `bson:",inline"`
	Sequence                 *uint32           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type                     []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Date                     *FHIRDateTime     `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCodeableConcept *CodeableConcept  `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference        `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
	Udi                      []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ExplanationOfBenefitInsuranceComponent struct {
	BackboneElement `bson:",inline"`
	Focal           *bool      `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage        *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	PreAuthRef      []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}

type ExplanationOfBenefitAccidentComponent struct {
	BackboneElement   `bson:",inline"`
	Date              *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}

type ExplanationOfBenefitItemComponent struct {
	BackboneElement         `bson:",inline"`
	Sequence                *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamSequence        []uint32                                    `bson:"careTeamSequence,omitempty" json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []uint32                                    `bson:"diagnosisSequence,omitempty" json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []uint32                                    `bson:"procedureSequence,omitempty" json:"procedureSequence,omitempty"`
	InformationSequence     []uint32                                    `bson:"informationSequence,omitempty" json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept                            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category                *CodeableConcept                            `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService        *CodeableConcept                            `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier                []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                           `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *FHIRDateTime                               `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                     `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                            `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                                    `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference                                  `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi                     []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite                *CodeableConcept                            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept                           `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter               []Reference                                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	NoteNumber              []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication            []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitDetailComponent       `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ExplanationOfBenefitAdjudicationComponent struct {
	BackboneElement `bson:",inline"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Reason          *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount          *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
	Value           *float64         `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitDetailComponent struct {
	BackboneElement  `bson:",inline"`
	Sequence         *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue          *CodeableConcept                            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category         *CodeableConcept                            `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService *CodeableConcept                            `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier         []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode      []CodeableConcept                           `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity         *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice        *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor           *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net              *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi              []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber       []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication     []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail        []ExplanationOfBenefitSubDetailComponent    `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ExplanationOfBenefitSubDetailComponent struct {
	BackboneElement  `bson:",inline"`
	Sequence         *uint32                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue          *CodeableConcept                            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category         *CodeableConcept                            `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService *CodeableConcept                            `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier         []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode      []CodeableConcept                           `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity         *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice        *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor           *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net              *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	Udi              []Reference                                 `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber       []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication     []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitAddedItemComponent struct {
	BackboneElement         `bson:",inline"`
	ItemSequence            []uint32                                       `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	DetailSequence          []uint32                                       `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	SubDetailSequence       []uint32                                       `bson:"subDetailSequence,omitempty" json:"subDetailSequence,omitempty"`
	Provider                []Reference                                    `bson:"provider,omitempty" json:"provider,omitempty"`
	ProductOrService        *CodeableConcept                               `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier                []CodeableConcept                              `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                              `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *FHIRDateTime                                  `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                        `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                               `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                                       `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference                                     `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity                                      `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Quantity                                      `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *float64                                       `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Quantity                                      `bson:"net,omitempty" json:"net,omitempty"`
	BodySite                *CodeableConcept                               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept                              `bson:"subSite,omitempty" json:"subSite,omitempty"`
	NoteNumber              []uint32                                       `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication            []ExplanationOfBenefitAdjudicationComponent    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitAddedItemDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ExplanationOfBenefitAddedItemDetailComponent struct {
	BackboneElement  `bson:",inline"`
	ProductOrService *CodeableConcept                                        `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier         []CodeableConcept                                       `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity         *Quantity                                               `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice        *Quantity                                               `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor           *float64                                                `bson:"factor,omitempty" json:"factor,omitempty"`
	Net              *Quantity                                               `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber       []uint32                                                `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication     []ExplanationOfBenefitAdjudicationComponent             `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail        []ExplanationOfBenefitAddedItemDetailSubDetailComponent `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ExplanationOfBenefitAddedItemDetailSubDetailComponent struct {
	BackboneElement  `bson:",inline"`
	ProductOrService *CodeableConcept                            `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier         []CodeableConcept                           `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity         *Quantity                                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice        *Quantity                                   `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor           *float64                                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net              *Quantity                                   `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber       []uint32                                    `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication     []ExplanationOfBenefitAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitTotalComponent struct {
	BackboneElement `bson:",inline"`
	Category        *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Amount          *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}

type ExplanationOfBenefitPaymentComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Adjustment       *Quantity        `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date             *FHIRDateTime    `bson:"date,omitempty" json:"date,omitempty"`
	Amount           *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
	Identifier       *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}

type ExplanationOfBenefitNoteComponent struct {
	BackboneElement `bson:",inline"`
	Number          *uint32          `bson:"number,omitempty" json:"number,omitempty"`
	Type            string           `bson:"type,omitempty" json:"type,omitempty"`
	Text            string           `bson:"text,omitempty" json:"text,omitempty"`
	Language        *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}

type ExplanationOfBenefitBenefitBalanceComponent struct {
	BackboneElement `bson:",inline"`
	Category        *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Excluded        *bool                                  `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name            string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Network         *CodeableConcept                       `bson:"network,omitempty" json:"network,omitempty"`
	Unit            *CodeableConcept                       `bson:"unit,omitempty" json:"unit,omitempty"`
	Term            *CodeableConcept                       `bson:"term,omitempty" json:"term,omitempty"`
	Financial       []ExplanationOfBenefitBenefitComponent `bson:"financial,omitempty" json:"financial,omitempty"`
}

type ExplanationOfBenefitBenefitComponent struct {
	BackboneElement    `bson:",inline"`
	Type               *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	AllowedUnsignedInt *uint32          `bson:"allowedUnsignedInt,omitempty" json:"allowedUnsignedInt,omitempty"`
	AllowedString      string           `bson:"allowedString,omitempty" json:"allowedString,omitempty"`
	AllowedMoney       *Quantity        `bson:"allowedMoney,omitempty" json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *uint32          `bson:"usedUnsignedInt,omitempty" json:"usedUnsignedInt,omitempty"`
	UsedMoney          *Quantity        `bson:"usedMoney,omitempty" json:"usedMoney,omitempty"`
}

type ExplanationOfBenefitPlus struct {
	ExplanationOfBenefit                     `bson:",inline"`
	ExplanationOfBenefitPlusRelatedResources `bson:",inline"`
}

type ExplanationOfBenefitPlusRelatedResources struct {
	IncludedCoverageResourcesReferencedByCoverage                          *[]Coverage                   `bson:"_includedCoverageResourcesReferencedByCoverage,omitempty"`
	IncludedPractitionerResourcesReferencedByCareteam                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByCareteam,omitempty"`
	IncludedOrganizationResourcesReferencedByCareteam                      *[]Organization               `bson:"_includedOrganizationResourcesReferencedByCareteam,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByCareteam                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByCareteam,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByPayee                         *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPayee,omitempty"`
	IncludedOrganizationResourcesReferencedByPayee                         *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPayee,omitempty"`
	IncludedPatientResourcesReferencedByPayee                              *[]Patient                    `bson:"_includedPatientResourcesReferencedByPayee,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPayee                     *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPayee,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPayee                        *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPayee,omitempty"`
	IncludedPractitionerResourcesReferencedByProvider                      *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedOrganizationResourcesReferencedByProvider                      *[]Organization               `bson:"_includedOrganizationResourcesReferencedByProvider,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByProvider                  *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedDeviceResourcesReferencedByDetailudi                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByDetailudi,omitempty"`
	IncludedClaimResourcesReferencedByClaim                                *[]Claim                      `bson:"_includedClaimResourcesReferencedByClaim,omitempty"`
	IncludedPractitionerResourcesReferencedByEnterer                       *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByEnterer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByEnterer                   *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByEnterer,omitempty"`
	IncludedDeviceResourcesReferencedByProcedureudi                        *[]Device                     `bson:"_includedDeviceResourcesReferencedByProcedureudi,omitempty"`
	IncludedDeviceResourcesReferencedBySubdetailudi                        *[]Device                     `bson:"_includedDeviceResourcesReferencedBySubdetailudi,omitempty"`
	IncludedLocationResourcesReferencedByFacility                          *[]Location                   `bson:"_includedLocationResourcesReferencedByFacility,omitempty"`
	IncludedDeviceResourcesReferencedByItemudi                             *[]Device                     `bson:"_includedDeviceResourcesReferencedByItemudi,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor              *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof               *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                    *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref              *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                             *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                        *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                      *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                       *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                   *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource          *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                         *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                    *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                   *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor           *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof            *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2        *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                     *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                    *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor           *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof            *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2        *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                             *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                           *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                     *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                      *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                        *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                 *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                      *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                 *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject            *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedCoverageResourceReferencedByCoverage() (coverage *Coverage, err error) {
	if e.IncludedCoverageResourcesReferencedByCoverage == nil {
		err = errors.New("Included coverages not requested")
	} else if len(*e.IncludedCoverageResourcesReferencedByCoverage) > 1 {
		err = fmt.Errorf("Expected 0 or 1 coverage, but found %d", len(*e.IncludedCoverageResourcesReferencedByCoverage))
	} else if len(*e.IncludedCoverageResourcesReferencedByCoverage) == 1 {
		coverage = &(*e.IncludedCoverageResourcesReferencedByCoverage)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByCareteam() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByCareteam == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByCareteam))
	} else if len(*e.IncludedPractitionerResourcesReferencedByCareteam) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByCareteam)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByCareteam() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByCareteam == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByCareteam))
	} else if len(*e.IncludedOrganizationResourcesReferencedByCareteam) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByCareteam)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByCareteam() (practitionerRole *PractitionerRole, err error) {
	if e.IncludedPractitionerRoleResourcesReferencedByCareteam == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByCareteam) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*e.IncludedPractitionerRoleResourcesReferencedByCareteam))
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByCareteam) == 1 {
		practitionerRole = &(*e.IncludedPractitionerRoleResourcesReferencedByCareteam)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedEncounterResourcesReferencedByEncounter() (encounters []Encounter, err error) {
	if e.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else {
		encounters = *e.IncludedEncounterResourcesReferencedByEncounter
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPayee() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByPayee == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByPayee))
	} else if len(*e.IncludedPractitionerResourcesReferencedByPayee) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByPayee() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByPayee == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByPayee))
	} else if len(*e.IncludedOrganizationResourcesReferencedByPayee) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPatientResourceReferencedByPayee() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPayee == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPayee))
	} else if len(*e.IncludedPatientResourcesReferencedByPayee) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByPayee() (practitionerRole *PractitionerRole, err error) {
	if e.IncludedPractitionerRoleResourcesReferencedByPayee == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*e.IncludedPractitionerRoleResourcesReferencedByPayee))
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByPayee) == 1 {
		practitionerRole = &(*e.IncludedPractitionerRoleResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPayee() (relatedPerson *RelatedPerson, err error) {
	if e.IncludedRelatedPersonResourcesReferencedByPayee == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByPayee) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*e.IncludedRelatedPersonResourcesReferencedByPayee))
	} else if len(*e.IncludedRelatedPersonResourcesReferencedByPayee) == 1 {
		relatedPerson = &(*e.IncludedRelatedPersonResourcesReferencedByPayee)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByProvider() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByProvider == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByProvider))
	} else if len(*e.IncludedOrganizationResourcesReferencedByProvider) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByProvider)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByProvider() (practitionerRole *PractitionerRole, err error) {
	if e.IncludedPractitionerRoleResourcesReferencedByProvider == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*e.IncludedPractitionerRoleResourcesReferencedByProvider))
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByProvider) == 1 {
		practitionerRole = &(*e.IncludedPractitionerRoleResourcesReferencedByProvider)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatient))
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedDeviceResourcesReferencedByDetailudi() (devices []Device, err error) {
	if e.IncludedDeviceResourcesReferencedByDetailudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *e.IncludedDeviceResourcesReferencedByDetailudi
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedClaimResourceReferencedByClaim() (claim *Claim, err error) {
	if e.IncludedClaimResourcesReferencedByClaim == nil {
		err = errors.New("Included claims not requested")
	} else if len(*e.IncludedClaimResourcesReferencedByClaim) > 1 {
		err = fmt.Errorf("Expected 0 or 1 claim, but found %d", len(*e.IncludedClaimResourcesReferencedByClaim))
	} else if len(*e.IncludedClaimResourcesReferencedByClaim) == 1 {
		claim = &(*e.IncludedClaimResourcesReferencedByClaim)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByEnterer() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByEnterer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByEnterer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByEnterer))
	} else if len(*e.IncludedPractitionerResourcesReferencedByEnterer) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByEnterer)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByEnterer() (practitionerRole *PractitionerRole, err error) {
	if e.IncludedPractitionerRoleResourcesReferencedByEnterer == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByEnterer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*e.IncludedPractitionerRoleResourcesReferencedByEnterer))
	} else if len(*e.IncludedPractitionerRoleResourcesReferencedByEnterer) == 1 {
		practitionerRole = &(*e.IncludedPractitionerRoleResourcesReferencedByEnterer)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedDeviceResourcesReferencedByProcedureudi() (devices []Device, err error) {
	if e.IncludedDeviceResourcesReferencedByProcedureudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *e.IncludedDeviceResourcesReferencedByProcedureudi
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedDeviceResourcesReferencedBySubdetailudi() (devices []Device, err error) {
	if e.IncludedDeviceResourcesReferencedBySubdetailudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *e.IncludedDeviceResourcesReferencedBySubdetailudi
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedLocationResourceReferencedByFacility() (location *Location, err error) {
	if e.IncludedLocationResourcesReferencedByFacility == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResourcesReferencedByFacility))
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) == 1 {
		location = &(*e.IncludedLocationResourcesReferencedByFacility)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedDeviceResourcesReferencedByItemudi() (devices []Device, err error) {
	if e.IncludedDeviceResourcesReferencedByItemudi == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *e.IncludedDeviceResourcesReferencedByItemudi
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if e.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *e.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if e.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *e.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if e.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *e.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if e.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *e.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if e.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *e.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if e.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *e.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if e.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *e.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedCoverageResourcesReferencedByCoverage != nil {
		for idx := range *e.IncludedCoverageResourcesReferencedByCoverage {
			rsc := (*e.IncludedCoverageResourcesReferencedByCoverage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*e.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*e.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByCareteam {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*e.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*e.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPayee {
			rsc := (*e.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByPayee {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*e.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByProvider {
			rsc := (*e.IncludedOrganizationResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByProvider {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedByDetailudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedByDetailudi {
			rsc := (*e.IncludedDeviceResourcesReferencedByDetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedClaimResourcesReferencedByClaim != nil {
		for idx := range *e.IncludedClaimResourcesReferencedByClaim {
			rsc := (*e.IncludedClaimResourcesReferencedByClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*e.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByEnterer != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByEnterer {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedByProcedureudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedByProcedureudi {
			rsc := (*e.IncludedDeviceResourcesReferencedByProcedureudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedBySubdetailudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedBySubdetailudi {
			rsc := (*e.IncludedDeviceResourcesReferencedBySubdetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByFacility {
			rsc := (*e.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedByItemudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedByItemudi {
			rsc := (*e.IncludedDeviceResourcesReferencedByItemudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*e.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*e.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*e.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *e.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*e.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*e.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedCoverageResourcesReferencedByCoverage != nil {
		for idx := range *e.IncludedCoverageResourcesReferencedByCoverage {
			rsc := (*e.IncludedCoverageResourcesReferencedByCoverage)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByCareteam {
			rsc := (*e.IncludedPractitionerResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByCareteam {
			rsc := (*e.IncludedOrganizationResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByCareteam != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByCareteam {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *e.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*e.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByPayee {
			rsc := (*e.IncludedPractitionerResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByPayee {
			rsc := (*e.IncludedOrganizationResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPayee {
			rsc := (*e.IncludedPatientResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByPayee {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedRelatedPersonResourcesReferencedByPayee != nil {
		for idx := range *e.IncludedRelatedPersonResourcesReferencedByPayee {
			rsc := (*e.IncludedRelatedPersonResourcesReferencedByPayee)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByProvider {
			rsc := (*e.IncludedPractitionerResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedOrganizationResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByProvider {
			rsc := (*e.IncludedOrganizationResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByProvider != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByProvider {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByProvider)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *e.IncludedPatientResourcesReferencedByPatient {
			rsc := (*e.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedByDetailudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedByDetailudi {
			rsc := (*e.IncludedDeviceResourcesReferencedByDetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedClaimResourcesReferencedByClaim != nil {
		for idx := range *e.IncludedClaimResourcesReferencedByClaim {
			rsc := (*e.IncludedClaimResourcesReferencedByClaim)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerResourcesReferencedByEnterer != nil {
		for idx := range *e.IncludedPractitionerResourcesReferencedByEnterer {
			rsc := (*e.IncludedPractitionerResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedPractitionerRoleResourcesReferencedByEnterer != nil {
		for idx := range *e.IncludedPractitionerRoleResourcesReferencedByEnterer {
			rsc := (*e.IncludedPractitionerRoleResourcesReferencedByEnterer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedByProcedureudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedByProcedureudi {
			rsc := (*e.IncludedDeviceResourcesReferencedByProcedureudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedBySubdetailudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedBySubdetailudi {
			rsc := (*e.IncludedDeviceResourcesReferencedBySubdetailudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for idx := range *e.IncludedLocationResourcesReferencedByFacility {
			rsc := (*e.IncludedLocationResourcesReferencedByFacility)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.IncludedDeviceResourcesReferencedByItemudi != nil {
		for idx := range *e.IncludedDeviceResourcesReferencedByItemudi {
			rsc := (*e.IncludedDeviceResourcesReferencedByItemudi)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*e.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*e.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *e.RevIncludedConsentResourcesReferencingData {
			rsc := (*e.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*e.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *e.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*e.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*e.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedContractResourcesReferencingSubject {
			rsc := (*e.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*e.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*e.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *e.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*e.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*e.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*e.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *e.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*e.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*e.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*e.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*e.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*e.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *e.RevIncludedListResourcesReferencingItem {
			rsc := (*e.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *e.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*e.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*e.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*e.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *e.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*e.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*e.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *e.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*e.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *e.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*e.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *e.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*e.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*e.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *e.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*e.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*e.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*e.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
