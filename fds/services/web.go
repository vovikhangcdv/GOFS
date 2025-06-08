package services

import (
	"html/template"
	"log"
	"net/http"

	"token-monitor/models"

	"gorm.io/gorm"
)

type WebServer struct {
	db *gorm.DB
}

func NewWebServer(db *gorm.DB) *WebServer {
	return &WebServer{db: db}
}

func (s *WebServer) Start(port string) error {
	http.HandleFunc("/", s.handleHome)
	http.HandleFunc("/blacklist", s.handleBlacklist)
	http.HandleFunc("/suspicious", s.handleSuspicious)
	http.HandleFunc("/pending", s.handlePending)

	log.Printf("Starting web server on port %s", port)
	return http.ListenAndServe(":"+port, nil)
}

func (s *WebServer) handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Token Monitor Dashboard</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 20px; }
			.container { max-width: 1200px; margin: 0 auto; }
			.nav { margin-bottom: 20px; }
			.nav a { margin-right: 20px; text-decoration: none; color: #0066cc; }
			table { width: 100%; border-collapse: collapse; margin-top: 20px; }
			th, td { padding: 10px; border: 1px solid #ddd; text-align: left; }
			th { background-color: #f5f5f5; }
			tr:nth-child(even) { background-color: #f9f9f9; }
			.severity-high { color: #cc0000; }
			.severity-medium { color: #cc6600; }
			.severity-low { color: #006600; }
			.status-pending { color: #cc6600; }
			.status-confirmed { color: #006600; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Token Monitor Dashboard</h1>
			<div class="nav">
				<a href="/pending">Pending Transactions</a>
				<a href="/blacklist">Blacklisted Addresses</a>
				<a href="/suspicious">Suspicious Transactions</a>
			</div>
			<p>Welcome to the Token Monitor Dashboard. Use the navigation above to view pending transactions, blacklisted addresses, and suspicious transactions.</p>
		</div>
	</body>
	</html>`

	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func (s *WebServer) handleBlacklist(w http.ResponseWriter, r *http.Request) {
	var addresses []models.BlacklistedAddress
	if err := s.db.Order("created_at DESC").Find(&addresses).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Blacklisted Addresses</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 20px; }
			.container { max-width: 1200px; margin: 0 auto; }
			.nav { margin-bottom: 20px; }
			.nav a { margin-right: 20px; text-decoration: none; color: #0066cc; }
			table { width: 100%; border-collapse: collapse; margin-top: 20px; }
			th, td { padding: 10px; border: 1px solid #ddd; text-align: left; }
			th { background-color: #f5f5f5; }
			tr:nth-child(even) { background-color: #f9f9f9; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Blacklisted Addresses</h1>
			<div class="nav">
				<a href="/">Home</a>
				<a href="/suspicious">Suspicious Transactions</a>
			</div>
			<table>
				<tr>
					<th>Address</th>
					<th>Transaction Hash</th>
					<th>Block Number</th>
					<th>Reason</th>
					<th>Date Added</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.Address}}</td>
					<td>{{.TxHash}}</td>
					<td>{{.BlockNumber}}</td>
					<td>{{.Reason}}</td>
					<td>{{.CreatedAt.Format "2006-01-02 15:04:05"}}</td>
				</tr>
				{{end}}
			</table>
		</div>
	</body>
	</html>`

	t, err := template.New("blacklist").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, addresses)
}

func (s *WebServer) handleSuspicious(w http.ResponseWriter, r *http.Request) {
	var transactions []models.SuspiciousTransfer
	if err := s.db.Order("created_at DESC").Find(&transactions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Suspicious Transactions</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 20px; }
			.container { max-width: 1200px; margin: 0 auto; }
			.nav { margin-bottom: 20px; }
			.nav a { margin-right: 20px; text-decoration: none; color: #0066cc; }
			table { width: 100%; border-collapse: collapse; margin-top: 20px; }
			th, td { padding: 10px; border: 1px solid #ddd; text-align: left; }
			th { background-color: #f5f5f5; }
			tr:nth-child(even) { background-color: #f9f9f9; }
			.severity-high { color: #cc0000; }
			.severity-medium { color: #cc6600; }
			.severity-low { color: #006600; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Suspicious Transactions</h1>
			<div class="nav">
				<a href="/">Home</a>
				<a href="/blacklist">Blacklisted Addresses</a>
			</div>
			<table>
				<tr>
					<th>Transaction Hash</th>
					<th>From</th>
					<th>To</th>
					<th>Amount</th>
					<th>Severity</th>
					<th>Reason</th>
					<th>Date Detected</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.TxHash}}</td>
					<td>{{.From}}</td>
					<td>{{.To}}</td>
					<td>{{.Amount}}</td>
					<td class="severity-{{.Severity}}">{{.Severity}}</td>
					<td>{{.Reason}}</td>
					<td>{{.CreatedAt.Format "2006-01-02 15:04:05"}}</td>
				</tr>
				{{end}}
			</table>
		</div>
	</body>
	</html>`

	t, err := template.New("suspicious").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, transactions)
}

func (s *WebServer) handlePending(w http.ResponseWriter, r *http.Request) {
	var transactions []models.PendingTransaction
	if err := s.db.Where("status = ?", "pending").Order("created_at DESC").Find(&transactions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Pending Transactions</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 20px; }
			.container { max-width: 1200px; margin: 0 auto; }
			.nav { margin-bottom: 20px; }
			.nav a { margin-right: 20px; text-decoration: none; color: #0066cc; }
			table { width: 100%; border-collapse: collapse; margin-top: 20px; }
			th, td { padding: 10px; border: 1px solid #ddd; text-align: left; }
			th { background-color: #f5f5f5; }
			tr:nth-child(even) { background-color: #f9f9f9; }
			.status-pending { color: #cc6600; }
			.status-confirmed { color: #006600; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Pending Transactions</h1>
			<div class="nav">
				<a href="/">Home</a>
				<a href="/blacklist">Blacklisted Addresses</a>
				<a href="/suspicious">Suspicious Transactions</a>
			</div>
			<table>
				<tr>
					<th>Transaction Hash</th>
					<th>From</th>
					<th>To</th>
					<th>Amount</th>
					<th>Status</th>
					<th>Time Pending</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.Hash}}</td>
					<td>{{.From}}</td>
					<td>{{.To}}</td>
					<td>{{.Value}}</td>
					<td class="status-{{.Status}}">{{.Status}}</td>
					<td>{{.CreatedAt.Format "2006-01-02 15:04:05"}}</td>
				</tr>
				{{end}}
			</table>
		</div>
	</body>
	</html>`

	t, err := template.New("pending").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, transactions)
}
