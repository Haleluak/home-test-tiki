package parser

import (
	"bufio"
	"errors"
	"fmt"
	"home-test-tiki/core"
	"home-test-tiki/util"
	"os"
)

// updateIfLastAction will update existing lines based on the situation
func updateIfLastAction(w *util.Writer, fd *os.File, lineIndex int64, oldLine string, newLine string) error {

	// Parse our line
	old := NewPhoneDetails(oldLine)
	new := NewPhoneDetails(newLine)
	if old == nil || new == nil {
		core.PrintAndExit(errors.New(fmt.Sprintf("Input not valid : %s", newLine)))
	}

	if isMoreRecentOwner(old, new) {
		// If it's a more recent owner we overwrite our existing line
		w.UpdateLine(fd, lineIndex, newLine)
		return nil
	}
	if isPlanUpdate(old, new) {
		// If owner simply updated it's plan, we update deactivate date of our existing line
		w.UpdateDeactivateDate(fd, lineIndex, new.DeactivationDate)
		return nil
	}
	if isPreviousPlanUpdate(old, new) {
		// If owner previously updated it's plan, we update activate date of our existing line
		w.UpdateActivateDate(fd, lineIndex, new.ActivationDate)
		return nil
	}
	return nil
}
// return true if newLine.activateDate > oldLine.activateDate && newLine.DeactivateDate != new.ActivateDate
func isMoreRecentOwner(old *PhoneNumberInfo, new *PhoneNumberInfo) bool {
	newActivationDate := *(new.getActivationDate())
	oldActivationDate := *(old.getActivationDate())
	if newActivationDate.After(oldActivationDate) && old.DeactivationDate != new.ActivationDate {
		return true
	}
	return false
}

// return true if newLine.activateDate > oldLine.activateDate && oldline.deactivateDate = new.activateDate
func isPlanUpdate(old *PhoneNumberInfo, new *PhoneNumberInfo) bool {
	newActivationDate := *(new.getActivationDate())
	oldActivationDate := *(old.getActivationDate())

	if newActivationDate.After(oldActivationDate) && old.DeactivationDate == new.ActivationDate {
		return true
	}
	return false
}

// return if newLine.activateDate < oldline.activateDate && newline.deactivateDate = oldLine.activateDate
func isPreviousPlanUpdate(old *PhoneNumberInfo, new *PhoneNumberInfo) bool {
	newActivationDate := *(new.getActivationDate())
	oldActivationDate := *(old.getActivationDate())

	if newActivationDate.Before(oldActivationDate) && new.DeactivationDate == old.ActivationDate {
		return true
	}
	return false
}

// indexOf will read the file and return the index of the phone number if existing
// Else -1 is return
func indexOf(fd *os.File, phoneNumber string) (string, int64, error) {
	fd.Seek(0, 0)
	reader := bufio.NewReader(fd)
	var line string
	var cursorPos int64 = 0
	var err error
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		if isDuplicate(line, phoneNumber) {
			return line, cursorPos, nil
		}
		cursorPos += int64(len(line))
	}
	return "", -1, nil
}

// Compare a line and a phonenumber to see if it match
func isDuplicate(line string, phoneNumber string) bool {
	return getNumber(line) == phoneNumber
}
