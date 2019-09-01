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

type RegulatedAuthorization struct {
	DomainResource              `bson:",inline"`
	Identifier                  []Identifier                                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject                     *Reference                                   `bson:"subject,omitempty" json:"subject,omitempty"`
	Type                        *CodeableConcept                             `bson:"type,omitempty" json:"type,omitempty"`
	Description                 string                                       `bson:"description,omitempty" json:"description,omitempty"`
	Region                      []CodeableConcept                            `bson:"region,omitempty" json:"region,omitempty"`
	Status                      *CodeableConcept                             `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate                  *FHIRDateTime                                `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	ValidityPeriod              *Period                                      `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	Basis                       []CodeableConcept                            `bson:"basis,omitempty" json:"basis,omitempty"`
	RelatedDate                 []RegulatedAuthorizationRelatedDateComponent `bson:"relatedDate,omitempty" json:"relatedDate,omitempty"`
	JurisdictionalAuthorization []Reference                                  `bson:"jurisdictionalAuthorization,omitempty" json:"jurisdictionalAuthorization,omitempty"`
	Holder                      *Reference                                   `bson:"holder,omitempty" json:"holder,omitempty"`
	Regulator                   *Reference                                   `bson:"regulator,omitempty" json:"regulator,omitempty"`
	Case                        *RegulatedAuthorizationCaseComponent         `bson:"case,omitempty" json:"case,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *RegulatedAuthorization) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "RegulatedAuthorization"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to RegulatedAuthorization), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *RegulatedAuthorization) GetBSON() (interface{}, error) {
	x.ResourceType = "RegulatedAuthorization"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "regulatedAuthorization" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type regulatedAuthorization RegulatedAuthorization

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *RegulatedAuthorization) UnmarshalJSON(data []byte) (err error) {
	x2 := regulatedAuthorization{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = RegulatedAuthorization(x2)
		return x.checkResourceType()
	}
	return
}

func (x *RegulatedAuthorization) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "RegulatedAuthorization"
	} else if x.ResourceType != "RegulatedAuthorization" {
		return errors.New(fmt.Sprintf("Expected resourceType to be RegulatedAuthorization, instead received %s", x.ResourceType))
	}
	return nil
}

type RegulatedAuthorizationRelatedDateComponent struct {
	BackboneElement `bson:",inline"`
	DatePeriod      *Period          `bson:"datePeriod,omitempty" json:"datePeriod,omitempty"`
	DateDateTime    *FHIRDateTime    `bson:"dateDateTime,omitempty" json:"dateDateTime,omitempty"`
	Type            *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}

type RegulatedAuthorizationCaseComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type            *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	Status          *CodeableConcept                      `bson:"status,omitempty" json:"status,omitempty"`
	DatePeriod      *Period                               `bson:"datePeriod,omitempty" json:"datePeriod,omitempty"`
	DateDateTime    *FHIRDateTime                         `bson:"dateDateTime,omitempty" json:"dateDateTime,omitempty"`
	Application     []RegulatedAuthorizationCaseComponent `bson:"application,omitempty" json:"application,omitempty"`
}

type RegulatedAuthorizationPlus struct {
	RegulatedAuthorization                     `bson:",inline"`
	RegulatedAuthorizationPlusRelatedResources `bson:",inline"`
}

