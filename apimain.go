package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func Getdb() *sql.DB {
	// Connect to database

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root@123", "127.0.0.1", 3306, "Dev")

	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("Error Connecting to DB %s\n", err.Error())
		return db
	}
	log.Println("DB Connected!\n")

	return db
}

//--------------------------------------------------------------------
// function to record program start and end time details into db
//--------------------------------------------------------------------
/*func RecordRunDetails(db *sql.DB, id int, runType string, programName string, count int, cmt string) (int, error) {
	insertedID := 0
	if runType == INSERT {
		insertString := "INSERT INTO SchedulerRunDetails(StartTime,ProgramName,RecordCount,comment)  values($1,$2,$3,$4);SELECT SCOPE_IDENTITY() "
		inserterr := db.QueryRow(insertString, time.Now(), programName, count, cmt).Scan(&insertedID)
		if inserterr != nil {
			return insertedID, fmt.Errorf("Error while inserting SchedulerRunDetails: ", inserterr.Error())
			//LogError(Panic, inserterr.Error())
		}
	} else if runType == UPDATE {
		insertedID = id
		updateString := "UPDATE SchedulerRunDetails  WITH (UPDLOCK)  SET EndTime=$1,RecordCount=$2,comment= w$3here id=$4 "

		_, updateerr := db.Exec(updateString, time.Now(), count, cmt, insertedID)
		if updateerr != nil {
			//log.Println(updateerr.Error())
			return insertedID, fmt.Errorf("Error while updating SchedulerRunDetails: ", updateerr.Error())

		}
	}
	return insertedID, nil

}*/

func location(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//construct query to get the parameter values
	fullpath := req.URL.Path + "?" + req.URL.RawQuery
	//parse parameter values
	_, err := url.Parse(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	//get parameter values
	//	q := u.Query()

	db := Getdb()    //create db session
	defer db.Close() //close db session

	locString := "SELECT * FROM location"

	type records struct {
		LocId   int
		LocName string
	}
	var rec records
	var recArray []records
	rows, err := db.Query(locString)
	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			err := rows.Scan(&rec.LocId, &rec.LocName)
			if err != nil {
				log.Println(err)
			} else {
				recArray = append(recArray, rec)
			}
		}
	}
	log.Println(recArray)
	data, err := json.Marshal(recArray)
	if err != nil {
		fmt.Fprintf(w, "Error taking data")
		return
	} else {
		fmt.Fprintf(w, string(data))
	}
	//fmt.Println(string(b))
}

func department(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//construct query to get the parameter values
	fullpath := req.URL.Path + "?" + req.URL.RawQuery
	//parse parameter values
	_, err := url.Parse(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	//get parameter values
	//	q := u.Query()

	db := Getdb()    //create db session
	defer db.Close() //close db session

	deptString := "SELECT * FROM department"

	type records struct {
		DeptId   int
		DeptName string
	}
	var rec records
	var recArray []records
	rows, err := db.Query(deptString)
	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			err := rows.Scan(&rec.DeptId, &rec.DeptName)
			if err != nil {
				log.Println(err)
			} else {
				recArray = append(recArray, rec)
			}
		}
	}
	log.Println(recArray)
	data, err := json.Marshal(recArray)
	if err != nil {
		fmt.Fprintf(w, "Error taking data")
		return
	} else {
		fmt.Fprintf(w, string(data))
	}
	//fmt.Println(string(b))
}

func updateEmployee(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("üpdate called")
	//construct query to get the parameter values
	fullpath := req.URL.Path + "?" + req.URL.RawQuery
	//parse parameter values
	u, err := url.Parse(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	//get parameter values
	q := u.Query()
	vId := ""
	vName := ""
	vDob := ""
	vdeptId := ""
	vLocId := ""
	if q.Get("id") != "undefined" && q.Get("id") != "" {
		vId = q.Get("id")
	}
	if q.Get("name") != "undefined" && q.Get("name") != "" {
		vName = q.Get("name")
	}
	if q.Get("dob") != "undefined" && q.Get("dob") != "" {
		vDob = q.Get("dob")
	}
	if q.Get("dept") != "undefined" && q.Get("dept") != "" {
		vdeptId = q.Get("dept")
	}
	if q.Get("location") != "undefined" && q.Get("location") != "" {
		vLocId = q.Get("location")
	}

	db := Getdb()    //create db session
	defer db.Close() //close db session
	log.Println(vName)
	log.Println(vDob)
	log.Println(vdeptId)
	log.Println(vLocId)
	log.Println(vId)
	updateString := "UPDATE employee SET NAME = '" + vName + "',dob= '" + vDob + "', dept= '" + vdeptId + "', location= '" + vLocId + "'  WHERE id = '" + vId + "'"
	log.Println(updateString)

	_, updateerr := db.Exec(updateString)
	if updateerr != nil {
		log.Println(updateerr.Error())
		fmt.Fprintf(w, "Error updating data "+updateerr.Error())

	}

	fmt.Fprintf(w, "200")
}

func employee(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//construct query to get the parameter values
	fullpath := req.URL.Path + "?" + req.URL.RawQuery
	//parse parameter values
	u, err := url.Parse(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	//get parameter values
	q := u.Query()
	qwhere := ""
	if q.Get("id") != "undefined" && q.Get("id") != "" {
		qwhere = " AND e.id=" + q.Get("id")
	}

	db := Getdb()    //create db session
	defer db.Close() //close db session

	sqlString := `SELECT e.id,e.name,if(e.dob is null,"",e.dob),d.name DeptName,l.name LocName,d.id deptId, l.id LocId
		FROM employee e, department d, location l
		WHERE e.dept = d.id
		AND e.location = l.id` + qwhere

	type records struct {
		Id       int
		Name     string
		Dob      string
		DeptName string
		LocName  string
		DeptId   int
		LocId    int
	}
	var rec records
	var recArray []records
	rows, err := db.Query(sqlString)
	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			err := rows.Scan(&rec.Id, &rec.Name, &rec.Dob, &rec.DeptName, &rec.LocName, &rec.DeptId, &rec.LocId)
			if err != nil {
				log.Println(err)
			} else {
				recArray = append(recArray, rec)
			}
		}
	}
	log.Println(recArray)
	data, err := json.Marshal(recArray)
	if err != nil {
		fmt.Fprintf(w, "Error taking data")
		return
	} else {
		fmt.Fprintf(w, string(data))
	}
	//fmt.Println(string(b))
}

func dept(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/employee", employee)
	http.HandleFunc("/employee/update/", updateEmployee)
	http.HandleFunc("/department", department)
	http.HandleFunc("/location", location)

	http.ListenAndServe(":28090", nil)
}
