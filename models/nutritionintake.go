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

type NutritionIntake struct {
	DomainResource    `bson:",inline"`
	Identifier        []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference                               `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf            []Reference                               `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status            string                                    `bson:"status,omitempty" json:"status,omitempty"`
	StatusReason      []CodeableConcept                         `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category          []CodeableConcept                         `bson:"category,omitempty" json:"category,omitempty"`
	ConsumedItem      []NutritionIntakeConsumedItemComponent    `bson:"consumedItem,omitempty" json:"consumedItem,omitempty"`
	IngredientLabel   []NutritionIntakeIngredientLabelComponent `bson:"ingredientLabel,omitempty" json:"ingredientLabel,omitempty"`
	Subject           *Reference                                `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference                                `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime *FHIRDateTime                             `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	DateAsserted      *FHIRDateTime                             `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	InformationSource *Reference                                `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	DerivedFrom       []Reference                               `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	ReasonCode        []CodeableConcept                         `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference   []Reference                               `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note              []Annotation                              `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *NutritionIntake) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "NutritionIntake"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to NutritionIntake), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *NutritionIntake) GetBSON() (interface{}, error) {
	x.ResourceType = "NutritionIntake"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "nutritionIntake" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type nutritionIntake NutritionIntake

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *NutritionIntake) UnmarshalJSON(data []byte) (err error) {
	x2 := nutritionIntake{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = NutritionIntake(x2)
		return x.checkResourceType()
	}
	return
}

func (x *NutritionIntake) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "NutritionIntake"
	} else if x.ResourceType != "NutritionIntake" {
		return errors.New(fmt.Sprintf("Expected resourceType to be NutritionIntake, instead received %s", x.ResourceType))
	}
	return nil
}

type NutritionIntakeConsumedItemComponent struct {
	BackboneElement   `bson:",inline"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	NutritionProduct  *CodeableConcept `bson:"nutritionProduct,omitempty" json:"nutritionProduct,omitempty"`
	Schedule          *Timing          `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Amount            *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
	Rate              *Quantity        `bson:"rate,omitempty" json:"rate,omitempty"`
	NotConsumed       *bool            `bson:"notConsumed,omitempty" json:"notConsumed,omitempty"`
	NotConsumedReason *CodeableConcept `bson:"notConsumedReason,omitempty" json:"notConsumedReason,omitempty"`
}

