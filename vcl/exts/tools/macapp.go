package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	infoplist = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>zh_CN</string>
	<key>CFBundleExecutable</key>
	<string>%s</string>
	<key>CFBundleName</key>
	<string>%s</string>
	<key>CFBundleIdentifier</key>
	<string>ying32.%s</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleSignature</key>
	<string>test</string>
	<key>CFBundleShortVersionString</key>
	<string>0.1</string>
	<key>CFBundleVersion</key>
	<string>1</string>
	<key>CSResourcesFileMapped</key>
	<true/>
	<key>CFBundleIconFile</key>
	<string>%s.icns</string>
	<key>CFBundleDocumentTypes</key>
	<array>
		<dict>
			<key>CFBundleTypeRole</key>
			<string>Viewer</string>
			<key>CFBundleTypeExtensions</key>
			<array>
				<string>*</string>
			</array>
			<key>CFBundleTypeOSTypes</key>
			<array>
				<string>fold</string>
				<string>disk</string>
				<string>****</string>
			</array>
		</dict>
	</array>
	<key>NSHighResolutionCapable</key>
	<true/>
    <key>NSHumanReadableCopyright</key>
	<string>copyright 2017-2018 ying32.com</string>
</dict>
</plist>`
)

var (
	pkgInfo = []byte{0x41, 0x50, 0x50, 0x4C, 0x3F, 0x3F, 0x3F, 0x3F, 0x0D, 0x0A}
)

func copyFile(src, dest string) error {
	filedest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer filedest.Close()
	filesrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer filesrc.Close()
	_, err = io.Copy(filedest, filesrc)
	return err
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func getdylib() string {
	env := os.Getenv("GOPATH")
	if env == "" {
		return ""
	}
	for _, s := range strings.Split(env, ":") {
		s += "/bin/liblcl.dylib"
		if fileExists(s) {
			return s
		}
	}
	return ""
}

func createExecLink(srcFileName, linkFileName string) error {
	if fileExists(linkFileName) {
		return nil
	}
	cmd := exec.Command("ln", "-s", srcFileName, linkFileName)
	return cmd.Run()
}

// 以一个Mac下的app形式运行
// 调试下使用，正式发布的时候虽然可 以不用去掉，但也不咋好
func RunWithMacOSApp() error {
	if runtime.GOOS == "darwin" {
		if len(os.Args) == 1 {
			if strings.Contains(os.Args[0], "Contents/MacOS") {
				return fmt.Errorf("_")
			}

			execName := os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
			macContentsDir := os.Args[0] + ".app/Contents"
			macOSDir := macContentsDir + "/MacOS"
			macResources := macContentsDir + "/Resources"
			execFile := macOSDir + "/" + execName
			if !fileExists(macOSDir) {
				if err := os.MkdirAll(macOSDir, 0755); err != nil {
					return err
				}
			}

			if !fileExists(macResources) {
				os.MkdirAll(macResources, 0755)
			}

			liblclFileName := macOSDir + "/liblcl.dylib"
			if !fileExists(liblclFileName) {
				libFileName := getdylib()
				if fileExists(libFileName) {
					copyFile(libFileName, liblclFileName)
				}
			}

			plistFileName := macContentsDir + "/Info.plist"
			if !fileExists(plistFileName) {
				ioutil.WriteFile(plistFileName, []byte(fmt.Sprintf(infoplist, execName, execName, execName, execName)), 0666)
			}

			pkgInfoFileName := macContentsDir + "/PkgInfo"
			if !fileExists(pkgInfoFileName) {
				ioutil.WriteFile(pkgInfoFileName, pkgInfo, 0666)
			}

			if copyFile(os.Args[0], execFile) == nil {
				//			if createExecLink(os.Args[0], execFile) == nil {
				os.Chmod(execFile, 0755)
				cmd := exec.Command("open", os.Args[0]+".app")
				fmt.Println(cmd.Run())
				os.Exit(0)
			}
		}
	}
	return nil
}
