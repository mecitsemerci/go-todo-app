package utils

import (
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilsTestSuite struct {
	suite.Suite
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

func (s *UtilsTestSuite) SetupTest() {}

func (s *UtilsTestSuite) Test_IsEmptyOrWhitespace_Should_Return_True() {
	// Given
	strList := []string{"", " ", "\t\t\t  \n"}

	for _, str := range strList {
		// When
		result := IsEmptyOrWhiteSpace(str)

		// Then
		assert.Equal(s.T(), true, result)
	}
}

func (s *UtilsTestSuite) Test_IsEmptyOrWhitespace_Should_Return_False() {
	// Given
	str := "Foo"

	// When
	result := IsEmptyOrWhiteSpace(str)

	// Then
	assert.Equal(s.T(), false, result)
}
