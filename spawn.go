package main
import("os/exec"
       "fmt"
       "encoding/json"
       "io/ioutil"
       "sync")

var wg sync.WaitGroup

func main(){

    data, err := ioutil.ReadFile("/var/openbox-spawn/config.json")

    if err != nil{
        fmt.Println(err)
        return
    }
    var cmds [][]string

    err = json.Unmarshal(data, &cmds)

    if err != nil{
        fmt.Println(err)
        return
    }
/*
    cmds = append(cmds, []string{"/usr/bin/tint2"})
    cmds = append(cmds, []string{"/usr/bin/nm-applet"})
    cmds = append(cmds, []string{"/usr/bin/xset", "-dpms"})
    cmds = append(cmds, []string{"/usr/bin/xset", "s", "off"})
    cmds = append(cmds, []string{"/usr/bin/xsetroot", "-solid", "#0d0d0d"})
*/
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
