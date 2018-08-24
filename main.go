package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

type MobileDataResult struct {
	Prefix      string ` json:”Prefix” `
	GatewayName string ` json:”GatewayName” `
	IpAddress   string ` json:”IpAddress” `
}

func main() {
	r := mux.NewRouter()
	fmt.Println("same thing")
	r.HandleFunc("/numberTrace/{mobileNumber}", MobileBook).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func MobileBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside1")
	vars := mux.Vars(r)
	mobileNumber := vars["mobileNumber"]
	fmt.Println("mobile", mobileNumber)

	prefix := mobileNumber[:4]

	dbStatus, Result := InsertIntoDb(prefix)
	switch dbStatus {
	case true:

		json.NewEncoder(w).Encode(Result)
		jsonResp, _ := json.Marshal(Result)
		fmt.Println(string(jsonResp))
		result := string(jsonResp)
		fmt.Fprintf(w, result)
	}

}

func InsertIntoDb(prefixString string) (bool, MobileDataResult) {

	/*db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	name :="MobileData"
	_,err = db.Exec("CREATE DATABASE  IF NOT EXISTS "+name)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("USE "+name)
	if err != nil {
		panic(err)
	}
	fmt.Println("achu 2")
	*/
	db, err := sql.Open("mysql", "user:password@tcp(host)/dbname")

	//db, err := sql.Open("mysql", "root:"+os.Getenv("mySQLPassword")+"@tcp("+os.Getenv("mySQLIPAddress")+":"+ os.Getenv("mySQLIPPort") +")/")

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	fmt.Println("welcome! to database ")

	// Last thing to do, close connection
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		fmt.Println("error in database connection")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	/*_, err = db.Exec("CREATE DATABASE IF NOT EXISTS MobileData")
	if err!=nil{
		log.Fatal(err)
	}
	*/

	//// Use our database
	_, err = db.Exec("USE MobileData")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("continuee")
	//create prefix table
	_, err = db.Exec("CREATE  TABLE IF NOT EXISTS `PrefixTable`(`prefix` VARCHAR(8),`gateway` VARCHAR (30))")
	if err != nil {
		panic(err.Error())
	}
	prefixArray := []int{123, 1234, 9194}
	gateWay := []string{"airtel", "vodafone", "tata"}
	for i := 0; i < len(prefixArray); i++ {
		prefix := prefixArray[i]
		gateway := gateWay[i]
		insForm, err := db.Prepare("INSERT INTO PrefixTable(prefix,gateway) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(prefix, gateway)
	}
	//defer db.Close()
	//create ip table
	_, err = db.Exec("CREATE  TABLE IF NOT EXISTS `IpAddress`(`gateway` VARCHAR(20),`ip` VARCHAR (50))")
	if err != nil {
		panic(err.Error())
	}
	ipArray := []string{"12.12.12.12, 13.13.13.13", "14.14.14.14 ", "15.15.15.15, 16.16.16.16"}
	gateWayData := []string{"airtel", "vodafone", "tata"}
	for i := 0; i < len(gateWayData); i++ {
		ip := ipArray[i]
		gateway := gateWayData[i]
		insForm, err := db.Prepare("INSERT INTO IpAddress(gateway,ip) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(gateway, ip)
	}
	//find gateway

	preDB, err := db.Query("SELECT prefix,gateway FROM PrefixTable")
	var prefixNo string
	var gatewayNo string
	var ResultGateway string
	if err != nil {
		panic(err.Error())
	}
	for preDB.Next() {
		err := preDB.Scan(&prefixNo, &gatewayNo)
		if err != nil {
			log.Fatal(err)
		}
		if prefixNo == prefixString {

			ResultGateway = gatewayNo
		}
	}
	if len(ResultGateway) == 0 {
		ResultGateway = "airtel"
		prefixString = "123"
	}
	//select ipaddress
	ipDB, err := db.Query("SELECT * FROM IpAddress WHERE gateway=?", ResultGateway)
	if err != nil {
		panic(err.Error())
	}
	Result := MobileDataResult{}
	var resultedIp string

	for ipDB.Next() {
		var gateWayIp string
		var resultIp string
		err = ipDB.Scan(&gateWayIp, &resultIp)
		if err != nil {
			panic(err.Error())
		}
		resultedIp = resultIp
	}
	Result.GatewayName = ResultGateway
	Result.IpAddress = resultedIp
	Result.Prefix = prefixString
	fmt.Println("result", Result)

	return true, Result
}
