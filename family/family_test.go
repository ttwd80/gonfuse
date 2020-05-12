package family

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttwd80/gonfuse/mocks"
)

func Test1(t *testing.T) {
	mockFamily := mocks.Family{}
	result1 := map[string]string{
		"1": "a",
	}
	mockFamily.On("SimpleFunctionWithAnUnusuallyLongName1").Return(result1, nil)
	actual1, err := mockFamily.SimpleFunctionWithAnUnusuallyLongName1()
	expected1 := map[string]string{
		"1": "a",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected1, actual1)
}

func Test2(t *testing.T) {
	mockFamily := mocks.Family{}
	result2 := map[string]string{
		"2": "b",
	}
	mockFamily.On("SimpleFunctionWithAnUnusuallyLongName2").Return(result2, nil)
	actual2, err := mockFamily.SimpleFunctionWithAnUnusuallyLongName2()
	expected2 := map[string]string{
		"2": "b",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected2, actual2)
}

func Test1And2(t *testing.T) {
	mockFamily := mocks.Family{}

	result1 := map[string]string{
		"1": "a",
	}
	mockFamily.On("SimpleFunctionWithAnUnusuallyLongName1").Return(result1, nil)
	actual1, err := mockFamily.SimpleFunctionWithAnUnusuallyLongName1()
	expected1 := map[string]string{
		"1": "a",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected1, actual1)

	result2 := map[string]string{
		"2": "b",
	}
	mockFamily.On("SimpleFunctionWithAnUnusuallyLongName2").Return(result2, nil)
	actual2, err := mockFamily.SimpleFunctionWithAnUnusuallyLongName2()
	expected2 := map[string]string{
		"2": "b",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected2, actual2)

}
