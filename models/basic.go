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

type Basic struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code           *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Subject        *Reference       `bson:"subject,omitempty" json:"subject,omitempty"`
	Created        *FHIRDateTime    `bson:"created,omitempty" json:"created,omitempty"`
	Author         *Reference       `bson:"author,omitempty" json:"author,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Basic) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Basic"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Basic), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Basic) GetBSON() (interface{}, error) {
	x.ResourceType = "Basic"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "basic" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type basic Basic

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Basic) UnmarshalJSON(data []byte) (err error) {
	x2 := basic{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Basic(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Basic) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Basic"
	} else if x.ResourceType != "Basic" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Basic, instead received %s", x.ResourceType))
	}
	return nil
}

type BasicPlus struct {
	Basic                     `bson:",inline"`
	BasicPlusRelatedResources `bson:",inline"`
}

type BasicPlusRelatedResources struct {
	IncludedPatientResourcesReferencedByPatient                            *[]Patient                    `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor                        *[]Practitioner               `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor                        *[]Organization               `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                             *[]Patient                    `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedPractitionerRoleResourcesReferencedByAuthor                    *[]PractitionerRole           `bson:"_includedPractitionerRoleResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor                       *[]RelatedPerson              `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
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

func (b *BasicPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if b.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*b.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*b.IncludedPatientResourcesReferencedByPatient))
	} else if len(*b.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*b.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (b *BasicPlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if b.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*b.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*b.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*b.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*b.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (b *BasicPlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if b.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*b.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*b.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*b.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*b.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (b *BasicPlusRelatedResources) GetIncludedPatientResourceReferencedByAuthor() (patient *Patient, err error) {
	if b.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else if len(*b.IncludedPatientResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*b.IncludedPatientResourcesReferencedByAuthor))
	} else if len(*b.IncludedPatientResourcesReferencedByAuthor) == 1 {
		patient = &(*b.IncludedPatientResourcesReferencedByAuthor)[0]
	}
	return
}

func (b *BasicPlusRelatedResources) GetIncludedPractitionerRoleResourceReferencedByAuthor() (practitionerRole *PractitionerRole, err error) {
	if b.IncludedPractitionerRoleResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitionerroles not requested")
	} else if len(*b.IncludedPractitionerRoleResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitionerRole, but found %d", len(*b.IncludedPractitionerRoleResourcesReferencedByAuthor))
	} else if len(*b.IncludedPractitionerRoleResourcesReferencedByAuthor) == 1 {
		practitionerRole = &(*b.IncludedPractitionerRoleResourcesReferencedByAuthor)[0]
	}
	return
}

func (b *BasicPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByAuthor() (relatedPerson *RelatedPerson, err error) {
	if b.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*b.IncludedRelatedPersonResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*b.IncludedRelatedPersonResourcesReferencedByAuthor))
	} else if len(*b.IncludedRelatedPersonResourcesReferencedByAuthor) == 1 {
		relatedPerson = &(*b.IncludedRelatedPersonResourcesReferencedByAuthor)[0]
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if b.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *b.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if b.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *b.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *b.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if b.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *b.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if b.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *b.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if b.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *b.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if b.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *b.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if b.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *b.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *b.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *b.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if b.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *b.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if b.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *b.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if b.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *b.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if b.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *b.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if b.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *b.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if b.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *b.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if b.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *b.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if b.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *b.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if b.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *b.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if b.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *b.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if b.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *b.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if b.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *b.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if b.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *b.RevIncludedListResourcesReferencingItem
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if b.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *b.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if b.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *b.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if b.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *b.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if b.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *b.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if b.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *b.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if b.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *b.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if b.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *b.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if b.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *b.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if b.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *b.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *b.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (b *BasicPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (b *BasicPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *b.IncludedPatientResourcesReferencedByPatient {
			rsc := (*b.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*b.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*b.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*b.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedPractitionerRoleResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedPractitionerRoleResourcesReferencedByAuthor {
			rsc := (*b.IncludedPractitionerRoleResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*b.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (b *BasicPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *b.RevIncludedConsentResourcesReferencingData {
			rsc := (*b.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*b.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *b.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*b.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*b.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingSubject {
			rsc := (*b.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *b.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*b.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*b.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*b.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*b.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *b.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*b.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*b.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*b.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*b.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedListResourcesReferencingItem {
			rsc := (*b.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*b.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*b.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*b.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*b.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *b.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*b.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*b.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*b.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*b.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*b.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (b *BasicPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if b.IncludedPatientResourcesReferencedByPatient != nil {
		for idx := range *b.IncludedPatientResourcesReferencedByPatient {
			rsc := (*b.IncludedPatientResourcesReferencedByPatient)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedPractitionerResourcesReferencedByAuthor {
			rsc := (*b.IncludedPractitionerResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedOrganizationResourcesReferencedByAuthor {
			rsc := (*b.IncludedOrganizationResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedPatientResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedPatientResourcesReferencedByAuthor {
			rsc := (*b.IncludedPatientResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedPractitionerRoleResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedPractitionerRoleResourcesReferencedByAuthor {
			rsc := (*b.IncludedPractitionerRoleResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for idx := range *b.IncludedRelatedPersonResourcesReferencedByAuthor {
			rsc := (*b.IncludedRelatedPersonResourcesReferencedByAuthor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*b.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *b.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*b.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *b.RevIncludedConsentResourcesReferencingData {
			rsc := (*b.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*b.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *b.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*b.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*b.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*b.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedContractResourcesReferencingSubject {
			rsc := (*b.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *b.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*b.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *b.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*b.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*b.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *b.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*b.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *b.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*b.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*b.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *b.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*b.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *b.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*b.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*b.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*b.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*b.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *b.RevIncludedListResourcesReferencingItem {
			rsc := (*b.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *b.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*b.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*b.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*b.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *b.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*b.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*b.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *b.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*b.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *b.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*b.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *b.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*b.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*b.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *b.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*b.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *b.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*b.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *b.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*b.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*b.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*b.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
