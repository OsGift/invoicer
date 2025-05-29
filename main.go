package main

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
)

// InvoiceData holds the data for the invoice
type InvoiceData struct {
	InvoiceNumber string
	InvoiceDate   string
	BillTo        string
	Items         []Item
	Subtotal      string
	Tax           string
	Total         string
}

// Item represents a line item on the invoice
type Item struct {
	Description string
	Quantity    int
	Price       float64
	Total       string
}

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Serve the HTML form
	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Handle form submission
	r.POST("/generate-pdf", func(c *gin.Context) {
		invoiceNumber := c.PostForm("invoiceNumber")
		invoiceDate := c.PostForm("invoiceDate")
		billTo := c.PostForm("billTo")

		var items []Item
		descriptions := c.PostFormArray("itemDescription[]")
		quantities := c.PostFormArray("itemQuantity[]")
		prices := c.PostFormArray("itemPrice[]")

		var subtotal float64
		for i := range descriptions {
			quantity, _ := strconv.Atoi(quantities[i])
			price, _ := strconv.ParseFloat(prices[i], 64)
			total := float64(quantity) * price
			subtotal += total
			items = append(items, Item{
				Description: descriptions[i],
				Quantity:    quantity,
				Price:       price,
				Total:       strconv.FormatFloat(total, 'f', 2, 64),
			})
		}

		tax := subtotal * 0.1 // Example: 10% tax
		total := subtotal + tax

		data := InvoiceData{
			InvoiceNumber: invoiceNumber,
			InvoiceDate:   invoiceDate,
			BillTo:        billTo,
			Items:         items,
			Subtotal:      strconv.FormatFloat(subtotal, 'f', 2, 64),
			Tax:           strconv.FormatFloat(tax, 'f', 2, 64),
			Total:         strconv.FormatFloat(total, 'f', 2, 64),
		}

		// Generate PDF
		pdfg, err := wkhtml.NewPDFGenerator()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating PDF generator: %v", err)
			return
		}

		// Load the HTML template
		tmpl, err := template.ParseFiles("templates/invoice.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading template: %v", err)
			return
		}

		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, data); err != nil {
			c.String(http.StatusInternalServerError, "Error executing template: %v", err)
			return
		}

		pdfg.AddPage(wkhtml.NewPageReader(&tpl))

		err = pdfg.Create()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating PDF: %v", err)
			return
		}

		c.Header("Content-Disposition", "attachment; filename=invoice.pdf")
		c.Data(http.StatusOK, "application/pdf", pdfg.Bytes())
	})

	r.Run(":8080")
}
