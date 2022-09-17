package stats

import (
	"math/rand"
	"reflect"
	"strconv"
)

type Stats struct {
	HP     int `invocation:"1" json:"hp"`
	MP     int `invocation:"2" json:"mp"`
	Allure int `invocation:"3" json:"allure"`
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
	numField := statsType.NumField()
	for i := 0; i < numField; i++ {
		statField := statsType.Field(i)
		stInvocation := statField.Tag.Get("invocation")
		invocation, err := strconv.Atoi(stInvocation)
		if err != nil {
			panic(err)
		}

		statVal := int64(getStatValue(id, invocation))

		reflect.ValueOf(&stats).Elem().Field(i).SetInt(statVal)
	}

	return stats
}

func getStatValue(id int64, invocation int) int {

	r := rand.New(rand.NewSource(id))

	var statVal int

	for i := 0; i <= invocation; i++ {
		v := r.Intn(20)
		if i == invocation {
			statVal = v
		}
	}

	return statVal
}
