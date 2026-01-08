package main

import {
    "fmt"
    "os"
    "bufio"
    "strings"
}

type Piece struct {
    Name string
    Color string
}

type Square struct {
    Piece *Piece
}

var board [8][8] string

func initialiseBoard(){
    // set up the initial chess board
    board = [8][8] string{
        {"r", "n", "b", "q", "k", "b", "n", "r"},
        {"p", "p", "p", "p", "p", "p", "p", "p"},
        {"", "", "", "", "", "", "", ""},
        {"", "", "", "", "", "", "", ""},
        {"", "", "", "", "", "", "", ""},
        {"", "", "", "", "", "", "", ""},
        {"P", "P", "P", "P", "P", "P", "P", "P"},
        {"R", "N", "B", "Q", "K", "B", "N", "R"},
    }
}

func updateBoard(move string){
    // Parse the move command into source and target squares
    fromSquare, toSquare := parseMove(move)

    // validate the move according to game rules
    if isValidMove(fromSquare, toSquare){
        executeMove(fromSquare, toSquare)

        // check for special moves, captures, game outcomes, etc.
        handleSpecialMoves()

        // display t feedback to the user
        fmt.Println("Move executed successfually")
        printBoard()
    } else {
        fmt.Println("Invalid move, please try again")
    }
}

func printBoard(){
    // display the current chess board
    for _, row := range board{
        for _, piece := range row{
            if piece == ""{
                fmt.Print(". ")
            } else {
                fmt.Print(piece + " ")
            }
        }
        fmt.Println(
    }
}

func main(){
    initialiseBoard()

    for{
        printBoard()

        fmt.Println("Enter your move (from square to square, e.g. 'e2,e4') : ")

        reader := bufino.NewReader(os.Stdin)
        move, _ := reader.ReadString('\n')
        move = strings.TrimSpace(move)

        if move == "exit" {
            fmt.Println("Exiting the game.")
            break
        }

        updateBoard(move)
    }
}
