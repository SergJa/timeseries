package timeseries

import (
	"testing"
	"time"
)

func assertEqualSlice(t *testing.T, slice1, slice2 []int) {
	if len(slice1) != len(slice2) {
		t.Error("slice1 and slice2 have different lengths")
	}
	for i:= range slice1 {
		if slice1[i] != slice2[i] {
			t.Error("slice1 and slice2 are different")
		}
	}
}

func assertEqualMap(t *testing.T, map1, map2 map[time.Time]int) {
	if len(map1) != len(map2) {
		t.Error("map1 and map2 have different lengths")
	}
	for i:= range map2 {
		if map1[i] != map2[i] {
			t.Errorf("map1 and map2 are different; time: %v, values: %v, %v", i, map1[i], map2[i])
		}
	}
}


func assertEqualInt(t *testing.T, v1, v2 int) {
	if v1 != v2 {
		t.Errorf("Values are not equal: %v, %v", v1, v2)
	}
}

func assertEqualTime(t *testing.T, t1, t2 time.Time) {
	if !t1.Equal(t2) {
		t.Errorf("Times are not equal: %v, %v", t1, t2)
	}
}

func TestSerie(t *testing.T) {
	quant := time.Second
	ts := NewTimeSeriesInt(quant)
	now := time.Now()
	ts.Add(now, 10)
	ts.Add(now.Add(quant), 15)
	ts.Add(now, 15)
	ts.Add(now.Add(quant*4), 45)

	testStart := now.Add(-quant)
	testEnd := now.Add(quant*7)

	t.Log("Testing GetIntervalSerieSlice")
	assertEqualSlice(t, ts.GetIntervalSerieSlice(testStart, testEnd), []int{0, 25, 15, 0, 0, 45, 0, 0})

	t.Log("Testing GetIntervalSerieMap")
	testMap := map[time.Time]int {
		now.Truncate(quant): 25,
		now.Add(quant).Truncate(quant) : 15,
		now.Add(quant*4).Truncate(quant) : 45,
	}
	assertEqualMap(t, ts.GetIntervalSerieMap(testStart, testEnd), testMap)

	t.Log("Testing AvgPerSecond")
	assertEqualInt(t, ts.AvgPerSecond(testStart, testEnd), 10)

	t.Log("Testing FitstLastTimeTime")
	first, last := ts.FitstLastTimeTime()
	assertEqualTime(t, first, now.Truncate(quant))
	assertEqualTime(t, last, now.Add(quant*4).Truncate(quant))

	t.Log("Test executed")
}
