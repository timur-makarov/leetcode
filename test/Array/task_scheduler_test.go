package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/task_scheduler"
	"log"
	"testing"
)

func TestTaskScheduler(t *testing.T) {
	tasksString := "AAABBB"
	n := 2
	answer := 8
	result := task_scheduler.RunSolution([]byte(tasksString), n)
	assert.Equal(t, answer, result)

	log.Println("NEXT TEST")

	tasksString = "ACABDB"
	n = 1
	answer = 6
	result = task_scheduler.RunSolution([]byte(tasksString), n)
	assert.Equal(t, answer, result)

	log.Println("NEXT TEST")

	tasksString = "AAABBB"
	n = 3
	answer = 10
	result = task_scheduler.RunSolution([]byte(tasksString), n)
	assert.Equal(t, answer, result)
}
