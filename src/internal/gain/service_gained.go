package gain

import (
	"context"
	"errors"
	"github.com/ctfloyd/hazelmere-api/src/pkg/api"
	"github.com/ctfloyd/hazelmere-api/src/pkg/client"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
	"time"
)

var ErrHiscoreTimeout = errors.New("hiscore timeout")

type GainedService interface {
	CalculateUserGains(username string, duration int, unit TimeUnit) (UserGains, error)
}
type gainedService struct {
	logger    hz_logger.Logger
	hazelmere *client.Hazelmere
}

func NewGainedService(logger hz_logger.Logger, hazelmere *client.Hazelmere) GainedService {
	return &gainedService{logger: logger, hazelmere: hazelmere}
}

func (gs *gainedService) CalculateUserGains(username string, duration int, unit TimeUnit) (UserGains, error) {
	userResponse, err := gs.hazelmere.User.GetAllUsers()
	if err != nil {
		return UserGains{}, err
	}

	userId := ""
	for _, user := range userResponse.Users {
		if strings.EqualFold(user.RunescapeName, username) {
			userId = user.Id
		}
	}

	if userId == "" {
		return UserGains{}, errors.New("user not found")
	}

	var d time.Duration
	if unit == TimeUnitDays {
		d = time.Hour * 24
	} else if unit == TimeUnitHours {
		d = time.Hour
	}

	var newSnapshot api.HiscoreSnapshot
	latest, err := gs.hazelmere.Worker.GenerateSnapshotOnDemand(userId)
	if err != nil {
		if errors.Is(err, client.ErrHiscoreTimeout) || errors.Is(err, context.DeadlineExceeded) {
			return UserGains{}, ErrHiscoreTimeout
		}
		return UserGains{}, errors.Join(errors.New("generate snapshot on demand"), err)
	}
	newSnapshot = latest.Snapshot

	var oldSnapshot api.HiscoreSnapshot
	timestamp := time.Now().Add(time.Duration(-1*duration) * d)
	oldResponse, err := gs.hazelmere.Snapshot.GetSnapshotForUserNearestTimestamp(userId, timestamp.UnixMilli())
	if err != nil {
		return UserGains{}, err
	}
	oldSnapshot = oldResponse.Snapshot

	gains := UserGains{
		Skills: make(map[api.ActivityType]Gain),
		Bosses: make(map[api.ActivityType]Gain),
	}

	p := message.NewPrinter(language.English)
	for _, skill := range newSnapshot.Skills {
		var oldSkill *api.SkillSnapshot

		for _, os := range oldSnapshot.Skills {
			if os.ActivityType == skill.ActivityType {
				oldSkill = &os
			}
		}

		if oldSkill == nil {
			continue
		}

		if skill.Experience-oldSkill.Experience > 0 {
			gains.Skills[skill.ActivityType] = Gain{
				Name:   skill.Name,
				Amount: p.Sprintf("%d", skill.Experience-oldSkill.Experience),
			}
		}
	}

	for _, boss := range newSnapshot.Bosses {
		var oldBoss *api.BossSnapshot

		for _, ob := range oldSnapshot.Bosses {
			if ob.ActivityType == boss.ActivityType {
				oldBoss = &ob
			}
		}

		if oldBoss == nil {
			continue
		}

		if boss.KillCount-oldBoss.KillCount > 0 {
			gains.Bosses[boss.ActivityType] = Gain{
				Name:   boss.Name,
				Amount: p.Sprintf("%d", boss.KillCount-oldBoss.KillCount),
			}
		}
	}

	return gains, nil
}

func TimeUnitFromString(timeUnit string) TimeUnit {
	if strings.EqualFold(timeUnit, "hour") || strings.EqualFold(timeUnit, "hours") {
		return TimeUnitHours
	}

	if strings.EqualFold(timeUnit, "day") || strings.EqualFold(timeUnit, "days") {
		return TimeUnitDays
	}

	return TimeUnitDays
}
