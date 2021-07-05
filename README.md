# GuestApplication

1. Assumptions
 -  When a guest is removed from  the party (arrived guest), the guest cannot arrive  or re-join the party again.Also the guest is not deleted from the guest list 
 - If guest brings to many  accompanying guests at the party (when the guest arrives ), the guest try again with a smaller entourage 
 - The guest and his/her entourage must fit in one table,no seperation
 - Multipe guests and their respective entourage can share a table

2. Disclaimer
 - You have to create a database named event or create a database and  set the name using os, the os variable is DB_NAME(Pleass see main.go)
 - Also the password,user and IP address for the mysql are set using os:
    os.Setenv("USER", "root")
	os.Setenv("PASSWORD", "turing221997")
	os.Setenv("IP_ADDRESS", "localhost")
	os.Setenv("DB_NAME", "localhost")
    os.Setenv("PORT_NUM", "3306")
	os.Setenv("HTTP_PORT", ":8001")
 - The Database refreshes everytime the program runs, so all the tables are cleared.
 - When the database is created, initial 10 entries to the "tables" table are created.But new "tables" can be added to the tables table

3. Requirments
 - Check the go.mod file for external packages used
 - Go and mysql should be installed

4. Endpoints
 - POST /guest_list/name
        Body:{
            "table": int,
            "accompanying_guests": int
        }
 - GET /guest_list
 - PUT /guests/name
        Body:{
            "table": int,
            "accompanying_guests": int
        }
 - DELETE /guests/name
 - GET /guests
 - GET /seats_empty
 - PUT /table adding a table 
         Body:{
            "table": int,
            "capacity": int
        }


5. Running the program: go run $PATH/main.go

6. Tests
 - Tests are in the package testing and for every test the os variables are set
 - to run test, navigate to the testing folder and via terminal type: ginkgo run . (to run all of the tests)
