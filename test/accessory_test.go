package test

import (
	"github.com/mahmudaZaman/commonutil/comutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccessoryTestSuit struct {
	suite.Suite
}

func (s *AccessoryTestSuit) Test01() {
	var str1 string = ""
	var str2 string = "text-1"
	var str3 string = "text-2"
	comutil.FirstNotNullString(str1, str2, str3)
	assert.Equal(s.T(), str2, "text-1", "Should return text-1")
}

func TestAccessoryTestSuit(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AccessoryTestSuit))
}
