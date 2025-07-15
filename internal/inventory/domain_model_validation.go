package inventory

import (
	"regexp"

	"github.com/ehharvey/malleus/internal/validation"
)

// initialize check functions here!
var createDomainModelCheckFunctions = [...]validation.ModelValidationFunction[*CreateDomainParams]{
	checkValidDomainNameFormat,
	checkValidDomainNameLength,
}

// --

var domainRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

var checkValidDomainNameFormatMessages = map[bool]string{
	true:  "check passed",
	false: "domain has incorrect format",
}

func checkValidDomainNameFormat(createDomainParams *CreateDomainParams) validation.ModelValidationCheckResult {
	check := domainRegex.MatchString(createDomainParams.Name)

	return validation.ModelValidationCheckResult{
		Name:      "checkValidDomainNameFormat",
		Succeeded: check,
		Field:     "Name",
		Value:     createDomainParams.Name,
		Message:   checkValidDomainNameFormatMessages[check],
	}
}

type checkValidDomainNameLengthResultCode uint

const (
	checkValidDomainNameLengthInvalidResult = iota
	checkValidDomainNameLengthTooLong
	checkValidDomainNameLengthTooShort
	checkValidDomainNameLengthCorrectLength
)

var checkValidDomainNameLengthMessages = map[checkValidDomainNameLengthResultCode]string{
	checkValidDomainNameLengthInvalidResult: "check failed. This is likely a bug!",
	checkValidDomainNameLengthTooLong:       "domain is too long",
	checkValidDomainNameLengthTooShort:      "domain is too short",
	checkValidDomainNameLengthCorrectLength: "check passed",
}

func checkValidDomainNameLength(createDomainParams *CreateDomainParams) validation.ModelValidationCheckResult {
	check_long := len(createDomainParams.Name) > 253
	check_short := len(createDomainParams.Name) <= 0
	var check_code checkValidDomainNameLengthResultCode = checkValidDomainNameLengthInvalidResult

	if check_long {
		check_code = checkValidDomainNameLengthTooLong
	} else if check_short {
		check_code = checkValidDomainNameLengthTooShort
	} else {
		check_code = checkValidDomainNameLengthCorrectLength
	}

	return validation.ModelValidationCheckResult{
		Name:      "checkValidDomainNameLength",
		Succeeded: check_code == checkValidDomainNameLengthCorrectLength,
		Field:     "Name",
		Value:     createDomainParams.Name,
		Message:   checkValidDomainNameLengthMessages[check_code],
	}
}
