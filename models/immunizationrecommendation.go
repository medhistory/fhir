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

type ImmunizationRecommendation struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier                                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient        *Reference                                          `bson:"patient,omitempty" json:"patient,omitempty"`
	Date           *FHIRDateTime                                       `bson:"date,omitempty" json:"date,omitempty"`
	Authority      *Reference                                          `bson:"authority,omitempty" json:"authority,omitempty"`
	Recommendation []ImmunizationRecommendationRecommendationComponent `bson:"recommendation,omitempty" json:"recommendation,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ImmunizationRecommendation"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ImmunizationRecommendation), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ImmunizationRecommendation) GetBSON() (interface{}, error) {
	x.ResourceType = "ImmunizationRecommendation"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "immunizationRecommendation" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type immunizationRecommendation ImmunizationRecommendation

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ImmunizationRecommendation) UnmarshalJSON(data []byte) (err error) {
	x2 := immunizationRecommendation{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ImmunizationRecommendation(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ImmunizationRecommendation) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ImmunizationRecommendation"
	} else if x.ResourceType != "ImmunizationRecommendation" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ImmunizationRecommendation, instead received %s", x.ResourceType))
	}
	return nil
}

type ImmunizationRecommendationRecommendationComponent struct {
	BackboneElement              `bson:",inline"`
	VaccineCode                  []CodeableConcept                                                `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	TargetDisease                []CodeableConcept                                                `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	ContraindicatedVaccineCode   []CodeableConcept                                                `bson:"contraindicatedVaccineCode,omitempty" json:"contraindicatedVaccineCode,omitempty"`
	ForecastStatus               *CodeableConcept                                                 `bson:"forecastStatus,omitempty" json:"forecastStatus,omitempty"`
	ForecastReason               []CodeableConcept                                                `bson:"forecastReason,omitempty" json:"forecastReason,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterionComponent `bson:"dateCriterion,omitempty" json:"dateCriterion,omitempty"`
	Description                  string                                                           `bson:"description,omitempty" json:"description,omitempty"`
	Series                       string                                                           `bson:"series,omitempty" json:"series,omitempty"`
	DoseNumberPositiveInt        *uint32                                                          `bson:"doseNumberPositiveInt,omitempty" json:"doseNumberPositiveInt,omitempty"`
	DoseNumberString             string                                                           `bson:"doseNumberString,omitempty" json:"doseNumberString,omitempty"`
	SeriesDosesPositiveInt       *uint32                                                          `bson:"seriesDosesPositiveInt,omitempty" json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString            string                                                           `bson:"seriesDosesString,omitempty" json:"seriesDosesString,omitempty"`
	SupportingImmunization       []Reference                                                      `bson:"supportingImmunization,omitempty" json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                                      `bson:"supportingPatientInformation,omitempty" json:"supportingPatientInformation,omitempty"`
}

type ImmunizationRecommendationRecommendationDateCriterionComponent struct {
	BackboneElement `bson:",inline"`
	Code            *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Value           *FHIRDateTime    `bson:"value,omitempty" json:"value,omitempty"`
}

type ImmunizationRecommendationPlus struct {
	ImmunizationRecommendation                     `bson:",inline"`
	ImmunizationRecommendationPlusRelatedResources `bson:",inline"`
}

type ImmunizationRecommendationPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedImmunizationResourcesReferencedBySupport                       *[]Immunization               `bson:"_includedImmunizationResourcesReferencedBySupport,omitempty"`
	IncludedImmunizationEvaluationResourcesReferencedBySupport             *[]ImmunizationEvaluation     `bson:"_includedImmunizationEvaluationResourcesReferencedBySupport,omitempty"`
	RevIncludedAppointmentResourcesReferencingReasonreference              *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingReasonreference,omitempty"`
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
	RevIncludedEncounterResourcesReferencingReasonreference                *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingReasonreference,omitempty"`
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
	RevIncludedCarePlanResourcesReferencingActivityreference               *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingBasedon                      *[]Observation                `bson:"_revIncludedObservationResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                      *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                       *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon             *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                            *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingBasedon                 *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingBasedon,omitempty"`
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

