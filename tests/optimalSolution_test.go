package test

import (
	"demo/utils/algorithm"
	"errors"
	"fmt"
	"testing"
)

func TestGenerateEnemies(t *testing.T) {
	tests := []struct {
		numEnemies int
		batchSizes []int
		wantErr    error
	}{
		{10, []int{4, 2, 2, 4}, nil},
		{20, []int{1, 2, 5, 10}, nil},
		{30, []int{1, 2, 6, 10}, nil},
		{0, []int{1, 2, 3}, errors.New("numEnemies must be a positive integer")},
		{10, []int{1, 2, 0}, errors.New("batchSize must be a positive integer")},
		{10, []int{1, 2, 3, 5}, errors.New("batchSize must be an even number")},
	}

	for _, test := range tests {
		_, err := algorithm.GenerateEnemies(test.numEnemies, test.batchSizes)
		if !errors.Is(err, test.wantErr) {
			t.Errorf("GenerateEnemies(%d, %v) returned error %v, want %v", test.numEnemies, test.batchSizes, err, test.wantErr)
		}
	}

	// Test for adjacent X coordinates and same Y coordinate
	for _, test := range tests[:3] {
		enemies, err := algorithm.GenerateEnemies(test.numEnemies, test.batchSizes)
		if err != nil {
			t.Errorf("GenerateEnemies(%d, %v) returned error: %v", test.numEnemies, test.batchSizes, err)
			continue
		}

		if len(enemies) != test.numEnemies {
			t.Errorf("GenerateEnemies(%d, %v) returned %d enemies instead of %d", test.numEnemies, test.batchSizes, len(enemies), test.numEnemies)
			continue
		}

		for i := 1; i < test.numEnemies; i++ {
			if enemies[i].X-enemies[i-1].X != 2 || enemies[i].Y != enemies[i-1].Y {
				t.Errorf("GenerateEnemies(%d, %v): enemies %d and %d are not adjacent or have different Y coordinates", test.numEnemies, test.batchSizes, i-1, i)
			}
		}
	}

	// Test for invalid X coordinates
	_, err := algorithm.GenerateEnemies(10, []int{2, 2, 2, 2})
	if err != nil {
		t.Errorf("GenerateEnemies(10, []int{2, 2, 2, 2}) returned error: %v", err)
	}
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			wantErr := fmt.Errorf("enemy %d has invalid X coordinate %d (not even)", i, i)
			_, err := algorithm.GenerateEnemies(10, []int{2, 2, 2, 2})
			if !errors.Is(err, wantErr) {
				t.Errorf("GenerateEnemies(10, []int{2, 2, 2, 2}) returned error %v for enemy %d, want %v", err, i, wantErr)
			}
		}
	}
}
