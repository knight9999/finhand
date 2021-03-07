package finhand

import (
	"sort"
	"sync"

	"github.com/google/uuid"
)

const (
	// Version is package version
	Version = "0.1.0"
)

// Handler is Finalizer Handler Object
type Handler struct {
	name     string
	priority int
	f        func()
}

var (
	handlersLock sync.RWMutex
	handlers     map[string]*Handler
	once         sync.Once
)

// OnceInit initialize the handlers.
func OnceInit() {
	once.Do(ClearHandlers)
}

// ClearHandlers remove all handlers.
func ClearHandlers() {
	handlers = map[string]*Handler{}
}

// GetHandler gets the registered Handler
func GetHandler(name string) (*Handler, bool) {
	OnceInit()
	handlersLock.Lock()
	defer handlersLock.Unlock()
	hand, ok := handlers[name]
	return hand, ok
}

// AddHandler is adding new hander.
func AddHandler(f func()) string {
	return AddHandlerWithPriority(0, f)
}

// AddHandlerWithPriority is adding new hander with priority
func AddHandlerWithPriority(priority int, f func()) string {
	var name string
	for {
		uuid, _ := uuid.NewUUID()
		name = uuid.String()
		if _, ok := GetHandler(name); !ok {
			break
		}
	}
	AddHandlerWithNameAndPriority(name, priority, f)
	return name
}

// AddHandlerWithNameAndPriority is adding new hander. This will override any handler with the same name if it exists.
func AddHandlerWithNameAndPriority(name string, priority int, f func()) {
	OnceInit()
	handlersLock.Lock()
	defer handlersLock.Unlock()
	hand := Handler{name: name, priority: priority, f: f}
	handlers[name] = &hand
}

// RemoveHandler is removing handler. Do nothing if the specified handler does not exist.
func RemoveHandler(name string) {
	OnceInit()
	handlersLock.Lock()
	defer handlersLock.Unlock()
	delete(handlers, name)
}

// ExecuteHandlers is execute all handler and clear all handlers
func ExecuteHandlers() {
	OnceInit()
	handlersLock.Lock()
	defer handlersLock.Unlock()
	values := []*Handler{}
	for _, v := range handlers {
		values = append(values, v)
	}
	sort.SliceStable(values, func(i int, j int) bool {
		return values[i].priority > values[j].priority
	})
	for len(values) > 0 {
		hand := values[0]
		hand.f()
		values = values[1:]
	}
}

// RunHandlers is execute all handlers once.
func RunHandlers() {
	ExecuteHandlers()
	ClearHandlers()
}
