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

type Endpoint struct {
	DomainResource       `bson:",inline"`
	Identifier           []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               string            `bson:"status,omitempty" json:"status,omitempty"`
	ConnectionType       *Coding           `bson:"connectionType,omitempty" json:"connectionType,omitempty"`
	Name                 string            `bson:"name,omitempty" json:"name,omitempty"`
	ManagingOrganization *Reference        `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Contact              []ContactPoint    `bson:"contact,omitempty" json:"contact,omitempty"`
	Period               *Period           `bson:"period,omitempty" json:"period,omitempty"`
	PayloadType          []CodeableConcept `bson:"payloadType,omitempty" json:"payloadType,omitempty"`
	PayloadMimeType      []string          `bson:"payloadMimeType,omitempty" json:"payloadMimeType,omitempty"`
	Address              string            `bson:"address,omitempty" json:"address,omitempty"`
	Header               []string          `bson:"header,omitempty" json:"header,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Endpoint) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Endpoint"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Endpoint), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Endpoint) GetBSON() (interface{}, error) {
	x.ResourceType = "Endpoint"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "endpoint" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type endpoint Endpoint

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Endpoint) UnmarshalJSON(data []byte) (err error) {
	x2 := endpoint{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Endpoint(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Endpoint) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Endpoint"
	} else if x.ResourceType != "Endpoint" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Endpoint, instead received %s", x.ResourceType))
	}
	return nil
}

type EndpointPlus struct {
	Endpoint                     `bson:",inline"`
	EndpointPlusRelatedResources `bson:",inline"`
}

type EndpointPlusRelatedResources struct {
	IncludedOrganizationResourcesReferencedByOrganization                  *[]Organization               `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
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
	RevIncludedPractitionerRoleResourcesReferencingEndpoint                *[]PractitionerRole           `bson:"_revIncludedPractitionerRoleResourcesReferencingEndpoint,omitempty"`
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
	RevIncludedOrganizationResourcesReferencingEndpoint                    *[]Organization               `bson:"_revIncludedOrganizationResourcesReferencingEndpoint,omitempty"`
	RevIncludedImplementationGuideResourcesReferencingResource             *[]ImplementationGuide        `bson:"_revIncludedImplementationGuideResourcesReferencingResource,omitempty"`
	RevIncludedImagingStudyResourcesReferencingEndpointPath1               *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingEndpointPath1,omitempty"`
	RevIncludedImagingStudyResourcesReferencingEndpointPath2               *[]ImagingStudy               `bson:"_revIncludedImagingStudyResourcesReferencingEndpointPath2,omitempty"`
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
	RevIncludedInsurancePlanResourcesReferencingEndpoint                   *[]InsurancePlan              `bson:"_revIncludedInsurancePlanResourcesReferencingEndpoint,omitempty"`
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
	RevIncludedOrganizationAffiliationResourcesReferencingEndpoint         *[]OrganizationAffiliation    `bson:"_revIncludedOrganizationAffiliationResourcesReferencingEndpoint,omitempty"`
	RevIncludedHealthcareServiceResourcesReferencingEndpoint               *[]HealthcareService          `bson:"_revIncludedHealthcareServiceResourcesReferencingEndpoint,omitempty"`
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
	RevIncludedLocationResourcesReferencingEndpoint                        *[]Location                   `bson:"_revIncludedLocationResourcesReferencingEndpoint,omitempty"`
}

