package timeseries

import (
	"fmt"
	"math"
	"time"
)

type TimeSerieInt struct {
	data map[int64]int
	quantization int64 // quantization interval in nanoseconds
}

// Creates new TimeSeriesInt, implementing integer time series
// quant means step of aggregating
func NewTimeSeriesInt(quant time.Duration) *TimeSerieInt {
	return &TimeSerieInt{
		map[int64]int{},
		quant.Nanoseconds(),
	}
}

func (ts *TimeSerieInt) AvgPerSecond(from, to time.Time) int {
	return ts.Avg(from, to, time.Second)
}

// Adds value to serie, time rounds to quant duration
func (ts *TimeSerieInt) Add(t time.Time, value int) {
	timeInt64 := ts.GetRoundedUnixTime(t)
	ts.data[timeInt64] = ts.data[timeInt64] + value
}

// returns average per quant value
func (ts *TimeSerieInt) Avg(from, to time.Time, quant time.Duration) int {
	x := ts.GetRoundedUnixTime(to)
	y := ts.GetRoundedUnixTime(from)
	quantCount := (x - y) / quant.Nanoseconds()
	if quantCount == 0 {
		quantCount = 1
	}
	return ts.Sum(from, to) / int(quantCount)
}

// returns sum between to and from times
// to, from rounds to quant
func (ts *TimeSerieInt) Sum(from, to time.Time) int {
	fromUnix := ts.GetRoundedUnixTime(from)
	toUnix := ts.GetRoundedUnixTime(to)
	if fromUnix > toUnix {
		return -1
	}
	var sum = 0
	for i:=fromUnix; i<toUnix; i += ts.quantization {
		sum +=ts.data[i]
	}
	return sum
}

func (ts *TimeSerieInt) GetIntervalSerieSlice(from, to time.Time) (serie []int) {
	fromUnix := ts.GetRoundedUnixTime(from)
	toUnix := ts.GetRoundedUnixTime(to)
	if fromUnix > toUnix {
		return
	}
	serie = make([]int, (toUnix - fromUnix) / ts.quantization)
	i2 := 0
	for i:=fromUnix; i<toUnix; i += ts.quantization {
		serie[i2] = ts.data[i]
		i2++
	}
	return
}

func (ts *TimeSerieInt) GetIntervalSerieMap(from, to time.Time) (serie map[time.Time]int) {
	fromUnix := ts.GetRoundedUnixTime(from)
	toUnix := ts.GetRoundedUnixTime(to)
	if fromUnix > toUnix {
		return
	}
	serie = map[time.Time]int{}
	for i:=fromUnix; i<toUnix; i += ts.quantization {
		if value, ok := ts.data[i]; ok {
			serie[time.Unix(0,i)] = value
		}
	}
	return
}

func (ts *TimeSerieInt) FitstLastTimeTime() (first, last time.Time) {
	var minTime int64 = math.MaxInt64
	var maxTime int64 = 0
	for t := range ts.data {
		if minTime > t {
			minTime = t
		}
		if maxTime < t {
			maxTime = t
		}
	}
	return time.Unix(0, minTime), time.Unix(0, maxTime)
}

func (ts *TimeSerieInt) PrettyPrint(showZeroes bool) {
	first, last := ts.FitstLastTimeTime()
	fmt.Println("[")
	for {
		if first.After(last) {
			break
		}
		value := ts.data[ts.GetRoundedUnixTime(first)]
		if showZeroes || value != 0 {
			fmt.Println("   ", first, ":", ts.data[ts.GetRoundedUnixTime(first)])
		}
		first = first.Add(time.Duration(ts.quantization))
	}

	fmt.Println("]")
}

func (ts *TimeSerieInt) GetRoundedUnixTime(t time.Time) int64 {
	//unixNano := (t.Unix()*1000000000 + t.UnixNano())
	unixNano :=  t.UnixNano()
	return unixNano - (unixNano % ts.quantization)
}