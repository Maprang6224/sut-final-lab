package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNeagativeName(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "", //wrong
		Email:      "Maprang@gmail.com",
		CustomerID: "L1234567",
	}
	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("กรุณากรอกชื่อ"))

}
