package day7

import (
	"fmt"
	"testing"
)

var example = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func Test_run(t *testing.T) {
	run()
}

func Test_camelCardsPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{input: example}, want: 6440},
		{name: "example 2", args: args{[]string{
			"2345A 1",
			"Q2KJJ 13",
			"Q2Q2Q 19",
			"T3T3J 17",
			"T3Q33 11",
			"2345J 3",
			"J345A 2",
			"32T3K 5",
			"T55J5 29",
			"KK677 7",
			"KTJJT 34",
			"QQQJA 31",
			"JJJJJ 37",
			"JAAAA 43",
			"AAAAJ 59",
			"AAAAA 61",
			"2AAAA 23",
			"2JJJJ 53",
			"JJJJ2 41",
		}}, want: 6592},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelCardsPart1(tt.args.input); got != tt.want {
				t.Errorf("camelCardsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_getType(t *testing.T) {
	tests := []struct {
		cards string
		want  HandType
	}{
		{cards: "AAAAA", want: fiveOfKind},
		{cards: "AAAAQ", want: fourOfKind},
		{cards: "QAAAA", want: fourOfKind},
		{cards: "QQAAA", want: fullHouse},
		{cards: "AAAQQ", want: fullHouse},
		{cards: "QJAAA", want: threeOfKind},
		{cards: "QQQJA", want: threeOfKind},
		{cards: "QQJJA", want: twoPair},
		{cards: "AAJJQ", want: twoPair},
		{cards: "AKQJJ", want: onePair},
		{cards: "AKQJ9", want: highCard},
	}
	for _, tt := range tests {
		t.Run(tt.cards, func(t *testing.T) {
			var cardsArray [5]CardType
			for i, c := range tt.cards {
				cardsArray[i] = CardType(c)
			}
			h := &Hand{
				cards: cardsArray,
			}
			if got := h.getType(false); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_camelCardsPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		}}, want: 5905},
		{name: "example 2", args: args{[]string{
			"2345A 1",
			"Q2KJJ 13",
			"Q2Q2Q 19",
			"T3T3J 17",
			"T3Q33 11",
			"2345J 3",
			"J345A 2",
			"32T3K 5",
			"T55J5 29",
			"KK677 7",
			"KTJJT 34",
			"QQQJA 31",
			"JJJJJ 37",
			"JAAAA 43",
			"AAAAJ 59",
			"AAAAA 61",
			"2AAAA 23",
			"2JJJJ 53",
			"JJJJ2 41",
		}}, want: 6839},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelCardsPart2(tt.args.input); got != tt.want {
				t.Errorf("camelCardsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_getTypeWithJokers(t *testing.T) {
	tests := []struct {
		card string
		want HandType
	}{
		{card: "JJJJJ", want: fiveOfKind},
		{card: "AAAAA", want: fiveOfKind},
		{card: "AAAAJ", want: fiveOfKind},
		{card: "AAAAQ", want: fourOfKind},
		{card: "QQAAA", want: fullHouse},
		{card: "AABBJ", want: fullHouse},
		{card: "2233J", want: fullHouse},
		{card: "AABBJ", want: fullHouse},
		{card: "JJJ34", want: fourOfKind},
		{card: "JJAAA", want: fiveOfKind},
		{card: "AJJQQ", want: fourOfKind},
		{card: "QJAAA", want: fourOfKind},
		{card: "QQQJA", want: fourOfKind},
		{card: "QQJJA", want: fourOfKind},
		{card: "AAJJQ", want: fourOfKind},
		{card: "AAKKJ", want: fullHouse},
		{card: "AKQJJ", want: threeOfKind},
		{card: "AKQ8J", want: onePair},
		{card: "2345J", want: onePair},
		{card: "AKQ89", want: highCard},
		{card: "12345", want: highCard},
		{card: "12344", want: onePair},
		{card: "1234J", want: onePair},
		{card: "12233", want: twoPair},
		{card: "12333", want: threeOfKind},
		{card: "1233J", want: threeOfKind},
		{card: "123JJ", want: threeOfKind},
		{card: "11222", want: fullHouse},
		{card: "1122J", want: fullHouse},
		{card: "12222", want: fourOfKind},
		{card: "1222J", want: fourOfKind},
		{card: "122JJ", want: fourOfKind},
		{card: "12JJJ", want: fourOfKind},
		{card: "11111", want: fiveOfKind},
		{card: "1111J", want: fiveOfKind},
		{card: "111JJ", want: fiveOfKind},
		{card: "11JJJ", want: fiveOfKind},
		{card: "1JJJJ", want: fiveOfKind},
		{card: "JJJJJ", want: fiveOfKind},
	}
	for _, tt := range tests {
		t.Run(tt.card, func(t *testing.T) {
			var cardsArray [5]CardType
			for i, c := range tt.card {
				cardsArray[i] = CardType(c)
			}
			h := &Hand{cards: cardsArray}
			if got := h.getType(true); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_compare(t *testing.T) {
	type args struct {
		cards  string
		jokers bool
	}
	tests := []struct {
		cards string
		args  args
		want  bool
	}{
		{cards: "AAAAA", args: args{cards: "AAAAA", jokers: false}, want: false},
		{cards: "JJJJJ", args: args{cards: "AAAAA", jokers: true}, want: true},
		{cards: "JJJJJ", args: args{cards: "AAAAA", jokers: false}, want: true},
		{cards: "AAAAA", args: args{cards: "AAAAA", jokers: true}, want: false},
		{cards: "22222", args: args{cards: "AAAAA", jokers: false}, want: true},
		{cards: "22222", args: args{cards: "AAAAA", jokers: true}, want: true},
		{cards: "AAAAQ", args: args{cards: "AAAAA", jokers: true}, want: true},
		{cards: "AAAAJ", args: args{cards: "AAAAA", jokers: true}, want: true},
		{cards: "AAAAJ", args: args{cards: "AAAAQ", jokers: true}, want: false},
		{cards: "KKKKK", args: args{cards: "AAAAA", jokers: true}, want: true},
		{cards: "KKKJJ", args: args{cards: "AAAAQ", jokers: true}, want: false},
		{cards: "2233J", args: args{cards: "22333", jokers: true}, want: true},
		{cards: "223JJ", args: args{cards: "22333", jokers: true}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s:%s(j:%v)", tt.cards, tt.args.cards, tt.args.jokers), func(t *testing.T) {

			var cardsArray [5]CardType
			for i, c := range tt.cards {
				cardsArray[i] = CardType(c)
			}

			var cardsArray2 [5]CardType
			for i, c := range tt.args.cards {
				cardsArray2[i] = CardType(c)
			}

			h := &Hand{cards: cardsArray}
			h2 := &Hand{cards: cardsArray2}
			if got := h.compare(*h2, tt.args.jokers); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
