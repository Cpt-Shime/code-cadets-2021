package fizzbuzz_test

import (
	_ "code-cadests-2021/homework_1/fizzbuzz/fizzbuzz_test"
	fizzbuzz "code-cadests-2021/homework_1/fizzbuzz/fizzbuzz_test"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestPlayFizzBuzz(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := fizzbuzz.Fizzbuzzgame(tc.inputStart, tc.inputEnd)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}