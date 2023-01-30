package main

import (
	"testing"
)

func TestCircle(t *testing.T) {
	circle := Circle{
		radius: 35,
	}

	expectedArea := 3848.0
	expectedPerimiter := 220.0

	gotArea := circle.area()
	gotPerimiter := circle.perimiter()

	if gotArea != expectedArea {
		t.Errorf("Expected %v got %v", expectedArea, gotArea)
	}

	if gotPerimiter != expectedPerimiter {
		t.Errorf("Expected %v got %v", expectedPerimiter, expectedPerimiter)
	}

}
