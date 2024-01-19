package main
import(
	"time"
        "github.com/faiface/beep"
        "github.com/faiface/beep/mp3"
        "github.com/faiface/beep/speaker"
	"fmt"
	"os"
	"path/filepath"
	"github.com/manifoldco/promptui"
)
func selectfile(dir string)string{
	f, _ := os.Open(dir)
	files, _ := f.Readdir(0)
	items := []string{}
	for _, v := range files {
		if v.IsDir() == true{
		}else if filepath.Ext(v.Name()) == ".mp3"{
		items = append(items,v.Name())
		}
	}
	prompt := promptui.Select{
		Label:"Select file",
		Items:items,
	}
	_,result,err := prompt.Run()
	if err != nil{
		panic(err)
	}
	return dir +"/" + result
}
func playmusic(path string){
	f, err := os.Open(path)
        if err != nil {
		panic(err)
	}
	st, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	defer st.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(st, beep.Callback(func() {
		done <- true
	})))
	<-done
}
func main() {
	result := selectfile("/home/vmyu/Downloads")
	fmt.Println(result)
	playmusic(result)
}
