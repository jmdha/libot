package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Engine struct {
	cmd *exec.Cmd
	in  io.WriteCloser
	out io.ReadCloser
}

func NewEngine(str string) Engine {
	return Engine{
		cmd: exec.Command(str),
	}
}

func (engine *Engine) Init() {
	engine.out, _ = engine.cmd.StdoutPipe()
	engine.in, _ = engine.cmd.StdinPipe()
	engine.cmd.Start()
}

func (engine *Engine) WaitForReady() {
	io.WriteString(engine.in, "isready\n")
	scanner := bufio.NewScanner(engine.out)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "readyok") {
			return
		}
	}
}

func (engine *Engine) NewGame() {
	engine.WaitForReady()
	io.WriteString(engine.in, "ucinewgame\n")
}

func (engine *Engine) BestMove(fen string, moves string, wtime int, btime int) string {
	engine.WaitForReady()
	var str string
	if fen == "startpos" {
		str = "position startpos"
	} else {
		str = "position fen " + fen
	}
	if len(moves) > 0 {
		str = str + " moves " + moves
	}
	io.WriteString(engine.in, str+"\n")
	engine.WaitForReady()
	str = fmt.Sprintf("go wtime %d btime %d", wtime, btime)
	io.WriteString(engine.in, str+"\n")
	scanner := bufio.NewScanner(engine.out)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "bestmove") {
			return strings.TrimSpace(strings.Split(scanner.Text(), " ")[1])
		}
	}
	return ""
}
