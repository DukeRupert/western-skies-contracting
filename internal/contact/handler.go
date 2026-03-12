package contact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Config struct {
	PostmarkToken string
	FromEmail     string
	ToEmail       string
}

type postmarkEmail struct {
	From     string `json:"From"`
	To       string `json:"To"`
	Subject  string `json:"Subject"`
	TextBody string `json:"TextBody"`
	ReplyTo  string `json:"ReplyTo,omitempty"`
}

func Handler(cfg Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formError)
			return
		}

		name := strings.TrimSpace(r.FormValue("name"))
		phone := strings.TrimSpace(r.FormValue("phone"))
		email := strings.TrimSpace(r.FormValue("email"))
		message := strings.TrimSpace(r.FormValue("message"))

		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, formError)
			return
		}

		body := fmt.Sprintf("Name: %s\nPhone: %s\nEmail: %s\n\nMessage:\n%s", name, phone, email, message)

		pm := postmarkEmail{
			From:     cfg.FromEmail,
			To:       cfg.ToEmail,
			Subject:  fmt.Sprintf("New contact form submission from %s", name),
			TextBody: body,
		}
		if email != "" {
			pm.ReplyTo = email
		}

		payload, err := json.Marshal(pm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formError)
			return
		}

		req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, "https://api.postmarkapp.com/email", bytes.NewReader(payload))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formError)
			return
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Postmark-Server-Token", cfg.PostmarkToken)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, formError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formSuccess)
	}
}

const formSuccess = `<div class="form-message form-message--success">Got it. We'll be in touch within one business day.</div>`

const formError = `<div class="form-message form-message--error">Something went wrong sending that. Try again or call us at 406-546-1018.</div>`
