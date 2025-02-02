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

type CarePlan struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string                    `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                    `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference                 `bson:"replaces,omitempty" json:"replaces,omitempty"`
	PartOf                []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                string                      `bson:"status,omitempty" json:"status,omitempty"`
	Intent                string                      `bson:"intent,omitempty" json:"intent,omitempty"`
	Category              []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Title                 string                      `bson:"title,omitempty" json:"title,omitempty"`
	Description           string                      `bson:"description,omitempty" json:"description,omitempty"`
	Subject               *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter             *Reference                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Period                *Period                     `bson:"period,omitempty" json:"period,omitempty"`
	Created               *FHIRDateTime               `bson:"created,omitempty" json:"created,omitempty"`
	Author                *Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Contributor           []Reference                 `bson:"contributor,omitempty" json:"contributor,omitempty"`
	CareTeam              []Reference                 `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	AddressesCode         []CodeableConcept           `bson:"addressesCode,omitempty" json:"addressesCode,omitempty"`
	AddressesReference    []Reference                 `bson:"addressesReference,omitempty" json:"addressesReference,omitempty"`
	SupportingInfo        []Reference                 `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Goal                  []Reference                 `bson:"goal,omitempty" json:"goal,omitempty"`
	Activity              []CarePlanActivityComponent `bson:"activity,omitempty" json:"activity,omitempty"`
	Note                  []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CarePlan) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "CarePlan"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to CarePlan), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *CarePlan) GetBSON() (interface{}, error) {
	x.ResourceType = "CarePlan"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "carePlan" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type carePlan CarePlan

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CarePlan) UnmarshalJSON(data []byte) (err error) {
	x2 := carePlan{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CarePlan(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CarePlan) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CarePlan"
	} else if x.ResourceType != "CarePlan" {
		return errors.New(fmt.Sprintf("Expected resourceType to be CarePlan, instead received %s", x.ResourceType))
	}
	return nil
}

type CarePlanActivityComponent struct {
	BackboneElement        `bson:",inline"`
	OutcomeCodeableConcept []CodeableConcept                `bson:"outcomeCodeableConcept,omitempty" json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference                      `bson:"outcomeReference,omitempty" json:"outcomeReference,omitempty"`
	Progress               []Annotation                     `bson:"progress,omitempty" json:"progress,omitempty"`
	Reference              *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CarePlanActivityDetailComponent struct {
	BackboneElement        `bson:",inline"`
	Kind                   string            `bson:"kind,omitempty" json:"kind,omitempty"`
	InstantiatesCanonical  []string          `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string          `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Code                   *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	ReasonCode             []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference        []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Goal                   []Reference       `bson:"goal,omitempty" json:"goal,omitempty"`
	Status                 string            `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason           *CodeableConcept  `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	DoNotPerform           *bool             `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	ScheduledTiming        *Timing           `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period           `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledString        string            `bson:"scheduledString,omitempty" json:"scheduledString,omitempty"`
	Location               *Reference        `bson:"location,omitempty" json:"location,omitempty"`
	ReportedBoolean        *bool             `bson:"reportedBoolean,omitempty" json:"reportedBoolean,omitempty"`
	ReportedReference      *Reference        `bson:"reportedReference,omitempty" json:"reportedReference,omitempty"`
	Performer              []Reference       `bson:"performer,omitempty" json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept  `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference        `bson:"productReference,omitempty" json:"productReference,omitempty"`
	DailyAmount            *Quantity         `bson:"dailyAmount,omitempty" json:"dailyAmount,omitempty"`
	Quantity               *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Description            string            `bson:"description,omitempty" json:"description,omitempty"`
}

type CarePlanPlus struct {
	CarePlan                     `bson:",inline"`
	CarePlanPlusRelatedResources `bson:",inline"`
}

