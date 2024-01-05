package expirenza

// HallPlan - обʼєкт залу
type HallPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Table - обʼєкт столика
type Table struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	HallPlanID string `json:"hallPlanId"`
	Number     int    `json:"number"`
}

// User - обʼєкт користувача
type User struct {
	ID        string `json:"id"`
	CellPhone string `json:"cellPhone"`
	Name      string `json:"name"`
	Roles     string `json:"roles"`
}

// Modifier - обʼєкт модифікатора
type Modifier struct {
	ID       string `json:"id"`
	GroupID  string `json:"groupId"`
	Quantity int    `json:"quantity"`
}

// Item - обʼєкт в меню
type Item struct {
	ID       string      `json:"id"`
	Quantity float64     `json:"quantity"`
	Modifier []*Modifier `json:"modifier"`
}

type Modifiers struct {
	LoadModifiers string `json:"load_modifiers"`
}

type Position struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}
