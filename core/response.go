package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/caffeines/choto/log"
	"github.com/fatih/color"
)

type response struct {
	Status    int         `json:"-"`
	Title     string      `json:"title,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Errors    error       `json:"errors,omitempty"`
	StartTime time.Time   `json:"-"`
}

// Response will return response instance
func Response() *response {
	r := response{StartTime: time.Now()}
	return &r
}

func printRequest(resp response, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Print(now)
	FgCyan := color.New(color.FgCyan, color.Bold)
	FgCyan.Printf(" - %s ", r.Method)
	fmt.Print(r.RequestURI)
	FgGreen := color.New(color.Bold, color.FgGreen)
	FgRed := color.New(color.Bold, color.FgRed)
	if resp.Status == 200 {
		FgGreen.Printf(" %d", resp.Status)
	} else {
		FgRed.Printf(" %d", resp.Status)
	}
	FgHiYellow := color.New(color.Bold, color.FgHiYellow)
	duration := time.Since(resp.StartTime)
	FgHiYellow.Println(" ~ ", duration)
	respLog := fmt.Sprintf("%s - %s %s %d ~ %v", now, r.Method, r.RequestURI, resp.Status, duration)
	log.Log().Infoln(respLog)
}

// SendResponse ...
func (r response) SendResponse(w http.ResponseWriter, req *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	printRequest(r, req)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		printRequest(response{Status: http.StatusInternalServerError}, req)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return nil
}
