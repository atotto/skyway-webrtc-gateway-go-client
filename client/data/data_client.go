// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new data API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Data dataの待受ポート開放要求を送りますs

DataConnectionで転送するデータを受け渡すためのUDPポート開放要求を送信します
*/
func (a *Client) Data(params *DataParams) (*DataCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "data",
		Method:             "POST",
		PathPattern:        "/data",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DataCreated), nil

}

/*
DataConnectionClose data connectionを解放します

DataConnectionを解放して、関連する全ての接続を切断します。他のDataConnectionで利用していないDataの待受ポートがあればそれらもクローズします
*/
func (a *Client) DataConnectionClose(params *DataConnectionCloseParams) (*DataConnectionCloseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataConnectionCloseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "data_connection_close",
		Method:             "DELETE",
		PathPattern:        "/data/connections/{data_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataConnectionCloseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DataConnectionCloseNoContent), nil

}

/*
DataConnectionEvents data connectionオブジェクトからイベントを取得するのに利用します

Long Pollでイベントを監視するのに利用します。連続でイベントが発火する場合があるため常に監視するようにしてください
*/
func (a *Client) DataConnectionEvents(params *DataConnectionEventsParams) (*DataConnectionEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataConnectionEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "data_connection_events",
		Method:             "GET",
		PathPattern:        "/data/connections/{data_connection_id}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataConnectionEventsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DataConnectionEventsOK), nil

}

/*
DataConnectionPut data connectionの動作を変更します

DataConnectionで相手側から受信したデータの転送先指定と、DataConnectionに与えるデータの指定/変更を行います
*/
func (a *Client) DataConnectionPut(params *DataConnectionPutParams) (*DataConnectionPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataConnectionPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "data_connection_put",
		Method:             "PUT",
		PathPattern:        "/data/connections/{data_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataConnectionPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DataConnectionPutOK), nil

}

/*
DataConnectionStatus data connectionの状態を取得します

DataConnectionの状態を取得します
*/
func (a *Client) DataConnectionStatus(params *DataConnectionStatusParams) (*DataConnectionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataConnectionStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "data_connection_status",
		Method:             "GET",
		PathPattern:        "/data/connections/{data_connection_id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataConnectionStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DataConnectionStatusOK), nil

}

/*
DataConnectionsCreate リモートのs peerへの data connection connection確立を開始します

リモートのPeerへ発信し、DataConnectionの確立を試みます。接続確立は非同期で行われるため、/peers/{peer_id}/eventsを監視してください
*/
func (a *Client) DataConnectionsCreate(params *DataConnectionsCreateParams) (*DataConnectionsCreateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataConnectionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "data_connections_create",
		Method:             "POST",
		PathPattern:        "/data/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataConnectionsCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DataConnectionsCreateAccepted), nil

}

/*
DataDelete dataの待受ポートの閉鎖要求を送りますs

DataConnectionで転送するデータを受け渡すためのUDPポートの閉鎖要求を送信します。activeなDataConnectionで利用されている場合はエラーが発生します
*/
func (a *Client) DataDelete(params *DataDeleteParams) (*DataDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "data_delete",
		Method:             "DELETE",
		PathPattern:        "/data/{data_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DataDeleteNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
