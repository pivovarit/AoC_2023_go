package day1

import "testing"

func Test_run(t *testing.T) {
	run()
}

func Test_trebuchet_part_1(t *testing.T) {
	result := trebuchet_part_1([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})

	if result != 142 {
		t.Fatalf("result should be equal to 142 and not %d", result)
	}
}

func Test_trebuchet_part_2(t *testing.T) {
	result := trebuchet_part_2([]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	})

	if result != 281 {
		t.Fatalf("result should be equal to 281 and not %d", result)
	}
}
