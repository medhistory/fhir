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

type CapabilityStatement2 struct {
	DomainResource      `bson:",inline"`
	Url                 string                                       `bson:"url,omitempty" json:"url,omitempty"`
	Version             string                                       `bson:"version,omitempty" json:"version,omitempty"`
	Name                string                                       `bson:"name,omitempty" json:"name,omitempty"`
	Title               string                                       `bson:"title,omitempty" json:"title,omitempty"`
	Status              string                                       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental        *bool                                        `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                *FHIRDateTime                                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher           string                                       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact             []ContactDetail                              `bson:"contact,omitempty" json:"contact,omitempty"`
	Description         string                                       `bson:"description,omitempty" json:"description,omitempty"`
	UseContext          []UsageContext                               `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                            `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose             string                                       `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright           string                                       `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Kind                string                                       `bson:"kind,omitempty" json:"kind,omitempty"`
	Instantiates        []string                                     `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	Imports             []string                                     `bson:"imports,omitempty" json:"imports,omitempty"`
	Software            *CapabilityStatement2SoftwareComponent       `bson:"software,omitempty" json:"software,omitempty"`
	Implementation      *CapabilityStatement2ImplementationComponent `bson:"implementation,omitempty" json:"implementation,omitempty"`
	FhirVersion         string                                       `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Format              []string                                     `bson:"format,omitempty" json:"format,omitempty"`
	PatchFormat         []string                                     `bson:"patchFormat,omitempty" json:"patchFormat,omitempty"`
	ImplementationGuide []string                                     `bson:"implementationGuide,omitempty" json:"implementationGuide,omitempty"`
	Rest                []CapabilityStatement2RestComponent          `bson:"rest,omitempty" json:"rest,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CapabilityStatement2) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "CapabilityStatement2"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to CapabilityStatement2), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *CapabilityStatement2) GetBSON() (interface{}, error) {
	x.ResourceType = "CapabilityStatement2"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "capabilityStatement2" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type capabilityStatement2 CapabilityStatement2

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CapabilityStatement2) UnmarshalJSON(data []byte) (err error) {
	x2 := capabilityStatement2{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CapabilityStatement2(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CapabilityStatement2) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CapabilityStatement2"
	} else if x.ResourceType != "CapabilityStatement2" {
		return errors.New(fmt.Sprintf("Expected resourceType to be CapabilityStatement2, instead received %s", x.ResourceType))
	}
	return nil
}

type CapabilityStatement2SoftwareComponent struct {
	BackboneElement `bson:",inline"`
	Name            string        `bson:"name,omitempty" json:"name,omitempty"`
	Version         string        `bson:"version,omitempty" json:"version,omitempty"`
	ReleaseDate     *FHIRDateTime `bson:"releaseDate,omitempty" json:"releaseDate,omitempty"`
}

type CapabilityStatement2ImplementationComponent struct {
	BackboneElement `bson:",inline"`
	Description     string     `bson:"description,omitempty" json:"description,omitempty"`
	Url             string     `bson:"url,omitempty" json:"url,omitempty"`
	Custodian       *Reference `bson:"custodian,omitempty" json:"custodian,omitempty"`
}

type CapabilityStatement2RestComponent struct {
	BackboneElement `bson:",inline"`
	Mode            string                                                 `bson:"mode,omitempty" json:"mode,omitempty"`
	Documentation   string                                                 `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Resource        []CapabilityStatement2RestResourceComponent            `bson:"resource,omitempty" json:"resource,omitempty"`
	Interaction     []CapabilityStatement2SystemInteractionComponent       `bson:"interaction,omitempty" json:"interaction,omitempty"`
	SearchParam     []CapabilityStatement2RestResourceSearchParamComponent `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	Operation       []CapabilityStatement2RestResourceOperationComponent   `bson:"operation,omitempty" json:"operation,omitempty"`
	Compartment     []string                                               `bson:"compartment,omitempty" json:"compartment,omitempty"`
}

type CapabilityStatement2RestResourceComponent struct {
	BackboneElement  `bson:",inline"`
	Type             string                                                 `bson:"type,omitempty" json:"type,omitempty"`
	Profile          string                                                 `bson:"profile,omitempty" json:"profile,omitempty"`
	SupportedProfile []string                                               `bson:"supportedProfile,omitempty" json:"supportedProfile,omitempty"`
	Documentation    string                                                 `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Interaction      []CapabilityStatement2ResourceInteractionComponent     `bson:"interaction,omitempty" json:"interaction,omitempty"`
	SearchParam      []CapabilityStatement2RestResourceSearchParamComponent `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	Operation        []CapabilityStatement2RestResourceOperationComponent   `bson:"operation,omitempty" json:"operation,omitempty"`
}