type NutritionIntakeIngredientLabelComponent struct {
	BackboneElement `bson:",inline"`
	Nutrient        *CodeableConcept `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Amount          *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}

type NutritionIntakePlus struct {
	NutritionIntake                     `bson:",inline"`
	NutritionIntakePlusRelatedResources `bson:",inline"`
}

type NutritionIntakePlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                              *[]Group                      `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                            *[]Patient                    `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedObservationResourcesReferencedByPartof                         *[]Observation                `bson:"_includedObservationResourcesReferencedByPartof,omitempty"`
	IncludedProcedureResourcesReferencedByPartof                           *[]Procedure                  `bson:"_includedProcedureResourcesReferencedByPartof,omitempty"`
	IncludedNutritionIntakeResourcesReferencedByPartof                     *[]NutritionIntake            `bson:"_includedNutritionIntakeResourcesReferencedByPartof,omitempty"`
	IncludedPractitionerResourcesReferencedBySource                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedBySource,omitempty"`
	IncludedOrganizationResourcesReferencedBySource                        *[]Organization               `bson:"_includedOrganizationResourcesReferencedBySource,omitempty"`
	IncludedPatientResourcesReferencedBySource                             *[]Patient                    `bson:"_includedPatientResourcesReferencedBySource,omitempty"`
	IncludedPractitionerRoleResourcesReferencedBySource                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedBySource,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySource                       *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedBySource,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter                        *[]Encounter                  `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
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
	RevIncludedNutritionIntakeResourcesReferencingPartof                   *[]NutritionIntake            `bson:"_revIncludedNutritionIntakeResourcesReferencingPartof,omitempty"`
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if n.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*n.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*n.IncludedGroupResourcesReferencedBySubject))
	} else if len(*n.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*n.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if n.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*n.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*n.IncludedPatientResourcesReferencedBySubject))
	} else if len(*n.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*n.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if n.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*n.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*n.IncludedPatientResourcesReferencedByPatient))
	} else if len(*n.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*n.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedObservationResourcesReferencedByPartof() (observations []Observation, err error) {
	if n.IncludedObservationResourcesReferencedByPartof == nil {
		err = errors.New("Included observations not requested")
	} else {
		observations = *n.IncludedObservationResourcesReferencedByPartof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedProcedureResourcesReferencedByPartof() (procedures []Procedure, err error) {
	if n.IncludedProcedureResourcesReferencedByPartof == nil {
		err = errors.New("Included procedures not requested")
	} else {
		procedures = *n.IncludedProcedureResourcesReferencedByPartof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedNutritionIntakeResourcesReferencedByPartof() (nutritionIntakes []NutritionIntake, err error) {
	if n.IncludedNutritionIntakeResourcesReferencedByPartof == nil {
		err = errors.New("Included nutritionIntakes not requested")
	} else {
		nutritionIntakes = *n.IncludedNutritionIntakeResourcesReferencedByPartof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedPractitionerResourceReferencedBySource() (practitioner *Practitioner, err error) {
	if n.IncludedPractitionerResourcesReferencedBySource == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*n.IncludedPractitionerResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*n.IncludedPractitionerResourcesReferencedBySource))
	} else if len(*n.IncludedPractitionerResourcesReferencedBySource) == 1 {
		practitioner = &(*n.IncludedPractitionerResourcesReferencedBySource)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedOrganizationResourceReferencedBySource() (organization *Organization, err error) {
	if n.IncludedOrganizationResourcesReferencedBySource == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*n.IncludedOrganizationResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*n.IncludedOrganizationResourcesReferencedBySource))
	} else if len(*n.IncludedOrganizationResourcesReferencedBySource) == 1 {
		organization = &(*n.IncludedOrganizationResourcesReferencedBySource)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedPatientResourceReferencedBySource() (patient *Patient, err error) {
	if n.IncludedPatientResourcesReferencedBySource == nil {
		err = errors.New("Included patients not requested")
	} else if len(*n.IncludedPatientResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*n.IncludedPatientResourcesReferencedBySource))
	} else if len(*n.IncludedPatientResourcesReferencedBySource) == 1 {
		patient = &(*n.IncludedPatientResourcesReferencedBySource)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedBySource() (practitionerRole *PractitionerRole, err error) {
	if n.IncludedPractitionerRoleResourcesReferencedBySource == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*n.IncludedPractitionerRoleResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*n.IncludedPractitionerRoleResourcesReferencedBySource))
	} else if len(*n.IncludedPractitionerRoleResourcesReferencedBySource) == 1 {
		practitionerRole = &(*n.IncludedPractitionerRoleResourcesReferencedBySource)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySource() (relatedPerson *RelatedPerson, err error) {
	if n.IncludedRelatedPersonResourcesReferencedBySource == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*n.IncludedRelatedPersonResourcesReferencedBySource) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*n.IncludedRelatedPersonResourcesReferencedBySource))
	} else if len(*n.IncludedRelatedPersonResourcesReferencedBySource) == 1 {
		relatedPerson = &(*n.IncludedRelatedPersonResourcesReferencedBySource)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if n.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*n.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*n.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*n.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*n.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *n.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if n.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *n.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *n.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if n.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *n.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *n.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *n.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if n.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *n.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if n.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *n.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *n.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *n.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if n.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *n.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if n.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *n.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if n.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *n.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if n.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *n.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if n.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *n.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *n.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if n.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *n.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if n.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *n.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if n.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *n.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if n.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *n.RevIncludedListResourcesReferencingItem
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *n.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if n.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *n.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if n.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *n.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *n.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if n.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *n.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if n.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *n.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if n.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *n.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if n.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *n.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if n.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *n.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *n.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedNutritionIntakeResourcesReferencingPartof() (nutritionIntakes []NutritionIntake, err error) {
	if n.RevIncludedNutritionIntakeResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded nutritionIntakes not requested")
	} else {
		nutritionIntakes = *n.RevIncludedNutritionIntakeResourcesReferencingPartof
	}
	return
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *n.IncludedGroupResourcesReferencedBySubject {
			rsc := (*n.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *n.IncludedPatientResourcesReferencedBySubject {
			rsc := (*n.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *n.IncludedPatientResourcesReferencedByPatient {
			rsc := (*n.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *n.IncludedObservationResourcesReferencedByPartof {
			rsc := (*n.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *n.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*n.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedNutritionIntakeResourcesReferencedByPartof != nil {
		for idx := range *n.IncludedNutritionIntakeResourcesReferencedByPartof {
			rsc := (*n.IncludedNutritionIntakeResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *n.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*n.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *n.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*n.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *n.IncludedPatientResourcesReferencedBySource {
			rsc := (*n.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *n.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*n.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *n.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*n.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *n.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*n.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (n *NutritionIntakePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *n.RevIncludedConsentResourcesReferencingData {
			rsc := (*n.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*n.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*n.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedContractResourcesReferencingSubject {
			rsc := (*n.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *n.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*n.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*n.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*n.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*n.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*n.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*n.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*n.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*n.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*n.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *n.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*n.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedNutritionIntakeResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedNutritionIntakeResourcesReferencingPartof {
			rsc := (*n.RevIncludedNutritionIntakeResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (n *NutritionIntakePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if n.IncludedGroupResourcesReferencedBySubject != nil {
		for idx := range *n.IncludedGroupResourcesReferencedBySubject {
			rsc := (*n.IncludedGroupResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedBySubject != nil {
		for idx := range *n.IncludedPatientResourcesReferencedBySubject {
			rsc := (*n.IncludedPatientResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *n.IncludedPatientResourcesReferencedByPatient {
			rsc := (*n.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedObservationResourcesReferencedByPartof != nil {
		for idx := range *n.IncludedObservationResourcesReferencedByPartof {
			rsc := (*n.IncludedObservationResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedProcedureResourcesReferencedByPartof != nil {
		for idx := range *n.IncludedProcedureResourcesReferencedByPartof {
			rsc := (*n.IncludedProcedureResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedNutritionIntakeResourcesReferencedByPartof != nil {
		for idx := range *n.IncludedNutritionIntakeResourcesReferencedByPartof {
			rsc := (*n.IncludedNutritionIntakeResourcesReferencedByPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPractitionerResourcesReferencedBySource != nil {
		for idx := range *n.IncludedPractitionerResourcesReferencedBySource {
			rsc := (*n.IncludedPractitionerResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedOrganizationResourcesReferencedBySource != nil {
		for idx := range *n.IncludedOrganizationResourcesReferencedBySource {
			rsc := (*n.IncludedOrganizationResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPatientResourcesReferencedBySource != nil {
		for idx := range *n.IncludedPatientResourcesReferencedBySource {
			rsc := (*n.IncludedPatientResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedPractitionerRoleResourcesReferencedBySource != nil {
		for idx := range *n.IncludedPractitionerRoleResourcesReferencedBySource {
			rsc := (*n.IncludedPractitionerRoleResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedRelatedPersonResourcesReferencedBySource != nil {
		for idx := range *n.IncludedRelatedPersonResourcesReferencedBySource {
			rsc := (*n.IncludedRelatedPersonResourcesReferencedBySource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.IncludedEncounterResourcesReferencedByEncounter != nil {
		for idx := range *n.IncludedEncounterResourcesReferencedByEncounter {
			rsc := (*n.IncludedEncounterResourcesReferencedByEncounter)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*n.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *n.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*n.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *n.RevIncludedConsentResourcesReferencingData {
			rsc := (*n.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*n.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *n.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*n.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*n.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*n.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedContractResourcesReferencingSubject {
			rsc := (*n.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *n.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*n.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *n.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*n.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*n.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *n.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*n.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *n.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*n.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*n.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *n.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*n.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *n.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*n.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*n.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*n.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*n.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *n.RevIncludedListResourcesReferencingItem {
			rsc := (*n.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *n.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*n.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*n.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*n.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *n.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*n.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*n.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *n.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*n.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *n.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*n.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *n.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*n.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*n.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *n.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*n.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *n.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*n.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *n.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*n.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*n.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*n.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if n.RevIncludedNutritionIntakeResourcesReferencingPartof != nil {
		for idx := range *n.RevIncludedNutritionIntakeResourcesReferencingPartof {
			rsc := (*n.RevIncludedNutritionIntakeResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
