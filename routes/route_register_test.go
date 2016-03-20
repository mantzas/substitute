package routes

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RouteRegister", func() {

	It("can register", func() {

		Register.Clear()

		err := Register.Register(http.MethodPost, JSON, "", "test respnse", http.StatusOK)

		Expect(err).NotTo(HaveOccurred())
		Expect(len(Register.routes)).To(Equal(1))
	})

	It("can clear", func() {

		Register.Clear()
		Register.Register(http.MethodPost, JSON, "", "test respnse", http.StatusOK)
		Expect(len(Register.routes)).To(Equal(1))

		Register.Clear()
		Expect(len(Register.routes)).To(Equal(0))
	})

	It("can match", func() {

		Register.Clear()
		Register.Register(http.MethodGet, JSON, `\/users$`, "users response", http.StatusOK)
		Register.Register(http.MethodGet, JSON, `\/users\/\d`, "specific user response", http.StatusOK)
		Register.Register(http.MethodPost, JSON, `\/users$`, "test response", http.StatusCreated)
		Expect(len(Register.routes)).To(Equal(3))

		matched, content, responseStatus := Register.Match(http.MethodGet, JSON, "/users/1")
		Expect(matched).To(BeTrue())
		Expect(content).To(Equal("specific user response"))
		Expect(responseStatus).To(Equal(http.StatusOK))
	})

	It("cannot match", func() {
		Register.Clear()

		matched, _, _ := Register.Match(http.MethodGet, JSON, "/users/1")
		Expect(matched).To(BeFalse())
	})
})
