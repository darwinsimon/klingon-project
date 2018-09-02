package stapi

// RESTURL of Stapi.co
const RESTURL = "http://stapi.co/api/v1/rest"

// If the result is bigger than tolerance, return ErrorTooManyResults
const maxToleranceResult = 10

// Types of errors for STAPI module
const (
	ErrorNone Error = iota
	ErrorTooManyResults
	ErrorCharacterNotFound
)
