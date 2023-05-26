package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		// Получаем текущее время
		now := time.Now()
		currentTime := now.Format("1504") // Форматируем время в HHMM

		// Удаляем ведущий ноль, если он есть (например, "1001" вместо "10001")
		currentTime = strings.TrimLeft(currentTime, "0")

		// Генерируем новый порт на основе текущего времени
		newPort, err := strconv.Atoi(currentTime)
		if err != nil {
			log.Fatalf("Ошибка при преобразовании порта: %v", err)
		}

		// Изменяем порт SSH
		command := fmt.Sprintf("sudo sed -i 's/Port .*/Port %d/' /etc/ssh/sshd_config.d/50-cloud-init.conf", newPort)
		cmd := exec.Command("bash", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Ошибка при изменении порта SSH: %v", err)
		}

		// Перезапускаем службу SSH для применения изменений
		cmd = exec.Command("sudo", "service", "ssh", "restart")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Ошибка при перезапуске службы SSH: %v", err)
		}

		fmt.Println("Порт SSH успешно изменен.")

		// Ждем 1 минуту перед следующей проверкой времени
		time.Sleep(time.Minute)
	}
}
