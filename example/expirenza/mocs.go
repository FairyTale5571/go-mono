package main

import "github.com/FairyTale5571/go-mono/expirenza"

var (
	tables = []*expirenza.Table{
		{
			ID:         "1",
			Name:       "Столик у кутку",
			HallPlanID: "1",
			Number:     1,
		},
		{
			ID:         "2",
			Name:       "В курятнику",
			HallPlanID: "2",
			Number:     2,
		},
		{
			ID:         "3",
			Name:       "Сидить на яйцях",
			HallPlanID: "2",
			Number:     3,
		},
		{
			ID:         "4",
			Name:       "Кукарекае",
			HallPlanID: "2",
			Number:     4,
		},
	}
	hallPlans = []*expirenza.HallPlan{
		{
			ID:   "1",
			Name: "Великий зал",
		},
		{
			ID:   "2",
			Name: "Курячий зал",
		},
	}
	waiters = []*expirenza.User{
		{
			ID:        "1",
			CellPhone: "+380123456789",
			Name:      "Олександр",
			Roles:     "waiter",
		},
		{
			ID:        "2",
			CellPhone: "+380987654321",
			Name:      "Василь",
			Roles:     "waiter",
		},
		{
			ID:        "3",
			CellPhone: "+380990679259",
			Name:      "Олександр",
			Roles:     "developer",
		},
	}
)