func (i *ImmunizationRecommendationPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if i.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*i.IncludedPatientResourcesReferencedByPatient))
	} else if len(*i.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*i.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetIncludedImmunizationResourcesReferencedBySupport() (immunizations []Immunization, err error) {
	if i.IncludedImmunizationResourcesReferencedBySupport == nil {
		err = errors.New("Included immunizations not requested")
	} else {
		immunizations = *i.IncludedImmunizationResourcesReferencedBySupport
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetIncludedImmunizationEvaluationResourcesReferencedBySupport() (immunizationEvaluations []ImmunizationEvaluation, err error) {
	if i.IncludedImmunizationEvaluationResourcesReferencedBySupport == nil {
		err = errors.New("Included immunizationEvaluations not requested")
	} else {
		immunizationEvaluations = *i.IncludedImmunizationEvaluationResourcesReferencedBySupport
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingReasonreference() (appointments []Appointment, err error) {
	if i.RevIncludedAppointmentResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *i.RevIncludedAppointmentResourcesReferencingReasonreference
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *i.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if i.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *i.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *i.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if i.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *i.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *i.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *i.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if i.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *i.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if i.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *i.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *i.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if i.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *i.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEncounterResourcesReferencingReasonreference() (encounters []Encounter, err error) {
	if i.RevIncludedEncounterResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *i.RevIncludedEncounterResourcesReferencingReasonreference
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if i.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *i.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if i.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *i.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if i.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *i.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *i.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if i.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *i.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if i.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *i.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if i.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *i.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if i.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *i.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if i.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *i.RevIncludedListResourcesReferencingItem
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *i.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if i.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *i.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingBasedon() (observations []Observation, err error) {
	if i.RevIncludedObservationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *i.RevIncludedObservationResourcesReferencingBasedon
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if i.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *i.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *i.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if i.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *i.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingBasedon() (diagnosticReports []DiagnosticReport, err error) {
	if i.RevIncludedDiagnosticReportResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *i.RevIncludedDiagnosticReportResourcesReferencingBasedon
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if i.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *i.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if i.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *i.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if i.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *i.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if i.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *i.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *i.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPatient {
			rsc := (*i.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedImmunizationResourcesReferencedBySupport != nil {
		for idx := range *i.IncludedImmunizationResourcesReferencedBySupport {
			rsc := (*i.IncludedImmunizationResourcesReferencedBySupport)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedImmunizationEvaluationResourcesReferencedBySupport != nil {
		for idx := range *i.IncludedImmunizationEvaluationResourcesReferencedBySupport {
			rsc := (*i.IncludedImmunizationEvaluationResourcesReferencedBySupport)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.RevIncludedAppointmentResourcesReferencingReasonreference != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingReasonreference {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*i.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*i.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEncounterResourcesReferencingReasonreference != nil {
		for idx := range *i.RevIncludedEncounterResourcesReferencingReasonreference {
			rsc := (*i.RevIncludedEncounterResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*i.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*i.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*i.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *i.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*i.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*i.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*i.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *i.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*i.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (i *ImmunizationRecommendationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if i.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *i.IncludedPatientResourcesReferencedByPatient {
			rsc := (*i.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedImmunizationResourcesReferencedBySupport != nil {
		for idx := range *i.IncludedImmunizationResourcesReferencedBySupport {
			rsc := (*i.IncludedImmunizationResourcesReferencedBySupport)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.IncludedImmunizationEvaluationResourcesReferencedBySupport != nil {
		for idx := range *i.IncludedImmunizationEvaluationResourcesReferencedBySupport {
			rsc := (*i.IncludedImmunizationEvaluationResourcesReferencedBySupport)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAppointmentResourcesReferencingReasonreference != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingReasonreference {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*i.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *i.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*i.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *i.RevIncludedConsentResourcesReferencingData {
			rsc := (*i.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*i.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *i.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*i.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*i.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*i.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedContractResourcesReferencingSubject {
			rsc := (*i.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *i.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*i.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *i.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*i.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEncounterResourcesReferencingReasonreference != nil {
		for idx := range *i.RevIncludedEncounterResourcesReferencingReasonreference {
			rsc := (*i.RevIncludedEncounterResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*i.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *i.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*i.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *i.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*i.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*i.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *i.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*i.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *i.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*i.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*i.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*i.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*i.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for idx := range *i.RevIncludedCarePlanResourcesReferencingActivityreference {
			rsc := (*i.RevIncludedCarePlanResourcesReferencingActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *i.RevIncludedListResourcesReferencingItem {
			rsc := (*i.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*i.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*i.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*i.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*i.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*i.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*i.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *i.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*i.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *i.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*i.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *i.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*i.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *i.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*i.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*i.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *i.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*i.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *i.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*i.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *i.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*i.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*i.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*i.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
