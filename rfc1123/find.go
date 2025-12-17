// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _ _ ____  _____ 
// |  _ \|  ___/ ___/ / |___ \|___ / 
// | |_) | |_ | |   | | | __) | |_ \ 
// |  _ <|  _|| |___| | |/ __/ ___) |
// |_| \_\_|   \____|_|_|_____|____/ 

package rfc1123
import "strings"
import "libsisimai.org/mailer-goemon/moji"

var prefix0x32 = []string{"(", "[", "<"}
var suffix0x32 = []string{")", "]", ">", ":", ";"}

// Find returns a valid internet hostname found from the argument.
//   Arguments:
//     - text (string): String including hostnames.
//   Returns:
//     - (string): Valid internet hostname found in the argument.
func Find(text string) string {
	if text == "" { return "" }

	// Replace some string for splitting by " "
	// - mx.example.net[192.0.2.1] => mx.example.net [192.0.2.1]
	// - mx.example.jp:[192.0.2.1] => mx.example.jp :[192.0.2.1]
	sourcetext := strings.ToLower(text)
	for _, e := range prefix0x32 { sourcetext = strings.ReplaceAll(sourcetext, e, " " + e) }
	for _, e := range suffix0x32 { sourcetext = strings.ReplaceAll(sourcetext, e, e + " ") }

	sourcelist := make([]string, 0, 25)
	foundtoken := make([]string, 0,  2)
	thelongest := uint8(0)
	hostnameis := ""

	MAKELIST: for {
		for _, e := range sandwiched {
			// Check a hostname exists between the e[0] and e[1] at slice "sandwiched"
			// Each slice in Sandwich have 2 elements
			if moji.Aligned(sourcetext, e) == false { continue }
			p1 := strings.Index(sourcetext, e[0])
			p2 := strings.Index(sourcetext, e[1]); cw := len(e[0]); if p1 + cw >= p2 { continue }

			sourcelist = strings.Split(sourcetext[p1 + cw:p2], " ")
			break MAKELIST
		}

		// Check other patterns which are not sandwiched
		for _, e := range startafter {
			// startafter have some strings, not a slice([]string).
			if strings.Contains(sourcetext, e) == false { continue }
			sourcelist = strings.Split(moji.Select(sourcetext + moji.RHS, e, "", 0), " ")
			break MAKELIST
		}

		for _, e := range existuntil {
			// existuntil have some strings, not a slice([]string).
			if strings.Contains(sourcetext, e) == false { continue }
			sourcelist = strings.Split(moji.Select(moji.LHS + sourcetext, "", e, 0), " ")
			break MAKELIST
		}

		if len(sourcelist) == 0 { sourcelist = strings.Split(sourcetext, " ") }
		break MAKELIST
	}

	for _, e := range sourcelist {
		// Pick some strings which have 4 or more length, is including "." character
		e = strings.TrimRight(e, ".") // Remove "." at the end of the string
		for _, f := range prefix0x32 { e = strings.ReplaceAll(e, f, "") }
		for _, f := range suffix0x32 { e = strings.ReplaceAll(e, f, "") }

		if len(e) < 4 || strings.IndexByte(e, '.') < 0 || IsInternetHost(e) == false { continue }
		foundtoken = append(foundtoken, e)
	}
	if len(foundtoken) == 0 { return ""            }
	if len(foundtoken) == 1 { return foundtoken[0] }

	for _, e := range foundtoken {
		// Returns the longest hostname
		if cw := uint8(len(e)); thelongest < cw { hostnameis = e; thelongest = cw }
	}
	return hostnameis
}

