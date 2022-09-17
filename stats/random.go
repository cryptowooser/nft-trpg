package stats

import (
	"math/rand"
	"reflect"
	"strconv"
)

type Stats struct {
	HP int `statIndex:"1" json:"hp"`
	MP int `statIndex:"2" json:"mp"`
}

func NewStats(hp int, mp int) Stats {
	return Stats{
		HP: hp,
		MP: mp,
	}
}

func GetStats(id int64) Stats {

	stats := Stats{}
	statsType := reflect.TypeOf(stats)

	for i := 0; i <= statsType.NumField(); i++ {
		statField := statsType.Field(i)
		strIndex := statField.Tag.Get("statIndex")
		statIndex, err := strconv.Atoi(strIndex)
		if err != nil {
			panic(err)
		}

		statVal := getStatValue(id, statIndex)

		reflect.ValueOf(stats).Field(i).SetInt(int64(statVal))
	}

	return stats
}

func getStatValue(id int64, statIndex int) int {

	r := rand.New(rand.NewSource(id))

	var statVal int

	for i := 1; i <= statIndex; i++ {
		v := r.Intn(20)
		if i == statIndex {
			statVal = v
		}
	}

	return statVal
}
