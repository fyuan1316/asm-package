package pkg

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
	"text/template"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	asm_bundle     = "asm/asm-operator-bundle"
	flagger_bundle = "asm/flagger-operator-bundle"
	registry       = "build-harbor.alauda.cn"
)

var UnauthorizedErr = errors.New("401 Unauthorized: registry build-harbor.alauda.cn")

//go:embed template.yaml
var scripts []byte

func render(typ string, params map[string]string) ([]byte, error) {
	var m map[string]string
	err := yaml.Unmarshal(scripts, &m)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("").Parse(m[typ])
	if err != nil {
		return nil, err
	}
	//
	//params := map[string]string{
	//	"Registry":             "build-harbor.alauda.cn",
	//	"User":                 "Jian_Liao",
	//	"Password":             "Asm@1234",
	//	"HelmBin":              "helm3",
	//	"DockerBin":            "docker",
	//	"AsmBundleVersion":     "v3.7-13-ge53b7de",
	//	"FlaggerBundleVersion": "v3.7-3-ga0a14d5",
	//	"Destination":          ".",
	//}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, params)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Download(typ string, params map[string]string) error {
	str, err := render(typ, params)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("/bin/bash", []string{"-c", string(str)}...)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}

func hasEle(s string, list []string) bool {
	for _, ig := range list {
		if s == ig {
			return true
		}
	}
	return false
}
func genImgList(values *ValuesContent, registry string) []string {
	var imageList []string
	ignores := []string{"1.6.5-"}
	var tmp string
	for _, img := range values.Global.Images {
		if hasEle(img.Prefix, ignores) {
			continue
		}
		tmp = fmt.Sprintf("%s/%s:%s", registry, img.Repository, img.Tag)
		imageList = append(imageList, tmp)
	}
	return imageList
}

func DownloadImages(valuesPath, saveTo string) {
	values := loadValues(valuesPath + "/values.yaml")
	imageList := genImgList(values, registry)
	dockerPull(imageList)
	dockerSave(imageList, saveTo+"/asm-images.tar")
}

func loadValues(path string) *ValuesContent {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	v := &ValuesContent{}
	err = yaml.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
	return v
}

func dockerPull(imgList []string) {
	var wg sync.WaitGroup
	for _, img := range imgList {
		wg.Add(1)
		fmt.Printf("pulling %s\n", img)
		img := img
		go func() {
			defer wg.Done()
			_ = doPull(img)
			fmt.Printf("%s pulled\n", img)
		}()
	}
	wg.Wait()
	fmt.Println("all images pulled.")
}
func doPull(img string) error {
	dockerCmd := viper.GetString("dockerCmd")
	var args []string
	args = append(args, "pull")
	args = append(args, img)
	cmd := exec.Command(dockerCmd, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	return nil
}

func dockerSave(imgList []string, saveTo string) {
	dockerCmd := viper.GetString("dockerCmd")
	// docker save -o images.tar postgres:9.6 mongo:3.4
	var args []string
	args = append(args, "save")
	args = append(args, "-o")
	args = append(args, saveTo)
	args = append(args, imgList...)
	cmd := exec.Command(dockerCmd, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		log.Fatalln(err)
	}
	fmt.Printf("all images saved to %s\n.", saveTo)
}

//func DownloadBundle(asmVersion, flaggerVersion string) {
//	str, err := render()
//	if err != nil {
//		log.Fatal(err)
//	}
//	cmd := exec.Command("/bin/bash", []string{"-c", string(str)}...)
//	out, err := cmd.CombinedOutput()
//	fmt.Println(string(out))
//
//	var b = fmt.Sprintf("%s/%s:%s", registry, asm_bundle, asmVersion)
//	if err := downloadBundle(b); err != nil {
//		log.Fatal(err)
//	}
//	b = fmt.Sprintf("%s/%s:%s", registry, flagger_bundle, flaggerVersion)
//	if err := downloadBundle(b); err != nil {
//		log.Fatal(err)
//	}
//}
//func downloadBundle(bundle string) error {
//	command := "helm3"
//	pullArgs := []string{"chart", "pull"}
//	exportArgs := []string{"chart", "export"}
//
//	// pull
//
//	authorized, err := ExecCmd(command, append(pullArgs, bundle)...)
//	if !authorized {
//		fmt.Println("login registry build-harbor.alauda.cn/asm firstly.")
//		return UnauthorizedErr
//	}
//	if err != nil {
//		panic(err)
//	}
//	// export
//	authorized, err = ExecCmd(command, append(exportArgs, bundle)...)
//	if !authorized {
//		fmt.Println("login registry build-harbor.alauda.cn/asm firstly.")
//		return UnauthorizedErr
//	}
//	if err != nil {
//		panic(err)
//	}
//	return nil
//}

//func ExecCmd(command string, args ...string) (authorized bool, err error) {
//	//cmd := exec.Command("/bin/bash", []string{"-c", "docker", "login", "build-harbor.alauda.cn", "-u", "Jian_Liao", "-p", "Asm@1234"}...)
//	cmd := exec.Command("/bin/bash", []string{"-c", "docker login build-harbor.alauda.cn -u Jian_Liao -p Asm@1234"}...)
//	out, err := cmd.CombinedOutput()
//	fmt.Println(string(out))
//	if strings.Contains(string(out), "Login Succeeded") {
//		authorized = true
//	}
//	if !authorized || err != nil {
//		return false, err
//	}
//
//	//cmd = exec.Command(command, args...)
//	//cmd = exec.Command("/bin/bash", []string{"-c", "helm3", "chart", "pull", "build-harbor.alauda.cn/asm/flagger-operator-bundle:v3.7-13-ge53b7de"}...)
//	cmd = exec.Command("/bin/bash", []string{"-c", "helm3 chart pull build-harbor.alauda.cn/asm/flagger-operator-bundle:v3.7-13-ge53b7de"}...)
//
//	cmd.Env = append(cmd.Env, "HELM_EXPERIMENTAL_OCI=1")
//
//	out, err = cmd.CombinedOutput()
//	fmt.Println(string(out))
//	if !strings.Contains(string(out), "401 Unauthorized") {
//		authorized = true
//	}
//	return
//}
