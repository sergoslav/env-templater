package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Получаем имена файлов и префикс из аргументов командной строки
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run env-templater.go <source_file> <dest_file> <prefix>")
		os.Exit(1)
	}

	sourceFile := os.Args[1]
	destFile := os.Args[2]
	prefix := os.Args[3] + "-"

	// Открываем исходный файл
	source, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	// Создаем или перезаписываем новый файл
	dest, err := os.Create(destFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	// Проходим по каждой строке исходного файла и записываем измененные строки в новый файл
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, prefix) {
			line = strings.TrimPrefix(line, prefix)
		}
		fmt.Fprintln(dest, line)
	}

	// Проверяем наличие ошибок при сканировании исходного файла
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Env generated. prefix:`" + prefix + "`")
}
