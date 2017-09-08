// Code generated by gocc; DO NOT EDIT.

package parser

import ( "github.com/sirkon/ldetool/ast" )

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Rules	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Rules : Rules Rule	<< ast.RuleSeq(X[0], X[1]) >>`,
		Id:         "Rules",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.RuleSeq(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Rules : Rule	<< ast.LastItem(X[0]) >>`,
		Id:         "Rules",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.LastItem(X[0])
		},
	},
	ProdTabEntry{
		String: `Rule : identifier "=" Action ";"	<< ast.Rule(X[0], X[2]) >>`,
		Id:         "Rule",
		NTType:     2,
		Index:      3,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Rule(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Action : "(" Action ")" Action	<< ast.ActionSeq(X[1], X[3]) >>`,
		Id:         "Action",
		NTType:     3,
		Index:      4,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ActionSeq(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Action : "(" Action ")"	<< X[1], nil >>`,
		Id:         "Action",
		NTType:     3,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Action : AtomicAction Action	<< ast.ActionSeq(X[0], X[1]) >>`,
		Id:         "Action",
		NTType:     3,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ActionSeq(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Action : AtomicAction	<< ast.Action(X[0]) >>`,
		Id:         "Action",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Action(X[0])
		},
	},
	ProdTabEntry{
		String: `Action : stress Action	<< ast.MatchRequired(X[1]) >>`,
		Id:         "Action",
		NTType:     3,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.MatchRequired(X[1])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "^" string_lit	<< ast.StartsWithString(X[1]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StartsWithString(X[1])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "^" char_lit	<< ast.StartsWithChar(X[1]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StartsWithChar(X[1])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "^" "??" string_lit	<< ast.MayBeStartsWithString(X[2]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.MayBeStartsWithString(X[2])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "^" "??" char_lit	<< ast.MayBeStartsWithChar(X[2]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.MayBeStartsWithChar(X[2])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "_" "[" int_lit ":" "]"	<< ast.PassFirst(X[2]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      13,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.PassFirst(X[2])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "_" Target	<< ast.PassUntilTarget(X[1]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      14,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.PassUntilTarget(X[1])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "_" "??" Target	<< ast.PassUntilTargetOrIgnore(X[2]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.PassUntilTargetOrIgnore(X[2])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : identifier "(" identifier ")" Target	<< ast.TakeUntilTarget(X[0], X[2], X[4]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      16,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.TakeUntilTarget(X[0], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : identifier "(" identifier ")" "??" Target	<< ast.TakeUntilTargetOrRest(X[0], X[2], X[5]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      17,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.TakeUntilTargetOrRest(X[0], X[2], X[5])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "?" identifier "(" Action ")"	<< ast.Option(X[1], X[3]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      18,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Option(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : identifier "(" identifier ")"	<< ast.TakeTheRest(X[0], X[2]) >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      19,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.TakeTheRest(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `AtomicAction : "$"	<< ast.AtTheEnd() >>`,
		Id:         "AtomicAction",
		NTType:     4,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AtTheEnd()
		},
	},
	ProdTabEntry{
		String: `Target : string_lit "[" ":" int_lit "]"	<< ast.LimitedScopeStringTarget(X[0], X[3]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      21,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.LimitedScopeStringTarget(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `Target : string_lit "[" int_lit ":" int_lit "]"	<< ast.BoundedScopeStringTarget(X[0], X[2], X[4]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      22,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.BoundedScopeStringTarget(X[0], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `Target : string_lit "[" int_lit "]"	<< ast.FixedStringTarget(X[0], X[2]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      23,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.FixedStringTarget(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Target : string_lit "[" "]"	<< ast.CloseStringTarget(X[0]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.CloseStringTarget(X[0])
		},
	},
	ProdTabEntry{
		String: `Target : string_lit	<< ast.StringTarget(X[0]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.StringTarget(X[0])
		},
	},
	ProdTabEntry{
		String: `Target : char_lit "[" ":" int_lit "]"	<< ast.LimitedScopeCharTarget(X[0], X[3]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      26,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.LimitedScopeCharTarget(X[0], X[3])
		},
	},
	ProdTabEntry{
		String: `Target : char_lit "[" int_lit ":" int_lit "]"	<< ast.BoundedScopeCharTarget(X[0], X[2], X[4]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      27,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.BoundedScopeCharTarget(X[0], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `Target : char_lit "[" int_lit "]"	<< ast.FixedCharTarget(X[0], X[2]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      28,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.FixedCharTarget(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Target : char_lit "[" "]"	<< ast.CloseCharTarget(X[0]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.CloseCharTarget(X[0])
		},
	},
	ProdTabEntry{
		String: `Target : char_lit	<< ast.CharTarget(X[0]) >>`,
		Id:         "Target",
		NTType:     5,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.CharTarget(X[0])
		},
	},
}
