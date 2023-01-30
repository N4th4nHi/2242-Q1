package main

import ("testing")

func TestSquare(t *testing.T) {

	area, perimiter := square(20)

	expectedArea := 400.0
	expectedPerimiter := 80.0

	if area != expectedArea {
		t.Errorf("Expected %v got %v", expectedArea, area)
	}

	if perimiter != expectedPerimiter {
		t.Errorf("Expected %v got %v", expectedPerimiter, perimiter)
	}
}
