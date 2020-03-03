## Simple timeseries library
Golang Timeseries simple library. For time parameter should be used ```time.Time```, value is ```int```.

## Get Package
```bash
$ go get github.com/SergJa/timeseries
```

## Import Package
```go
import "github.com/SergJa/timeseries"
```

## Create a TimeSeries Struct
```go
ts := timeseries.NewTimeSeriesIntquantDuration)
```

quantDuration must be ```time.Duration```

## Manipulate with data
Add Data with Time
```go
ts.Add(time.Now(), 12)
```

Get values from ```timeStart``` to ```timeStop``` as slice
```go
rangeVals = ts.GetIntervalSerieSlice(timeStart, timeStop)
```
Zero all values from ```timeStart``` to ```timeStop``` 
```go
ts.ClearInterval(timeStart, timeStop)
```
Get values from ```timeStart``` to ```timeStop``` as as map, excluding zero values
```go
rangeVals = ts.GetIntervalSerieMap(timeStart, timeStop)
```
Get first and last times
```
first, last := ts.FitstLastTimeTime()
```
Get average per ```time.Duration``` 
```
avg = ts.Avg(from, to, quant)
```
where from, to - ```time.Time```, quant - ```time.Duration```

Get average per second
```
avgPerSecond = ts.AvgPerSecond(from, to, quant)
```
Show all data (for debug usage)
```
ts.PrettyPrint(true)
```
boolean parameter indicates whether show zero values

## Example
```go
import (
    "github.com/SergJa/timeseries"
    "fmt"
    "time"
)

func main(){
	ts := timeseries.NewTimeSeriesInt(time.Second)
	ts.Add(now, 10)
	ts.Add(now.Add(time.Second), 15)
	ts.Add(now, 15)
	ts.Add(now.Add(time.Second*4), 45)
	ts.PrettyPrint(false)


	testStart := now.Add(-time.Second)
	testEnd := now.Add(time.Second*7)

	fmt.Println(ts.GetIntervalSerieSlice(testStart, testEnd))
	fmt.Println(ts.GetIntervalSerieMap(testStart, testEnd))
	fmt.Println(ts.AvgPerSecond(testStart, testEnd))
	fmt.Println(ts.FitstLastTimeTime())
}
```