package acquiring

import (
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

type Acquiring struct {
	// token Токен з особистого кабінету https://web.monobank.ua/ або тестовий токен з https://api.monobank.ua/
	token string

	// cmsName Назва CMS, якщо ви розробляєте платіжний модуль для CMS
	cmsName string

	// cmsVersion Версія CMS, якщо ви розробляєте платіжний модуль для CMS
	cmsVersion string

	// apiClient Клієнт для відправки запитів
	apiClient *api.APIClient
}

func (a *Acquiring) sign() string {
	return "acquiring"
}

// SetCMS Встановлює назву CMS
func (a *Acquiring) SetCMS(cms string) *Acquiring {
	a.cmsName = cms
	return a
}

// SetCMSVersion Встановлює версію CMS
func (a *Acquiring) SetCMSVersion(version string) *Acquiring {
	a.cmsVersion = version
	return a
}

// Opts Налаштування
type Opts struct {
	Token  string
	Client *http.Client
}

const (
	// acquiringApiUrl URL для відправки запитів
	acquiringApiUrl = "https://api.monobank.ua/api/merchant"
)

// NewAcquiring Створює новий екземпляр класу
func NewAcquiring(opts Opts) *Acquiring {
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	return &Acquiring{
		token: opts.Token,
		apiClient: &api.APIClient{
			BaseURL: acquiringApiUrl,
			Client:  opts.Client,
		},
	}
}
