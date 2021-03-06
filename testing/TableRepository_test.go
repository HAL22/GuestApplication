package testing

import (
	"os"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/GG_Backend_tech_challenge/src/repository"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Table Repository test", func() {
	os.Setenv("USER", "root")
	os.Setenv("PASSWORD", "turing221997")
	os.Setenv("IP_ADDRESS", "localhost")
	os.Setenv("DB_NAME", "event")
	os.Setenv("PORT_NUM", "3306")
	var (
		DB        *gorm.DB
		tableRepo repository.TableRepo
		user      string = os.Getenv("USER")
		password  string = os.Getenv("PASSWORD")
		host      string = os.Getenv("IP_ADDRESS")
		port      string = os.Getenv("PORT_NUM")
		db        string = os.Getenv("DB_NAME")
	)
	Context("testing DoesTableExist", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			tableRepo = repository.TableRepo{
				DB: DB,
			}

		})
		It("table exists", func() {
			table := model.Table{ID: 12, Capacity: 15, Sizeofguests: 0, Emptyseats: 15}
			DB.Create(&table)
			bool_answer, table_answer := tableRepo.DoesTableExist(int(table.ID))

			Expect(bool_answer).To(BeTrue())
			Expect(table_answer.ID).To(Equal(table.ID))

		})
		It("table does not exist", func() {
			bool_answer, _ := tableRepo.DoesTableExist(int(24))

			Expect(bool_answer).To(BeFalse())

		})

	})
	Context("testing AssignTableToGuest", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			tableRepo = repository.TableRepo{
				DB: DB,
			}

		})
		It("Right number of guests", func() {
			var table model.Table
			DB.Where("id = ?", 10).Find(&table)
			guest := model.Guest{Name: "Tim", AccompanyingGuests: 5, TableID: -1}
			bool_answer, size := tableRepo.AssignTableToGuest(table, &guest)

			Expect(bool_answer).To(BeTrue())
			Expect(size).To(Equal(9))
			Expect(guest.TableID).To(Equal(int(table.ID)))

		})
		It("To many guests", func() {
			var table model.Table
			DB.Where("id = ?", 10).Find(&table)
			guest := model.Guest{Name: "Tim", AccompanyingGuests: 100, TableID: -1}
			bool_answer, size := tableRepo.AssignTableToGuest(table, &guest)

			Expect(bool_answer).To(BeFalse())
			Expect(size).To(Equal(0))
			Expect(guest.TableID).NotTo(Equal(int(table.ID)))

		})
	})
	Context("testing RemoveGuestFromTable", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			tableRepo = repository.TableRepo{
				DB: DB,
			}
		})
		It("It removes the guest", func() {
			var table model.Table
			var guest model.Guest
			DB.Where("id = ?", 10).Find(&table)
			emptyseats := table.Emptyseats
			sizeofguests := table.Sizeofguests
			guest = model.Guest{Name: "Tim", AccompanyingGuests: 3, TableID: -1}
			tableRepo.AssignTableToGuest(table, &guest)
			DB.Create(&guest)
			DB.Where("name = ?", "Tim").Find(&guest)
			DB.Where("id = ?", 10).Find(&table)
			_, seats := tableRepo.RemoveGuestFromTable(guest, table)
			DB.Where("id = ?", 10).Find(&table)
			DB.Where("name = ?", "Tim").Find(&guest)
			Expect(seats).To(Equal(emptyseats))
			Expect(sizeofguests).To(Equal(table.Sizeofguests))
			Expect(guest.TableID).To(Equal(-1))

		})

	})
	Context("testing IncreaseGuestSeats", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			tableRepo = repository.TableRepo{
				DB: DB,
			}
		})
		It("Can add more guests", func() {
			var table model.Table
			var guest model.Guest
			DB.Where("id = ?", 10).Find(&table)
			guest = model.Guest{Name: "Tim", AccompanyingGuests: 3, TableID: -1}
			tableRepo.AssignTableToGuest(table, &guest)
			DB.Create(&guest)
			DB.Where("id = ?", 10).Find(&table)
			sizeofguests := table.Sizeofguests
			empty_seats := table.Emptyseats
			additionalGuests := 5
			bool_answer, _ := tableRepo.IncreaseGuestSeats(table, additionalGuests)
			DB.Where("id = ?", 10).Find(&table)

			Expect(table.Emptyseats).To(Equal(empty_seats - additionalGuests))
			Expect(table.Sizeofguests).To(Equal(sizeofguests + additionalGuests))
			Expect(bool_answer).To(BeTrue())

		})
		It("Cannot add more guests", func() {
			var table model.Table
			var guest model.Guest
			DB.Where("id = ?", 10).Find(&table)
			guest = model.Guest{Name: "Tim", AccompanyingGuests: 3, TableID: -1}
			tableRepo.AssignTableToGuest(table, &guest)
			DB.Create(&guest)
			DB.Where("id = ?", 10).Find(&table)
			sizeofguests := table.Sizeofguests
			empty_seats := table.Emptyseats
			additionalGuests := 500
			bool_answer, _ := tableRepo.IncreaseGuestSeats(table, additionalGuests)
			DB.Where("id = ?", 10).Find(&table)

			Expect(table.Emptyseats).To(Equal(empty_seats))
			Expect(table.Sizeofguests).To(Equal(sizeofguests))
			Expect(bool_answer).To(BeFalse())

		})
	})
	Context("testing DecreaseGuestSeats", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			tableRepo = repository.TableRepo{
				DB: DB,
			}
		})
		It("Can decrease guest and increase empty seats", func() {
			var table model.Table
			var guest model.Guest
			DB.Where("id = ?", 10).Find(&table)
			guest = model.Guest{Name: "Tim", AccompanyingGuests: 3, TableID: -1}
			tableRepo.AssignTableToGuest(table, &guest)
			DB.Create(&guest)
			DB.Where("id = ?", 10).Find(&table)
			sizeofguests := table.Sizeofguests
			empty_seats := table.Emptyseats
			additionalGuests := 2
			bool_answer, _ := tableRepo.DecreaseGuestSeats(table, additionalGuests)
			DB.Where("id = ?", 10).Find(&table)

			Expect(table.Emptyseats).To(Equal(empty_seats + additionalGuests))
			Expect(table.Sizeofguests).To(Equal(sizeofguests - additionalGuests))
			Expect(bool_answer).To(BeTrue())

		})
		It("Cannot decrease", func() {
			var table model.Table
			var guest model.Guest
			DB.Where("id = ?", 10).Find(&table)
			guest = model.Guest{Name: "Tim", AccompanyingGuests: 3, TableID: -1}
			tableRepo.AssignTableToGuest(table, &guest)
			DB.Create(&guest)
			DB.Where("id = ?", 10).Find(&table)
			sizeofguests := table.Sizeofguests
			empty_seats := table.Emptyseats
			additionalGuests := 5
			bool_answer, _ := tableRepo.DecreaseGuestSeats(table, additionalGuests)
			DB.Where("id = ?", 10).Find(&table)
			Expect(table.Emptyseats).To(Equal(empty_seats))
			Expect(table.Sizeofguests).To(Equal(sizeofguests))
			Expect(bool_answer).To(BeFalse())
		})
	})
	Context("testing GetEmptySeats", func() {
		BeforeEach(func() {
			DB = repository.GetDataBaseConnectionWithTablesAndData(user, password, host, port, db)
			tableRepo = repository.TableRepo{
				DB: DB,
			}
		})
		It("Get empty seats", func() {
			var table model.Table
			var guest model.Guest
			beforeEmptySeats := tableRepo.GetEmptySeats()
			DB.Where("id = ?", 10).Find(&table)
			guest = model.Guest{Name: "Tim", AccompanyingGuests: 3, TableID: -1}
			tableRepo.AssignTableToGuest(table, &guest)
			DB.Create(&guest)
			numberOfguests := 4
			afterEmptySeats := tableRepo.GetEmptySeats()

			Expect(afterEmptySeats).To(Equal(beforeEmptySeats - numberOfguests))

		})

	})

})
