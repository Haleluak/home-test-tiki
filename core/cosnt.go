package core

const (
	// RegexpFilename for valid csv file name
	RegexpFilename = "[a-z0-9.-_].csv"
	// DefaultPageSize is the default value for 'x' param
	DefaultPageSize = 1000
	// DefaultOverwriteOutput is the default value for the 'o' param
	DefaultOverwriteOutput = false
	// DefaultDisplayclock is the default value for the 'c' param
	DefaultDisplayclock = false

	// PhoneNumberLength used to parse line details
	PhoneNumberLength int = 10
	// PhoneDetailsLength used to parse line details
	PhoneDetailsLength int = 21
	// DateLength used to parse line details
	DateLength = 10
	// TimeFormat used cast string to time
	TimeFormat = "2006-02-01"
	// LineFullSize used to check line details
	LineFullSize = 32

	// CsvHeader is the first line of the input file
	CsvHeader = "PHONE_NUMBER,ACTIVATION_DATE,DEACTIVATION_DATE\n"
	// UndefinedDeactivateDate is the default value for empty deactivated date
	UndefinedDeactivateDate = "9999-99-99"

	// StartActivateDate is the pos where ActivateDate starts
	StartActivateDate = 11
	// StartDeactivateDate is the pos where deactivateDate starts
	StartDeactivateDate = 22
)
