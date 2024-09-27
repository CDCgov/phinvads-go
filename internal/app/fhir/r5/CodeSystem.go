package r5

import (
	"fmt"

	"github.com/google/fhir/go/proto/google/fhir/proto/r5/core/codes_go_proto"
	datatypespb "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/datatypes_go_proto"
	r5pb "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/resources/code_system_go_proto"
	"github.com/CDCgov/phinvads-go/internal/database/models/xo"
)

func SerializeCodeSystemToFhir(cs *xo.CodeSystem) (*r5pb.CodeSystem, error) {
	fhirCS := &r5pb.CodeSystem{
		Id:           newId(cs.Oid),
		Status:       &r5pb.CodeSystem_StatusCode{Value: codes_go_proto.PublicationStatusCode_DRAFT},
		Version:      newString(cs.Version),
		Name:         newString(cs.Name),
		Description:  newNullableMarkdown(cs.Definitiontext),
		Experimental: newBoolean(cs.Legacyflag),
		Url:          newNullableUri(cs.Sourceurl),
		Date:         newDateTime(cs.Statusdate),
		Publisher:    newNullableString(cs.Distributionsourceversionname),
		Title:        newNullableString(cs.Assigningauthorityversionname),
		Content:      &r5pb.CodeSystem_ContentCode{Value: 4},
		// TODO: Count
		// TODO: Concept
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

	fhirCS.Text = &datatypespb.Narrative{
		Status: &datatypespb.Narrative_StatusCode{Value: codes_go_proto.NarrativeStatusCode_GENERATED},
		Div:    newXhtml("<div>Your narrative text here</div>"),
	}

	fhirCS.Contact = []*datatypespb.ContactDetail{
		{
			Name: newString("PHIN Vocabulary Services"),
			Telecom: []*datatypespb.ContactPoint{
				{
					System: &datatypespb.ContactPoint_SystemCode{Value: 5},
					Value:  newString("https://www.cdc.gov/phin/php/phinvads/index.html"),
				},
				{
					System: &datatypespb.ContactPoint_SystemCode{Value: 3},
					Value:  newString("phinvs@cdc.gov"),
				},
			},
		},
	}

	return fhirCS, nil
}
