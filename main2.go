package main
//
//import (
//	"flag"
//	"fmt"
//	"gopkg.in/yaml.v3"
//	"io/ioutil"
//	"log"
//	"os"
//	"os/exec"
//	"sync"
//)
//
//var (
//	valuesPath string
//	registry   string
//	saveTo     string
//)
//
//func main() {
//	flag.StringVar(&valuesPath, "path", "", "chart values file path")
//	flag.StringVar(&registry, "registry", "build-harbor.alauda.cn", "chart registry")
//	flag.StringVar(&saveTo, "savePath", "/tmp/chart-global-asm-images.tar", "images tar file path")
//	flag.Parse()
//	if len(valuesPath) == 0 {
//		fmt.Println("Usage: main.go -path <chart values file path> -registry <chart registry> -savePath <images tar file path>")
//		flag.PrintDefaults()
//		os.Exit(1)
//	}
//	values := loadValues(valuesPath)
//	imageList := genImgList(values, registry)
//	// imageList = imageList[0:2]
//	dockerPull(imageList)
//	dockerSave(imageList)
//}
//func hasEle(s string, list []string) bool {
//	for _, ig := range list {
//		if s == ig {
//			return true
//		}
//	}
//	return false
//}
//func genImgList(values *ValuesContent, registry string) []string {
//	var imageList []string
//	ignores := []string{"1.6.5-"}
//	var tmp string
//	for _, img := range values.Global.Images {
//		if hasEle(img.Prefix, ignores) {
//			continue
//		}
//		tmp = fmt.Sprintf("%s/%s:%s", registry, img.Repository, img.Tag)
//		imageList = append(imageList, tmp)
//	}
//	return imageList
//}
//
//func loadValues(path string) *ValuesContent {
//	data, err := ioutil.ReadFile(path)
//	if err != nil {
//		panic(err)
//	}
//	v := &ValuesContent{}
//	err = yaml.Unmarshal(data, v)
//	if err != nil {
//		panic(err)
//	}
//	return v
//}
//
//func dockerPull(imgList []string) {
//	var wg sync.WaitGroup
//	for _, img := range imgList {
//		wg.Add(1)
//		fmt.Printf("pulling %s\n", img)
//		img := img
//		go func() {
//			defer wg.Done()
//			_ = doPull(img)
//			fmt.Printf("%s pulled\n", img)
//		}()
//	}
//	wg.Wait()
//	fmt.Println("all images pulled.")
//}
//func doPull(img string) error {
//	var args []string
//	args = append(args, "pull")
//	args = append(args, img)
//	cmd := exec.Command("docker", args...)
//	out, err := cmd.CombinedOutput()
//	if err != nil {
//		fmt.Println(out)
//		return err
//	}
//	return nil
//}
//
//func dockerSave(imgList []string) {
//	// docker save -o images.tar postgres:9.6 mongo:3.4
//	var args []string
//	args = append(args, "save")
//	args = append(args, "-o")
//	args = append(args, saveTo)
//	args = append(args, imgList...)
//	cmd := exec.Command("docker", args...)
//	out, err := cmd.CombinedOutput()
//	if err != nil {
//		fmt.Println(out)
//		log.Fatalln(err)
//	}
//	fmt.Printf("all images saved to %s\n.", saveTo)
//}
