package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func getBuses() (out []Bus) {
	buses := strings.Split(busesStr, ",")
	for idx, busStr := range buses {
		if busStr == "x" {
			continue
		}
		out = append(out, Bus{BusID: doParseI64(busStr), Offset: int64(idx)})
	}
	return
}

func getBusNextDepartureTime(departureTime, busID int64) (nextDepartureTime int64) {
	mul := int64(math.Ceil(float64(departureTime) / float64(busID)))
	return busID * mul
}

func part1() {
	departureTime := int64(1004098)
	buses := getBuses()
	nextDepartureTime := getBusNextDepartureTime(departureTime, buses[0].BusID)
	minBusID := buses[0].BusID
	for _, bus := range buses {
		busNextDepartureTime := getBusNextDepartureTime(departureTime, bus.BusID)
		if busNextDepartureTime < nextDepartureTime {
			nextDepartureTime = busNextDepartureTime
			minBusID = bus.BusID
		}
	}
	waitTime := nextDepartureTime - departureTime
	fmt.Println(minBusID * waitTime)
}

func solve(bus1, bus2 Bus) (out int64) {
	a := bus2.Offset - bus1.Offset
	for {
		if a%bus2.BusID == 0 {
			return out
		}
		a += bus1.BusID
		out++
	}
}

func mergeBuses(bus1, bus2 Bus) Bus {
	x1 := solve(bus1, bus2)
	return Bus{Offset: bus1.Offset - (bus1.BusID * x1), BusID: bus1.BusID * bus2.BusID}
}

func part2() {
	buses := getBuses()
	var res, bus1, bus2 Bus
	for {
		bus1, bus2, buses = buses[0], buses[1], buses[2:]
		res = mergeBuses(bus1, bus2)
		buses = append(buses, res)
		if len(buses) == 1 {
			break
		}
	}
	fmt.Println(-res.Offset)
}

func main() {
	part1()
	part2()
}
