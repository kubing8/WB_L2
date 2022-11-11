package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

// CommandI - интерфес команды (Все команды должны реализовывоть метод Exec( ...string))
type CommandI interface {
	Exec(args ...string) ([]byte, error)
}

// echoCmd выполняет unix-команду echo и возвращает результат в байтах
type echoCmd struct {
}

func (e *echoCmd) Exec(args ...string) ([]byte, error) {
	return []byte(strings.Join(args, " ")), nil
}

// cdCmd - изменяет директорию
type cdCmd struct {
}

func (c *cdCmd) Exec(args ...string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return []byte(dir), nil
}

// pwdCmd - выводит путь директории в которой находится терминал
type pwdCmd struct {
}

func (p *pwdCmd) Exec(args ...string) ([]byte, error) {
	dir, err := os.Getwd()
	return []byte(dir), err
}

// killCmd - убивает запущенный процесс
type killCmd struct {
}

func (k *killCmd) Exec(args ...string) ([]byte, error) {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = process.Kill()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return []byte("killed"), nil
}

// psCmd - Выводит работающие процессы
type psCmd struct {
}

func (p *psCmd) Exec(args ...string) ([]byte, error) {
	//fmt.Println("LOG: ", exec.Command("tasklist"))
	return exec.Command("tasklist").Output()
}

// Shell - UNIX-шелл-утилита с поддержкой ряда простейших команд
type Shell struct {
	command CommandI
	output  io.Writer
}

// SetCommand - метод установки выбранной команды
func (s *Shell) SetCommand(cmd CommandI) {
	s.command = cmd
}

// Run - выполнение конкретной команды
func (s *Shell) run(args ...string) {
	b, err := s.command.Exec(args...)
	//fmt.Println("LOG: ", b)
	_, err = fmt.Fprintln(s.output, string(b))
	if err != nil {
		fmt.Println("[err]", err.Error())
		return
	}
}

// ExecuteCommands Исполняет команды, которые ввел пользователь
func (s *Shell) ExecuteCommands(cmds []string) {
	for _, command := range cmds {
		args := strings.Split(command, " ")

		com := args[0]
		com = strings.ToLower(com)
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case "echo":
			cmd := &echoCmd{}
			s.SetCommand(cmd)

		case "cd":
			cmd := &cdCmd{}
			s.SetCommand(cmd)

		case "kill":
			cmd := &killCmd{}
			s.SetCommand(cmd)

		case "pwd":
			cmd := &pwdCmd{}
			s.SetCommand(cmd)

		case "ps":
			cmd := &psCmd{}
			s.SetCommand(cmd)

		case "quit", "exit": // завершение программы
			_, err := fmt.Fprintln(s.output, "Stop program...")
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
			os.Exit(0)

		default:
			fmt.Println("Команда еще в разработке")
			continue
		}

		s.run(args...)
	}
}

func main() {
	// Читает из стандартного ввода
	scan := bufio.NewScanner(os.Stdin)

	// устанавливается общий вывод результата команд
	var output = os.Stdout

	shell := &Shell{output: output}
	for {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("LOG: error os path")
			os.Exit(1)
		}
		fmt.Printf("%v>", dir)

		if scan.Scan() {
			line := scan.Text()
			cmds := strings.Split(line, " | ")

			shell.ExecuteCommands(cmds)
		}
	}
}
