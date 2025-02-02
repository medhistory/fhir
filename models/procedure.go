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

type Procedure struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string                        `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                        `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference                     `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                []Reference                     `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                string                          `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason          *CodeableConcept                `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category              []CodeableConcept               `bson:"category,omitempty" json:"category,omitempty"`
	Code                  *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	Subject               *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter             *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime    *FHIRDateTime                   `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                         `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceString      string                          `bson:"occurrenceString,omitempty" json:"occurrenceString,omitempty"`
	OccurrenceAge         *Quantity                       `bson:"occurrenceAge,omitempty" json:"occurrenceAge,omitempty"`
	OccurrenceRange       *Range                          `bson:"occurrenceRange,omitempty" json:"occurrenceRange,omitempty"`
	OccurrenceTiming      *Timing                         `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Recorded              *FHIRDateTime                   `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Recorder              *Reference                      `bson:"recorder,omitempty" json:"recorder,omitempty"`
	ReportedBoolean       *bool                           `bson:"reportedBoolean,omitempty" json:"reportedBoolean,omitempty"`
	ReportedReference     *Reference                      `bson:"reportedReference,omitempty" json:"reportedReference,omitempty"`
	Performer             []ProcedurePerformerComponent   `bson:"performer,omitempty" json:"performer,omitempty"`
	Location              *Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode            []CodeableConcept               `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference                     `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	BodySite              []CodeableConcept               `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Outcome               *CodeableConcept                `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report                []Reference                     `bson:"report,omitempty" json:"report,omitempty"`
	Complication          []CodeableConcept               `bson:"complication,omitempty" json:"complication,omitempty"`
	ComplicationDetail    []Reference                     `bson:"complicationDetail,omitempty" json:"complicationDetail,omitempty"`
	FollowUp              []CodeableConcept               `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Note                  []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
	FocalDevice           []ProcedureFocalDeviceComponent `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	UsedReference         []Reference                     `bson:"usedReference,omitempty" json:"usedReference,omitempty"`
	UsedCode              []CodeableConcept               `bson:"usedCode,omitempty" json:"usedCode,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Procedure) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Procedure"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Procedure), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Procedure) GetBSON() (interface{}, error) {
	x.ResourceType = "Procedure"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "procedure" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type procedure Procedure

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Procedure) UnmarshalJSON(data []byte) (err error) {
	x2 := procedure{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Procedure(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Procedure) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Procedure"
	} else if x.ResourceType != "Procedure" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Procedure, instead received %s", x.ResourceType))
	}
	return nil
}

type ProcedurePerformerComponent struct {
	BackboneElement `bson:",inline"`
	Function        *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor           *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf      *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

type ProcedureFocalDeviceComponent struct {
	BackboneElement `bson:",inline"`
	Action          *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated     *Reference       `bson:"manipulated,omitempty" json:"manipulated,omitempty"`
}

type ProcedurePlus struct {
	Procedure                     `bson:",inline"`
	ProcedurePlusRelatedResources `bson:",inline"`
}

type ProcedurePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPerformer                     *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByPerformer,omitempty"`
	IncludedOrganizationResourcesReferencedByPerformer                     *[]Organization               `bson:"_includedOrganizationResourcesReferencedByPerformer,omitempty"`
	IncludedCareTeamResourcesReferencedByPerformer                         *[]CareTeam                   `bson:"_includedCareTeamResourcesReferencedByPerformer,omitempty"`
	IncludedDeviceResourcesReferencedByPerformer                           *[]Device                     `bson:"_includedDeviceResourcesReferencedByPerformer,omitempty"`
	IncludedPatientResourcesReferencedByPerformer                          *[]Patient                    `bson:"_includedPatientResourcesReferencedByPerformer,omitempty"`
	IncludedHealthcareServiceResourcesReferencedByPerformer                *[]HealthcareService          `bson:"_includedHealthcareServiceResourcesReferencedByPerformer,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByPerformer                 *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByPerformer,omitempty"`
	IncludedRelatedPersonResourcesReferencedByPerformer                    *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByPerformer,omitempty"`
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedQuestionnaireResourcesReferencedByInstantiatescanonical        *[]Questionnaire              `bson:"_includedQuestionnaireResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedMeasureResourcesReferencedByInstantiatescanonical              *[]Measure                    `bson:"_includedMeasureResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical       *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical  *[]OperationDefinition        `bson:"_includedOperationDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical   *[]ActivityDefinition         `bson:"_includedActivityDefinitionResourcesReferencedByInstantiatescanonical,omitempty"`
	IncludedObservationResourcesReferencedByPartof                         *[]Observation                `bson:"_includedObservationResourcesReferencedByPartof,omitempty"`
	IncludedProcedureResourcesReferencedByPartof                           *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByPartof,omitempty"`
	IncludedMedicationAdministrationResourcesReferencedByPartof            *[]MedicationAdministration   `bson:"_includedMedicationAdministrationResourcesReferencedByPartof,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedCarePlanResourcesReferencedByBasedon                           *[]CarePlan                   `bson:"_includedCarePlanResourcesReferencedByBasedon,omitempty"`
	IncludedServiceRequestResourcesReferencedByBasedon                     *[]ServiceRequest             `bson:"_includedServiceRequestResourcesReferencedByBasedon,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedConditionResourcesReferencedByReasonreference                  *[]Condition                  `bson:"_includedConditionResourcesReferencedByReasonreference,omitempty"`
	IncludedObservationResourcesReferencedByReasonreference                *[]Observation                `bson:"_includedObservationResourcesReferencedByReasonreference,omitempty"`
	IncludedProcedureResourcesReferencedByReasonreference                  *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByReasonreference,omitempty"`
	IncludedDiagnosticReportResourcesReferencedByReasonreference           *[]DiagnosticReport           `bson:"_includedDiagnosticReportResourcesReferencedByReasonreference,omitempty"`
	IncludedDocumentReferenceResourcesReferencedByReasonreference          *[]DocumentReference          `bson:"_includedDocumentReferenceResourcesReferencedByReasonreference,omitempty"`
	IncludedLocationResourcesReferencedByLocation                          *[]Location                   `bson:"_includedLocationResourcesReferencedByLocation,omitempty"`
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
	RevIncludedChargeItemResourcesReferencingService                       *[]ChargeItem                 `bson:"_revIncludedChargeItemResourcesReferencingService,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingSuccessor      *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingPredecessor    *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingComposedof     *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 *[]ResearchElementDefinition  `bson:"_revIncludedResearchElementDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedEncounterResourcesReferencingDiagnosis                      *[]Encounter                  `bson:"_revIncludedEncounterResourcesReferencingDiagnosis,omitempty"`
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
	RevIncludedMedicationUsageResourcesReferencingPartof                   *[]MedicationUsage            `bson:"_revIncludedMedicationUsageResourcesReferencingPartof,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingBasedon                    *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingBasedon,omitempty"`
	RevIncludedDeviceRequestResourcesReferencingPriorrequest               *[]DeviceRequest              `bson:"_revIncludedDeviceRequestResourcesReferencingPriorrequest,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingFocus                      *[]MessageHeader              `bson:"_revIncludedMessageHeaderResourcesReferencingFocus,omitempty"`
	RevIncludedImmunizationRecommendationResourcesReferencingInformation   *[]ImmunizationRecommendation `bson:"_revIncludedImmunizationRecommendationResourcesReferencingInformation,omitempty"`
	RevIncludedProvenanceResourcesReferencingEntity                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingEntity,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget                        *[]Provenance                 `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedTaskResourcesReferencingSubject                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingSubject,omitempty"`
	RevIncludedTaskResourcesReferencingFocus                               *[]Task                       `bson:"_revIncludedTaskResourcesReferencingFocus,omitempty"`
	RevIncludedTaskResourcesReferencingBasedon                             *[]Task                       `bson:"_revIncludedTaskResourcesReferencingBasedon,omitempty"`
	RevIncludedProcedureResourcesReferencingPartof                         *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingPartof,omitempty"`
	RevIncludedProcedureResourcesReferencingReasonreference                *[]Procedure                  `bson:"_revIncludedProcedureResourcesReferencingReasonreference,omitempty"`
	RevIncludedListResourcesReferencingItem                                *[]List                       `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingSuccessor               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingSuccessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDerivedfrom             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingPredecessor             *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingPredecessor,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingComposedof              *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingComposedof,omitempty"`
	RevIncludedEvidenceVariableResourcesReferencingDependson               *[]EvidenceVariable           `bson:"_revIncludedEvidenceVariableResourcesReferencingDependson,omitempty"`
	RevIncludedFlagResourcesReferencingSubject                             *[]Flag                       `bson:"_revIncludedFlagResourcesReferencingSubject,omitempty"`
	RevIncludedAdverseEventResourcesReferencingSubstance                   *[]AdverseEvent               `bson:"_revIncludedAdverseEventResourcesReferencingSubstance,omitempty"`
	RevIncludedObservationResourcesReferencingFocus                        *[]Observation                `bson:"_revIncludedObservationResourcesReferencingFocus,omitempty"`
	RevIncludedObservationResourcesReferencingPartof                       *[]Observation                `bson:"_revIncludedObservationResourcesReferencingPartof,omitempty"`
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
	RevIncludedQuestionnaireResponseResourcesReferencingPartof             *[]QuestionnaireResponse      `bson:"_revIncludedQuestionnaireResponseResourcesReferencingPartof,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingSupportinginfo        *[]ClinicalImpression         `bson:"_revIncludedClinicalImpressionResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
	RevIncludedNutritionIntakeResourcesReferencingPartof                   *[]NutritionIntake            `bson:"_revIncludedNutritionIntakeResourcesReferencingPartof,omitempty"`
}

