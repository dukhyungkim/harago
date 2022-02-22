package common

import "time"

const (
	DefaultTimeout = 5 * time.Second

	SharedActionSubject          = "harago.shared.action"
	CompanyActionSubject         = "harago.company.action"
	SpecificCompanyActionSubject = "harago.%s.action"
)
