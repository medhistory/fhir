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

type DeviceDefinition struct {
	DomainResource          `bson:",inline"`
	Identifier              []Identifier                                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	UdiDeviceIdentifier     []DeviceDefinitionUdiDeviceIdentifierComponent `bson:"udiDeviceIdentifier,omitempty" json:"udiDeviceIdentifier,omitempty"`
	ManufacturerString      string                                         `bson:"manufacturerString,omitempty" json:"manufacturerString,omitempty"`
	ManufacturerReference   *Reference                                     `bson:"manufacturerReference,omitempty" json:"manufacturerReference,omitempty"`
	DeviceName              []DeviceDefinitionDeviceNameComponent          `bson:"deviceName,omitempty" json:"deviceName,omitempty"`
	ModelNumber             string                                         `bson:"modelNumber,omitempty" json:"modelNumber,omitempty"`
	Type                    *CodeableConcept                               `bson:"type,omitempty" json:"type,omitempty"`
	Specialization          []DeviceDefinitionSpecializationComponent      `bson:"specialization,omitempty" json:"specialization,omitempty"`
	Version                 []string                                       `bson:"version,omitempty" json:"version,omitempty"`
	Safety                  []CodeableConcept                              `bson:"safety,omitempty" json:"safety,omitempty"`
	ShelfLifeStorage        []ProductShelfLife                             `bson:"shelfLifeStorage,omitempty" json:"shelfLifeStorage,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic                            `bson:"physicalCharacteristics,omitempty" json:"physicalCharacteristics,omitempty"`
	LanguageCode            []CodeableConcept                              `bson:"languageCode,omitempty" json:"languageCode,omitempty"`
	Capability              []DeviceDefinitionCapabilityComponent          `bson:"capability,omitempty" json:"capability,omitempty"`
	Property                []DeviceDefinitionPropertyComponent            `bson:"property,omitempty" json:"property,omitempty"`
	Owner                   *Reference                                     `bson:"owner,omitempty" json:"owner,omitempty"`
	Contact                 []ContactPoint                                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Url                     string                                         `bson:"url,omitempty" json:"url,omitempty"`
	OnlineInformation       string                                         `bson:"onlineInformation,omitempty" json:"onlineInformation,omitempty"`
	Note                    []Annotation                                   `bson:"note,omitempty" json:"note,omitempty"`
	Quantity                *Quantity                                      `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ParentDevice            *Reference                                     `bson:"parentDevice,omitempty" json:"parentDevice,omitempty"`
	Material                []DeviceDefinitionMaterialComponent            `bson:"material,omitempty" json:"material,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DeviceDefinition) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DeviceDefinition"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DeviceDefinition), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DeviceDefinition) GetBSON() (interface{}, error) {
	x.ResourceType = "DeviceDefinition"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "deviceDefinition" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type deviceDefinition DeviceDefinition

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DeviceDefinition) UnmarshalJSON(data []byte) (err error) {
	x2 := deviceDefinition{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DeviceDefinition(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DeviceDefinition) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DeviceDefinition"
	} else if x.ResourceType != "DeviceDefinition" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DeviceDefinition, instead received %s", x.ResourceType))
	}
	return nil
}

type DeviceDefinitionUdiDeviceIdentifierComponent struct {
	BackboneElement  `bson:",inline"`
	DeviceIdentifier string `bson:"deviceIdentifier,omitempty" json:"deviceIdentifier,omitempty"`
	Issuer           string `bson:"issuer,omitempty" json:"issuer,omitempty"`
	Jurisdiction     string `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
}

type DeviceDefinitionDeviceNameComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
}

type DeviceDefinitionSpecializationComponent struct {
	BackboneElement `bson:",inline"`
	SystemType      string `bson:"systemType,omitempty" json:"systemType,omitempty"`
	Version         string `bson:"version,omitempty" json:"version,omitempty"`
}

type DeviceDefinitionCapabilityComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Description     []CodeableConcept `bson:"description,omitempty" json:"description,omitempty"`
}

