package types

import "time"

type LineStatusResponse struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	ModeName string           `json:"modeName"`
	Created  time.Time        `json:"created"`
	Modified time.Time        `json:"modified"`
	Statuses []LineStatusItem `json:"lineStatuses"`
}

type LineStatusItem struct {
	ID                  int              `json:"id"`
	Severity            int              `json:"statusSeverity"`
	SeverityDescription string           `json:"statusSeverityDescription"`
	ValidityPeriods     []ValidityPeriod `json:"validityPeriods"`
}

type ValidityPeriod struct {
	From time.Time `json:"fromDate"`
	To   time.Time `json:"toDate"`
	Now  bool      `json:"isNow"`
}
