// Copyright (C) 2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           __                                            _ 
//  ___ _ __ ___ | |_ _ __   / /__ ___  _ __ ___  _ __ ___   __ _ _ __   __| |
// / __| '_ ` _ \| __| '_ \ / / __/ _ \| '_ ` _ \| '_ ` _ \ / _` | '_ \ / _` |
// \__ \ | | | | | |_| |_) / / (_| (_) | | | | | | | | | | | (_| | | | | (_| |
// |___/_| |_| |_|\__| .__/_/ \___\___/|_| |_| |_|_| |_| |_|\__,_|_| |_|\__,_|
//                   |_|                                                      

// Package "smtp/command" provides functions related to SMTP commands.
package command
import "strings"
import "libsisimai.org/mailer-goemon/moji"

const (
	CeHELO = "HELO"
	CeEHLO = "EHLO"
	CeMAIL = "MAIL"
	CeRCPT = "RCPT"
	CeDATA = "DATA"
	CeQUIT = "QUIT"
	CeRSET = "RSET"
	CeNOOP = "NOOP"
	CeVRFY = "VRFY"
	CeETRN = "ETRN"
	CeEXPN = "EXPN"
	CeHELP = "HELP"
	CeAUTH = "AUTH"
	CeTTLS = "STARTTLS"
	CeXFWD = "XFORWARD"
)

var availables = []string{
	CeHELO, CeEHLO, CeMAIL, CeRCPT, CeDATA, CeQUIT, CeRSET, CeNOOP, CeVRFY, CeETRN, CeEXPN, CeHELP,
	CeAUTH, CeTTLS, CeXFWD,
}
var detectable = []string{
	CeHELO, CeEHLO, CeTTLS, CeAUTH + " PLAIN", CeAUTH + " LOGIN", CeAUTH + " CRAM-", CeAUTH + " DIGEST-",
	CeMAIL + " F", CeRCPT, CeRCPT + " T", CeDATA, CeQUIT, CeXFWD,
}
var ExceptDATA = []string{CeEHLO, CeHELO, CeMAIL, CeRCPT}

// Test checks that an SMTP command in the argument is valid or not.
//   Arguments:
//     - comm (string): An SMTP command.
//   Returns:
//     - (bool): true if the argument is a valid SMTP command.
func Test(comm string) bool {
	if len(comm) < 4                                       { return false }
	if moji.ContainsAny(strings.ToUpper(comm), availables) { return true  }
	return false
}

