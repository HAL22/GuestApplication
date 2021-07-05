package testing

import (
	"github.com/GG_Backend_tech_challenge/src/mocks"
	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/GG_Backend_tech_challenge/src/repository"
	"github.com/GG_Backend_tech_challenge/src/service"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service Test", func() {
	var (
		ctrl              *gomock.Controller
		tableRepo         *mocks.MockTableRepository
		guestRepo         *mocks.MockGuestRepository
		guestArrivalsRepo *mocks.MockGuestArrivalsRepository
		DB                *gorm.DB
	)
	Context("testing AddGuestToGuestList", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			tableRepo = mocks.NewMockTableRepository(ctrl)
			guestRepo = mocks.NewMockGuestRepository(ctrl)
			DB = repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
		})
		It("Can add guests", func() {
			name := "Tim"
			accompany_guests := 3
			guest := model.Guest{Name: name, AccompanyingGuests: accompany_guests}
			guestRepo.EXPECT().GetGuestByName(gomock.Any()).Return(false, guest)
			table := model.Table{ID: 2, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
			tableRepo.EXPECT().DoesTableExist(gomock.Any()).Return(true, table)
			tableRepo.EXPECT().AssignTableToGuest(gomock.Any(), gomock.Any()).Return(true, 3)
			guest.TableID = int(table.ID)
			guestRepo.EXPECT().AddGuest(gomock.Any()).Return(true)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok, response := EventService.AddGuestToGuestList(name, 2, accompany_guests, tableRepo, guestRepo)
			Expect(ok).To(BeTrue())
			Expect(response.Name).To(Equal(name))

		})
		It("Cannot add guests", func() {
			name := "Tim"
			accompany_guests := 3
			guest := model.Guest{Name: name, AccompanyingGuests: accompany_guests}
			guestRepo.EXPECT().GetGuestByName(gomock.Any()).Return(true, guest)
			table := model.Table{ID: 2, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
			tableRepo.EXPECT().DoesTableExist(gomock.Any()).Return(true, table)
			tableRepo.EXPECT().AssignTableToGuest(gomock.Any(), gomock.Any()).Return(true, 3)
			guest.TableID = int(table.ID)
			guestRepo.EXPECT().AddGuest(gomock.Any()).Return(true)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok, _ := EventService.AddGuestToGuestList(name, 2, accompany_guests, tableRepo, guestRepo)
			Expect(ok).To(BeFalse())
		})
	})
	Context("testing GetGuests", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			tableRepo = mocks.NewMockTableRepository(ctrl)
			guestRepo = mocks.NewMockGuestRepository(ctrl)
			DB = repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
		})
		It("Gets guests", func() {
			guests := []model.Guest{{Name: "Tim"}, {Name: "Peter"}, {Name: "Rummy"}}
			guestRepo.EXPECT().GetGuests().Return(true, guests)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok, response := EventService.GetGuests(guestRepo)
			Expect(ok).To(BeTrue())
			Expect(len(response.Guests)).To(Equal(3))
		})

	})

	Context("testing AddGuestToArrivedGuests", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			tableRepo = mocks.NewMockTableRepository(ctrl)
			guestRepo = mocks.NewMockGuestRepository(ctrl)
			guestArrivalsRepo = mocks.NewMockGuestArrivalsRepository(ctrl)
			DB = repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
		})
		It("Can add guest to arrival guests 1", func() {
			name := "Tim"
			accompany_guests := 3
			guest := model.Guest{Name: name, AccompanyingGuests: accompany_guests}
			guestRepo.EXPECT().GetGuestByName(gomock.Any()).Return(true, guest)
			table := model.Table{ID: 2, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
			tableRepo.EXPECT().DoesTableExist(gomock.Any()).Return(true, table)
			tableRepo.EXPECT().AssignTableToGuest(gomock.Any(), gomock.Any()).Return(true, 3)
			guestArrivalsRepo.EXPECT().AddArrivedGuest(gomock.Any(), gomock.Any()).Return(true)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok, _ := EventService.AddGuestToArrivedGuests(name, accompany_guests, tableRepo, guestRepo, guestArrivalsRepo)

			Expect(ok).To(BeTrue())

		})
		It("Can add guest to arrival guests 2", func() {
			name := "Tim"
			accompany_guests := 3
			guest := model.Guest{Name: name, AccompanyingGuests: 1}
			guestRepo.EXPECT().GetGuestByName(gomock.Any()).Return(true, guest)
			table := model.Table{ID: 2, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
			tableRepo.EXPECT().IncreaseGuestSeats(gomock.Any(), gomock.Any()).Return(true, 1)
			tableRepo.EXPECT().DoesTableExist(gomock.Any()).Return(true, table)
			tableRepo.EXPECT().AssignTableToGuest(gomock.Any(), gomock.Any()).Return(true, 3)
			guestArrivalsRepo.EXPECT().AddArrivedGuest(gomock.Any(), gomock.Any()).Return(true)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok, _ := EventService.AddGuestToArrivedGuests(name, accompany_guests, tableRepo, guestRepo, guestArrivalsRepo)

			Expect(ok).To(BeTrue())

		})
		It("Cannot add guest to arrival guests", func() {
			name := "Tim"
			accompany_guests := 3
			guest := model.Guest{Name: name, AccompanyingGuests: 1}
			guestRepo.EXPECT().GetGuestByName(gomock.Any()).Return(true, guest)
			table := model.Table{ID: 2, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
			tableRepo.EXPECT().IncreaseGuestSeats(gomock.Any(), gomock.Any()).Return(false, 1)
			tableRepo.EXPECT().DoesTableExist(gomock.Any()).Return(true, table)
			tableRepo.EXPECT().AssignTableToGuest(gomock.Any(), gomock.Any()).Return(true, 3)
			guestArrivalsRepo.EXPECT().AddArrivedGuest(gomock.Any(), gomock.Any()).Return(true)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok, _ := EventService.AddGuestToArrivedGuests(name, accompany_guests, tableRepo, guestRepo, guestArrivalsRepo)

			Expect(ok).To(BeFalse())

		})

	})
	Context("testing DeleteArrivedGuest", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			tableRepo = mocks.NewMockTableRepository(ctrl)
			guestRepo = mocks.NewMockGuestRepository(ctrl)
			DB = repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
		})
		It("Deleting", func() {
			name := "Tim"
			guest := model.Guest{Name: name, AccompanyingGuests: 1}
			guestRepo.EXPECT().GetGuestByName(gomock.Any()).Return(true, guest)
			table := model.Table{ID: 2, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
			tableRepo.EXPECT().DoesTableExist(gomock.Any()).Return(true, table)
			tableRepo.EXPECT().RemoveGuestFromTable(gomock.Any(), gomock.Any()).Return(true, 2)
			guestArrivalsRepo.EXPECT().DeleteArrivedGuestByGuestName(gomock.Any()).Return(true)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok := EventService.DeleteArrivedGuest(name, tableRepo, guestRepo, guestArrivalsRepo)

			Expect(ok).To(BeTrue())
		})
	})
	Context("testing GetArrivedGuests", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			tableRepo = mocks.NewMockTableRepository(ctrl)
			guestRepo = mocks.NewMockGuestRepository(ctrl)
			DB = repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
		})
		It("Gets arrived guests", func() {
			arrived_guests := []model.GuestArrivals{{Name: "Tim"}, {Name: "Peter"}, {Name: "Rummy"}}
			guestArrivalsRepo.EXPECT().GetArrivedGuests().Return(true, arrived_guests)
			EventService := service.Eventservice{
				DB: DB,
			}
			ok, response := EventService.GetArrivedGuests(guestArrivalsRepo)
			Expect(ok).To(BeTrue())
			Expect(len(response.Guests)).To(Equal(3))
		})

	})
	Context("testing GetEmptySeats", func() {
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			tableRepo = mocks.NewMockTableRepository(ctrl)
			guestRepo = mocks.NewMockGuestRepository(ctrl)
			DB = repository.GetDataBaseConnectionWithTablesAndData("root", "turing221997", "localhost", 3306, "event")
		})
		It("Seats", func() {
			tableRepo.EXPECT().GetEmptySeats().Return(3)
			EventService := service.Eventservice{
				DB: DB,
			}
			seats := EventService.GetEmptySeats(tableRepo)
			Expect(seats.SeatsEmpty).To(Equal(3))

		})
	})

})
