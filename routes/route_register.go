package routes

import (
	"regexp"
)

// RequestType defines the type of request
type RequestType int

// Request Type enum
const (
	JSON RequestType = iota // JSON reuest type
	XML                     //XML request type
)

// RouteRegistration contains a registered route
type RouteRegistration struct {
	HTTPMethod string
	RouteRegex *regexp.Regexp
	Response   string
	Type       RequestType
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
func (rr *RouteRegister) Register(method string, requestType RequestType, routeRegex string, response string) error {

	r, err := regexp.Compile(routeRegex)
	if err != nil {
		return err
	}

	rr.routes = append(rr.routes, RouteRegistration{
		HTTPMethod: method,
		Response:   response,
		Type:       requestType,
		RouteRegex: r,
	})
	return nil
}

// Clear the register
func (rr *RouteRegister) Clear(r RouteRegistration) {
	rr.routes = nil
}

// Match a registration
func (rr *RouteRegister) Match(method string, requestType RequestType, route string) (bool, string) {

	for _, r := range rr.routes {

		if r.HTTPMethod == method || r.Type == requestType || r.RouteRegex.MatchString(route) {
			return true, r.Response
		}
	}

	return false, ""
}
