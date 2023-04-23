package algorithm

import (
	"errors"
	"fmt"
	"math/rand"
)

type Obstacle struct {
	X int
	Y int
}

type Enemy struct {
	X int
	Y int
}

func GenerateObstacles(numObstacles int, enemies []Enemy) []Obstacle {
	obstacles := make([]Obstacle, 0)

	for i := 0; i < numObstacles; i++ {
		var o Obstacle
		for {
			o = Obstacle{rand.Intn(10), rand.Intn(10)}
			if !containsObstacle(obstacles, o.X, o.Y) && !containsEnemy(enemies, Enemy{o.X, o.Y}) {
				break
			}
		}
		obstacles = append(obstacles, o)
	}

	return obstacles
}

func GenerateEnemies(numEnemies int, batchSizes []int) ([]Enemy, error) {
	// Check for invalid inputs
	if numEnemies <= 0 {
		return nil, errors.New("numEnemies must be a positive integer")
	}
	for _, batchSize := range batchSizes {
		if batchSize <= 0 {
			return nil, errors.New("batchSize must be a positive integer")
		}
		if batchSize%2 != 0 {
			return nil, errors.New("batchSize must be an even number")
		}
	}

	enemies := make([]Enemy, numEnemies)
	batchNum := len(batchSizes)

	startX := 0
	for i := 0; i < batchNum; i++ {
		batchSizeToUse := batchSizes[i]

		if i > 0 && i*batchSizes[i-1]-1 < numEnemies {
			lastEnemy := enemies[i*batchSizes[i-1]-1]
			startX = lastEnemy.X + 2
		}

		if startX+batchSizeToUse*2-1 > 10 {
			startX = 10 - batchSizeToUse*2 + 1
		}

		for j := 0; j < batchSizeToUse; j++ {
			if i*batchSizeToUse+j >= numEnemies {
				break
			}
			enemies[i*batchSizeToUse+j].X = startX + j*2
			if enemies[i*batchSizeToUse+j].X > 10 {
				return nil, fmt.Errorf("enemy %d has invalid X coordinate %d (exceeds 10)", i*batchSizeToUse+j, enemies[i*batchSizeToUse+j].X)
			}
			enemies[i*batchSizeToUse+j].Y = rand.Intn(10)
		}
	}

	// Check for invalid X coordinates
	for i := 0; i < numEnemies; i++ {
		if enemies[i].X%2 != 0 {
			return nil, fmt.Errorf("enemy %d has invalid X coordinate %d (not even)", i, enemies[i].X)
		}
	}

	return enemies, nil
}

func containsEnemy(enemies []Enemy, e Enemy) bool {
	for _, enemy := range enemies {
		if enemy.X == e.X && enemy.Y == e.Y {
			return true
		}
	}
	return false
}

func containsObstacle(obstacles []Obstacle, x int, y int) bool {
	for _, o := range obstacles {
		if o.X == x && o.Y == y {
			return true
		}
	}
	return false
}

func GenerateMatrix(numObstacles int, numEnemies int, batchSizes []int) [][]byte {
	if numObstacles == 0 || numEnemies == 0 {
		return nil
	}

	matrix := make([][]byte, 10)
	for i := range matrix {
		matrix[i] = make([]byte, 10)
	}

	obstacles := GenerateObstacles(numObstacles, nil)
	enemies, _ := GenerateEnemies(numEnemies, batchSizes)

	for _, o := range obstacles {
		if o.X >= 0 && o.X < 10 && o.Y >= 0 && o.Y < 10 {
			matrix[o.X][o.Y] = 'O'
		}
	}

	for _, e := range enemies {
		if e.X >= 0 && e.X < 10 && e.Y >= 0 && e.Y < 10 {
			matrix[e.X][e.Y] = 'E'
		}
	}

	return matrix
}
