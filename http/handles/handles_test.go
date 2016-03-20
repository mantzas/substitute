package handles

import (
	"net/http"
	"net/http/httptest"

	"github.com/mantzas/substitute/routes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("handles", func() {

	It("any handle with no content type", func() {

		request, _ := http.NewRequest(http.MethodGet, "/tests", nil)

		response := httptest.NewRecorder()

		AnyHandle(response, request, nil)

		Expect(response.Code).To(Equal(http.StatusInternalServerError))
		Expect(response.Body.String()).To(Equal("Content type not supported!"))
	})

	It("any handle route not matched", func() {

		request, _ := http.NewRequest(http.MethodGet, "/tests", nil)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		AnyHandle(response, request, nil)

		Expect(response.Body.String()).To(Equal("Route not matched!"))
		Expect(response.Code).To(Equal(http.StatusInternalServerError))
	})

	It("any handle route matched", func() {

		routes.Register.Register(http.MethodGet, routes.JSON, "/tests", "hello world!", http.StatusCreated)

		request, _ := http.NewRequest(http.MethodGet, "/tests", nil)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		AnyHandle(response, request, nil)

		Expect(response.Code).To(Equal(http.StatusCreated))
		Expect(response.Body.String()).To(Equal("hello world!"))
	})
})