func (e *EndpointPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if e.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *e.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if e.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *e.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if e.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *e.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if e.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *e.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *e.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPractitionerRoleResourcesReferencingEndpoint() (practitionerRoles []PractitionerRole, err error) {
	if e.RevIncludedPractitionerRoleResourcesReferencingEndpoint == nil {
		err = errors.New("RevIncluded practitionerRoles not requested")
	} else {
		practitionerRoles = *e.RevIncludedPractitionerRoleResourcesReferencingEndpoint
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if e.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *e.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if e.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *e.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *e.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedOrganizationResourcesReferencingEndpoint() (organizations []Organization, err error) {
	if e.RevIncludedOrganizationResourcesReferencingEndpoint == nil {
		err = errors.New("RevIncluded organizations not requested")
	} else {
		organizations = *e.RevIncludedOrganizationResourcesReferencingEndpoint
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if e.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *e.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingEndpointPath1() (imagingStudies []ImagingStudy, err error) {
	if e.RevIncludedImagingStudyResourcesReferencingEndpointPath1 == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *e.RevIncludedImagingStudyResourcesReferencingEndpointPath1
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingEndpointPath2() (imagingStudies []ImagingStudy, err error) {
	if e.RevIncludedImagingStudyResourcesReferencingEndpointPath2 == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *e.RevIncludedImagingStudyResourcesReferencingEndpointPath2
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *e.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if e.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *e.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *e.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedInsurancePlanResourcesReferencingEndpoint() (insurancePlans []InsurancePlan, err error) {
	if e.RevIncludedInsurancePlanResourcesReferencingEndpoint == nil {
		err = errors.New("RevIncluded insurancePlans not requested")
	} else {
		insurancePlans = *e.RevIncludedInsurancePlanResourcesReferencingEndpoint
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if e.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *e.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if e.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *e.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if e.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *e.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if e.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *e.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if e.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *e.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if e.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *e.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if e.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *e.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if e.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *e.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if e.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *e.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedOrganizationAffiliationResourcesReferencingEndpoint() (organizationAffiliations []OrganizationAffiliation, err error) {
	if e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint == nil {
		err = errors.New("RevIncluded organizationAffiliations not requested")
	} else {
		organizationAffiliations = *e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedHealthcareServiceResourcesReferencingEndpoint() (healthcareServices []HealthcareService, err error) {
	if e.RevIncludedHealthcareServiceResourcesReferencingEndpoint == nil {
		err = errors.New("RevIncluded healthcareServices not requested")
	} else {
		healthcareServices = *e.RevIncludedHealthcareServiceResourcesReferencingEndpoint
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if e.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *e.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if e.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *e.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *e.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (e *EndpointPlusRelatedResources) GetRevIncludedLocationResourcesReferencingEndpoint() (locations []Location, err error) {
	if e.RevIncludedLocationResourcesReferencingEndpoint == nil {
		err = errors.New("RevIncluded locations not requested")
	} else {
		locations = *e.RevIncludedLocationResourcesReferencingEndpoint
	}
	return
}

func (e *EndpointPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganization)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EndpointPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if e.RevIncludedPractitionerRoleResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedPractitionerRoleResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedPractitionerRoleResourcesReferencingEndpoint)[idx]
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
	if e.RevIncludedOrganizationResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedOrganizationResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedOrganizationResourcesReferencingEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingEndpointPath1 != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingEndpointPath1 {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingEndpointPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingEndpointPath2 != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingEndpointPath2 {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingEndpointPath2)[idx]
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
	if e.RevIncludedInsurancePlanResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedInsurancePlanResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedInsurancePlanResourcesReferencingEndpoint)[idx]
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
	if e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedHealthcareServiceResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedHealthcareServiceResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedHealthcareServiceResourcesReferencingEndpoint)[idx]
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
	if e.RevIncludedLocationResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedLocationResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedLocationResourcesReferencingEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (e *EndpointPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for idx := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			rsc := (*e.IncludedOrganizationResourcesReferencedByOrganization)[idx]
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
	if e.RevIncludedPractitionerRoleResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedPractitionerRoleResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedPractitionerRoleResourcesReferencingEndpoint)[idx]
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
	if e.RevIncludedOrganizationResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedOrganizationResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedOrganizationResourcesReferencingEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *e.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*e.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingEndpointPath1 != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingEndpointPath1 {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingEndpointPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedImagingStudyResourcesReferencingEndpointPath2 != nil {
		for idx := range *e.RevIncludedImagingStudyResourcesReferencingEndpointPath2 {
			rsc := (*e.RevIncludedImagingStudyResourcesReferencingEndpointPath2)[idx]
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
	if e.RevIncludedInsurancePlanResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedInsurancePlanResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedInsurancePlanResourcesReferencingEndpoint)[idx]
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
	if e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedOrganizationAffiliationResourcesReferencingEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if e.RevIncludedHealthcareServiceResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedHealthcareServiceResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedHealthcareServiceResourcesReferencingEndpoint)[idx]
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
	if e.RevIncludedLocationResourcesReferencingEndpoint != nil {
		for idx := range *e.RevIncludedLocationResourcesReferencingEndpoint {
			rsc := (*e.RevIncludedLocationResourcesReferencingEndpoint)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
