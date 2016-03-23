package routes

import (
	"fmt"
	"regexp"
	"strings"
)

// RequestType defines the type of request
type RequestType int

// Request Type enum
const (
	JSON   RequestType = iota // JSON reuest type
	SOAP                      // SOAP request type
	SOAP12                    // SOAP 1.2 request type
)

// ContentTypeToRequestType map content type to request type
func ContentTypeToRequestType(contentType string) (RequestType, error) {

	if strings.HasPrefix(contentType, "application/json") {

		return JSON, nil
	}

	if strings.HasPrefix(contentType, "text/xml") {

		return SOAP, nil
	}

	if strings.HasPrefix(contentType, "application/soap+xml") {

		return SOAP12, nil
	}

	return -1, fmt.Errorf("Unable to map %s to request type!", contentType)
}

// RouteRegistration contains a registered route
type RouteRegistration struct {
	HTTPMethod string
	RouteRegex *regexp.Regexp
	Request    string
	Response   string
	Type       RequestType
	Status     int
}

// Register contains the route registrations
var Register = RouteRegister{
	routes: make([]RouteRegistration, 0),
}

// RouteRegister contains the registered routes
type RouteRegister struct {
	routes []RouteRegistration
}

// Register a route
func (rr *RouteRegister) Register(method string, requestType RequestType, routeRegex string,
	request string, response string, status int) error {

	r, err := regexp.Compile(routeRegex)
	if err != nil {
		return err
	}

	rr.routes = append(rr.routes, RouteRegistration{
		HTTPMethod: method,
		Response:   response,
		Request:    request,
		Type:       requestType,
		RouteRegex: r,
		Status:     status,
	})
	return nil
}

// Clear the register
func (rr *RouteRegister) Clear() {
	rr.routes = nil
}

// Match a registration
func (rr *RouteRegister) Match(method string, requestType RequestType, route string,
	request string) (bool, string, int) {

	for _, r := range rr.routes {

		if r.HTTPMethod == method && r.Type == requestType && r.RouteRegex.MatchString(route) {
			return true, r.Response, r.Status
		}
	}

	return false, "", 0
}
