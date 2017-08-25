package main
import("os/exec"
       "fmt"
       "sync")

var wg sync.WaitGroup

func main(){
    var cmds [][]string

    cmds = append(cmds, []string{"/usr/bin/tint2"})
    cmds = append(cmds, []string{"/usr/bin/nm-applet"})
    cmds = append(cmds, []string{"/usr/bin/xset", "-dpms"})
    cmds = append(cmds, []string{"/usr/bin/xset", "s", "off"})

    wg.Add(len(cmds))
    for _,v := range cmds{
        go waitExec(v)
    }
    wg.Wait()
}

func waitExec(args []string){
    var cmd exec.Cmd
    if len(args) > 0{
        cmd.Path = args[0]
    }
    cmd.Args = args
    err := cmd.Run()
    if err != nil{
        fmt.Println(err)
    }
    wg.Done()
}
