1. Клонируем себе репозиторий `https://github.com/rarinaferrari/ssh-port-changer.git`
2. Устанавливаем Golang `apt install golang`
3. `go buld ssh.go`
4. `cp ssh-port-changer.service /etc/systemd/system/` Не забудьте сравнить пути в демоне
5. `echo "Port 22" >> /etc/ssh/sshd.config.d/50-cloud-init.conf` Обычно там нет  порта по умолчанию
6. `systemctl daemon-reload`
7. `systemctl start ssh-port-changer.service` 
