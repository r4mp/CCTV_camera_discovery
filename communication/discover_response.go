package communication

import (
	"encoding/xml"
	"log"
)

type DiscoveryResponse struct{}

type Envelope struct {
	Header SoapHeader
	Body   SoapBody
}

type SoapHeader struct {
	MessageID   string
	RelatesTo   string
	To          string
	Action      string
	AppSequence string
}

type SoapBody struct {
	ProbeMatches ProbeMatches
}

type ProbeMatches struct {
	ProbeMatch []ProbeMatch
}

type ProbeMatch struct {
	EndpointReference EndpointReference
	Types             string
	Scopes            string
	XAddrs            string
	MetaVersion       string
}

type EndpointReference struct {
	Address string
}

func (d DiscoveryResponse) UnMarshall(data string) (interface{}, error) {

	envelope := Envelope{}

	err := xml.Unmarshal([]byte(data), &envelope)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return envelope, nil
}
