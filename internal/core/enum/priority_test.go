package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PriorityEnumTestSuite struct {
	suite.Suite
}

func TestEnumsTestSuite(t *testing.T) {
	suite.Run(t, new(PriorityEnumTestSuite))
}

func (s *PriorityEnumTestSuite) Test_PriorityLevel_Should_Same_Value_When_Created_From_String() {
	//Given
	values := []int{0, 1, 2, 3}

	//When
	var priorityLevels = make([]PriorityLevel, 0)
	for _, v := range values {
		priorityLevels = append(priorityLevels, PriorityLevel(v))
	}

	//Then
	assert.NotEmpty(s.T(), priorityLevels)
	assert.Equal(s.T(), PriorityNone, priorityLevels[0])
	assert.Equal(s.T(), PriorityNormal, priorityLevels[1])
	assert.Equal(s.T(), PriorityHigh, priorityLevels[2])
	assert.Equal(s.T(), PriorityCritical, priorityLevels[3])
}
