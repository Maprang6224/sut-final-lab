package entity

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "Maprang",
		Email:      "MAprang@gmail.com",
		CustomerID: "L1234567",
	}
	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

	fmt.Println(err)
}
