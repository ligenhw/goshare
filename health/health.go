package health

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Checker interface {
	Check() error
}

type CheckFunc func() error

func (cf CheckFunc) Check() error {
	return cf()
}

type Registry struct {
	mu               sync.RWMutex
	registeredChecks map[string]Checker
}

func NewRegistry() *Registry {
	return &Registry{
		registeredChecks: make(map[string]Checker),
	}
}

var DefaultRegistry *Registry

type Updater interface {
	Checker

	Update(status error)
}

type updater struct {
	mu     sync.RWMutex
	status error
}

func (u *updater) Check() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.status
}

func (u *updater) Update(status error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.status = status
}

func NewStatusUpdater() Updater {
	return &updater{}
}

type thresholdUpdater struct {
	mu        sync.RWMutex
	status    error
	threshold int
	count     int
}

func (t *thresholdUpdater) Check() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.count < t.threshold {
		return nil
	}

	return t.status
}

func (t *thresholdUpdater) Update(status error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if status == nil {
		t.count = 0
	} else {
		t.count++
	}

	t.status = status
}

func NewThresholdStatusUpdater(t int) *thresholdUpdater {
	return &thresholdUpdater{threshold: t}
}

func PeriodicChecker(checker Checker, period time.Duration) Checker {
	u := NewStatusUpdater()
	go func() {
		t := time.NewTicker(period)
		for {
			<-t.C
			u.Update(checker.Check())
		}
	}()

	return u
}

func PeriodicThresholdChecker(checker Checker, period time.Duration, threshold int) Checker {
	u := NewThresholdStatusUpdater(threshold)
	go func() {
		t := time.NewTicker(period)
		for {
			<-t.C
			u.Update(checker.Check())
		}
	}()

	return u
}

func (registry *Registry) CheckStatus() map[string]string {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	statusKeys := make(map[string]string)
	for k, v := range registry.registeredChecks {
		err := v.Check()
		if err != nil {
			statusKeys[k] = err.Error()
		}
	}

	return statusKeys
}

func CheckStatus() map[string]string {
	return DefaultRegistry.CheckStatus()
}

func (registry *Registry) Register(name string, checker Checker) {
	if registry == nil {
		registry = DefaultRegistry
	}
	registry.mu.Lock()
	defer registry.mu.Unlock()
	_, ok := registry.registeredChecks[name]
	if ok {
		panic("Check already exists: " + name)
	}
	registry.registeredChecks[name] = checker
}

func Register(name string, checker Checker) {
	DefaultRegistry.Register(name, checker)
}

func (registry *Registry) RegisterFunc(name string, checker func() error) {
	registry.Register(name, CheckFunc(checker))
}

func RegisterFunc(name string, checker func() error) {
	DefaultRegistry.RegisterFunc(name, checker)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		checks := CheckStatus()
		status := http.StatusOK

		if len(checks) != 0 {
			status = http.StatusServiceUnavailable
		}
		statusResponse(w, r, status, checks)
	} else {
		http.NotFound(w, r)
	}
}

func statusResponse(w http.ResponseWriter, r *http.Request, status int, checks map[string]string) {
	p, err := json.Marshal(checks)
	if err != nil {
		log.Println(err)
		p, err = json.Marshal(struct {
			ServerStatus string `json:"server_error"`
		}{
			ServerStatus: "Could not parse error message",
		})

		status = http.StatusInternalServerError
		if err != nil {
			log.Println(err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(p)))
	w.WriteHeader(status)
	if _, err := w.Write(p); err != nil {
		log.Println(err)
	}
}

func init() {
	DefaultRegistry = NewRegistry()
	http.HandleFunc("/debug/health", StatusHandler)
}
