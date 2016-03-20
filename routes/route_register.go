package routes

import (
	"fmt"
	"regexp"
)

// RequestType defines the type of request
type RequestType int

// Request Type enum
const (
	JSON RequestType = iota // JSON reuest type
	XML                     //XML request type
)

// ContentTypeToRequestType map content type to request type
func ContentTypeToRequestType(contentType string) (RequestType, error) {

	switch contentType {
	case "application/json":
		return JSON, nil
	case "application/xml":
		return XML, nil
	default:
		return -1, fmt.Errorf("Unable to map %s to request type!", contentType)
	}
}

// RouteRegistration contains a registered route
type RouteRegistration struct {
	HTTPMethod string
	RouteRegex *regexp.Regexp
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
func (rr *RouteRegister) Register(method string, requestType RequestType, routeRegex string, response string, status int) error {

	r, err := regexp.Compile(routeRegex)
	if err != nil {
		return err
	}

	rr.routes = append(rr.routes, RouteRegistration{
		HTTPMethod: method,
		Response:   response,
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
func (rr *RouteRegister) Match(method string, requestType RequestType, route string) (bool, string, int) {

	for _, r := range rr.routes {

		if r.HTTPMethod == method || r.Type == requestType || r.RouteRegex.MatchString(route) {
			return true, r.Response, r.Status
		}
	}

	return false, "", 0
}
