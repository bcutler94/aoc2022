package main

import (
	"testing"
)

type TestArr []struct{ word string; len int; want int }

func TestGetIndex(t *testing.T) {

	var m = TestArr{
		{word: "bvwbjplbgvbhsrlpgdmjqwftvncz", len: 4, want: 5}, 
		{word: "nppdvjthqldpwncqszvftbrmjlhg", len: 4, want: 6}, 
		{word: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", len: 4, want: 10}, 
		{word: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", len: 4, want: 11},
		{word: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", len: 14, want: 19}, 
		{word: "bvwbjplbgvbhsrlpgdmjqwftvncz", len: 14, want: 23}, 
		{word: "nppdvjthqldpwncqszvftbrmjlhg", len: 14, want: 23}, 
		{word: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", len: 14, want: 29},
		{word: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", len: 14, want: 26}, 
	}

	for _, testInfo := range m {
		want := testInfo.want
		got := GetIndex(testInfo.word, testInfo.len)
		if want != got {
			t.Fatalf(`GetIndex(%v, %v) = %v, but wanted %v`, testInfo.word, testInfo.len, got, want)
		}
	}

}