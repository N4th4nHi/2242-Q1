package main

import ("testing")

func TestTriangle(t *testing.T) {
	triangle := Triangle{
		base: 25, height: 29,
	}

	expectedArea := 363.0
	expectedPerimiter := 92.0

	gotArea := triangle.area()
	gotPerimiter := triangle.perimiter()

	if gotArea != expectedArea {
		t.Errorf("Expected %v got %v", expectedArea, gotArea)
	}

	if gotPerimiter != expectedPerimiter {
		t.Errorf("Expected %v got %v", expectedPerimiter, expectedPerimiter)
	}
}