type CarePlanPlusRelatedResources struct {
	IncludedCareTeamResourcesReferencedByCareteam                            *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByCareteam,omitempty"`
	IncludedPractitionerResourcesReferencedByPerformer                       *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                       *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedCareTeamResourcesReferencedByPerformer                           *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByPerformer,omitempty"`
	IncludedDeviceResourcesReferencedByPerformer                             *[]Device                     `bson:"_includedDeviceResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByPerformer                  *[]HealthcareService          `bson:"_includedHealthcareServiceResourcesReferencedByPerformer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPerformer                   *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                      *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGoalResourcesReferencedByGoal                                    *[]Goal                       `bson:"_includedGoalResourcesReferencedByGoal,omitempty"`
	IncludedGroupResourcesReferencedBySubject                                *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                              *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedCarePlanResourcesReferencedByReplaces                            *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByReplaces,omitempty"`
	IncludedQuestionnaireResourcesReferencedByInstantiatescanonical          *[]Questionnaire              `bson:"_includedQuestionnaireResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedMeasureResourcesReferencedByInstantiatescanonical                *[]Measure                    `bson:"_includedMeasureResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical         *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical    *[]OperationDefinition        `bson:"_includedOperationDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical     *[]ActivityDefinition         `bson:"_includedActivityDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedCarePlanResourcesReferencedByPartof                              *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByPartof,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                          *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedAppointmentResourcesReferencedByActivityreference                *[]Appointment                `bson:"_includedAppointmentResourcesReferencedByActivityreference,omitempty"`
	IncludedMedicationRequestResourcesReferencedByActivityreference          *[]MedicationRequest          `bson:"_includedMedicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedTaskResourcesReferencedByActivityreference                       *[]Task                       `bson:"_includedTaskResourcesReferencedByActivityreference,omitempty"`
	IncludedNutritionOrderResourcesReferencedByActivityreference             *[]NutritionOrder             `bson:"_includedNutritionOrderResourcesReferencedByActivityreference,omitempty"`
	IncludedRequestGroupResourcesReferencedByActivityreference               *[]RequestGroup               `bson:"_includedRequestGroupResourcesReferencedByActivityreference,omitempty"`
	IncludedVisionPrescriptionResourcesReferencedByActivityreference         *[]VisionPrescription         `bson:"_includedVisionPrescriptionResourcesReferencedByActivityreference,omitempty"`
	IncludedDeviceRequestResourcesReferencedByActivityreference              *[]DeviceRequest              `bson:"_includedDeviceRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedServiceRequestResourcesReferencedByActivityreference             *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedCommunicationRequestResourcesReferencedByActivityreference       *[]CommunicationRequest       `bson:"_includedCommunicationRequestResourcesReferencedByActivityreference,omitempty"`
	IncludedImmunizationRecommendationResourcesReferencedByActivityreference *[]ImmunizationRecommendation `bson:"_includedImmunizationRecommendationResourcesReferencedByActivityreference,omitempty"`
	IncludedConditionResourcesReferencedByCondition                          *[]Condition                  `bson:"_includedConditionResourcesReferencedByCondition,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                             *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                              *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo                 *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingSuccessor                  *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDerivedfrom                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingPredecessor                *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingComposedof                 *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedEventDefinitionResourcesReferencingDependson                  *[]EventDefinition            `bson:"_revIncludedEventDefinitionResourcesReferencingDependson,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingItem                      *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingItem,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref                *[]DocumentManifest           `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedConsentResourcesReferencingData                               *[]Consent                    `bson:"_revIncludedConsentResourcesReferencingData,omitempty"`
	RevIncludedMeasureResourcesReferencingSuccessor                          *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingSuccessor,omitempty"`
	RevIncludedMeasureResourcesReferencingDerivedfrom                        *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedMeasureResourcesReferencingPredecessor                        *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingPredecessor,omitempty"`
	RevIncludedMeasureResourcesReferencingComposedof                         *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingComposedof,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath1                     *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedMeasureResourcesReferencingDependsonPath2                     *[]Measure                    `bson:"_revIncludedMeasureResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingBasedon                  *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingBasedon,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelated                  *[]DocumentReference          `bson:"_revIncludedDocumentReferenceResourcesReferencingRelated,omitempty"`
	RevIncludedMeasureReportResourcesReferencingEvaluatedresource            *[]MeasureReport              `bson:"_revIncludedMeasureReportResourcesReferencingEvaluatedresource,omitempty"`
	RevIncludedServiceRequestResourcesReferencingBasedon                     *[]ServiceRequest             `bson:"_revIncludedServiceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedVerificationResultResourcesReferencingTarget                  *[]VerificationResult         `bson:"_revIncludedVerificationResultResourcesReferencingTarget,omitempty"`
	RevIncludedContractResourcesReferencingSubject                           *[]Contract                   `bson:"_revIncludedContractResourcesReferencingSubject,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest                      *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse                     *[]PaymentNotice              `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingSuccessor               *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDerivedfrom             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingPredecessor             *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingComposedof              *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath1          *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchDefinitionResourcesReferencingDependsonPath2          *[]ResearchDefinition         `bson:"_revIncludedResearchDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource               *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingBasedon                       *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingBasedon,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor        *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof       *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1   *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2   *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedCommunicationResourcesReferencingPartof                       *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingPartof,omitempty"`
	RevIncludedCommunicationResourcesReferencingBasedon                      *[]Communication              `bson:"_revIncludedCommunicationResourcesReferencingBasedon,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingSuccessor               *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDerivedfrom             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingPredecessor             *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingComposedof              *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath1          *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedActivityDefinitionResourcesReferencingDependsonPath2          *[]ActivityDefinition         `bson:"_revIncludedActivityDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedLinkageResourcesReferencingItem                               *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingItem,omitempty"`
	RevIncludedLinkageResourcesReferencingSource                             *[]Linkage                    `bson:"_revIncludedLinkageResourcesReferencingSource,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                      *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest                 *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                        *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation     *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                          *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                          *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                                 *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedCarePlanResourcesReferencingReplaces                          *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingReplaces,omitempty"`
	RevIncludedCarePlanResourcesReferencingPartof                            *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingPartof,omitempty"`
	RevIncludedCarePlanResourcesReferencingBasedon                           *[]CarePlan                   `bson:"_revIncludedCarePlanResourcesReferencingBasedon,omitempty"`
	RevIncludedProcedureResourcesReferencingBasedon                          *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingBasedon,omitempty"`
	RevIncludedListResourcesReferencingItem                                  *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor                 *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof                *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson                 *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                          *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingBasedon                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingBasedon,omitempty"`
	RevIncludedLibraryResourcesReferencingSuccessor                          *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingSuccessor,omitempty"`
	RevIncludedLibraryResourcesReferencingDerivedfrom                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedLibraryResourcesReferencingPredecessor                        *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingPredecessor,omitempty"`
	RevIncludedLibraryResourcesReferencingComposedof                         *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingComposedof,omitempty"`
	RevIncludedLibraryResourcesReferencingDependson                          *[]Library                    `bson:"_revIncludedLibraryResourcesReferencingDependson,omitempty"`
	RevIncludedCommunicationRequestResourcesReferencingBasedon               *[]CommunicationRequest       `bson:"_revIncludedCommunicationRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                              *[]Basic                      `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingBasedon                   *[]DiagnosticReport           `bson:"_revIncludedDiagnosticReportResourcesReferencingBasedon,omitempty"`
	RevIncludedEvidenceResourcesReferencingSuccessor                         *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingDerivedfrom                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceResourcesReferencingPredecessor                       *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceResourcesReferencingComposedof                        *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceResourcesReferencingDependson                         *[]Evidence                   `bson:"_revIncludedEvidenceResourcesReferencingDependson,omitempty"`
	RevIncludedAuditEventResourcesReferencingEntity                          *[]AuditEvent                 `bson:"_revIncludedAuditEventResourcesReferencingEntity,omitempty"`
	RevIncludedConditionResourcesReferencingEvidencedetail                   *[]Condition                  `bson:"_revIncludedConditionResourcesReferencingEvidencedetail,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject                        *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry                          *[]Composition                `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated                   *[]DetectedIssue              `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingBasedon              *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingBasedon,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject              *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo          *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                   *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                  *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1              *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2              *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (c *CarePlanPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByCareteam() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByCareteam == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByCareteam
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByPerformer() (practitioners []Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *c.IncludedPractitionerResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByPerformer() (organizations []Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *c.IncludedOrganizationResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCareTeamResourcesReferencedByPerformer() (careTeams []CareTeam, err error) {
	if c.IncludedCareTeamResourcesReferencedByPerformer == nil {
		err = errors.New("Included careTeams not requested")
	} else {
		careTeams = *c.IncludedCareTeamResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedDeviceResourcesReferencedByPerformer() (devices []Device, err error) {
	if c.IncludedDeviceResourcesReferencedByPerformer == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *c.IncludedDeviceResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourcesReferencedByPerformer() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedHealthcareServiceResourcesReferencedByPerformer() (healthcareServices []HealthcareService, err error) {
	if c.IncludedHealthcareServiceResourcesReferencedByPerformer == nil {
		err = errors.New("Included healthcareServices not requested")
	} else {
		healthcareServices = *c.IncludedHealthcareServiceResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPractitionerRoleResourcesReferencedByPerformer() (practitionerRoles []PractitionerRole, err error) {
	if c.IncludedPractitionerRoleResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitionerRoles not requested")
	} else {
		practitionerRoles = *c.IncludedPractitionerRoleResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByPerformer() (relatedPeople []RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *c.IncludedRelatedPersonResourcesReferencedByPerformer
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedGoalResourcesReferencedByGoal() (goals []Goal, err error) {
	if c.IncludedGoalResourcesReferencedByGoal == nil {
		err = errors.New("Included goals not requested")
	} else {
		goals = *c.IncludedGoalResourcesReferencedByGoal
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedBySubject))
	} else if len(*c.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySubject))
	} else if len(*c.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByReplaces() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByReplaces == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByReplaces
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedQuestionnaireResourcesReferencedByInstantiatescanonical() (questionnaires []Questionnaire, err error) {
	if c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included questionnaires not requested")
	} else {
		questionnaires = *c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedMeasureResourcesReferencedByInstantiatescanonical() (measures []Measure, err error) {
	if c.IncludedMeasureResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included measures not requested")
	} else {
		measures = *c.IncludedMeasureResourcesReferencedByInstantiatescanonical
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical() (planDefinitions []PlanDefinition, err error) {
	if c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedOperationDefinitionResourcesReferencedByInstantiatescanonical() (operationDefinitions []OperationDefinition, err error) {
	if c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included operationDefinitions not requested")
	} else {
		operationDefinitions = *c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical() (activityDefinitions []ActivityDefinition, err error) {
	if c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByPartof() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByPartof == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByPartof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if c.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*c.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*c.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*c.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedAppointmentResourceReferencedByActivityreference() (appointment *Appointment, err error) {
	if c.IncludedAppointmentResourcesReferencedByActivityreference == nil {
		err = errors.New("Included appointments not requested")
	} else if len(*c.IncludedAppointmentResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 appointment, but found %d", len(*c.IncludedAppointmentResourcesReferencedByActivityreference))
	} else if len(*c.IncludedAppointmentResourcesReferencedByActivityreference) == 1 {
		appointment = &(*c.IncludedAppointmentResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedMedicationRequestResourceReferencedByActivityreference() (medicationRequest *MedicationRequest, err error) {
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included medicationrequests not requested")
	} else if len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicationRequest, but found %d", len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedMedicationRequestResourcesReferencedByActivityreference) == 1 {
		medicationRequest = &(*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedTaskResourceReferencedByActivityreference() (task *Task, err error) {
	if c.IncludedTaskResourcesReferencedByActivityreference == nil {
		err = errors.New("Included tasks not requested")
	} else if len(*c.IncludedTaskResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 task, but found %d", len(*c.IncludedTaskResourcesReferencedByActivityreference))
	} else if len(*c.IncludedTaskResourcesReferencedByActivityreference) == 1 {
		task = &(*c.IncludedTaskResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedNutritionOrderResourceReferencedByActivityreference() (nutritionOrder *NutritionOrder, err error) {
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference == nil {
		err = errors.New("Included nutritionorders not requested")
	} else if len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 nutritionOrder, but found %d", len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference))
	} else if len(*c.IncludedNutritionOrderResourcesReferencedByActivityreference) == 1 {
		nutritionOrder = &(*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedRequestGroupResourceReferencedByActivityreference() (requestGroup *RequestGroup, err error) {
	if c.IncludedRequestGroupResourcesReferencedByActivityreference == nil {
		err = errors.New("Included requestgroups not requested")
	} else if len(*c.IncludedRequestGroupResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 requestGroup, but found %d", len(*c.IncludedRequestGroupResourcesReferencedByActivityreference))
	} else if len(*c.IncludedRequestGroupResourcesReferencedByActivityreference) == 1 {
		requestGroup = &(*c.IncludedRequestGroupResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedVisionPrescriptionResourceReferencedByActivityreference() (visionPrescription *VisionPrescription, err error) {
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference == nil {
		err = errors.New("Included visionprescriptions not requested")
	} else if len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 visionPrescription, but found %d", len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference))
	} else if len(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference) == 1 {
		visionPrescription = &(*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedDeviceRequestResourceReferencedByActivityreference() (deviceRequest *DeviceRequest, err error) {
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included devicerequests not requested")
	} else if len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceRequest, but found %d", len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedDeviceRequestResourcesReferencedByActivityreference) == 1 {
		deviceRequest = &(*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedServiceRequestResourceReferencedByActivityreference() (serviceRequest *ServiceRequest, err error) {
	if c.IncludedServiceRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included servicerequests not requested")
	} else if len(*c.IncludedServiceRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 serviceRequest, but found %d", len(*c.IncludedServiceRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedServiceRequestResourcesReferencedByActivityreference) == 1 {
		serviceRequest = &(*c.IncludedServiceRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCommunicationRequestResourceReferencedByActivityreference() (communicationRequest *CommunicationRequest, err error) {
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference == nil {
		err = errors.New("Included communicationrequests not requested")
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 communicationRequest, but found %d", len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference))
	} else if len(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference) == 1 {
		communicationRequest = &(*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedImmunizationRecommendationResourceReferencedByActivityreference() (immunizationRecommendation *ImmunizationRecommendation, err error) {
	if c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference == nil {
		err = errors.New("Included immunizationrecommendations not requested")
	} else if len(*c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference) > 1 {
		err = fmt.Errorf("Expected 0 or 1 immunizationRecommendation, but found %d", len(*c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference))
	} else if len(*c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference) == 1 {
		immunizationRecommendation = &(*c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedConditionResourcesReferencedByCondition() (conditions []Condition, err error) {
	if c.IncludedConditionResourcesReferencedByCondition == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *c.IncludedConditionResourcesReferencedByCondition
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if c.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *c.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByPatient))
	} else if len(*c.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *c.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingBasedon() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedServiceRequestResourcesReferencingBasedon() (serviceRequests []ServiceRequest, err error) {
	if c.RevIncludedServiceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded serviceRequests not requested")
	} else {
		serviceRequests = *c.RevIncludedServiceRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if c.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *c.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingBasedon() (imagingStudies []ImagingStudy, err error) {
	if c.RevIncludedImagingStudyResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *c.RevIncludedImagingStudyResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingReplaces() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingReplaces == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingReplaces
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingPartof() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingPartof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingBasedon() (carePlans []CarePlan, err error) {
	if c.RevIncludedCarePlanResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *c.RevIncludedCarePlanResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedProcedureResourcesReferencingBasedon() (procedures []Procedure, err error) {
	if c.RevIncludedProcedureResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *c.RevIncludedProcedureResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if c.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *c.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedObservationResourcesReferencingBasedon() (observations []Observation, err error) {
	if c.RevIncludedObservationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *c.RevIncludedObservationResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *c.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingBasedon() (diagnosticReports []DiagnosticReport, err error) {
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *c.RevIncludedDiagnosticReportResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if c.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *c.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingBasedon() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CarePlanPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*c.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*c.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*c.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*c.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGoalResourcesReferencedByGoal != nil {
		for idx := range *c.IncludedGoalResourcesReferencedByGoal {
			rsc := (*c.IncludedGoalResourcesReferencedByGoal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByReplaces != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByReplaces {
			rsc := (*c.IncludedCarePlanResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPartof != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPartof {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByActivityreference {
			rsc := (*c.IncludedAppointmentResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedTaskResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedTaskResourcesReferencedByActivityreference {
			rsc := (*c.IncludedTaskResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRequestGroupResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedRequestGroupResourcesReferencedByActivityreference {
			rsc := (*c.IncludedRequestGroupResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedServiceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedServiceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedServiceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference {
			rsc := (*c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*c.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CarePlanPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedServiceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedServiceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*c.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*c.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*c.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*c.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*c.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingReplaces {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingPartof {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*c.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CarePlanPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedCareTeamResourcesReferencedByCareteam != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByCareteam {
			rsc := (*c.IncludedCareTeamResourcesReferencedByCareteam)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*c.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*c.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*c.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*c.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*c.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*c.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *c.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*c.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGoalResourcesReferencedByGoal != nil {
		for idx := range *c.IncludedGoalResourcesReferencedByGoal {
			rsc := (*c.IncludedGoalResourcesReferencedByGoal)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedGroupResourcesReferencedBySubject {
			rsc := (*c.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *c.IncludedPatientResourcesReferencedBySubject {
			rsc := (*c.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByReplaces != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByReplaces {
			rsc := (*c.IncludedCarePlanResourcesReferencedByReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*c.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByPartof != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByPartof {
			rsc := (*c.IncludedCarePlanResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *c.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*c.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedAppointmentResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedAppointmentResourcesReferencedByActivityreference {
			rsc := (*c.IncludedAppointmentResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedMedicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedMedicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedMedicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedTaskResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedTaskResourcesReferencedByActivityreference {
			rsc := (*c.IncludedTaskResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedNutritionOrderResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedNutritionOrderResourcesReferencedByActivityreference {
			rsc := (*c.IncludedNutritionOrderResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedRequestGroupResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedRequestGroupResourcesReferencedByActivityreference {
			rsc := (*c.IncludedRequestGroupResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedVisionPrescriptionResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedVisionPrescriptionResourcesReferencedByActivityreference {
			rsc := (*c.IncludedVisionPrescriptionResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedDeviceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedDeviceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedDeviceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedServiceRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedServiceRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedServiceRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCommunicationRequestResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedCommunicationRequestResourcesReferencedByActivityreference {
			rsc := (*c.IncludedCommunicationRequestResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference != nil {
		for idx := range *c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference {
			rsc := (*c.IncludedImmunizationRecommendationResourcesReferencedByActivityreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedConditionResourcesReferencedByCondition != nil {
		for idx := range *c.IncludedConditionResourcesReferencedByCondition {
			rsc := (*c.IncludedConditionResourcesReferencedByCondition)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *c.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*c.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *c.IncludedPatientResourcesReferencedByPatient {
			rsc := (*c.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*c.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*c.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *c.RevIncludedConsentResourcesReferencingData {
			rsc := (*c.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*c.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *c.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*c.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedServiceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedServiceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedServiceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*c.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedContractResourcesReferencingSubject {
			rsc := (*c.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*c.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *c.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*c.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImagingStudyResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedImagingStudyResourcesReferencingBasedon {
			rsc := (*c.RevIncludedImagingStudyResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*c.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *c.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*c.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*c.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*c.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*c.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*c.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*c.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*c.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*c.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingReplaces != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingReplaces {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingReplaces)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingPartof != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingPartof {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCarePlanResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCarePlanResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCarePlanResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedProcedureResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedProcedureResourcesReferencingBasedon {
			rsc := (*c.RevIncludedProcedureResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *c.RevIncludedListResourcesReferencingItem {
			rsc := (*c.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*c.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedObservationResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedObservationResourcesReferencingBasedon {
			rsc := (*c.RevIncludedObservationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*c.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*c.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*c.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*c.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDiagnosticReportResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedDiagnosticReportResourcesReferencingBasedon {
			rsc := (*c.RevIncludedDiagnosticReportResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *c.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*c.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *c.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*c.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *c.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*c.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*c.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *c.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*c.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*c.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*c.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
