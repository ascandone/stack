package interpreter

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSend(t *testing.T) {
	got := CompileFull(`send [COIN 42] (
		source = @src
		destination = @dest	
	)`)

	if len(got.Errors) != 0 {
		t.Fatalf(`Unexpected compilation errors = %#v`, got.Errors)
	}

	expected := Program{
		Statements: []SendStatement{
			{
				Amount:      42,
				Source:      &AccountSrc{"src"},
				Destination: &AccountDest{"dest"},
			},
		},
	}

	require.Equal(t, &expected, got.Program, "Program should be the same.")

}

func TestParseSeq(t *testing.T) {
	got := CompileFull(`send [COIN 42] (
		source = {
			@s1
			@s2
		}
		destination = @dest	
	)`)

	if len(got.Errors) != 0 {
		t.Fatalf(`Unexpected compilation errors = %#v`, got.Errors)
	}

	expected := Program{
		Statements: []SendStatement{
			{
				Amount: 42,
				Source: &SeqSrc{[]Source{
					&AccountSrc{Name: "s1"},
					&AccountSrc{Name: "s2"},
				}},
				Destination: &AccountDest{"dest"},
			},
		},
	}

	require.Equal(t, &expected, got.Program, "Program should be the same.")
}

func TestParseAllottedSource(t *testing.T) {
	got := CompileFull(`send [COIN 42] (
		source = {
			1/3 from @s1
			2/3 from @s2
		}
		destination = @dest	
	)`)

	if len(got.Errors) != 0 {
		t.Fatalf(`Unexpected compilation errors = %#v`, got.Errors)
	}

	expected := Program{
		Statements: []SendStatement{
			{
				Amount: 42,
				Source: &AllottedSrc{[]Allotment[Source]{
					{*big.NewRat(1, 3), &AccountSrc{Name: "s1"}},
					{*big.NewRat(2, 3), &AccountSrc{Name: "s2"}},
				}},
				Destination: &AccountDest{"dest"},
			},
		},
	}

	require.Equal(t, &expected, got.Program, "Program should be the same.")
}

func TestParseAllottedPercSource(t *testing.T) {
	got := CompileFull(`send [COIN 42] (
		source = {
			80% from @s1
			20% from @s2
		}
		destination = @dest	
	)`)

	if len(got.Errors) != 0 {
		t.Fatalf(`Unexpected compilation errors = %#v`, got.Errors)
	}

	expected := Program{
		Statements: []SendStatement{
			{
				Amount: 42,
				Source: &AllottedSrc{[]Allotment[Source]{
					{*big.NewRat(80, 100), &AccountSrc{Name: "s1"}},
					{*big.NewRat(20, 100), &AccountSrc{Name: "s2"}},
				}},
				Destination: &AccountDest{"dest"},
			},
		},
	}

	require.Equal(t, &expected, got.Program, "Program should be the same.")
}

func TestParseAllottedDest(t *testing.T) {
	got := CompileFull(`send [COIN 42] (
		source = @src
		destination = {
			1/3 to @d1
			2/3 to @d2
		}
	)`)

	if len(got.Errors) != 0 {
		t.Fatalf(`Unexpected compilation errors = %#v`, got.Errors)
	}

	expected := Program{
		Statements: []SendStatement{
			{
				Amount: 42,
				Source: &AccountSrc{"src"},
				Destination: &AllottedDest{[]Allotment[Destination]{
					{*big.NewRat(1, 3), &AccountDest{Name: "d1"}},
					{*big.NewRat(2, 3), &AccountDest{Name: "d2"}},
				}},
			},
		},
	}

	require.Equal(t, &expected, got.Program, "Program should be the same.")
}

func TestParseMaxedSource(t *testing.T) {
	got := CompileFull(`send [COIN 42] (
		source = max [COIN 100] from @src
		destination = @dest
	)`)

	if len(got.Errors) != 0 {
		t.Fatalf(`Unexpected compilation errors = %#v`, got.Errors)
	}

	expected := Program{
		Statements: []SendStatement{
			{
				Amount: 42,
				Source: &CappedSrc{
					Cap:    100,
					Source: &AccountSrc{"src"}},
				Destination: &AccountDest{"dest"},
			},
		},
	}

	require.Equal(t, &expected, got.Program, "Program should be the same.")
}

func TestParseMaxedDestination(t *testing.T) {
	got := CompileFull(`send [COIN 42] (
		source = @src
		destination = {
			max [COIN 100] to @d1
			remaining to @d2
		}
	)`)

	if len(got.Errors) != 0 {
		t.Fatalf(`Unexpected compilation errors = %#v`, got.Errors)
	}

	expected := Program{
		Statements: []SendStatement{
			{
				Amount: 42,
				Source: &AccountSrc{"src"},
				Destination: &SeqDest{
					[]Destination{
						&CappedDest{
							Cap:         100,
							Destination: &AccountDest{"d1"},
						},
						&AccountDest{"d2"},
					},
				},
			},
		},
	}

	require.Equal(t, &expected, got.Program, "Program should be the same.")
}
