package gcode

import "net/http"

// https://rfc-editor.org/rfc/rfc7231.html#section-4.3
const (
	HttpStatusNotSpecified = HttpStatusCode(-1)

	HttpStatusContinue           = HttpStatusCode(http.StatusContinue)
	HttpStatusSwitchingProtocols = HttpStatusCode(http.StatusSwitchingProtocols)
	HttpStatusProcessing         = HttpStatusCode(http.StatusProcessing)
	HttpStatusEarlyHints         = HttpStatusCode(http.StatusEarlyHints)

	HttpStatusOK                   = HttpStatusCode(http.StatusOK)
	HttpStatusCreated              = HttpStatusCode(http.StatusCreated)
	HttpStatusAccepted             = HttpStatusCode(http.StatusAccepted)
	HttpStatusNonAuthoritativeInfo = HttpStatusCode(http.StatusNonAuthoritativeInfo)
	HttpStatusNoContent            = HttpStatusCode(http.StatusNoContent)
	HttpStatusResetContent         = HttpStatusCode(http.StatusResetContent)
	HttpStatusPartialContent       = HttpStatusCode(http.StatusPartialContent)
	HttpStatusMultiStatus          = HttpStatusCode(http.StatusMultiStatus)
	HttpStatusAlreadyReported      = HttpStatusCode(http.StatusAlreadyReported)
	HttpStatusIMUsed               = HttpStatusCode(http.StatusIMUsed)

	HttpStatusMultipleChoices   = HttpStatusCode(http.StatusMultipleChoices)
	HttpStatusMovedPermanently  = HttpStatusCode(http.StatusMovedPermanently)
	HttpStatusFound             = HttpStatusCode(http.StatusFound)
	HttpStatusSeeOther          = HttpStatusCode(http.StatusSeeOther)
	HttpStatusNotModified       = HttpStatusCode(http.StatusNotModified)
	HttpStatusUseProxy          = HttpStatusCode(http.StatusUseProxy)
	HttpStatusTemporaryRedirect = HttpStatusCode(http.StatusTemporaryRedirect)
	HttpStatusPermanentRedirect = HttpStatusCode(http.StatusPermanentRedirect)

	HttpStatusBadRequest                   = HttpStatusCode(http.StatusBadRequest)
	HttpStatusUnauthorized                 = HttpStatusCode(http.StatusUnauthorized)
	HttpStatusPaymentRequired              = HttpStatusCode(http.StatusPaymentRequired)
	HttpStatusForbidden                    = HttpStatusCode(http.StatusForbidden)
	HttpStatusNotFound                     = HttpStatusCode(http.StatusNotFound)
	HttpStatusMethodNotAllowed             = HttpStatusCode(http.StatusMethodNotAllowed)
	HttpStatusNotAcceptable                = HttpStatusCode(http.StatusNotAcceptable)
	HttpStatusProxyAuthRequired            = HttpStatusCode(http.StatusProxyAuthRequired)
	HttpStatusRequestTimeout               = HttpStatusCode(http.StatusRequestTimeout)
	HttpStatusConflict                     = HttpStatusCode(http.StatusConflict)
	HttpStatusGone                         = HttpStatusCode(http.StatusGone)
	HttpStatusLengthRequired               = HttpStatusCode(http.StatusLengthRequired)
	HttpStatusPreconditionFailed           = HttpStatusCode(http.StatusPreconditionFailed)
	HttpStatusRequestEntityTooLarge        = HttpStatusCode(http.StatusRequestEntityTooLarge)
	HttpStatusRequestURITooLong            = HttpStatusCode(http.StatusRequestURITooLong)
	HttpStatusUnsupportedMediaType         = HttpStatusCode(http.StatusUnsupportedMediaType)
	HttpStatusRequestedRangeNotSatisfiable = HttpStatusCode(http.StatusRequestedRangeNotSatisfiable)
	HttpStatusExpectationFailed            = HttpStatusCode(http.StatusExpectationFailed)
	HttpStatusTeapot                       = HttpStatusCode(http.StatusTeapot)
	HttpStatusMisdirectedRequest           = HttpStatusCode(http.StatusMisdirectedRequest)
	HttpStatusUnprocessableEntity          = HttpStatusCode(http.StatusUnprocessableEntity)
	HttpStatusLocked                       = HttpStatusCode(http.StatusLocked)
	HttpStatusFailedDependency             = HttpStatusCode(http.StatusFailedDependency)
	HttpStatusTooEarly                     = HttpStatusCode(http.StatusTooEarly)
	HttpStatusUpgradeRequired              = HttpStatusCode(http.StatusUpgradeRequired)
	HttpStatusPreconditionRequired         = HttpStatusCode(http.StatusPreconditionRequired)
	HttpStatusTooManyRequests              = HttpStatusCode(http.StatusTooManyRequests)
	HttpStatusRequestHeaderFieldsTooLarge  = HttpStatusCode(http.StatusRequestHeaderFieldsTooLarge)
	HttpStatusUnavailableForLegalReasons   = HttpStatusCode(http.StatusUnavailableForLegalReasons)

	HttpStatusInternalServerError           = HttpStatusCode(http.StatusInternalServerError)
	HttpStatusNotImplemented                = HttpStatusCode(http.StatusNotImplemented)
	HttpStatusBadGateway                    = HttpStatusCode(http.StatusBadGateway)
	HttpStatusServiceUnavailable            = HttpStatusCode(http.StatusServiceUnavailable)
	HttpStatusGatewayTimeout                = HttpStatusCode(http.StatusGatewayTimeout)
	HttpStatusHTTPVersionNotSupported       = HttpStatusCode(http.StatusHTTPVersionNotSupported)
	HttpStatusVariantAlsoNegotiates         = HttpStatusCode(http.StatusVariantAlsoNegotiates)
	HttpStatusInsufficientStorage           = HttpStatusCode(http.StatusInsufficientStorage)
	HttpStatusLoopDetected                  = HttpStatusCode(http.StatusLoopDetected)
	HttpStatusNotExtended                   = HttpStatusCode(http.StatusNotExtended)
	HttpStatusNetworkAuthenticationRequired = HttpStatusCode(http.StatusNetworkAuthenticationRequired)
)

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code HttpStatusCode) string {
	return http.StatusText(int(code))
}
