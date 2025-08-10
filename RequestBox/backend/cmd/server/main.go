package main

// КОНФИГУРАЦИЯ
type Configuration struct {
	Port           string
	LogLevel       string
	DBHost         string
	DBName         string
	DBPort         string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
}

// ЛОГИРОВАНИЕ
type Log struct {
	ID  string
	msg string
}

type Logging struct {
	Error   Log
	Debug   Log
	Info    Log
	Warning Log
	Fatal   Log
}

// СЕРВЕР
type Server struct {
	Config Configuration
}

func NewServer() *Server {
	cfg := Configuration{
		Port:           "1488",
		LogLevel:       "---",
		DBHost:         "---",
		DBName:         "---",
		DBPort:         "---",
		ReadTimeout:    0,
		WriteTimeout:   0,
		MaxHeaderBytes: 0,
	}

	srv := &Server{
		Config: cfg,
	}
	return srv
}

func main() {

	//МАРШРУТИЗАТОР

	//ОБРАБОТЧИК

	//ОТВЕТЧИК
}
