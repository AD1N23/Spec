// Код сгенерирован на основе диаграммы классов classes.puml.
// Содержит определения типов, интерфейсов и сигнатуры методов,
// которые служат основой для разработки программного решения.
//
// Язык: Go 1.21+
// Проект: Telegram-бот учёта питания

package tgbot

import (
	"sync"
	"time"
)

// ── Config ────────────────────────────────────────────────────────────────────

type BotConfig struct {
	Token string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Config struct {
	Bot BotConfig
	DB  DBConfig
}

// ── Models ────────────────────────────────────────────────────────────────────

type FoodEntry struct {
	ID            int
	ChatID        int64
	FoodName      string
	Calories      int
	Proteins      float64
	Fats          float64
	Carbohydrates float64
	Portion       int
	Date          time.Time
}

type Food struct {
	Calories      uint16
	Proteins      uint16
	Fats          uint16
	Carbohydrates uint16
}

// ── Domain ────────────────────────────────────────────────────────────────────

type UserState string

const (
	StateNone               UserState = ""
	StateWaitingForFoodName UserState = "waiting_for_food_name"
)

// ── Repository ────────────────────────────────────────────────────────────────

// FoodRepository — интерфейс хранилища записей о питании.
// Позволяет подменять реализацию (PostgreSQL, in-memory) без изменения бизнес-логики.
type FoodRepository interface {
	SaveEntry(entry *FoodEntry) error
	GetUserEntries(chatID int64, date *time.Time) ([]*FoodEntry, error)
	RemoveEntry(id int) error
}

// PostgresFoodRepo — продакшн-реализация на PostgreSQL.
type PostgresFoodRepo struct {
	db interface{} // *sql.DB
}

func (r *PostgresFoodRepo) SaveEntry(entry *FoodEntry) error        { panic("not implemented") }
func (r *PostgresFoodRepo) GetUserEntries(chatID int64, date *time.Time) ([]*FoodEntry, error) {
	panic("not implemented")
}
func (r *PostgresFoodRepo) RemoveEntry(id int) error { panic("not implemented") }

// MemoryFoodRepo — in-memory реализация для тестов.
type MemoryFoodRepo struct {
	mu      sync.Mutex
	entries map[int64][]FoodEntry
	nextID  int
}

func (r *MemoryFoodRepo) SaveEntry(entry *FoodEntry) error        { panic("not implemented") }
func (r *MemoryFoodRepo) GetUserEntries(chatID int64, date *time.Time) ([]*FoodEntry, error) {
	panic("not implemented")
}
func (r *MemoryFoodRepo) RemoveEntry(id int) error { panic("not implemented") }

// ── Services ──────────────────────────────────────────────────────────────────

// MessageService — интерфейс отправки и редактирования сообщений в Telegram.
type MessageService interface {
	SendSimpleMessage(chatID int64, text string, keyboard interface{})
	SendSimpleRyplyToMessage(chatID int64, messageID int, text string)
	DeleteMessage(chatID int64, messageID int)
	GenerateDailyReport(entries []*FoodEntry, reportDate time.Time) string
	EditLogMessage(chatID int64, text string, messageID int)
}

// TelegramMessageService — реализация MessageService через Telegram Bot API.
type TelegramMessageService struct {
	bot           interface{} // *tgbotapi.BotAPI
	logMessageIDs map[int64]int
	repo          FoodRepository
}

func (s *TelegramMessageService) SendSimpleMessage(chatID int64, text string, keyboard interface{}) {
	panic("not implemented")
}
func (s *TelegramMessageService) SendSimpleRyplyToMessage(chatID int64, messageID int, text string) {
	panic("not implemented")
}
func (s *TelegramMessageService) DeleteMessage(chatID int64, messageID int) {
	panic("not implemented")
}
func (s *TelegramMessageService) GenerateDailyReport(entries []*FoodEntry, reportDate time.Time) string {
	panic("not implemented")
}
func (s *TelegramMessageService) EditLogMessage(chatID int64, text string, messageID int) {
	panic("not implemented")
}

// KeyboardService — генерация inline-клавиатур Telegram.
type KeyboardService struct {
	keyboard interface{} // tgbotapi.InlineKeyboardMarkup
}

func (s *KeyboardService) StartMenu() interface{}                            { panic("not implemented") }
func (s *KeyboardService) AddEditButton() interface{}                        { panic("not implemented") }
func (s *KeyboardService) GenerateEditKeyboard(entries []*FoodEntry) interface{} {
	panic("not implemented")
}
func (s *KeyboardService) GenerateChoiceKeyboard(entries []FoodEntry) interface{} {
	panic("not implemented")
}

// FoodService — бизнес-логика поиска и расчёта КБЖУ.
type FoodService struct {
	foodData map[string]Food
	repo     FoodRepository
}

func (s *FoodService) PrepareFoodEntry(chatID int64, input string, weight int) ([]FoodEntry, error) {
	panic("not implemented")
}
func (s *FoodService) findBestMatch(input string) (string, int) { panic("not implemented") }
func (s *FoodService) CalculatePortion(weight int, base Food) Food { panic("not implemented") }

// ── Handlers ──────────────────────────────────────────────────────────────────

// Handler — точка входа для всех входящих Update от Telegram.
type Handler struct {
	bot           interface{} // *tgbotapi.BotAPI
	msgService    MessageService
	repo          FoodRepository
	kbService     KeyboardService
	foodService   *FoodService
	userStates    map[int64]UserState
	pendingChoice map[int64][]FoodEntry
}

func (h *Handler) CommandHandler(update interface{})     { panic("not implemented") }
func (h *Handler) CallbackHandler(update interface{})    { panic("not implemented") }
func (h *Handler) HandleTextMessage(update interface{})  { panic("not implemented") }

// ── Parser ────────────────────────────────────────────────────────────────────

// ParseFoodInput разбирает строку вида "Куриная грудка 150"
// на название продукта и вес в граммах.
func ParseFoodInput(text string) (name string, weight int) { panic("not implemented") }

// ── App ───────────────────────────────────────────────────────────────────────

// App — корневой объект приложения, связывает все компоненты.
type App struct {
	db               interface{} // *sql.DB
	bot              interface{} // *tgbotapi.BotAPI
	repo             *PostgresFoodRepo
	msgService       MessageService
	kbService        KeyboardService
	foodService      *FoodService
	userStates       map[int64]UserState
	foodLogMessageID map[int64]int
	handler          *Handler
}

// BotRun запускает главный цикл обработки входящих Update.
func (a *App) BotRun(updates interface{}) { panic("not implemented") }
