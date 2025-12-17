// Copyright (C) 2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           __                                            _ 
//  ___ _ __ ___ | |_ _ __   / /__ ___  _ __ ___  _ __ ___   __ _ _ __   __| |
// / __| '_ ` _ \| __| '_ \ / / __/ _ \| '_ ` _ \| '_ ` _ \ / _` | '_ \ / _` |
// \__ \ | | | | | |_| |_) / / (_| (_) | | | | | | | | | | | (_| | | | | (_| |
// |___/_| |_| |_|\__| .__/_/ \___\___/|_| |_| |_|_| |_| |_|\__,_|_| |_|\__,_|
//                   |_|                                                      

package command
import "strings"
import "slices"
import "libsisimai.org/mailer-goemon/moji"

// Find returns an SMTP command found in the argument.
//   Arguments:
//     - text (string): Text including SMTP command.
//   Returns:
//     - (string): Found SMTP command.
func Find(text string) string {
	if Test(text) == false { return "" }

	commandset := make([]string, 0, 4)
	commandmap := map[string]string{"STAR": CeTTLS, "XFOR": CeXFWD}
	issuedcode := " " + text + " "

	for _, e := range detectable {
		// Find an SMTP command from the given string
		p0 := strings.Index(text, e); if p0 < 0 { continue }
		if strings.IndexByte(e, ' ') < 0 {
			// For example, "RCPT T" does not appear in an email address or a domain name
			cx, cw := true, len(e) + 1
			ca, cz := []byte(issuedcode[p0:p0 + 1])[0], []byte(issuedcode[p0 + cw:p0 + cw + 1])[0]
			switch {
				// Exclude an SMTP command in the part of an email address, a domain name, such as
				// DATABASE@EXAMPLE.JP, EMAIL.EXAMPLE.COM, and so on.
				case ca > 47 && ca <  58 || cz > 47 && cz <  58: // 0-9
				case ca > 63 && ca <  91 || cz > 63 && cz <  91: // @-Z
				case ca > 96 && ca < 123 || cz > 96 && cz < 123: // `-z
				default: cx = false
			}
			if cx == true { continue }
		}
		smtpc := e[0:4] // The first 4 characters of SMTP command found in the argument

		if moji.HasPrefixAny(smtpc, commandset) { continue }
		if slices.Contains([]string{"STAR", "XFOR"}, smtpc) { smtpc = commandmap[smtpc] }
		commandset = append(commandset, smtpc)
	}
	if len(commandset) == 0 { return "" }
	return commandset[len(commandset)-1]
}

