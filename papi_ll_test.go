// Copyright (C) 2011, Los Alamos National Security, LLC.
// Use of this source code is governed by a BSD-style license.

// This file tests the low-level features of the PAPI library.

package papi

import "testing"
import "time"


// Ensure that the real-time cycle counter is strictly increasing.
func TestRealCyc(t *testing.T) {
	time1 := GetRealCyc()
	time2 := GetRealCyc()
	if time1 >= time2 {
		t.Fatalf("Real-time cycles are not strictly increasing: %d >= %d",
			time1, time2)
	}
}


// Ensure that the real-time microsecond counter is increasing.
func TestRealUsec(t *testing.T) {
	const sleep_usecs = 10000
	time1 := GetRealUsec()
	time.Sleep(sleep_usecs * 1000)
	time2 := GetRealUsec()
	if time1 > time2 {
		t.Fatalf("Real-time microseconds are decreasing: %d >= %d",
			time1, time2)
	} else {
		if time2-time1 < sleep_usecs {
			t.Fatalf("Real-time microseconds increase too slowly: %d vs. %d",
				time2-time1, sleep_usecs)
		}
	}
}


// Ensure that the virtual-time cycle counter is increasing.  Ideally,
// it should be strictly increasing, but this doesn't seem to be the
// case on all systems.
func TestVirtCyc(t *testing.T) {
	const maxTimings = 1000000000
	time1 := GetVirtCyc()
	time2 := time1
	for i := 0; time1 == time2 && i < maxTimings; i++ {
		time2 = GetVirtCyc()
	}
	if time1 > time2 {
		t.Fatalf("Virtual-time cycles are decreasing: %d >= %d",
			time1, time2)
	}
}


// Ensure that the virtual-time microsecond counter is increasing.
func TestVirtUsec(t *testing.T) {
	const maxTimings = 1000000000
	time1 := GetVirtUsec()
	time2 := time1
	for i := 0; time1 == time2 && i < maxTimings; i++ {
		time2 = GetVirtUsec()
	}
	if time1 > time2 {
		t.Fatalf("Virtual-time microseconds are decreasing: %d >= %d",
			time1, time2)
	}
}