// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           __              _       
//  ___ _ __ ___ | |_ _ __   / / __ ___ _ __ | |_   _ 
// / __| '_ ` _ \| __| '_ \ / / '__/ _ \ '_ \| | | | |
// \__ \ | | | | | |_| |_) / /| | |  __/ |_) | | |_| |
// |___/_| |_| |_|\__| .__/_/ |_|  \___| .__/|_|\__, |
//                   |_|               |_|      |___/ 

package reply
import "strings"
import "libsisimai.org/mailer-goemon/moji"

// Find returns an SMTP reply code found from the given string.
//   Arguments:
//     - logs (string): String including SMTP reply code like 550.
//     - hint (string): SMTP status code like "5.1.1", or the 1st digit of the code like "2", "4", or "5".
//   Returns:
//     - (string): SMTP reply code found in the 1st argument.
func Find(logs, hint string) string {
	if len(logs) < 3 || strings.Contains(strings.ToUpper(logs), "X-UNIX") { return "" }
	if len(hint) == 0 { hint = "0" }

	esmtperror := " " + logs + " "
	replycodes := make([]string, 0, 50)
	if statuscode := hint[0:1]; statuscode == "2" || statuscode == "4" || statuscode == "5" {
		// The first character of the 2nd argument is 2 or 4 or 5
		replycodes = codeofsmtp[statuscode]

	} else {
		// The first character of the 2nd argument is 0 or other values
		replycodes = append(replycodes, codeofsmtp["5"]...)
		replycodes = append(replycodes, codeofsmtp["4"]...)
		replycodes = append(replycodes, codeofsmtp["2"]...)
	}

	for _, e := range replycodes {
		// Try to find an SMTP Reply Code from the given string
		appearance := strings.Count(esmtperror, e); if appearance == 0 { continue }
		startingat := 1

		for j := 0; j < appearance; j++ {
			// Find all the reply codes in the error message
			replyindex := moji.IndexOnTheWay(esmtperror, e, startingat); if replyindex < 0 { break }
			formerchar := []byte(esmtperror[replyindex - 1:replyindex])[0]
			latterchar := []byte(esmtperror[replyindex + 3:replyindex + 4])[0]

			if formerchar > 45 && formerchar < 58 { startingat += replyindex + 3; continue } // '.' => '9'
			if latterchar > 45 && latterchar < 58 { startingat += replyindex + 3; continue } // '.' => '9'
			return e
		}
	}
	return ""
}