type DeviceDefinitionPropertyComponent struct {
	BackboneElement `bson:",inline"`
	Type            *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	ValueQuantity   []Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCode       []CodeableConcept `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
}

type DeviceDefinitionMaterialComponent struct {
	BackboneElement     `bson:",inline"`
	Substance           *CodeableConcept `bson:"substance,omitempty" json:"substance,omitempty"`
	Alternate           *bool            `bson:"alternate,omitempty" json:"alternate,omitempty"`
	AllergenicIndicator *bool            `bson:"allergenicIndicator,omitempty" json:"allergenicIndicator,omitempty"`
}

type DeviceDefinitionPlus struct {
	DeviceDefinition                     `bson:",inline"`
	DeviceDefinitionPlusRelatedResources `bson:",inline"`
}

type DeviceDefinitionPlusRelatedResources struct {
	IncludedDeviceDefinitionResourcesReferencedByParent                    *[]DeviceDefinition           `bson:"_includedDeviceDefinitionResourcesReferencedByParent,omitempty"`
	RevIncludedAppointmentResourcesReferencingSupportinginfo               *[]Appointment                `bson:"_revIncludedAppointmentResourcesReferencingSupportinginfo,omitempty"`
	RevIncludedCatalogEntryResourcesReferencingReferenceditem              *[]CatalogEntry               `bson:"_revIncludedCatalogEntryResourcesReferencingReferenceditem,omitempty"`
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
	RevIncludedClinicalUseIssueResourcesReferencingSubject                 *[]ClinicalUseIssue           `bson:"_revIncludedClinicalUseIssueResourcesReferencingSubject,omitempty"`
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
	RevIncludedRegulatedAuthorizationResourcesReferencingSubject           *[]RegulatedAuthorization     `bson:"_revIncludedRegulatedAuthorizationResourcesReferencingSubject,omitempty"`
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
	RevIncludedDeviceDefinitionResourcesReferencingParent                  *[]DeviceDefinition           `bson:"_revIncludedDeviceDefinitionResourcesReferencingParent,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingSuccessor                 *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingSuccessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDerivedfrom               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDerivedfrom,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingPredecessor               *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingPredecessor,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingComposedof                *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingComposedof,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath1            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath1,omitempty"`
	RevIncludedPlanDefinitionResourcesReferencingDependsonPath2            *[]PlanDefinition             `bson:"_revIncludedPlanDefinitionResourcesReferencingDependsonPath2,omitempty"`
}

func (d *DeviceDefinitionPlusRelatedResources) GetIncludedDeviceDefinitionResourceReferencedByParent() (deviceDefinition *DeviceDefinition, err error) {
	if d.IncludedDeviceDefinitionResourcesReferencedByParent == nil {
		err = errors.New("Included devicedefinitions not requested")
	} else if len(*d.IncludedDeviceDefinitionResourcesReferencedByParent) > 1 {
		err = fmt.Errorf("Expected 0 or 1 deviceDefinition, but found %d", len(*d.IncludedDeviceDefinitionResourcesReferencedByParent))
	} else if len(*d.IncludedDeviceDefinitionResourcesReferencedByParent) == 1 {
		deviceDefinition = &(*d.IncludedDeviceDefinitionResourcesReferencedByParent)[0]
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *d.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedCatalogEntryResourcesReferencingReferenceditem() (catalogEntries []CatalogEntry, err error) {
	if d.RevIncludedCatalogEntryResourcesReferencingReferenceditem == nil {
		err = errors.New("RevIncluded catalogEntries not requested")
	} else {
		catalogEntries = *d.RevIncludedCatalogEntryResourcesReferencingReferenceditem
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if d.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *d.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if d.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *d.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *d.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedClinicalUseIssueResourcesReferencingSubject() (clinicalUseIssues []ClinicalUseIssue, err error) {
	if d.RevIncludedClinicalUseIssueResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded clinicalUseIssues not requested")
	} else {
		clinicalUseIssues = *d.RevIncludedClinicalUseIssueResourcesReferencingSubject
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if d.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *d.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if d.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *d.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if d.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *d.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if d.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *d.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if d.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *d.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if d.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *d.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if d.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *d.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *d.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedRegulatedAuthorizationResourcesReferencingSubject() (regulatedAuthorizations []RegulatedAuthorization, err error) {
	if d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded regulatedAuthorizations not requested")
	} else {
		regulatedAuthorizations = *d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if d.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *d.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if d.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *d.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *d.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if d.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *d.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if d.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *d.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedDeviceDefinitionResourcesReferencingParent() (deviceDefinitions []DeviceDefinition, err error) {
	if d.RevIncludedDeviceDefinitionResourcesReferencingParent == nil {
		err = errors.New("RevIncluded deviceDefinitions not requested")
	} else {
		deviceDefinitions = *d.RevIncludedDeviceDefinitionResourcesReferencingParent
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (d *DeviceDefinitionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedDeviceDefinitionResourcesReferencedByParent != nil {
		for idx := range *d.IncludedDeviceDefinitionResourcesReferencedByParent {
			rsc := (*d.IncludedDeviceDefinitionResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DeviceDefinitionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCatalogEntryResourcesReferencingReferenceditem != nil {
		for idx := range *d.RevIncludedCatalogEntryResourcesReferencingReferenceditem {
			rsc := (*d.RevIncludedCatalogEntryResourcesReferencingReferenceditem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalUseIssueResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedClinicalUseIssueResourcesReferencingSubject {
			rsc := (*d.RevIncludedClinicalUseIssueResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*d.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*d.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*d.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*d.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject {
			rsc := (*d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*d.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceDefinitionResourcesReferencingParent != nil {
		for idx := range *d.RevIncludedDeviceDefinitionResourcesReferencingParent {
			rsc := (*d.RevIncludedDeviceDefinitionResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (d *DeviceDefinitionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedDeviceDefinitionResourcesReferencedByParent != nil {
		for idx := range *d.IncludedDeviceDefinitionResourcesReferencedByParent {
			rsc := (*d.IncludedDeviceDefinitionResourcesReferencedByParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCatalogEntryResourcesReferencingReferenceditem != nil {
		for idx := range *d.RevIncludedCatalogEntryResourcesReferencingReferenceditem {
			rsc := (*d.RevIncludedCatalogEntryResourcesReferencingReferenceditem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*d.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*d.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *d.RevIncludedConsentResourcesReferencingData {
			rsc := (*d.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*d.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *d.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*d.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalUseIssueResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedClinicalUseIssueResourcesReferencingSubject {
			rsc := (*d.RevIncludedClinicalUseIssueResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*d.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*d.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedContractResourcesReferencingSubject {
			rsc := (*d.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*d.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *d.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*d.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*d.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *d.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*d.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *d.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*d.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*d.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *d.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*d.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*d.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*d.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*d.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*d.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *d.RevIncludedListResourcesReferencingItem {
			rsc := (*d.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject {
			rsc := (*d.RevIncludedRegulatedAuthorizationResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *d.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*d.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*d.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*d.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *d.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*d.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*d.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *d.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*d.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *d.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*d.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *d.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*d.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*d.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *d.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*d.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*d.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*d.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*d.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedDeviceDefinitionResourcesReferencingParent != nil {
		for idx := range *d.RevIncludedDeviceDefinitionResourcesReferencingParent {
			rsc := (*d.RevIncludedDeviceDefinitionResourcesReferencingParent)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*d.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
