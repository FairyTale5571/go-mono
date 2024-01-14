package expirenza

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/FairyTale5571/go-mono/internal/api"
	"github.com/FairyTale5571/go-mono/internal/wss"
	"github.com/gorilla/websocket"
)

type Expirenza struct {
	clientMu  *sync.RWMutex
	wssClient *wss.WebsocketClient
	apiClient *api.APIClient

	debug bool

	handlersMu   *sync.RWMutex
	handlers     map[Operation][]*eventHandlerInstance
	onceHandlers map[Operation][]*eventHandlerInstance
	listener     chan any

	restoID   string
	secretKey string
}

type Settings struct {
	RestoID   string
	SecretKey string
	Debug     bool

	Client    *http.Client
	WssClient *websocket.Dialer
}

// NewExpirenzaClient - Створення нового клієнта для Expirenza aka ShakeToPay API
func NewExpirenzaClient(settings Settings) *Expirenza {
	if settings.Client == nil {
		settings.Client = http.DefaultClient
	}
	if settings.WssClient == nil {
		settings.WssClient = websocket.DefaultDialer
	}
	return &Expirenza{
		wssClient: wss.NewWssClient(wss.Settings{
			WssURL: "wss://api.shaketopay.com.ua/restaurantEntryPoint",
			Dialer: settings.WssClient,
			Headers: http.Header{
				"Authorization": []string{"Basic " + base64.StdEncoding.EncodeToString([]byte(settings.RestoID+":"+settings.SecretKey))},
			},
		}),
		apiClient: &api.APIClient{
			Client:          settings.Client,
			BaseURL:         "https://api.shaketopay.com.ua",
			ErrorParserFunc: errorParser,
		},
		clientMu:     &sync.RWMutex{},
		handlersMu:   &sync.RWMutex{},
		handlers:     make(map[Operation][]*eventHandlerInstance),
		onceHandlers: make(map[Operation][]*eventHandlerInstance),
		restoID:      settings.RestoID,
		secretKey:    settings.SecretKey,
		debug:        settings.Debug,
	}
}

// Open - Відкриття сесії
// При відкритті сесії відбувається підключення до WebSocket сервера
// та запуск прослуховування подій
func (e *Expirenza) Open(ctx context.Context) error {
	if err := e.wssClient.Connect(ctx); err != nil {
		return err
	}
	go func() {
		e.listener = make(chan any)
		e.listen()
	}()

	return nil
}

func (e *Expirenza) listen() {
	e.log("start listening")
	for {
		messageType, message, err := e.wssClient.ReadMessage()
		if err != nil {
			// TODO: reconnect
			return
		}

		select {
		case <-e.listener:
			return
		default:
			e.log("type: %d message: %v", messageType, string(message))
			e.onEvent(messageType, message)
		}
	}
}

func (e *Expirenza) onEvent(_ int, message []byte) {
	var msg map[string]interface{}
	if err := json.Unmarshal(message, &msg); err != nil {
		return
	}

	operationStr, ok := msg["operation"].(string)
	if !ok {
		return
	}
	operation := Operation(operationStr)
	getOperationRequest := func() any {
		var request any
		switch operation {
		case OperationGetBill:
			request = &GetBillEvent{}
		case OperationPayBill:
			request = &PayBillRequest{}
		case OperationCreateOrderOnTable:
			request = &CreateOrderOnTableEvent{}
		case OperationCheckProductsRestrictions:
			request = &CheckProductsRestrictionsEvent{}
		case OperationSplitOrder:
			request = &SplitOrderEvent{}
		case OperationTablesInfo:
			request = &TablesInfoEvent{}
		case OperationUsersInfo:
			request = &UsersInfoEvent{}
		case OperationCategoriesInfo:
			request = &CategoriesInfoEvent{}
		case OperationHallPlansInfo:
			request = &HallPlansInfoEvent{}
		case OperationCheckWork:
			request = &CheckWorkEvent{}
		case OperationSetSettings:
			request = &SetSettingsEvent{}
		case OperationStopList:
			request = &StopListEvent{}
		case OperationVersionInfo:
			request = &VersionInfoEvent{}
		case OperationMenuInfo:
			request = &MenuInfoEvent{}
		default:
			return nil
		}
		e.log("catch event: %s", string(operation))
		if err := json.Unmarshal(message, request); err != nil {
			return nil
		}
		return request
	}
	request := getOperationRequest()
	e.handlersMu.RLock()
	defer e.handlersMu.RUnlock()

	go func() {
		if e.handlers[operation] != nil {
			for _, handler := range e.handlers[operation] {
				e.log("handler: %v request: %v", handler, request)
				go handler.eventHandler.Handle(e, request)
			}
		}
	}()

	go func() {
		if e.onceHandlers[operation] != nil {
			for _, handler := range e.onceHandlers[operation] {
				e.log("handler: %v request: %v", handler, request)
				go handler.eventHandler.Handle(e, request)
			}
			e.onceHandlers[operation] = nil
		}
	}()
}

func (e *Expirenza) StopSession() error {
	return e.wssClient.Close()
}

func (e *Expirenza) Ping() error {
	return e.wssClient.SendMessage(fmt.Sprintf("Resto %s ping", e.restoID))
}
