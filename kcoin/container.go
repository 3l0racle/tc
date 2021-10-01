package main

import (
  "os"
  "fmt"
  "os/exec"
  "syscall"
)

func main(){
  switch os.Args[1]{
  case "run":
    run()
  case "child":
    child()
  default:
    panic("[-] What the hell dude!!")
  }
}
func run(){
  cmd := exec.Command("/proc/self/exe",append([]string{"child"},os.Args[2:]...)...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
  }
  Must(cmd.Run())
}
func child(){
  fmt.Printf("[+] Running %v as PID %d\n",os.Args[2:],os.Getpid())
  cmd := exec.Command(os.Args[2],os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  Must(syscall.Chroot("/home/rootfs"))
  Must(os.Chdir("/"))
  Must(syscall.Mount("proc","proc","proc",0,""))
  Must(cmd.Run())
}
func Must(err error){
  if err != nil {
      panic(err)
    }
}
