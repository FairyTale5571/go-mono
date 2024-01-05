package wss

import (
	"context"
	"net/http"
	"sync"
	
	"github.com/gorilla/websocket"
)

type WebsocketClient struct {
	mu      *sync.RWMutex
	connect *websocket.Conn
	wssURL  string
	dialer  *websocket.Dialer
	headers http.Header
}

type Settings struct {
	WssURL  string
	Dialer  *websocket.Dialer
	Headers http.Header
}

func NewWssClient(sets Settings) *WebsocketClient {
	if sets.Dialer == nil {
		sets.Dialer = websocket.DefaultDialer
	}
	return &WebsocketClient{
		mu:      &sync.RWMutex{},
		wssURL:  sets.WssURL,
		headers: sets.Headers,
		dialer:  sets.Dialer,
	}
}

func (w *WebsocketClient) Connect(ctx context.Context) error {
	var err error

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.connect != nil {
		return nil
	}

	w.connect, _, err = w.dialer.DialContext(ctx, w.wssURL, w.headers)
	if err != nil {
		w.connect = nil
		return err
	}

	defer func() {
		if err != nil {
			_ = w.connect.Close()
			w.connect = nil
		}
	}()
	return nil
}

func (w *WebsocketClient) SetCloseHandler(handler func(code int, text string) error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.connect == nil {
		return
	}
	w.connect.SetCloseHandler(handler)
}

func (w *WebsocketClient) SendMessage(str string) error {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.connect == nil {
		return ErrWSNotFound
	}
	return w.connect.WriteMessage(websocket.TextMessage, []byte(str))
}

func (w *WebsocketClient) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.connect == nil {
		return nil
	}

	err := w.connect.Close()
	if err != nil {
		return err
	}
	w.connect = nil
	return nil
}

func (w *WebsocketClient) ReadMessage() (int, []byte, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.connect == nil {
		return 0, nil, ErrWSNotFound
	}
	return w.connect.ReadMessage()
}

func (w *WebsocketClient) ReadJSON(v any) error {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if w.connect == nil {
		return ErrWSNotFound
	}
	return w.connect.ReadJSON(v)
}

func (w *WebsocketClient) IsConnected() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.connect != nil
}
