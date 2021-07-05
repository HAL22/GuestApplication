package testing

import (
	"time"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/GG_Backend_tech_challenge/src/repository"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Guest Arrivals Repository Test", func() {
	var (
		DB                *gorm.DB
		guestArrivalsRepo repository.GuestArrivalsRepo
		user              string = "root"
		password          string = "turing221997"
		host              string = "localhost"
		port              int    = 3306
		db                string = "event"
	)
	Context("testing AddArrivedGuest", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			guestArrivalsRepo = repository.GuestArrivalsRepo{
				DB: DB,
			}
		})
		It("Adding arrived guest", func() {
			var arrivedguest model.GuestArrivals
			guest := model.Guest{ID: 0, Name: "Tim", AccompanyingGuests: 10, TableID: -1}
			arrivetime := time.Now()
			bool_answer := guestArrivalsRepo.AddArrivedGuest(guest, arrivetime)
			DB.Where("name = ?", guest.Name).First(&arrivedguest)

			Expect(bool_answer).To(BeTrue())
			Expect(arrivedguest.GuestID).To(Equal(guest.ID))
		})
	})
	Context("testing DeleteArrivedGuestByGuestName", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			guestArrivalsRepo = repository.GuestArrivalsRepo{
				DB: DB,
			}
		})
		It("Guest exist can be deleted", func() {
			guest := model.Guest{ID: 0, Name: "Tim", AccompanyingGuests: 10, TableID: -1}
			arrivetime := time.Now()
			guestArrivalsRepo.AddArrivedGuest(guest, arrivetime)
			bool_answer := guestArrivalsRepo.DeleteArrivedGuestByGuestName(guest.Name)

			Expect(bool_answer).To(BeTrue())
		})
		It("Guest does not exisits and  cannot  be deleted", func() {
			bool_answer := guestArrivalsRepo.DeleteArrivedGuestByGuestName("Phil")
			Expect(bool_answer).To(BeFalse())
		})

	})
	Context("testing GetArrivedGuests", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			guestArrivalsRepo = repository.GuestArrivalsRepo{
				DB: DB,
			}
		})
		It("Get a list Guests that have arrived", func() {
			arrived1 := model.GuestArrivals{Name: "Shil", AccompanyingGuests: 10, ArrivalTime: time.Now(), GuestID: 1}
			arrived2 := model.GuestArrivals{Name: "Phil", AccompanyingGuests: 10, ArrivalTime: time.Now(), GuestID: 2}
			arrived3 := model.GuestArrivals{Name: "Nil", AccompanyingGuests: 15, ArrivalTime: time.Now(), GuestID: 3}
			arrived4 := model.GuestArrivals{Name: "Bil", AccompanyingGuests: 5, ArrivalTime: time.Now(), GuestID: 4}
			DB.Create(&arrived1)
			DB.Create(&arrived2)
			DB.Create(&arrived3)
			DB.Create(&arrived4)
			bool_answer, arrivedguest := guestArrivalsRepo.GetArrivedGuests()
			Expect(bool_answer).To(BeTrue())
			Expect(len(arrivedguest)).To(Equal(4))

		})
	})

})
