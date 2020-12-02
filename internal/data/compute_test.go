package data

import "testing"

func TestComputeExpenseReport(t *testing.T) {
	testData := []int{1721, 979, 366, 299, 675, 1456}
	expectedResult := 514579
	result := ComputeExpenseReport(testData)

	if result != expectedResult {
		t.Errorf("unexpected product %d; expected %d", result, expectedResult)
	}
}
