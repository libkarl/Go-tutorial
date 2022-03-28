package calculator_test

import (
	"gotest/calculator"
	"testing"
)

type TestCase struct {
	value int
	expected bool
	actual bool
}

// očekávám, že vrátí true
func TestCalculateIsArmstrong(t *testing.T) {
	

	t.Run("It should return true for value 371", func(t *testing.T) {
		testCase := TestCase {
			value: 371,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)

		if testCase.actual != testCase.expected {
		t.Fail()
		}
	})	

	t.Run("It should return true for value 370", func(t *testing.T) {
		testCase := TestCase {
			value: 370,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)

		if testCase.actual != testCase.expected {
		t.Fail()
		}
	})	
}
// testuji s očekáváním, že dostanu fail
func TestNegativeCalculateIsArmstrong(t *testing.T)  {
	testCase := TestCase {
		value: 350,
		expected: false,
	}
	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)

	if testCase.actual != testCase.expected {
		t.Fail()
	}
}