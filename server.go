package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	account "github.com/khanaphoz/myapi/account/type"
)

var acc = make(map[string]*account.AccStr)
var id = 0

type Accounties []account.AccStr

func getAccHandler(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"myid": "id"})
	var tid int
	status := c.Query("status")
	fmt.Println("status", status)
	//tempAcc := []*account.AccStr{}
	//ac := new(account.AccStr)
	//c.JSON(http.StatusOK, acc)
	var accounties Accounties

	if status != "" {

		for _, _ = range acc {

			tid++
			newid := strconv.Itoa(tid)

			if acc[newid].AccStatus == status {
				//c.JSON(http.StatusOK, acc[newid].AccStatus)
				accounties = append(accounties, account.AccStr{
					AccID:     acc[newid].AccID,
					AccName:   acc[newid].AccName,
					AccNo:     acc[newid].AccNo,
					AccStatus: acc[newid].AccStatus,
				})

			}
		}
		c.JSON(http.StatusOK, accounties)
		return

	}
	c.JSON(http.StatusOK, acc)

}

func getAccByIdHandler(c *gin.Context) {
	id := c.Param("id")

	if acc[id] != nil {
		//		c.JSON(http.StatusOK, acc[id].AccName)
		c.JSON(http.StatusOK, acc[id])
		return
	}
	c.String(http.StatusBadRequest, "Account not found")

}

func createAccHandler(c *gin.Context) {
	//var item account.AccStr
	var item account.AccStr
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}
	id++
	item.AccID = strconv.Itoa(id)
	//item.AccStatus = "Active"
	acc[item.AccID] = &item
	c.String(http.StatusOK, "Create Account Name :  "+item.AccName+"  Completed")

}
func updateAccHandler(c *gin.Context) {
	id := c.Param("id")
	var item account.AccStr
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.AccID = id
	acc[id] = &item
	c.JSON(http.StatusOK, "UPDATE COMPLETE")
}
func deleteAccHandler(c *gin.Context) {
	id := c.Param("id")
	acc[id].AccID = ""
	acc[id].AccName = ""
	acc[id].AccNo = ""
	acc[id].AccStatus = ""
	c.JSON(http.StatusOK, "DELETE COMPLETE")

}
func main() {
	r := gin.Default()
	r.GET("/account", getAccHandler)
	r.GET("/account/:id", getAccByIdHandler)
	r.POST("/account", createAccHandler)
	r.PUT("/account/:id", updateAccHandler)
	r.DELETE("/account/:id", deleteAccHandler)
	r.Run(":8787")

}
