package gain

import "github.com/ctfloyd/hazelmere-api/src/pkg/api"

type UserGains struct {
	Skills map[api.ActivityType]Gain
	Bosses map[api.ActivityType]Gain
}
type Gain struct {
	Name   string
	Amount string
}

type TimeUnit string

const (
	TimeUnitHours TimeUnit = "hours"
	TimeUnitDays  TimeUnit = "days"
)
