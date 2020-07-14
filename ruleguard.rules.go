// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

// This is a collection of rules for ruleguard: https://github.com/quasilyte/go-ruleguard

// Remove extra conversions: mdempsky/unconvert
func unconvert(m fluent.Matcher) {
	m.Match("int($x)").Where(m["x"].Type.Is("int") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	m.Match("float32($x)").Where(m["x"].Type.Is("float32") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("float64($x)").Where(m["x"].Type.Is("float64") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	// m.Match("byte($x)").Where(m["x"].Type.Is("byte")).Report("unnecessary conversion").Suggest("$x")
	// m.Match("rune($x)").Where(m["x"].Type.Is("rune")).Report("unnecessary conversion").Suggest("$x")
	m.Match("bool($x)").Where(m["x"].Type.Is("bool") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	m.Match("int8($x)").Where(m["x"].Type.Is("int8") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("int16($x)").Where(m["x"].Type.Is("int16") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("int32($x)").Where(m["x"].Type.Is("int32") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("int64($x)").Where(m["x"].Type.Is("int64") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")

	m.Match("uint8($x)").Where(m["x"].Type.Is("uint8") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("uint16($x)").Where(m["x"].Type.Is("uint16") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("uint32($x)").Where(m["x"].Type.Is("uint32") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
	m.Match("uint64($x)").Where(m["x"].Type.Is("uint64") && !m["x"].Const).Report("unnecessary conversion").Suggest("$x")
}

// Don't use == or != with time.Time
// https://github.com/dominikh/go-tools/issues/47 : Wontfix
func timeeq(m fluent.Matcher) {
	m.Match("$t0 == $t1").Where(m["t0"].Type.Is("time.Time")).Report("using == with time.Time")
	m.Match("$t0 != $t1").Where(m["t0"].Type.Is("time.Time")).Report("using != with time.Time")
}

// Wrong err in error check
func wrongerr(m fluent.Matcher) {
	m.Match("if $*_, $err0 := $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 := $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 := $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 := $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("if $*_, $err0 = $*_; $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 := $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 != nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text == "err" && m["err0"].Type.Is("error") && m["err1"].Text != "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")

	m.Match("$*_, $err0 = $*_; if $err1 == nil { $*_ }").
		Where(m["err0"].Text != "err" && m["err0"].Type.Is("error") && m["err1"].Text == "err" && m["err1"].Type.Is("error")).
		Report("maybe wrong err in error check")
}

// err but no an error
func errnoterror(m fluent.Matcher) {

	// Would be easier to check for all err identifiers instead, but then how do we get the type from m[] ?

	m.Match(
		"if $*_, err := $x; $err != nil { $*_ } else if $_ { $*_ }",
		"if $*_, err := $x; $err != nil { $*_ } else { $*_ }",
		"if $*_, err := $x; $err != nil { $*_ }",

		"if $*_, err = $x; $err != nil { $*_ } else if $_ { $*_ }",
		"if $*_, err = $x; $err != nil { $*_ } else { $*_ }",
		"if $*_, err = $x; $err != nil { $*_ }",

		"$*_, err := $x; if $err != nil { $*_ } else if $_ { $*_ }",
		"$*_, err := $x; if $err != nil { $*_ } else { $*_ }",
		"$*_, err := $x; if $err != nil { $*_ }",

		"$*_, err = $x; if $err != nil { $*_ } else if $_ { $*_ }",
		"$*_, err = $x; if $err != nil { $*_ } else { $*_ }",
		"$*_, err = $x; if $err != nil { $*_ }",
	).
		Where(m["err"].Text == "err" && !m["err"].Type.Is("error") && m["x"].Text != "recover()").
		Report("err variable not error type")
}

// Identical if and else bodies
func ifbodythenbody(m fluent.Matcher) {
	m.Match("if $*_ { $body } else { $body }").
		Report("identical if and else bodies")

	// Lots of false positives.
	// m.Match("if $*_ { $body } else if $*_ { $body }").
	//	Report("identical if and else bodies")
}

// Odd inequality: A - B < 0 instead of !=
// Too many false positives.
/*
func subtractnoteq(m fluent.Matcher) {
	m.Match("$a - $b < 0").Report("consider $a != $b")
	m.Match("$a - $b > 0").Report("consider $a != $b")
	m.Match("0 < $a - $b").Report("consider $a != $b")
	m.Match("0 > $a - $b").Report("consider $a != $b")
}
*/

// Self-assignment
func selfassign(m fluent.Matcher) {
	m.Match("$x = $x").Report("useless self-assignment")
}

// Odd nested ifs
func oddnestedif(m fluent.Matcher) {
	m.Match("if $x { if $x { $*_ }; $*_ }",
		"if $x == $y { if $x != $y {$*_ }; $*_ }",
		"if $x != $y { if $x == $y {$*_ }; $*_ }",
		"if $x { if !$x { $*_ }; $*_ }",
		"if !$x { if $x { $*_ }; $*_ }").
		Report("odd nested ifs")

	m.Match("for $x { if $x { $*_ }; $*_ }",
		"for $x == $y { if $x != $y {$*_ }; $*_ }",
		"for $x != $y { if $x == $y {$*_ }; $*_ }",
		"for $x { if !$x { $*_ }; $*_ }",
		"for !$x { if $x { $*_ }; $*_ }").
		Report("odd nested for/ifs")
}

// odd bitwise expressions
func oddbitwise(m fluent.Matcher) {
	m.Match("$x | $x",
		"$x | ^$x",
		"^$x | $x").
		Report("odd bitwise OR")

	m.Match("$x & $x",
		"$x & ^$x",
		"^$x & $x").
		Report("odd bitwise AND")
}

// odd sequence of if tests with return
func ifreturn(m fluent.Matcher) {
	m.Match("if $x { return $*_ }; if $x {$*_ }").Report("odd sequence of if test")
	m.Match("if $x { return $*_ }; if !$x {$*_ }").Report("odd sequence of if test")
	m.Match("if !$x { return $*_ }; if $x {$*_ }").Report("odd sequence of if test")
	m.Match("if $x == $y { return $*_ }; if $x != $y {$*_ }").Report("odd sequence of if test")
	m.Match("if $x != $y { return $*_ }; if $x == $y {$*_ }").Report("odd sequence of if test")

}

func oddifsequence(m fluent.Matcher) {
	/*
		m.Match("if $x { $*_ }; if $x {$*_ }").Report("odd sequence of if test")

		m.Match("if $x == $y { $*_ }; if $y == $x {$*_ }").Report("odd sequence of if tests")
		m.Match("if $x != $y { $*_ }; if $y != $x {$*_ }").Report("odd sequence of if tests")

		m.Match("if $x < $y { $*_ }; if $y > $x {$*_ }").Report("odd sequence of if tests")
		m.Match("if $x <= $y { $*_ }; if $y >= $x {$*_ }").Report("odd sequence of if tests")

		m.Match("if $x > $y { $*_ }; if $y < $x {$*_ }").Report("odd sequence of if tests")
		m.Match("if $x >= $y { $*_ }; if $y <= $x {$*_ }").Report("odd sequence of if tests")
	*/
}

// odd sequence of nested if tests
func nestedifsequence(m fluent.Matcher) {
	/*
		m.Match("if $x < $y { if $x >= $y {$*_ }; $*_ }").Report("odd sequence of nested if tests")
		m.Match("if $x <= $y { if $x > $y {$*_ }; $*_ }").Report("odd sequence of nested if tests")
		m.Match("if $x > $y { if $x <= $y {$*_ }; $*_ }").Report("odd sequence of nested if tests")
		m.Match("if $x >= $y { if $x < $y {$*_ }; $*_ }").Report("odd sequence of nested if tests")
	*/
}

// odd sequence of assignments
func identicalassignments(m fluent.Matcher) {
	m.Match("$x  = $y; $y = $x").Report("odd sequence of assignments")
}

func oddcompoundop(m fluent.Matcher) {
	m.Match("$x += $x + $_",
		"$x += $x - $_").
		Report("odd += expression")

	m.Match("$x -= $x + $_",
		"$x -= $x - $_").
		Report("odd -= expression")
}

func constswitch(m fluent.Matcher) {
	m.Match("switch $x { $*_ }").
		Where(m["x"].Const && !m["x"].Text.Matches(`^runtime\.`)).
		Report("constant switch")
}
