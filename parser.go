package main

import (
    "fmt"
    "flag"
)

var (
    WORKERS       int       = 2            // кол-во "потоков"
    REPORT_PERIOD int       = 10           // частота отчетов (сек)
    DUP_TO_STOP   int       = 500          // максимум повторов до останова
    HASH_FILE     string    = "hash.bin"   // файл с хешами
    QUOTES_FILE   string    = "quotes.txt" // файл с цитатами
)

func main() {
    fmt.Println("Hello")

    initFlags()
}

func initFlags() {
    //Задаем правила разбора:
    flag.IntVar(&WORKERS, "w", WORKERS, "количество потоков")
    flag.IntVar(&REPORT_PERIOD, "r", REPORT_PERIOD, "частота отчетов (сек)")
    flag.IntVar(&DUP_TO_STOP, "d", DUP_TO_STOP, "кол-во дубликатов для остановки")
    flag.StringVar(&HASH_FILE, "hf", HASH_FILE, "файл хешей")
    flag.StringVar(&QUOTES_FILE, "qf", QUOTES_FILE, "файл записей")
    //И запускаем разбор аргументов
    flag.Parse()
}