package stats

import (
	"math/rand"
	"reflect"
	"strconv"
)

type Stats struct {
	HP     int `invocation:"1" json:"hp"`
	MP     int `invocation:"2" json:"mp"`
	Attack int `invocation:"3" json:"attack"`
	Magic  int `invocation:"4" json:"magic"`
	Dodge  int `invocation:"5" json:"dodge"`
}

type ClassInfo struct {
	HPModfier     int `json:"hpModfier"`
	MPModfier     int `json:"mpModfier"`
	AttackModfier int `json:"attackModfier"`
	MagicModfier  int `json:"magicModfier"`
	DodgeModfier  int `json:"dodgeModfier"`
}

func NewStats(hp int, mp int) Stats {
	return Stats{
		HP: hp,
		MP: mp,
	}
}

func NewClassInfo(hpModfier int, mpModfier int, attackModifier int, magicModifier int, dodgeModifier int) ClassInfo {
	return ClassInfo{
		HPModfier:     hpModfier,
		MPModfier:     mpModfier,
		AttackModfier: attackModifier,
		MagicModfier:  magicModifier,
		DodgeModfier:  dodgeModifier,
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
		statVal = r.Intn(20)
	}

	return statVal
}

func GetClassInfo(className string) ClassInfo {

	var classInfo ClassInfo

	switch className {
	case "StreetNinja":
		classInfo = NewClassInfo(0, -3, 2, -3, 5)
	case "CyberMage":
		classInfo = NewClassInfo(-3, 3, -2, 5, -5)
	case "CyborgEnforcer":
		classInfo = NewClassInfo(5, -5, 5, -3, 5)
	default:
		classInfo = NewClassInfo(-999, -999, -999, -999, -999)
	}
	return classInfo
}
