package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sajjad3k/ginbankapiex/models"
)

//var customers []models.Customer

/*
 customers = []models.customer{
	{Id: "first@1", Name: "first", Branch: "bangfir", Balance: 20000, City: "bengaluru"},
	{Id: "second@2", Name: "second", Branch: "bangsec", Balance: 30000, City: "bengaluru"},
	{Id: "third@3", Name: "third", Branch: "bangthird", Balance: 50000, City: "bengaluru"},
}

var resp models.response */

//Fetch all the customers
func Getallcustomers(c *gin.Context) {

	p, err := models.Askdata()
	var resp models.Response
	if err != nil {
		resp.Status = "Failed"
		resp.Message = "No records Found"
		c.JSON(http.StatusNotFound, resp)
		//c.AbortWithStatus(http.StatusNotFound)
	} else {
		resp.Status = "Success"
		resp.Data = append(resp.Data, p...)
		resp.Message = "fetched all the records"
		c.JSON(http.StatusOK, resp)
	}

}

//Fetch the customer by ID
func Getcustomerbyid(c *gin.Context) {

	id := c.Params.ByName("id")
	p, err := models.Askdata()
	var resp models.Response
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}
	if id != "" && p != nil {
		for _, val := range p {
			if val.Id == id {
				resp.Status = "success"
				resp.Data = append(resp.Data, val)
				resp.Message = "The customer is found with given id"
				c.JSON(http.StatusOK, resp)
			}
		}
	} else {
		resp.Status = "Error"
		resp.Message = "Error occured"
		c.JSON(http.StatusBadRequest, resp)
	}
}

//Create a new customer
func Createcustomer(c *gin.Context) {

	var resp models.Response
	var customer models.Customer
	/*if c.Params == nil {
		resp.Status = "Error"
		resp.Message = "the request parameters are empty"
		c.JSON(http.StatusBadRequest, resp)
		c.AbortWithStatus(http.StatusBadRequest)
	} else { */
	c.BindJSON(&customer)
	da, err := models.Askdata()
	if da != nil || err != nil {
		//log.Fatal("empty data") //c.AbortWithStatus(http.StatusInternalServerError)
		//   if da != nil {
		da = append(da, customer)
		models.Setdata(da)
		resp.Status = "success"
		resp.Data = append(resp.Data, customer)
		resp.Message = "new customer created"
		c.JSON(http.StatusOK, resp)
	}
	//}

}

//Delete a customer entry
func Deletecustomer(c *gin.Context) {
	var resp models.Response
	var customer models.Customer
	var out []models.Customer
	var flag bool = false
	id := c.Params.ByName("id")
	if id != "" {
		p, err := models.Askdata()
		if err != nil {
			resp.Status = "Info"
			resp.Message = "The customers list is empty"
			c.JSON(http.StatusInternalServerError, resp)
			//c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			for _, val := range p {
				if val.Id == id {
					flag = true
					customer = val
					continue
				}
				out = append(out, val)
			}
			p := out
			if flag != false {
				resp.Data = append(resp.Data, customer)
				resp.Status = "Success"
				resp.Message = "Deleted the Customer entry"
				models.Setdata(p)
				c.JSON(http.StatusOK, resp)
			} else {
				c.AbortWithStatus(http.StatusNotFound)
			}
		}

	} else {
		resp.Status = "error"
		resp.Message = "The request is not correct"
		c.JSON(http.StatusBadRequest, resp)
	}

}

//Update a customer details
func Updatecustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer models.Customer
	var resp models.Response
	var flag bool = false
	if id != "" {
		c.BindJSON(&customer)
		p, err := models.Askdata()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {

			for i, val := range p {
				if val.Id == id {
					p[i].Name = customer.Name
					p[i].City = customer.City
					p[i].Branch = customer.Branch
					flag = true
				}
			}
			if flag == true {
				models.Setdata(p)
				resp.Status = "success"
				resp.Message = "The customer entry is updated successfully"
				resp.Data = append(resp.Data, customer)
				c.JSON(http.StatusOK, resp)
			}

		}

	} else {
		resp.Status = "Error"
		resp.Message = "The request is not correct"
		c.JSON(http.StatusBadRequest, resp)
	}

}

//Update the Customer balance
func Updatebalance(c *gin.Context) {

	id := c.Params.ByName("id")
	var resp models.Response
	var flag bool = false
	var customer models.Customer
	newbalance, errs := strconv.Atoi(c.Params.ByName("amount"))
	if errs != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if id != "" {
		p, err := models.Askdata()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			for i, val := range p {
				if val.Id == id {
					p[i].Balance = newbalance
					customer = p[i]
					flag = true
					break
				}
			}

			if flag == true {
				resp.Status = "success"
				resp.Message = "new balance updated"
				resp.Data = append(resp.Data, customer)
				c.JSON(http.StatusOK, resp)
			} else {
				resp.Status = "error"
				resp.Message = "Customer does not exist"
				c.JSON(http.StatusBadRequest, resp)
			}
		}
	}

}

//Transfer the money from one customer to another customer
func Transfermoney(c *gin.Context) {
	var resp models.Response
	var flag1 bool = false
	var flag2 bool = false
	var cust1, cust2 models.Customer
	id := c.Params.ByName("id")
	toid := c.Params.ByName("toid")
	amount, err := strconv.Atoi(c.Params.ByName("amount"))
	if id == "" || toid == "" || err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		p, err := models.Askdata()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			for i, val := range p {
				if val.Id == id {
					p[i].Balance = p[i].Balance - amount
					cust1 = p[i]
					flag1 = true
					continue
				} else if val.Id == toid {
					p[i].Balance = p[i].Balance + amount
					cust2 = p[i]
					flag2 = true
				} else {
					continue
				}
			}
			if flag1 == true && flag2 == true {
				resp.Status = "success"
				resp.Message = "The amount transeferred successfully"
				resp.Data = append(resp.Data, cust1)
				resp.Data = append(resp.Data, cust2)
				c.JSON(http.StatusOK, resp)
			} else {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}

	}

}

//To check if the Customer ID available to create
func Checkidavailable(c *gin.Context) {

	var resp models.Response
	id := c.Params.ByName("id")
	if id != "" {
		p, err := models.Askdata()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			for _, val := range p {
				if val.Id == id {
					resp.Message = "The Customer ID exists ,Try a different one"
					break
				}
			}
			if resp.Message == "" {
				resp.Message = "The Customer ID is available"
			}
			c.JSON(http.StatusOK, resp)
		}
	} else {
		resp.Message = "enter a Customer ID to check"
		c.JSON(http.StatusOK, resp)
	}
}
