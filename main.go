package main

import (
  "fmt"
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"

)

const DRIVER = "mysql"
const DSN = "golang-test-user:golang-test-pass@tcp(go_api_db:3306)/golang-test-database"

func main() {
  e := echo.New()
  // Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  // Routes
  e.GET("/", getAllRecord)
  e.POST("/", createRecord)
  e.PATCH("/", updateRecord)
  e.DELETE("/:date", deleteRecord)
  // e.GET("/users/:Id", getUser)
  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

type Record struct{
  Date string `json:"date"`
  Weight float64 `json:"weight"`
  Bfp float64 `json:"bfp"`
  Mm float64 `json:"mm"`
}

func getAllRecord(c echo.Context) (err error){
  db, err := sql.Open(DRIVER, DSN)
  if err != nil{
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  defer db.Close()

  rows, err := db.Query(`SELECT * FROM records`)
    if err != nil {
        fmt.Println(err)
    }

    defer rows.Close()

    records := []Record{}
    record := Record{}
    for rows.Next() {
        err = rows.Scan(&record.Date, &record.Weight, &record.Bfp, &record.Mm)
        if err != nil {
            fmt.Println(err)
        }
        records = append(records, record)
    }
    fmt.Println(records)
    return c.JSON(http.StatusOK, records)

}

func createRecord(c echo.Context) (err error){
  db, err := sql.Open(DRIVER, DSN)
  if err != nil{
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  defer db.Close()

  date := c.FormValue("date")
  weight := c.FormValue("weight")
  bfp := c.FormValue("bfp")
  mm := c.FormValue("mm")

  ins, err := db.Prepare("INSERT INTO records(date, weight, bfp, mm) VALUES(?,?,?,?)")
    if err != nil {
        fmt.Println(err)
    }
    ins.Exec(date, weight, bfp, mm)
	return c.String(http.StatusCreated, "record created")
}

func updateRecord(c echo.Context) (err error){
  db, err := sql.Open(DRIVER, DSN)
  if err != nil{
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  defer db.Close()

  date := c.QueryParam("date")
  weight := c.QueryParam("weight")
  bfp := c.QueryParam("bfp")
  mm := c.QueryParam("mm")
  fmt.Println(date, weight, bfp, mm)
  if weight != ""{
    upd, err := db.Prepare("UPDATE records set weight = ? where date = ? ")
      if err != nil {
          fmt.Println(err)
      }
      upd.Exec(weight, date)
  }
  if bfp != ""{
    upd, err := db.Prepare("UPDATE records set bfp = ? where date = ? ")
      if err != nil {
          fmt.Println(err)
      }
      upd.Exec(bfp, date)
  }
  if mm != ""{
    upd, err := db.Prepare("UPDATE records set mm = ? where date = ? ")
      if err != nil {
          fmt.Println(err)
      }
      upd.Exec(mm, date)
  }

  return c.String(http.StatusOK, "Record updated")
}


func deleteRecord(c echo.Context) (err error){
  db, err := sql.Open(DRIVER, DSN)
  if err != nil{
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  defer db.Close()

  date := c.Param("date")
  del, err := db.Prepare("DELETE FROM records where date = ?")
    if err != nil {
        fmt.Println(err)
    }
    del.Exec(date)

  return c.String(http.StatusOK, "Record deleted")
}
