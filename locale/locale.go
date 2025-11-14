package locale

import "fmt"

// Locale represents a set of messages for a specific language
type Locale struct {
	Language string
	Messages map[string]string
}

// Manager manages the current locale and language switching
type Manager struct {
	currentLocale *Locale
	locales       map[string]*Locale
}

// NewManager creates a locale manager with available languages
func NewManager() *Manager {
	manager := &Manager{
		locales: make(map[string]*Locale),
	}

	// Register available languages
	manager.RegisterLocale("en", English())
	manager.RegisterLocale("ru", Russian())

	// Set default language
	manager.SetLocale("en")

	return manager
}

// RegisterLocale registers a new locale
func (m *Manager) RegisterLocale(code string, locale *Locale) {
	m.locales[code] = locale
}

// SetLocale sets the current locale
func (m *Manager) SetLocale(code string) error {
	locale, exists := m.locales[code]
	if !exists {
		return fmt.Errorf("locale %s not found", code)
	}
	m.currentLocale = locale
	return nil
}

// Get returns a message by key with enhanced fallback
func (m *Manager) Get(key string) string {
	if m.currentLocale == nil {
		return fmt.Sprintf("[[NO_LOCALE:%s]]", key)
	}

	if msg, exists := m.currentLocale.Messages[key]; exists {
		return msg
	}

	// Enhanced fallback with more context
	return fmt.Sprintf("[[MISSING:%s@%s]]", key, m.currentLocale.Language)
}

// GetFormatted returns a formatted message
func (m *Manager) GetFormatted(key string, args ...interface{}) string {
	return fmt.Sprintf(m.Get(key), args...)
}

// AvailableLocales returns a list of available locales
func (m *Manager) AvailableLocales() []string {
	locales := make([]string, 0, len(m.locales))
	for code := range m.locales {
		locales = append(locales, code)
	}
	return locales
}

// CurrentLocale returns the code of current locale
func (m *Manager) CurrentLocale() string {
	return m.currentLocale.Language
}