type RegulatedAuthorizationPlusRelatedResources struct {
	IncludedMedicinalProductDefinitionResourcesReferencedBySubject         *[]MedicinalProductDefinition `bson:"_includedMedicinalProductDefinitionResourcesReferencedBySubject,omitempty"`
	IncludedDeviceDefinitionResourcesReferencedBySubject                   *[]DeviceDefinition           `bson:"_includedDeviceDefinitionResourcesReferencedBySubject,omitempty"`
	IncludedPlanDefinitionResourcesReferencedBySubject                     *[]PlanDefinition             `bson:"_includedPlanDefinitionResourcesReferencedBySubject,omitempty"`
	IncludedPackagedProductDefinitionResourcesReferencedBySubject          *[]PackagedProductDefinition  `bson:"_includedPackagedProductDefinitionResourcesReferencedBySubject,omitempty"`
	IncludedResearchStudyResourcesReferencedBySubject                      *[]ResearchStudy              `bson:"_includedResearchStudyResourcesReferencedBySubject,omitempty"`
	IncludedActivityDefinitionResourcesReferencedBySubject                 *[]ActivityDefinition         `bson:"_includedActivityDefinitionResourcesReferencedBySubject,omitempty"`
	IncludedOrganizationResourcesReferencedByHolder                        *[]Organization               `bson:"_includedOrganizationResourcesReferencedByHolder,omitempty"`
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

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedMedicinalProductDefinitionResourceReferencedBySubject() (medicinalProductDefinition *MedicinalProductDefinition, err error) {
	if r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject == nil {
		err = errors.New("Included medicinalproductdefinitions not requested")
	} else if len(*r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medicinalProductDefinition, but found %d", len(*r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject))
	} else if len(*r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject) == 1 {
		medicinalProductDefinition = &(*r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedDeviceDefinitionResourceReferencedBySubject() (deviceDefinition *DeviceDefinition, err error) {
	if r.IncludedDeviceDefinitionResourcesReferencedBySubject == nil {
		err = errors.New("Included devicedefinitions not requested")
	} else if len(*r.IncludedDeviceDefinitionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceDefinition, but found %d", len(*r.IncludedDeviceDefinitionResourcesReferencedBySubject))
	} else if len(*r.IncludedDeviceDefinitionResourcesReferencedBySubject) == 1 {
		deviceDefinition = &(*r.IncludedDeviceDefinitionResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedPlanDefinitionResourceReferencedBySubject() (planDefinition *PlanDefinition, err error) {
	if r.IncludedPlanDefinitionResourcesReferencedBySubject == nil {
		err = errors.New("Included plandefinitions not requested")
	} else if len(*r.IncludedPlanDefinitionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 planDefinition, but found %d", len(*r.IncludedPlanDefinitionResourcesReferencedBySubject))
	} else if len(*r.IncludedPlanDefinitionResourcesReferencedBySubject) == 1 {
		planDefinition = &(*r.IncludedPlanDefinitionResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedPackagedProductDefinitionResourceReferencedBySubject() (packagedProductDefinition *PackagedProductDefinition, err error) {
	if r.IncludedPackagedProductDefinitionResourcesReferencedBySubject == nil {
		err = errors.New("Included packagedproductdefinitions not requested")
	} else if len(*r.IncludedPackagedProductDefinitionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 packagedProductDefinition, but found %d", len(*r.IncludedPackagedProductDefinitionResourcesReferencedBySubject))
	} else if len(*r.IncludedPackagedProductDefinitionResourcesReferencedBySubject) == 1 {
		packagedProductDefinition = &(*r.IncludedPackagedProductDefinitionResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedResearchStudyResourceReferencedBySubject() (researchStudy *ResearchStudy, err error) {
	if r.IncludedResearchStudyResourcesReferencedBySubject == nil {
		err = errors.New("Included researchstudies not requested")
	} else if len(*r.IncludedResearchStudyResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 researchStudy, but found %d", len(*r.IncludedResearchStudyResourcesReferencedBySubject))
	} else if len(*r.IncludedResearchStudyResourcesReferencedBySubject) == 1 {
		researchStudy = &(*r.IncludedResearchStudyResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedActivityDefinitionResourceReferencedBySubject() (activityDefinition *ActivityDefinition, err error) {
	if r.IncludedActivityDefinitionResourcesReferencedBySubject == nil {
		err = errors.New("Included activitydefinitions not requested")
	} else if len(*r.IncludedActivityDefinitionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 activityDefinition, but found %d", len(*r.IncludedActivityDefinitionResourcesReferencedBySubject))
	} else if len(*r.IncludedActivityDefinitionResourcesReferencedBySubject) == 1 {
		activityDefinition = &(*r.IncludedActivityDefinitionResourcesReferencedBySubject)[0]
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedOrganizationResourceReferencedByHolder() (organization *Organization, err error) {
	if r.IncludedOrganizationResourcesReferencedByHolder == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*r.IncludedOrganizationResourcesReferencedByHolder) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*r.IncludedOrganizationResourcesReferencedByHolder))
	} else if len(*r.IncludedOrganizationResourcesReferencedByHolder) == 1 {
		organization = &(*r.IncludedOrganizationResourcesReferencedByHolder)[0]
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *r.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if r.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *r.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *r.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if r.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *r.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *r.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *r.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if r.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *r.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if r.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *r.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *r.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if r.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *r.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if r.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *r.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if r.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *r.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if r.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *r.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if r.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *r.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if r.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *r.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if r.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *r.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if r.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *r.RevIncludedListResourcesReferencingItem
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *r.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if r.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *r.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if r.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *r.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *r.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if r.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *r.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if r.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *r.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if r.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *r.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if r.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *r.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if r.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *r.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *r.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedDeviceDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedDeviceDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedDeviceDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPlanDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedPlanDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedPlanDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPackagedProductDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedPackagedProductDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedPackagedProductDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedResearchStudyResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedResearchStudyResourcesReferencedBySubject {
			rsc := (*r.IncludedResearchStudyResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedActivityDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedActivityDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedActivityDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedByHolder != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedByHolder {
			rsc := (*r.IncludedOrganizationResourcesReferencedByHolder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (r *RegulatedAuthorizationPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedMedicinalProductDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedDeviceDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedDeviceDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedDeviceDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPlanDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedPlanDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedPlanDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedPackagedProductDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedPackagedProductDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedPackagedProductDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedResearchStudyResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedResearchStudyResourcesReferencedBySubject {
			rsc := (*r.IncludedResearchStudyResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedActivityDefinitionResourcesReferencedBySubject != nil {
		for idx := range *r.IncludedActivityDefinitionResourcesReferencedBySubject {
			rsc := (*r.IncludedActivityDefinitionResourcesReferencedBySubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.IncludedOrganizationResourcesReferencedByHolder != nil {
		for idx := range *r.IncludedOrganizationResourcesReferencedByHolder {
			rsc := (*r.IncludedOrganizationResourcesReferencedByHolder)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*r.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *r.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*r.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *r.RevIncludedConsentResourcesReferencingData {
			rsc := (*r.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*r.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *r.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*r.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*r.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*r.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedContractResourcesReferencingSubject {
			rsc := (*r.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *r.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*r.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *r.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*r.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*r.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *r.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*r.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *r.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*r.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*r.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *r.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*r.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *r.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*r.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*r.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*r.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*r.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *r.RevIncludedListResourcesReferencingItem {
			rsc := (*r.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *r.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*r.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*r.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*r.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *r.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*r.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*r.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *r.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*r.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *r.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*r.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *r.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*r.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*r.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *r.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*r.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *r.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*r.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *r.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*r.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*r.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*r.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
