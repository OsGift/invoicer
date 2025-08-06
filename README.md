# Invoice PDF Generator in Go

A simple web application built with [Gin](https://github.com/gin-gonic/gin) and [`go-wkhtmltopdf`](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) that allows users to create invoices via a form and download them as PDFs.

---

## 🔧 Prerequisites

- Go 1.18 or newer  
- `wkhtmltopdf` installed and available in your system PATH

---

## 🚀 Installation & Usage

### Clone Repository

```bash
git clone https://github.com/OsGift/invoicer.git
cd invoice-pdf-generator
Install Dependencies

go mod tidy
Install wkhtmltopdf
Ubuntu / Debian:


sudo apt update
sudo apt install wkhtmltopdf
macOS (Homebrew):


brew install wkhtmltopdf
Windows:

Download from wkhtmltopdf.org and add to your PATH.

Run Application
go run main.go
Visit http://localhost:8080 in your browser.

🧾 Usage Guide
Access the web form at http://localhost:8080.

Enter invoice metadata:

Invoice Number

Invoice Date

Bill To (customer name or company)

Add one or more line items with description, quantity, and price.

Click Generate Invoice PDF.

The browser will automatically download the generated invoice.pdf.

⚙️ How It Works (Overview)
The server exposes a form at /, handled by Gin handlers.

Submitted form data is parsed into a Go struct (InvoiceData and Item).

Basic invoice calculations: subtotal, 10% tax, and total.

A Go HTML template renders the invoice layout.

go-wkhtmltopdf generates a PDF from the HTML and sends it to the client.

🐳 Docker (Optional)
To build and run using Docker, add a Dockerfile like below:

dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o invoice-app

FROM alpine:latest
RUN apk --no-cache add ca-certificates wkhtmltopdf
WORKDIR /root/
COPY --from=builder /app/invoice-app .
COPY templates templates
COPY static static
EXPOSE 8080
CMD ["./invoice-app"]
Build and start:


docker build -t invoice-app .
docker run -p 8080:8080 invoice-app
Then open http://localhost:8080.

📦 Dependencies
Gin: HTTP web framework

go-wkhtmltopdf: Go bindings for wkhtmltopdf

wkhtmltopdf: External CLI for rendering PDFs

🛠 Troubleshooting
wkhtmltopdf errors or not found: Ensure it’s correctly installed and accessible in PATH.

Browser loads but PDF doesn’t download: Check server logs and validate submitted form fields.

Styling issues: Confirm style.css and script.js paths are served under /static.

🎯 License & Author
MIT License

Created by Gift and Os / https://github.com/OsGift/invoicer