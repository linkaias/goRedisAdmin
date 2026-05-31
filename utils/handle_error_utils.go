package utils

import (
	"fmt"
	"time"
)

const (
	// HandleErrorLevelInfo is the informational severity level.
	HandleErrorLevelInfo = iota
	// HandleErrorLevelWarn is the warning severity level.
	HandleErrorLevelWarn
	// HandleErrorLevelError is the error severity level.
	HandleErrorLevelError
)

const (
	// HandleDescErrorTypeShow prints the error details only.
	HandleDescErrorTypeShow = iota
	// HandleDescErrorTypeLog prints the error details and writes logs.
	HandleDescErrorTypeLog
	// HandleDescErrorTypeSendMsg prints the error details and notifies by message/email.
	HandleDescErrorTypeSendMsg
)

// HandleError is a centralized error handling entry.
//
// Parameters:
//   - err: the original error
//   - level: severity level, one of HandleErrorLevel*
//   - descErrorType: output strategy, one of HandleDescErrorType*
//   - args: optional extra context values
//
// Current behavior:
//   - returns immediately when err is nil
//   - prints timestamped error details to stdout
//
// Placeholder:
//   - different handling branches (log persistence, notifications, etc.)
//     can be added in the marked section below.
func HandleError(err error, level int, descErrorType int, args ...interface{}) {
	if err == nil {
		return
	}
	// Print a minimal unified error line with timestamp and extra context.
	fmt.Printf("%s\t%v%v\n", time.Now().Format("2006-01-02 15:04:05"), err, args)

	// TODO: Handle by error type/severity (log, alerting, etc.).

}
