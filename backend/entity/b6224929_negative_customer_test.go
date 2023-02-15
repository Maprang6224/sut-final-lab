package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNeagativeCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	customers := []string{
		"A123456",
		"L1245",
		"M1455",
		"H14",
	}
	for _, customer := range customers {
		c := Customer{
			Name:       "Maprang", //wrong
			Email:      "Maprang@gmail.com",
			CustomerID: customer,
		}
		ok, err := govalidator.ValidateStruct(c)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("CustermerID not match"))

	}

}
