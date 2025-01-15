package unit

import (
	"testing"
	"week05/entity"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestMember(t *testing.T) {
	g := NewGomegaWithT(t)

	// t.Run(`valid case`, func(t *testing.T) {
	// 	member := entity.Member{
	// 		Username: "test",
	// 		Password: "12",
	// 		Email:    "B6526542@gmail.com",
	// 	}
	// 	ok, err := govalidator.ValidateStruct(member)
	// 	g.Expect(ok).To(BeTrue())
	// 	g.Expect(err).To(BeNil())
	// })

	// t.Run(`username required`, func(t *testing.T) {
	// 	member := entity.Member{
	// 		Username: "",
	// 		Password: "12",
	// 		Email:    "B6526542@gmail.com",
	// 	}
	// 	ok, err := govalidator.ValidateStruct(member)
	// 	g.Expect(ok).NotTo(BeTrue())
	// 	g.Expect(err).NotTo(BeNil())
	// 	g.Expect(err.Error()).To(Equal("Username is required."))
	// })

	t.Run(`invalid email`, func(t *testing.T) {
		member := entity.Member{
			Username: "trest",
			Password: "12",
			Email:    "B6526542",
		}
		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid."))
	})

}
