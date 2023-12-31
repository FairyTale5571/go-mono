package api

// ErrorParser это интерфейс для разбора тела ошибки
type ErrorParser interface {
	Error([]byte) error
}

// ErrorParserFunc тип функции, который соответствует интерфейсу ErrorParser
type ErrorParserFunc func([]byte) error

// Error реализация интерфейса Error для функций
func (f ErrorParserFunc) Error(body []byte) error {
	return f(body)
}
