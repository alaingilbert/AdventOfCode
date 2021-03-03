package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var departureTime int64 = 1004098

var busesStr = "23,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,509,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29,x,401,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,19"

//var busesStr = "7,13,x,x,59,x,31,19"

func doParseI64(s string) (out int64) {
	out, _ = strconv.ParseInt(s, 10, 64)
	return
}

type Bus struct {
	BusID  int64
	Offset int64
}

func getSortedBuses() (out []Bus) {
	buses := strings.Split(busesStr, ",")
	for idx, busStr := range buses {
		if busStr == "x" {
			continue
		}
		out = append(out, Bus{BusID: doParseI64(busStr), Offset: int64(idx)})
	}
	// Sort buses by bus ID desc
	sort.Slice(out, func(i, j int) bool { return out[i].BusID > out[j].BusID })
	return
}

func getBusNextDepartureTime(busID int64) (nextDepartureTime int64) {
	mul := int64(math.Ceil(float64(departureTime) / float64(busID)))
	return busID * mul
}

func part1() {
	buses := getSortedBuses()
	nextDepartureTime := getBusNextDepartureTime(buses[0].BusID)
	minBusID := buses[0].BusID
	for _, bus := range buses {
		busNextDepartureTime := getBusNextDepartureTime(bus.BusID)
		if busNextDepartureTime < nextDepartureTime {
			nextDepartureTime = busNextDepartureTime
			minBusID = bus.BusID
		}
	}
	waitTime := nextDepartureTime - departureTime
	fmt.Println(minBusID * waitTime)
}

/**
7,0   13,1   59,4   31,6   19,7
7      14     63     37     26

1068781 / 7
1068782 / 13
1068783
1068784
1068785 / 59
1068786
1068787 / 31
1068788 / 19



*/

func part2() {
	buses := getSortedBuses()
	maxBus := buses[0]
	fmt.Println(maxBus)

	var departureTime int64 = 100000000000016
	for {
		departureTime += maxBus.BusID
		if departureTime%100000000 == 0 {
			fmt.Println("progress:", departureTime)
		}
		found := true
		for _, bus := range buses {
			if ((departureTime-maxBus.Offset)+bus.Offset)%bus.BusID != 0 {
				found = false
				break
			}
		}
		if found {
			fmt.Println("found", departureTime-maxBus.Offset)
			break
		}
	}
}

func main() {
	part1()
	part2()
}
