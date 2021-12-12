package Pet

import "time"

type IPetFactory interface {
	Create(userId int, petName string, gender string, breed string, birthday time.Time, adoptaversary time.Time, memo string) (error error)
}