type CapabilityStatement2ResourceInteractionComponent struct {
	BackboneElement `bson:",inline"`
	Code            string `bson:"code,omitempty" json:"code,omitempty"`
	Documentation   string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type CapabilityStatement2RestResourceSearchParamComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Definition      string `bson:"definition,omitempty" json:"definition,omitempty"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Documentation   string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type CapabilityStatement2RestResourceOperationComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Definition      string `bson:"definition,omitempty" json:"definition,omitempty"`
	Documentation   string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type CapabilityStatement2SystemInteractionComponent struct {
	BackboneElement `bson:",inline"`
	Code            string `bson:"code,omitempty" json:"code,omitempty"`
	Documentation   string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type CapabilityStatement2Plus struct {
	CapabilityStatement2                     `bson:",inline"`
	CapabilityStatement2PlusRelatedResources `bson:",inline"`
}

type CapabilityStatement2PlusRelatedResources struct {
	IncludedStructureDefinitionResourcesReferencedByResourceprofile        *[]StructureDefinition        `bson:"_includedStructureDefinitionResourcesReferencedByResourceprofile,omitempty"`
	IncludedStructureDefinitionResourcesReferencedBySupportedprofile       *[]StructureDefinition        `bson:"_includedStructureDefinitionResourcesReferencedBySupportedprofile,omitempty"`
	IncludedImplementationGuideResourcesReferencedByGuide                  *[]ImplementationGuide        `bson:"_includedImplementationGuideResourcesReferencedByGuide,omitempty"`
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

func (c *CapabilityStatement2PlusRelatedResources) GetIncludedStructureDefinitionResourceReferencedByResourceprofile() (structureDefinition *StructureDefinition, err error) {
	if c.IncludedStructureDefinitionResourcesReferencedByResourceprofile == nil {
		err = errors.New("Included structuredefinitions not requested")
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedByResourceprofile) > 1 {
		err = fmt.Errorf("Expected 0 or 1 structureDefinition, but found %d", len(*c.IncludedStructureDefinitionResourcesReferencedByResourceprofile))
	} else if len(*c.IncludedStructureDefinitionResourcesReferencedByResourceprofile) == 1 {
		structureDefinition = &(*c.IncludedStructureDefinitionResourcesReferencedByResourceprofile)[0]
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetIncludedStructureDefinitionResourcesReferencedBySupportedprofile() (structureDefinitions []StructureDefinition, err error) {
	if c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile == nil {
		err = errors.New("Included structureDefinitions not requested")
	} else {
		structureDefinitions = *c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetIncludedImplementationGuideResourcesReferencedByGuide() (implementationGuides []ImplementationGuide, err error) {
	if c.IncludedImplementationGuideResourcesReferencedByGuide == nil {
		err = errors.New("Included implementationGuides not requested")
	} else {
		implementationGuides = *c.IncludedImplementationGuideResourcesReferencedByGuide
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if c.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *c.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if c.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *c.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if c.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *c.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if c.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *c.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *c.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if c.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *c.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *c.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if c.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *c.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *c.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if c.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *c.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *c.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if c.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *c.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if c.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *c.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if c.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *c.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if c.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *c.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if c.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *c.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if c.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *c.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if c.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *c.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if c.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *c.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if c.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *c.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if c.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *c.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *c.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (c *CapabilityStatement2PlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedStructureDefinitionResourcesReferencedByResourceprofile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedByResourceprofile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedByResourceprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedImplementationGuideResourcesReferencedByGuide != nil {
		for idx := range *c.IncludedImplementationGuideResourcesReferencedByGuide {
			rsc := (*c.IncludedImplementationGuideResourcesReferencedByGuide)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (c *CapabilityStatement2PlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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

func (c *CapabilityStatement2PlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedStructureDefinitionResourcesReferencedByResourceprofile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedByResourceprofile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedByResourceprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile != nil {
		for idx := range *c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile {
			rsc := (*c.IncludedStructureDefinitionResourcesReferencedBySupportedprofile)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if c.IncludedImplementationGuideResourcesReferencedByGuide != nil {
		for idx := range *c.IncludedImplementationGuideResourcesReferencedByGuide {
			rsc := (*c.IncludedImplementationGuideResourcesReferencedByGuide)[idx]
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
