package main
import("os/exec"
       "os/user"
       "fmt"
       "strings"
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
    var globalCmds [][]string

    err = json.Unmarshal(data, &globalCmds)

    if err != nil{
        fmt.Println(err)
        return
    }

    usr, err := user.Current()

    if err != nil{
        fmt.Println(err)
        return
    }
    wg.Add(len(globalCmds))

    var userCmds [][]string
    if len(usr.HomeDir) > 0{
        homeFolder := strings.TrimRight(usr.HomeDir, "/")
        data, err = ioutil.ReadFile(homeFolder + "/.config/openbox-spawn.json")
        if err != nil{
            fmt.Println(err)
        }else{
            err = json.Unmarshal(data, &userCmds)
            if err != nil{
                fmt.Println(err)
                return
            }
        }
        wg.Add(len(userCmds))
    }

    for _,v := range globalCmds{
        go waitExec(v)
    }
    for _,v := range userCmds{
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