func (p *ProcedurePlusRelatedResources) GetIncludedPractitionerResourceReferencedByPerformer() (practitioner *Practitioner, err error) {
	if p.IncludedPractitionerResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*p.IncludedPractitionerResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*p.IncludedPractitionerResourcesReferencedByPerformer))
	} else if len(*p.IncludedPractitionerResourcesReferencedByPerformer) == 1 {
		practitioner = &(*p.IncludedPractitionerResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedOrganizationResourceReferencedByPerformer() (organization *Organization, err error) {
	if p.IncludedOrganizationResourcesReferencedByPerformer == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*p.IncludedOrganizationResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*p.IncludedOrganizationResourcesReferencedByPerformer))
	} else if len(*p.IncludedOrganizationResourcesReferencedByPerformer) == 1 {
		organization = &(*p.IncludedOrganizationResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedCareTeamResourceReferencedByPerformer() (careTeam *CareTeam, err error) {
	if p.IncludedCareTeamResourcesReferencedByPerformer == nil {
		err = errors.New("Included careteams not requested")
	} else if len(*p.IncludedCareTeamResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 careTeam, but found %d", len(*p.IncludedCareTeamResourcesReferencedByPerformer))
	} else if len(*p.IncludedCareTeamResourcesReferencedByPerformer) == 1 {
		careTeam = &(*p.IncludedCareTeamResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedDeviceResourceReferencedByPerformer() (device *Device, err error) {
	if p.IncludedDeviceResourcesReferencedByPerformer == nil {
		err = errors.New("Included devices not requested")
	} else if len(*p.IncludedDeviceResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*p.IncludedDeviceResourcesReferencedByPerformer))
	} else if len(*p.IncludedDeviceResourcesReferencedByPerformer) == 1 {
		device = &(*p.IncludedDeviceResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedByPerformer() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByPerformer == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByPerformer))
	} else if len(*p.IncludedPatientResourcesReferencedByPerformer) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedHealthcareServiceResourceReferencedByPerformer() (healthcareService *HealthcareService, err error) {
	if p.IncludedHealthcareServiceResourcesReferencedByPerformer == nil {
		err = errors.New("Included healthcareservices not requested")
	} else if len(*p.IncludedHealthcareServiceResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 healthcareService, but found %d", len(*p.IncludedHealthcareServiceResourcesReferencedByPerformer))
	} else if len(*p.IncludedHealthcareServiceResourcesReferencedByPerformer) == 1 {
		healthcareService = &(*p.IncludedHealthcareServiceResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByPerformer() (practitionerRole *PractitionerRole, err error) {
	if p.IncludedPractitionerRoleResourcesReferencedByPerformer == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*p.IncludedPractitionerRoleResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*p.IncludedPractitionerRoleResourcesReferencedByPerformer))
	} else if len(*p.IncludedPractitionerRoleResourcesReferencedByPerformer) == 1 {
		practitionerRole = &(*p.IncludedPractitionerRoleResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByPerformer() (relatedPerson *RelatedPerson, err error) {
	if p.IncludedRelatedPersonResourcesReferencedByPerformer == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByPerformer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*p.IncludedRelatedPersonResourcesReferencedByPerformer))
	} else if len(*p.IncludedRelatedPersonResourcesReferencedByPerformer) == 1 {
		relatedPerson = &(*p.IncludedRelatedPersonResourcesReferencedByPerformer)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if p.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*p.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*p.IncludedGroupResourcesReferencedBySubject))
	} else if len(*p.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*p.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedBySubject))
	} else if len(*p.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedQuestionnaireResourcesReferencedByInstantiatescanonical() (questionnaires []Questionnaire, err error) {
	if p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included questionnaires not requested")
	} else {
		questionnaires = *p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedMeasureResourcesReferencedByInstantiatescanonical() (measures []Measure, err error) {
	if p.IncludedMeasureResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included measures not requested")
	} else {
		measures = *p.IncludedMeasureResourcesReferencedByInstantiatescanonical
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPlanDefinitionResourcesReferencedByInstantiatescanonical() (planDefinitions []PlanDefinition, err error) {
	if p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included planDefinitions not requested")
	} else {
		planDefinitions = *p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedOperationDefinitionResourcesReferencedByInstantiatescanonical() (operationDefinitions []OperationDefinition, err error) {
	if p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included operationDefinitions not requested")
	} else {
		operationDefinitions = *p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedActivityDefinitionResourcesReferencedByInstantiatescanonical() (activityDefinitions []ActivityDefinition, err error) {
	if p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical == nil {
		err = errors.New("Included activityDefinitions not requested")
	} else {
		activityDefinitions = *p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedObservationResourcesReferencedByPartof() (observations []Observation, err error) {
	if p.IncludedObservationResourcesReferencedByPartof == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *p.IncludedObservationResourcesReferencedByPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedProcedureResourcesReferencedByPartof() (procedures []Procedure, err error) {
	if p.IncludedProcedureResourcesReferencedByPartof == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *p.IncludedProcedureResourcesReferencedByPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedMedicationAdministrationResourcesReferencedByPartof() (medicationAdministrations []MedicationAdministration, err error) {
	if p.IncludedMedicationAdministrationResourcesReferencedByPartof == nil {
		err = errors.New("Included medicationAdministrations not requested")
	} else {
		medicationAdministrations = *p.IncludedMedicationAdministrationResourcesReferencedByPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if p.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*p.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*p.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*p.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*p.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedCarePlanResourcesReferencedByBasedon() (carePlans []CarePlan, err error) {
	if p.IncludedCarePlanResourcesReferencedByBasedon == nil {
		err = errors.New("Included carePlans not requested")
	} else {
		carePlans = *p.IncludedCarePlanResourcesReferencedByBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedServiceRequestResourcesReferencedByBasedon() (serviceRequests []ServiceRequest, err error) {
	if p.IncludedServiceRequestResourcesReferencedByBasedon == nil {
		err = errors.New("Included serviceRequests not requested")
	} else {
		serviceRequests = *p.IncludedServiceRequestResourcesReferencedByBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if p.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*p.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*p.IncludedPatientResourcesReferencedByPatient))
	} else if len(*p.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*p.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedConditionResourcesReferencedByReasonreference() (conditions []Condition, err error) {
	if p.IncludedConditionResourcesReferencedByReasonreference == nil {
		err = errors.New("Included conditions not requested")
	} else {
		conditions = *p.IncludedConditionResourcesReferencedByReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedObservationResourcesReferencedByReasonreference() (observations []Observation, err error) {
	if p.IncludedObservationResourcesReferencedByReasonreference == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *p.IncludedObservationResourcesReferencedByReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedProcedureResourcesReferencedByReasonreference() (procedures []Procedure, err error) {
	if p.IncludedProcedureResourcesReferencedByReasonreference == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *p.IncludedProcedureResourcesReferencedByReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedDiagnosticReportResourcesReferencedByReasonreference() (diagnosticReports []DiagnosticReport, err error) {
	if p.IncludedDiagnosticReportResourcesReferencedByReasonreference == nil {
		err = errors.New("Included diagnosticReports not requested")
	} else {
		diagnosticReports = *p.IncludedDiagnosticReportResourcesReferencedByReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedDocumentReferenceResourcesReferencedByReasonreference() (documentReferences []DocumentReference, err error) {
	if p.IncludedDocumentReferenceResourcesReferencedByReasonreference == nil {
		err = errors.New("Included documentReferences not requested")
	} else {
		documentReferences = *p.IncludedDocumentReferenceResourcesReferencedByReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedLocationResourceReferencedByLocation() (location *Location, err error) {
	if p.IncludedLocationResourcesReferencedByLocation == nil {
		err = errors.New("Included locations not requested")
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*p.IncludedLocationResourcesReferencedByLocation))
	} else if len(*p.IncludedLocationResourcesReferencedByLocation) == 1 {
		location = &(*p.IncludedLocationResourcesReferencedByLocation)[0]
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingReasonreference() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *p.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if p.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *p.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if p.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *p.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *p.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if p.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *p.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if p.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *p.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if p.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *p.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedChargeItemResourcesReferencingService() (chargeItems []ChargeItem, err error) {
	if p.RevIncludedChargeItemResourcesReferencingService == nil {
		err = errors.New("RevIncluded chargeItems not requested")
	} else {
		chargeItems = *p.RevIncludedChargeItemResourcesReferencingService
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEncounterResourcesReferencingDiagnosis() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingDiagnosis == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingDiagnosis
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEncounterResourcesReferencingReasonreference() (encounters []Encounter, err error) {
	if p.RevIncludedEncounterResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded encounters not requested")
	} else {
		encounters = *p.RevIncludedEncounterResourcesReferencingReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if p.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *p.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if p.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *p.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMedicationUsageResourcesReferencingPartof() (medicationUsages []MedicationUsage, err error) {
	if p.RevIncludedMedicationUsageResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded medicationUsages not requested")
	} else {
		medicationUsages = *p.RevIncludedMedicationUsageResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if p.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *p.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingPartof() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedProcedureResourcesReferencingReasonreference() (procedures []Procedure, err error) {
	if p.RevIncludedProcedureResourcesReferencingReasonreference == nil {
		err = errors.New("RevIncluded procedures not requested")
	} else {
		procedures = *p.RevIncludedProcedureResourcesReferencingReasonreference
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *p.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedFlagResourcesReferencingSubject() (flags []Flag, err error) {
	if p.RevIncludedFlagResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded flags not requested")
	} else {
		flags = *p.RevIncludedFlagResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedAdverseEventResourcesReferencingSubstance() (adverseEvents []AdverseEvent, err error) {
	if p.RevIncludedAdverseEventResourcesReferencingSubstance == nil {
		err = errors.New("RevIncluded adverseEvents not requested")
	} else {
		adverseEvents = *p.RevIncludedAdverseEventResourcesReferencingSubstance
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedObservationResourcesReferencingPartof() (observations []Observation, err error) {
	if p.RevIncludedObservationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *p.RevIncludedObservationResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if p.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *p.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *p.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if p.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *p.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if p.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *p.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingPartof() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedNutritionIntakeResourcesReferencingPartof() (nutritionIntakes []NutritionIntake, err error) {
	if p.RevIncludedNutritionIntakeResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded nutritionIntakes not requested")
	} else {
		nutritionIntakes = *p.RevIncludedNutritionIntakeResourcesReferencingPartof
	}
	return
}

func (p *ProcedurePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*p.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*p.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*p.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*p.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*p.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedGroupResourcesReferencedBySubject {
			rsc := (*p.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedPatientResourcesReferencedBySubject {
			rsc := (*p.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedObservationResourcesReferencedByPartof {
			rsc := (*p.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*p.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*p.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*p.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*p.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*p.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPatient {
			rsc := (*p.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedConditionResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedConditionResourcesReferencedByReasonreference {
			rsc := (*p.IncludedConditionResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedObservationResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedObservationResourcesReferencedByReasonreference {
			rsc := (*p.IncludedObservationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedProcedureResourcesReferencedByReasonreference {
			rsc := (*p.IncludedProcedureResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDiagnosticReportResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedDiagnosticReportResourcesReferencedByReasonreference {
			rsc := (*p.IncludedDiagnosticReportResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDocumentReferenceResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedDocumentReferenceResourcesReferencedByReasonreference {
			rsc := (*p.IncludedDocumentReferenceResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *ProcedurePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedAppointmentResourcesReferencingReasonreference != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingReasonreference {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingDiagnosis != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingDiagnosis {
			rsc := (*p.RevIncludedEncounterResourcesReferencingDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingReasonreference != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingReasonreference {
			rsc := (*p.RevIncludedEncounterResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*p.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationUsageResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedMedicationUsageResourcesReferencingPartof {
			rsc := (*p.RevIncludedMedicationUsageResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingReasonreference != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingReasonreference {
			rsc := (*p.RevIncludedProcedureResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*p.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPartof {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionIntakeResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedNutritionIntakeResourcesReferencingPartof {
			rsc := (*p.RevIncludedNutritionIntakeResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (p *ProcedurePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedPractitionerResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOrganizationResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedOrganizationResourcesReferencedByPerformer {
			rsc := (*p.IncludedOrganizationResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedCareTeamResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedCareTeamResourcesReferencedByPerformer {
			rsc := (*p.IncludedCareTeamResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDeviceResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedDeviceResourcesReferencedByPerformer {
			rsc := (*p.IncludedDeviceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPerformer {
			rsc := (*p.IncludedPatientResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedHealthcareServiceResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedHealthcareServiceResourcesReferencedByPerformer {
			rsc := (*p.IncludedHealthcareServiceResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPractitionerRoleResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedPractitionerRoleResourcesReferencedByPerformer {
			rsc := (*p.IncludedPractitionerRoleResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedRelatedPersonResourcesReferencedByPerformer != nil {
		for idx := range *p.IncludedRelatedPersonResourcesReferencedByPerformer {
			rsc := (*p.IncludedRelatedPersonResourcesReferencedByPerformer)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedGroupResourcesReferencedBySubject {
			rsc := (*p.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *p.IncludedPatientResourcesReferencedBySubject {
			rsc := (*p.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedQuestionnaireResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedMeasureResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedMeasureResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedMeasureResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedPlanDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedOperationDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical != nil {
		for idx := range *p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical {
			rsc := (*p.IncludedActivityDefinitionResourcesReferencedByInstantiatescanonical)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedObservationResourcesReferencedByPartof {
			rsc := (*p.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*p.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedMedicationAdministrationResourcesReferencedByPartof != nil {
		for idx := range *p.IncludedMedicationAdministrationResourcesReferencedByPartof {
			rsc := (*p.IncludedMedicationAdministrationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *p.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*p.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedCarePlanResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedCarePlanResourcesReferencedByBasedon {
			rsc := (*p.IncludedCarePlanResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedServiceRequestResourcesReferencedByBasedon != nil {
		for idx := range *p.IncludedServiceRequestResourcesReferencedByBasedon {
			rsc := (*p.IncludedServiceRequestResourcesReferencedByBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *p.IncludedPatientResourcesReferencedByPatient {
			rsc := (*p.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedConditionResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedConditionResourcesReferencedByReasonreference {
			rsc := (*p.IncludedConditionResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedObservationResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedObservationResourcesReferencedByReasonreference {
			rsc := (*p.IncludedObservationResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedProcedureResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedProcedureResourcesReferencedByReasonreference {
			rsc := (*p.IncludedProcedureResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDiagnosticReportResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedDiagnosticReportResourcesReferencedByReasonreference {
			rsc := (*p.IncludedDiagnosticReportResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedDocumentReferenceResourcesReferencedByReasonreference != nil {
		for idx := range *p.IncludedDocumentReferenceResourcesReferencedByReasonreference {
			rsc := (*p.IncludedDocumentReferenceResourcesReferencedByReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.IncludedLocationResourcesReferencedByLocation != nil {
		for idx := range *p.IncludedLocationResourcesReferencedByLocation {
			rsc := (*p.IncludedLocationResourcesReferencedByLocation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingReasonreference != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingReasonreference {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*p.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*p.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *p.RevIncludedConsentResourcesReferencingData {
			rsc := (*p.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*p.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *p.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*p.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*p.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*p.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedContractResourcesReferencingSubject {
			rsc := (*p.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*p.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *p.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*p.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedChargeItemResourcesReferencingService != nil {
		for idx := range *p.RevIncludedChargeItemResourcesReferencingService {
			rsc := (*p.RevIncludedChargeItemResourcesReferencingService)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingDiagnosis != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingDiagnosis {
			rsc := (*p.RevIncludedEncounterResourcesReferencingDiagnosis)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEncounterResourcesReferencingReasonreference != nil {
		for idx := range *p.RevIncludedEncounterResourcesReferencingReasonreference {
			rsc := (*p.RevIncludedEncounterResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*p.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *p.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*p.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMedicationUsageResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedMedicationUsageResourcesReferencingPartof {
			rsc := (*p.RevIncludedMedicationUsageResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *p.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*p.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*p.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *p.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*p.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*p.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*p.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*p.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*p.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingPartof {
			rsc := (*p.RevIncludedProcedureResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedProcedureResourcesReferencingReasonreference != nil {
		for idx := range *p.RevIncludedProcedureResourcesReferencingReasonreference {
			rsc := (*p.RevIncludedProcedureResourcesReferencingReasonreference)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *p.RevIncludedListResourcesReferencingItem {
			rsc := (*p.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedFlagResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedFlagResourcesReferencingSubject {
			rsc := (*p.RevIncludedFlagResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAdverseEventResourcesReferencingSubstance != nil {
		for idx := range *p.RevIncludedAdverseEventResourcesReferencingSubstance {
			rsc := (*p.RevIncludedAdverseEventResourcesReferencingSubstance)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*p.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedObservationResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedObservationResourcesReferencingPartof {
			rsc := (*p.RevIncludedObservationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*p.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*p.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *p.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*p.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*p.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *p.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*p.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *p.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*p.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *p.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*p.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*p.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *p.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*p.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*p.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedQuestionnaireResponseResourcesReferencingPartof {
			rsc := (*p.RevIncludedQuestionnaireResponseResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*p.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*p.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if p.RevIncludedNutritionIntakeResourcesReferencingPartof != nil {
		for idx := range *p.RevIncludedNutritionIntakeResourcesReferencingPartof {
			rsc := (*p.RevIncludedNutritionIntakeResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
