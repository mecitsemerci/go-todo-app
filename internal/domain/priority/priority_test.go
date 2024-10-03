package priority

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
	var priorityLevels = make([]Level, 0)
	for _, v := range values {
		priorityLevels = append(priorityLevels, Level(v))
	}

	//Then
	assert.NotEmpty(s.T(), priorityLevels)
	assert.Equal(s.T(), None, priorityLevels[0])
	assert.Equal(s.T(), Normal, priorityLevels[1])
	assert.Equal(s.T(), High, priorityLevels[2])
	assert.Equal(s.T(), Critical, priorityLevels[3])
}
