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

type TestScript struct {
	DomainResource `bson:",inline"`
	Url            string                           `bson:"url,omitempty" json:"url,omitempty"`
	Identifier     *Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                           `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                           `bson:"name,omitempty" json:"name,omitempty"`
	Title          string                           `bson:"title,omitempty" json:"title,omitempty"`
	Status         string                           `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date           *FHIRDateTime                    `bson:"date,omitempty" json:"date,omitempty"`
	Publisher      string                           `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []ContactDetail                  `bson:"contact,omitempty" json:"contact,omitempty"`
	Description    string                           `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []UsageContext                   `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction   []CodeableConcept                `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose        string                           `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright      string                           `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Origin         []TestScriptOriginComponent      `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination    []TestScriptDestinationComponent `bson:"destination,omitempty" json:"destination,omitempty"`
	Metadata       *TestScriptMetadataComponent     `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Fixture        []TestScriptFixtureComponent     `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Profile        []Reference                      `bson:"profile,omitempty" json:"profile,omitempty"`
	Variable       []TestScriptVariableComponent    `bson:"variable,omitempty" json:"variable,omitempty"`
	Setup          *TestScriptSetupComponent        `bson:"setup,omitempty" json:"setup,omitempty"`
	Test           []TestScriptTestComponent        `bson:"test,omitempty" json:"test,omitempty"`
	Teardown       *TestScriptTeardownComponent     `bson:"teardown,omitempty" json:"teardown,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *TestScript) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "TestScript"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to TestScript), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *TestScript) GetBSON() (interface{}, error) {
	x.ResourceType = "TestScript"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "testScript" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type testScript TestScript

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *TestScript) UnmarshalJSON(data []byte) (err error) {
	x2 := testScript{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = TestScript(x2)
		return x.checkResourceType()
	}
	return
}

func (x *TestScript) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "TestScript"
	} else if x.ResourceType != "TestScript" {
		return errors.New(fmt.Sprintf("Expected resourceType to be TestScript, instead received %s", x.ResourceType))
	}
	return nil
}

type TestScriptOriginComponent struct {
	BackboneElement `bson:",inline"`
	Index           *int32  `bson:"index,omitempty" json:"index,omitempty"`
	Profile         *Coding `bson:"profile,omitempty" json:"profile,omitempty"`
}

type TestScriptDestinationComponent struct {
	BackboneElement `bson:",inline"`
	Index           *int32  `bson:"index,omitempty" json:"index,omitempty"`
	Profile         *Coding `bson:"profile,omitempty" json:"profile,omitempty"`
}

type TestScriptMetadataComponent struct {
	BackboneElement `bson:",inline"`
	Link            []TestScriptMetadataLinkComponent       `bson:"link,omitempty" json:"link,omitempty"`
	Capability      []TestScriptMetadataCapabilityComponent `bson:"capability,omitempty" json:"capability,omitempty"`
}

type TestScriptMetadataLinkComponent struct {
	BackboneElement `bson:",inline"`
	Url             string `bson:"url,omitempty" json:"url,omitempty"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
}

type TestScriptMetadataCapabilityComponent struct {
	BackboneElement `bson:",inline"`
	Required        *bool    `bson:"required,omitempty" json:"required,omitempty"`
	Validated       *bool    `bson:"validated,omitempty" json:"validated,omitempty"`
	Description     string   `bson:"description,omitempty" json:"description,omitempty"`
	Origin          []int32  `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination     *int32   `bson:"destination,omitempty" json:"destination,omitempty"`
	Link            []string `bson:"link,omitempty" json:"link,omitempty"`
	Capabilities    string   `bson:"capabilities,omitempty" json:"capabilities,omitempty"`
}

type TestScriptFixtureComponent struct {
	BackboneElement `bson:",inline"`
	Autocreate      *bool      `bson:"autocreate,omitempty" json:"autocreate,omitempty"`
	Autodelete      *bool      `bson:"autodelete,omitempty" json:"autodelete,omitempty"`
	Resource        *Reference `bson:"resource,omitempty" json:"resource,omitempty"`
}

