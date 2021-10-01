/*
 * Simple pos with main server
*/
package main

import (
  "os"
  "io"
  "log"
  "fmt"
  "sync"
  "time"
  "bufio"
  "strconv"
  "math/rand"
  "encoding/hex"
  "encoding/json"
  "crypto/sha256"
  "github.com/joho/godotenv"
  "github.com/davecgh/go-spew/spew"
)

type Block struct {
  Index int
  Timestamp string
  BPM int
  Hash string
  PrevHash string
  Validator string
}

///Variales
//a series of  validated blocks
var Blockchain []Block
var tempBlocks []Blockvar
//Handler for incoming blocks for validation
var candidateBlocks = make(chan Block)
//a broadcast winning validator to all nodes
var announcements = make(chan string)
var mutex = &sync.Mutex{}
//keeps track of open validators and balance
var validators = make(map[string]int)

var t time.Now()
//Uility Functions
func CalaculateHash(s string) string{
  hash := sha256.New()
  hash.Write([]byte(s))
  hashed := hash.Sum(nil)
  return hex.EncodeToString(hashed)
}
func CalculateBlockHash(block Block) string{
  record := string(block.Index)+block.Timestamp+string(block.BPM)+ block.PrevHash
  return CalaculateHash(record)
}

//create a new block using old blocks hash
func GenerateBlock(oldBlock Block,BPM int,address string)(Block,error){
  var newBlock Block
  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.BPM = BPM
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash = CalculateBlockHash(newBlock)
  newBlock.Validator = address
  return newBlock,nil
}

func IsBlockValid(newBlock,oldBlock Block) bool{
  if oldBlock.Index + 1 != newBlock.Index {
    return false
  }
  if oldBlock.Hash != newBlock.PrevHash {
    return false
  }
  if CalculateBlockHash(newBlock) != newBlock.Hash {
    return false
  }
  return true
}

func HandleConn(conn net.Conn){
  defer conn.Close()
  go func(){
    for {
      msg := <-announcements
      io.WriteString(conn,msg)
    }
  }()
  var address string//valoidator address
  io.WriteString(conn,"Enter Token Balance: ")
  scanBalance := bufio.NewScanner(conn)
  for scanBalance.Scan() {
    balance,err := strconv.Atoi(scanBalance.Text())
    if err != nil {
      log.Printf("[+] %v not a number: %v",scanBalance.Text(),err)
      return
    }
    address = CalaculateHash(t.String())
    validators[address] = balance
    fmt.Println(validators)
    break
  }
  io.WriteString(conn,"\nEnter New BPM: ")
  scanBPM := bufio.NewScanner(conn)
  go func(){
    for {
      for scanBPM.Scan(){
        bpm,err := strconv.Atoi(scanBPM.Text())

      }
    }
  }()
  //Simulate the receiving broadcast
  for {
    time.Sleep(time.Minute)
    mutex.Lock()
    output,err := json.Marshal(Blockchain)
    mutex.Unlock()
    if err != nil {
      log.Fatal(err)
    }
    io.WriteString(conn,string(output)+"\n")
  }
}
