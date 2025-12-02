/*
Based on game_client.c by Maciej Gębala


*/

package main

import (
	"fmt"
	"net"
	"os"
	"player/eval"
	"strconv"
	"strings"
)

func main() {
	var evaluationMode int = 1
	if len(os.Args) == 7 {
		evaluationMode, _ = strconv.Atoi(os.Args[6])
		fmt.Println("Debug mode enabled: evaluation mode", evaluationMode)
	} else if len(os.Args) != 6 {
		fmt.Println("Wrong number of arguments")
		os.Exit(1)
	}

	eval.SetEvaluationMode(evaluationMode)

	serverIP := os.Args[1]
	serverPort := os.Args[2]
	playerStr := os.Args[3]
	gameID := os.Args[4]
	depth := os.Args[5]
	// Parse player number
	player, err := strconv.Atoi(playerStr)
	if err != nil {
		fmt.Printf("Invalid player number: %v\n", err)
		os.Exit(1)
	}

	depthInt, err := strconv.Atoi(depth)
	if err != nil {
		fmt.Printf("Invalid depth: %v\n", err)
		os.Exit(1)
	}

	// Connect to server
	conn, err := net.Dial("tcp", net.JoinHostPort(serverIP, serverPort))
	if err != nil {
		fmt.Printf("Unable to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected with server successfully")

	// Receive initial server message
	buffer := make([]byte, 16)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error while receiving server's message: %v\n", err)
		os.Exit(1)
	}
	serverMessage := strings.TrimSpace(string(buffer[:n]))
	fmt.Printf("Received: %s\n", serverMessage)

	// Send player information
	playerMessage := fmt.Sprintf("%s %s", playerStr, gameID)
	_, err = conn.Write([]byte(playerMessage))
	if err != nil {
		fmt.Printf("Unable to send message: %v\n", err)
		os.Exit(1)
	}

	var board eval.Board
	// Initialize board
	board.SetEmptyBoard()
	endGame := false

	// Main game loop
	for !endGame {
		// Receive server message
		buffer = make([]byte, 16)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error while receiving server's message: %v\n", err)
			os.Exit(1)
		}

		serverMessage = strings.TrimSpace(string(buffer[:n]))
		msg, err := strconv.Atoi(serverMessage)
		if err != nil {
			fmt.Printf("Error parsing server message: %v\n", err)
			continue
		}

		// Extract move and message code
		move := msg % 100
		msgCode := msg / 100

		// Update board with opponent's move
		if move != 0 {
			board.SetMove(move, 3-player)
		}

		// Handle message codes
		if msgCode == 0 || msgCode == 6 {
			myMove := eval.Move(board, depthInt, player)
			board.SetMove(myMove, player)

			// Send move to server
			moveStr := strconv.Itoa(myMove)
			_, err = conn.Write([]byte(moveStr))
			if err != nil {
				fmt.Printf("Unable to send message: %v\n", err)
				os.Exit(1)
			}
		} else {
			// Game ended
			endGame = true
			switch msgCode {
			case 1:
				fmt.Println("You won.")
			case 2:
				fmt.Println("You lost.")
			case 3:
				fmt.Println("Draw.")
			case 4:
				fmt.Println("You won. Opponent error.")
			case 5:
				fmt.Println("You lost. Your error.")
			}
		}
	}
}
