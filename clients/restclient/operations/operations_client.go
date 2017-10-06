package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetComponents get components API
*/
func (a *Client) GetComponents(params *GetComponentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetComponentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetComponents",
		Method:             "GET",
		PathPattern:        "/components",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetComponentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetComponentsOK), nil

}

/*
GetMta get mta API
*/
func (a *Client) GetMta(params *GetMtaParams, authInfo runtime.ClientAuthInfoWriter) (*GetMtaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMtaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMta",
		Method:             "GET",
		PathPattern:        "/components/{mtaId}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMtaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMtaOK), nil

}

/*
GetOperation get operation API
*/
func (a *Client) GetOperation(params *GetOperationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOperationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOperationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOperation",
		Method:             "GET",
		PathPattern:        "/operations/{processId}",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOperationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOperationOK), nil

}

/*
GetOperations get operations API
*/
func (a *Client) GetOperations(params *GetOperationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOperationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOperationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOperations",
		Method:             "GET",
		PathPattern:        "/operations",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOperationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOperationsOK), nil

}

/*
PurgeConfiguration purge configuration API
*/
func (a *Client) PurgeConfiguration(params *PurgeConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*PurgeConfigurationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurgeConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PurgeConfiguration",
		Method:             "POST",
		PathPattern:        "/configuration-entries/purge",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PurgeConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PurgeConfigurationNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