type TestScriptVariableComponent struct {
	BackboneElement `bson:",inline"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	DefaultValue    string `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	Description     string `bson:"description,omitempty" json:"description,omitempty"`
	Expression      string `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField     string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	Hint            string `bson:"hint,omitempty" json:"hint,omitempty"`
	Path            string `bson:"path,omitempty" json:"path,omitempty"`
	SourceId        string `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}

type TestScriptSetupComponent struct {
	BackboneElement `bson:",inline"`
	Action          []TestScriptSetupActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptSetupActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert          *TestScriptSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptSetupActionOperationComponent struct {
	BackboneElement  `bson:",inline"`
	Type             *Coding                                                `bson:"type,omitempty" json:"type,omitempty"`
	Resource         string                                                 `bson:"resource,omitempty" json:"resource,omitempty"`
	Label            string                                                 `bson:"label,omitempty" json:"label,omitempty"`
	Description      string                                                 `bson:"description,omitempty" json:"description,omitempty"`
	Accept           string                                                 `bson:"accept,omitempty" json:"accept,omitempty"`
	ContentType      string                                                 `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Destination      *int32                                                 `bson:"destination,omitempty" json:"destination,omitempty"`
	EncodeRequestUrl *bool                                                  `bson:"encodeRequestUrl,omitempty" json:"encodeRequestUrl,omitempty"`
	Method           string                                                 `bson:"method,omitempty" json:"method,omitempty"`
	Origin           *int32                                                 `bson:"origin,omitempty" json:"origin,omitempty"`
	Params           string                                                 `bson:"params,omitempty" json:"params,omitempty"`
	RequestHeader    []TestScriptSetupActionOperationRequestHeaderComponent `bson:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	RequestId        string                                                 `bson:"requestId,omitempty" json:"requestId,omitempty"`
	ResponseId       string                                                 `bson:"responseId,omitempty" json:"responseId,omitempty"`
	SourceId         string                                                 `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	TargetId         string                                                 `bson:"targetId,omitempty" json:"targetId,omitempty"`
	Url              string                                                 `bson:"url,omitempty" json:"url,omitempty"`
}

type TestScriptSetupActionOperationRequestHeaderComponent struct {
	BackboneElement `bson:",inline"`
	Field           string `bson:"field,omitempty" json:"field,omitempty"`
	Value           string `bson:"value,omitempty" json:"value,omitempty"`
}

type TestScriptSetupActionAssertComponent struct {
	BackboneElement           `bson:",inline"`
	Label                     string `bson:"label,omitempty" json:"label,omitempty"`
	Description               string `bson:"description,omitempty" json:"description,omitempty"`
	Direction                 string `bson:"direction,omitempty" json:"direction,omitempty"`
	CompareToSourceId         string `bson:"compareToSourceId,omitempty" json:"compareToSourceId,omitempty"`
	CompareToSourceExpression string `bson:"compareToSourceExpression,omitempty" json:"compareToSourceExpression,omitempty"`
	CompareToSourcePath       string `bson:"compareToSourcePath,omitempty" json:"compareToSourcePath,omitempty"`
	ContentType               string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Expression                string `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField               string `bson:"headerField,omitempty" json:"headerField,omitempty"`
	MinimumId                 string `bson:"minimumId,omitempty" json:"minimumId,omitempty"`
	NavigationLinks           *bool  `bson:"navigationLinks,omitempty" json:"navigationLinks,omitempty"`
	Operator                  string `bson:"operator,omitempty" json:"operator,omitempty"`
	Path                      string `bson:"path,omitempty" json:"path,omitempty"`
	RequestMethod             string `bson:"requestMethod,omitempty" json:"requestMethod,omitempty"`
	RequestURL                string `bson:"requestURL,omitempty" json:"requestURL,omitempty"`
	Resource                  string `bson:"resource,omitempty" json:"resource,omitempty"`
	Response                  string `bson:"response,omitempty" json:"response,omitempty"`
	ResponseCode              string `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	SourceId                  string `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	ValidateProfileId         string `bson:"validateProfileId,omitempty" json:"validateProfileId,omitempty"`
	Value                     string `bson:"value,omitempty" json:"value,omitempty"`
	WarningOnly               *bool  `bson:"warningOnly,omitempty" json:"warningOnly,omitempty"`
}

type TestScriptTestComponent struct {
	BackboneElement `bson:",inline"`
	Name            string                          `bson:"name,omitempty" json:"name,omitempty"`
	Description     string                          `bson:"description,omitempty" json:"description,omitempty"`
	Action          []TestScriptTestActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptTestActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert          *TestScriptSetupActionAssertComponent    `bson:"assert,omitempty" json:"assert,omitempty"`
}

type TestScriptTeardownComponent struct {
	BackboneElement `bson:",inline"`
	Action          []TestScriptTeardownActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

type TestScriptTeardownActionComponent struct {
	BackboneElement `bson:",inline"`
	Operation       *TestScriptSetupActionOperationComponent `bson:"operation,omitempty" json:"operation,omitempty"`
}

type TestScriptPlus struct {
	TestScript                     `bson:",inline"`
	TestScriptPlusRelatedResources `bson:",inline"`
}

type TestScriptPlusRelatedResources struct {
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
	RevIncludedTestReportResourcesReferencingTestscript                    *[]TestReport                 `bson:"_revIncludedTestReportResourcesReferencingTestscript,omitempty"`
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

func (t *TestScriptPlusRelatedResources) GetRevIncludedAppointmentResourcesReferencingSupportinginfo() (appointments []Appointment, err error) {
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded appointments not requested")
	} else {
		appointments = *t.RevIncludedAppointmentResourcesReferencingSupportinginfo
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingSuccessor() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDerivedfrom() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingPredecessor() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingComposedof() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEventDefinitionResourcesReferencingDependson() (eventDefinitions []EventDefinition, err error) {
	if t.RevIncludedEventDefinitionResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded eventDefinitions not requested")
	} else {
		eventDefinitions = *t.RevIncludedEventDefinitionResourcesReferencingDependson
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingItem() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingItem == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingItem
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *t.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedConsentResourcesReferencingData() (consents []Consent, err error) {
	if t.RevIncludedConsentResourcesReferencingData == nil {
		err = errors.New("RevIncluded consents not requested")
	} else {
		consents = *t.RevIncludedConsentResourcesReferencingData
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingSuccessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDerivedfrom() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingPredecessor() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingComposedof() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath1() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath1
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMeasureResourcesReferencingDependsonPath2() (measures []Measure, err error) {
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded measures not requested")
	} else {
		measures = *t.RevIncludedMeasureResourcesReferencingDependsonPath2
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelated() (documentReferences []DocumentReference, err error) {
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *t.RevIncludedDocumentReferenceResourcesReferencingRelated
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMeasureReportResourcesReferencingEvaluatedresource() (measureReports []MeasureReport, err error) {
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource == nil {
		err = errors.New("RevIncluded measureReports not requested")
	} else {
		measureReports = *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedVerificationResultResourcesReferencingTarget() (verificationResults []VerificationResult, err error) {
	if t.RevIncludedVerificationResultResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded verificationResults not requested")
	} else {
		verificationResults = *t.RevIncludedVerificationResultResourcesReferencingTarget
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedContractResourcesReferencingSubject() (contracts []Contract, err error) {
	if t.RevIncludedContractResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *t.RevIncludedContractResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *t.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingSuccessor() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDerivedfrom() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingPredecessor() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingComposedof() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath1() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchDefinitionResourcesReferencingDependsonPath2() (researchDefinitions []ResearchDefinition, err error) {
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchDefinitions not requested")
	} else {
		researchDefinitions = *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedImplementationGuideResourcesReferencingResource() (implementationGuides []ImplementationGuide, err error) {
	if t.RevIncludedImplementationGuideResourcesReferencingResource == nil {
		err = errors.New("RevIncluded implementationGuides not requested")
	} else {
		implementationGuides = *t.RevIncludedImplementationGuideResourcesReferencingResource
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingSuccessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingPredecessor() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingComposedof() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2() (researchElementDefinitions []ResearchElementDefinition, err error) {
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded researchElementDefinitions not requested")
	} else {
		researchElementDefinitions = *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingPartof() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingPartof == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingPartof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedCommunicationResourcesReferencingBasedon() (communications []Communication, err error) {
	if t.RevIncludedCommunicationResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communications not requested")
	} else {
		communications = *t.RevIncludedCommunicationResourcesReferencingBasedon
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingSuccessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDerivedfrom() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingPredecessor() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingComposedof() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath1() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedActivityDefinitionResourcesReferencingDependsonPath2() (activityDefinitions []ActivityDefinition, err error) {
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded activityDefinitions not requested")
	} else {
		activityDefinitions = *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingItem() (linkages []Linkage, err error) {
	if t.RevIncludedLinkageResourcesReferencingItem == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *t.RevIncludedLinkageResourcesReferencingItem
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedLinkageResourcesReferencingSource() (linkages []Linkage, err error) {
	if t.RevIncludedLinkageResourcesReferencingSource == nil {
		err = errors.New("RevIncluded linkages not requested")
	} else {
		linkages = *t.RevIncludedLinkageResourcesReferencingSource
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingBasedon() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingBasedon
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDeviceRequestResourcesReferencingPriorrequest() (deviceRequests []DeviceRequest, err error) {
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest == nil {
		err = errors.New("RevIncluded deviceRequests not requested")
	} else {
		deviceRequests = *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingFocus() (messageHeaders []MessageHeader, err error) {
	if t.RevIncludedMessageHeaderResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *t.RevIncludedMessageHeaderResourcesReferencingFocus
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedImmunizationRecommendationResourcesReferencingInformation() (immunizationRecommendations []ImmunizationRecommendation, err error) {
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation == nil {
		err = errors.New("RevIncluded immunizationRecommendations not requested")
	} else {
		immunizationRecommendations = *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingEntity() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingEntity
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if t.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *t.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedTaskResourcesReferencingSubject() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedTaskResourcesReferencingFocus() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingFocus
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedTaskResourcesReferencingBasedon() (tasks []Task, err error) {
	if t.RevIncludedTaskResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded tasks not requested")
	} else {
		tasks = *t.RevIncludedTaskResourcesReferencingBasedon
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if t.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *t.RevIncludedListResourcesReferencingItem
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingSuccessor() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDerivedfrom() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingPredecessor() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingComposedof() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceVariableResourcesReferencingDependson() (evidenceVariables []EvidenceVariable, err error) {
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidenceVariables not requested")
	} else {
		evidenceVariables = *t.RevIncludedEvidenceVariableResourcesReferencingDependson
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedObservationResourcesReferencingFocus() (observations []Observation, err error) {
	if t.RevIncludedObservationResourcesReferencingFocus == nil {
		err = errors.New("RevIncluded observations not requested")
	} else {
		observations = *t.RevIncludedObservationResourcesReferencingFocus
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingSuccessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDerivedfrom() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingPredecessor() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingComposedof() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedLibraryResourcesReferencingDependson() (libraries []Library, err error) {
	if t.RevIncludedLibraryResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded libraries not requested")
	} else {
		libraries = *t.RevIncludedLibraryResourcesReferencingDependson
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedCommunicationRequestResourcesReferencingBasedon() (communicationRequests []CommunicationRequest, err error) {
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon == nil {
		err = errors.New("RevIncluded communicationRequests not requested")
	} else {
		communicationRequests = *t.RevIncludedCommunicationRequestResourcesReferencingBasedon
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if t.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *t.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedTestReportResourcesReferencingTestscript() (testReports []TestReport, err error) {
	if t.RevIncludedTestReportResourcesReferencingTestscript == nil {
		err = errors.New("RevIncluded testReports not requested")
	} else {
		testReports = *t.RevIncludedTestReportResourcesReferencingTestscript
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingSuccessor() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDerivedfrom() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingPredecessor() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingComposedof() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedEvidenceResourcesReferencingDependson() (evidences []Evidence, err error) {
	if t.RevIncludedEvidenceResourcesReferencingDependson == nil {
		err = errors.New("RevIncluded evidences not requested")
	} else {
		evidences = *t.RevIncludedEvidenceResourcesReferencingDependson
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingEntity() (auditEvents []AuditEvent, err error) {
	if t.RevIncludedAuditEventResourcesReferencingEntity == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *t.RevIncludedAuditEventResourcesReferencingEntity
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedConditionResourcesReferencingEvidencedetail() (conditions []Condition, err error) {
	if t.RevIncludedConditionResourcesReferencingEvidencedetail == nil {
		err = errors.New("RevIncluded conditions not requested")
	} else {
		conditions = *t.RevIncludedConditionResourcesReferencingEvidencedetail
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if t.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *t.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *t.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingSupportinginfo() (clinicalImpressions []ClinicalImpression, err error) {
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingSuccessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDerivedfrom() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingPredecessor() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingComposedof() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingComposedof
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath1() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedPlanDefinitionResourcesReferencingDependsonPath2() (planDefinitions []PlanDefinition, err error) {
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 == nil {
		err = errors.New("RevIncluded planDefinitions not requested")
	} else {
		planDefinitions = *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2
	}
	return
}

func (t *TestScriptPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (t *TestScriptPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*t.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*t.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*t.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*t.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*t.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTestReportResourcesReferencingTestscript != nil {
		for idx := range *t.RevIncludedTestReportResourcesReferencingTestscript {
			rsc := (*t.RevIncludedTestReportResourcesReferencingTestscript)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}

func (t *TestScriptPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if t.RevIncludedAppointmentResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedAppointmentResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedAppointmentResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEventDefinitionResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEventDefinitionResourcesReferencingDependson {
			rsc := (*t.RevIncludedEventDefinitionResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingItem {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for idx := range *t.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			rsc := (*t.RevIncludedDocumentManifestResourcesReferencingRelatedref)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConsentResourcesReferencingData != nil {
		for idx := range *t.RevIncludedConsentResourcesReferencingData {
			rsc := (*t.RevIncludedConsentResourcesReferencingData)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedMeasureResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingComposedof {
			rsc := (*t.RevIncludedMeasureResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedMeasureResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedMeasureResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDocumentReferenceResourcesReferencingRelated != nil {
		for idx := range *t.RevIncludedDocumentReferenceResourcesReferencingRelated {
			rsc := (*t.RevIncludedDocumentReferenceResourcesReferencingRelated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource != nil {
		for idx := range *t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource {
			rsc := (*t.RevIncludedMeasureReportResourcesReferencingEvaluatedresource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedVerificationResultResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedVerificationResultResourcesReferencingTarget {
			rsc := (*t.RevIncludedVerificationResultResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedContractResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedContractResourcesReferencingSubject {
			rsc := (*t.RevIncludedContractResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingRequest {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingRequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for idx := range *t.RevIncludedPaymentNoticeResourcesReferencingResponse {
			rsc := (*t.RevIncludedPaymentNoticeResourcesReferencingResponse)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImplementationGuideResourcesReferencingResource != nil {
		for idx := range *t.RevIncludedImplementationGuideResourcesReferencingResource {
			rsc := (*t.RevIncludedImplementationGuideResourcesReferencingResource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedResearchElementDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingPartof != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingPartof {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingPartof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedActivityDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingItem {
			rsc := (*t.RevIncludedLinkageResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLinkageResourcesReferencingSource != nil {
		for idx := range *t.RevIncludedLinkageResourcesReferencingSource {
			rsc := (*t.RevIncludedLinkageResourcesReferencingSource)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDeviceRequestResourcesReferencingPriorrequest != nil {
		for idx := range *t.RevIncludedDeviceRequestResourcesReferencingPriorrequest {
			rsc := (*t.RevIncludedDeviceRequestResourcesReferencingPriorrequest)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedMessageHeaderResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedMessageHeaderResourcesReferencingFocus {
			rsc := (*t.RevIncludedMessageHeaderResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedImmunizationRecommendationResourcesReferencingInformation != nil {
		for idx := range *t.RevIncludedImmunizationRecommendationResourcesReferencingInformation {
			rsc := (*t.RevIncludedImmunizationRecommendationResourcesReferencingInformation)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingEntity {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for idx := range *t.RevIncludedProvenanceResourcesReferencingTarget {
			rsc := (*t.RevIncludedProvenanceResourcesReferencingTarget)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingSubject {
			rsc := (*t.RevIncludedTaskResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingFocus {
			rsc := (*t.RevIncludedTaskResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTaskResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedTaskResourcesReferencingBasedon {
			rsc := (*t.RevIncludedTaskResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedListResourcesReferencingItem != nil {
		for idx := range *t.RevIncludedListResourcesReferencingItem {
			rsc := (*t.RevIncludedListResourcesReferencingItem)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceVariableResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceVariableResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceVariableResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedObservationResourcesReferencingFocus != nil {
		for idx := range *t.RevIncludedObservationResourcesReferencingFocus {
			rsc := (*t.RevIncludedObservationResourcesReferencingFocus)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedLibraryResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingComposedof {
			rsc := (*t.RevIncludedLibraryResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedLibraryResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedLibraryResourcesReferencingDependson {
			rsc := (*t.RevIncludedLibraryResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCommunicationRequestResourcesReferencingBasedon != nil {
		for idx := range *t.RevIncludedCommunicationRequestResourcesReferencingBasedon {
			rsc := (*t.RevIncludedCommunicationRequestResourcesReferencingBasedon)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedBasicResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedBasicResourcesReferencingSubject {
			rsc := (*t.RevIncludedBasicResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedTestReportResourcesReferencingTestscript != nil {
		for idx := range *t.RevIncludedTestReportResourcesReferencingTestscript {
			rsc := (*t.RevIncludedTestReportResourcesReferencingTestscript)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingComposedof {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedEvidenceResourcesReferencingDependson != nil {
		for idx := range *t.RevIncludedEvidenceResourcesReferencingDependson {
			rsc := (*t.RevIncludedEvidenceResourcesReferencingDependson)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedAuditEventResourcesReferencingEntity != nil {
		for idx := range *t.RevIncludedAuditEventResourcesReferencingEntity {
			rsc := (*t.RevIncludedAuditEventResourcesReferencingEntity)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedConditionResourcesReferencingEvidencedetail != nil {
		for idx := range *t.RevIncludedConditionResourcesReferencingEvidencedetail {
			rsc := (*t.RevIncludedConditionResourcesReferencingEvidencedetail)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingSubject {
			rsc := (*t.RevIncludedCompositionResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedCompositionResourcesReferencingEntry != nil {
		for idx := range *t.RevIncludedCompositionResourcesReferencingEntry {
			rsc := (*t.RevIncludedCompositionResourcesReferencingEntry)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for idx := range *t.RevIncludedDetectedIssueResourcesReferencingImplicated {
			rsc := (*t.RevIncludedDetectedIssueResourcesReferencingImplicated)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for idx := range *t.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			rsc := (*t.RevIncludedQuestionnaireResponseResourcesReferencingSubject)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo != nil {
		for idx := range *t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo {
			rsc := (*t.RevIncludedClinicalImpressionResourcesReferencingSupportinginfo)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingSuccessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingSuccessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingSuccessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDerivedfrom)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingPredecessor != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingPredecessor {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingPredecessor)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingComposedof != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingComposedof {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingComposedof)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath1)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	if t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 != nil {
		for idx := range *t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2 {
			rsc := (*t.RevIncludedPlanDefinitionResourcesReferencingDependsonPath2)[idx]
			resourceMap[rsc.Id] = &rsc
		}
	}
	return resourceMap
}
