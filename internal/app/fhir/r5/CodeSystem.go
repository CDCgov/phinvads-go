package r5

import (
	"fmt"

	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
	"github.com/google/fhir/go/proto/google/fhir/proto/r5/core/codes_go_proto"
	datatypespb "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/datatypes_go_proto"
	r5pb "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/resources/code_system_go_proto"
)

func SerializeCodeSystemToFhir(cs *xo.CodeSystem, conceptCount int, concepts []*xo.CodeSystemConcept) (*r5pb.CodeSystem, error) {
	fhirCS := &r5pb.CodeSystem{
		Id:           newId(cs.Oid),
		Status:       &r5pb.CodeSystem_StatusCode{Value: codes_go_proto.PublicationStatusCode_ACTIVE},
		Version:      newString(cs.Version),
		Name:         newString(cs.Codesystemcode),
		Description:  newNullableMarkdown(cs.Definitiontext),
		Experimental: newBoolean(cs.Legacyflag),
		Url:          newUri(fmt.Sprintf("https://phinvads.cdc.gov/r5/CodeSystem/%s", cs.Oid)),
		Date:         newDateTime(cs.Statusdate),
		Publisher:    newNullableString(cs.Distributionsourceversionname),
		Title:        newString(cs.Name),
		Content:      &r5pb.CodeSystem_ContentCode{Value: codes_go_proto.CodeSystemContentModeCode_COMPLETE},
		Count:        newUnsignedInt(conceptCount),
	}

	fhirCS.Identifier = []*datatypespb.Identifier{
		{
			System: newUri("urn:ietf:rfc:3986"), // Assuming this system for oid mapping
			Value:  newString(fmt.Sprintf("urn:oid:%s", cs.Oid)),
		},
	}

	fhirCS.Meta = &datatypespb.Meta{
		Profile: []*datatypespb.Canonical{
			{Value: "http://hl7.org/fhir/StructureDefinition/shareablecodesystem"},
		},
	}

	definitionText := ""
	if cs.Definitiontext.Valid {
		definitionText = cs.Definitiontext.String
	}

	fhirCS.Text = &datatypespb.Narrative{
		Status: &datatypespb.Narrative_StatusCode{Value: codes_go_proto.NarrativeStatusCode_GENERATED},
		Div:    newXhtml(fmt.Sprintf("<div xmlns=\"http://www.w3.org/1999/xhtml\">%s</div>", definitionText)),
	}

	fhirCS.Contact = []*datatypespb.ContactDetail{
		{
			Name: newString("PHIN Vocabulary Services"),
			Telecom: []*datatypespb.ContactPoint{
				{
					System: &datatypespb.ContactPoint_SystemCode{Value: codes_go_proto.ContactPointSystemCode_URL},
					Value:  newString("https://www.cdc.gov/phin/php/phinvads/index.html"),
				},
				{
					System: &datatypespb.ContactPoint_SystemCode{Value: codes_go_proto.ContactPointSystemCode_EMAIL},
					Value:  newString("phinvs@cdc.gov"),
				},
			},
		},
	}

	conceptDefinitions := make([]*r5pb.CodeSystem_ConceptDefinition, 0, conceptCount)
	for _, concept := range concepts {
		conceptDefinition := &r5pb.CodeSystem_ConceptDefinition{
			Code:    &datatypespb.Code{Value: concept.Conceptcode},
			Display: newString(concept.Name),
		}
		conceptDefinitions = append(conceptDefinitions, conceptDefinition)
	}
	fhirCS.Concept = conceptDefinitions

	return fhirCS, nil
}
