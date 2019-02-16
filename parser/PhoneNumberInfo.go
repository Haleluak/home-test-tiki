package parser

import (
	"errors"
	"fmt"
	"github.com/Haleluak/home-test-tiki/core"
	"strings"
	"time"
)

type PhoneNumberInfo struct {
	PhoneNumber      string
	ActivationDate   string
	DeactivationDate string
	activationTime   *time.Time
	deactivationTime *time.Time
}

func NewPhoneDetails(line string) *PhoneNumberInfo {
	details := strings.Split(line, ",")

	if len(details) != 3 {
		// Data are not valid
		core.PrintAndExit(errors.New(fmt.Sprintf("Invalid line : %s\n", line)))
	}
	/// [0] = phone number
	/// [1] = ACTIVATION_DATE
	/// [2] = DEACTIVATION_DATE
	if details[2] == "" {
		// No deactivate date? we use default value
		details[2] = core.UndefinedDeactivateDate
	}
	return &PhoneNumberInfo{details[0][0:core.PhoneNumberLength], details[1][0:core.DateLength], details[2][0:core.DateLength], nil, nil}
}

// getActivateDate will return activationDate as time.Time
func (pd *PhoneNumberInfo) getActivationDate() *time.Time {
	if pd.activationTime == nil {
		t, e := time.Parse(core.TimeFormat, pd.ActivationDate)
		if e != nil {
			core.PrintAndExit(e)
		}
		pd.activationTime = &t
	}
	return pd.activationTime
}

// getDeactivationDate will return deactivationDate as time.Time or nil if not existing
func (pd *PhoneNumberInfo) getDeactivationDate() *time.Time {
	if pd.DeactivationDate == core.UndefinedDeactivateDate {
		return nil
	}
	if pd.deactivationTime == nil {
		t, e := time.Parse(core.TimeFormat, pd.DeactivationDate)
		if e != nil {
			core.PrintAndExit(e)
		}
		pd.deactivationTime = &t
	}
	return pd.deactivationTime
}