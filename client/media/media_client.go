// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new media API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for media API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Media mediaの待受ポート開放要求を送りますs

MediaConnectionで転送するMediaを受け渡すためのUDPポート開放要求を送信します
*/
func (a *Client) Media(params *MediaParams) (*MediaCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media",
		Method:             "POST",
		PathPattern:        "/media",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaCreated), nil

}

/*
MediaConnectionAnswer callに応答しますs

callにどのように応答するかMediaConstraintsを提供します
*/
func (a *Client) MediaConnectionAnswer(params *MediaConnectionAnswerParams) (*MediaConnectionAnswerAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaConnectionAnswerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_connection_answer",
		Method:             "POST",
		PathPattern:        "/media/connections/{media_connection_id}/answer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaConnectionAnswerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaConnectionAnswerAccepted), nil

}

/*
MediaConnectionClose media connectionを解放します

MediaConnectionを解放します。このMediaConnection以外で利用されていないMediaがあれば同時にクローズします
*/
func (a *Client) MediaConnectionClose(params *MediaConnectionCloseParams) (*MediaConnectionCloseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaConnectionCloseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_connection_close",
		Method:             "DELETE",
		PathPattern:        "/media/connections/{media_connection_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaConnectionCloseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaConnectionCloseNoContent), nil

}

/*
MediaConnectionCreate リモートのs peerへの media connection確立を開始します

リモートのPeerへ発信し、MediaConnectionの確立を試みます。接続確立は非同期で行われるため、/peers/{peer_id}/eventsを監視してください
*/
func (a *Client) MediaConnectionCreate(params *MediaConnectionCreateParams) (*MediaConnectionCreateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaConnectionCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_connection_create",
		Method:             "POST",
		PathPattern:        "/media/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaConnectionCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaConnectionCreateAccepted), nil

}

/*
MediaConnectionEvent media connectionオブジェクトからイベントを取得するのに利用します

Long Pollでイベントを監視するのに利用します。連続でイベントが発火する場合があるため常に監視するようにしてください
*/
func (a *Client) MediaConnectionEvent(params *MediaConnectionEventParams) (*MediaConnectionEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaConnectionEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_connection_event",
		Method:             "GET",
		PathPattern:        "/media/connections/{media_connection_id}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaConnectionEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaConnectionEventOK), nil

}

/*
MediaConnectionPli ps l iパケットを送信します

keyframe要求を相手に伝えるため、PLIパケットを送信します。どのメディアがkeyframe要求を必要としているか特定するため、メディアの転送先ポートとIPアドレスを指定します
*/
func (a *Client) MediaConnectionPli(params *MediaConnectionPliParams) (*MediaConnectionPliCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaConnectionPliParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_connection_pli",
		Method:             "POST",
		PathPattern:        "/media/connections/{media_connection_id}/pli",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaConnectionPliReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaConnectionPliCreated), nil

}

/*
MediaConnectionStatus media connectionの状態を取得します

MediaConnectionの状態を取得します
*/
func (a *Client) MediaConnectionStatus(params *MediaConnectionStatusParams) (*MediaConnectionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaConnectionStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_connection_status",
		Method:             "GET",
		PathPattern:        "/media/connections/{media_connection_id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaConnectionStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaConnectionStatusOK), nil

}

/*
MediaRtcpCreate rs tcpの待受ポート開放要求を送ります

MediaConnectionで転送するRTCPを受け渡すためのUDPポート開放要求を送信します
*/
func (a *Client) MediaRtcpCreate(params *MediaRtcpCreateParams) (*MediaRtcpCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaRtcpCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_rtcp_create",
		Method:             "POST",
		PathPattern:        "/media/rtcp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaRtcpCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaRtcpCreateCreated), nil

}

/*
MediaRtcpDelete rs tcp待ち受けポートの解放

MediaConnectionで転送するRTCPを受け渡すためのUDPポートの閉鎖要求を送信します。MediaConnectionの中で利用中であればエラーを返します
*/
func (a *Client) MediaRtcpDelete(params *MediaRtcpDeleteParams) (*MediaRtcpDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMediaRtcpDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "media_rtcp_delete",
		Method:             "DELETE",
		PathPattern:        "/media/rtcp/{rtcp_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MediaRtcpDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MediaRtcpDeleteNoContent), nil

}

/*
StreamsDelete mediaの解放s

MediaConnectionで転送するMediaを受け渡すためのUDPポートの閉鎖要求を送信します。MediaConnectionの中で利用中であればエラーを返します
*/
func (a *Client) StreamsDelete(params *StreamsDeleteParams) (*StreamsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStreamsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "streams_delete",
		Method:             "DELETE",
		PathPattern:        "/media/{media_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StreamsDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StreamsDeleteNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
