package testing

import (
	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/GG_Backend_tech_challenge/src/repository"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Guest Repository Test", func() {
	var (
		DB        *gorm.DB
		guestRepo repository.GuestRepo
		user      string = "root"
		password  string = "turing221997"
		host      string = "localhost"
		port      int    = 3306
		db        string = "event"
	)
	Context("testing GetGuestByName", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			guestRepo = repository.GuestRepo{
				DB: DB,
			}

		})
		It("Name exits", func() {
			guest := model.Guest{Name: "Tim", AccompanyingGuests: 100, TableID: -1}
			DB.Create(&guest)
			bool_answer, guest2 := guestRepo.GetGuestByName(guest.Name)
			Expect(bool_answer).To(BeTrue())
			Expect(guest2.Name).To(Equal(guest.Name))
		})
		It("Name does not exist", func() {
			guest := model.Guest{Name: "Tim", AccompanyingGuests: 100, TableID: -1}
			bool_answer, _ := guestRepo.GetGuestByName(guest.Name)
			Expect(bool_answer).To(BeFalse())
		})
	})
	Context("testing GetGuests", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			guestRepo = repository.GuestRepo{
				DB: DB,
			}

		})
		It("Get Guest Array", func() {
			guest1 := model.Guest{Name: "Tim", AccompanyingGuests: 10, TableID: 2}
			guest2 := model.Guest{Name: "Phil", AccompanyingGuests: 18, TableID: 3}
			guest3 := model.Guest{Name: "Fo", AccompanyingGuests: 2, TableID: 8}
			guest4 := model.Guest{Name: "Den", AccompanyingGuests: 1, TableID: 10}
			DB.Create(&guest1)
			DB.Create(&guest2)
			DB.Create(&guest3)
			DB.Create(&guest4)
			bool_answer, guests := guestRepo.GetGuests()
			Expect(bool_answer).To(BeTrue())
			Expect(len(guests)).To(Equal(4))

		})
	})
})
