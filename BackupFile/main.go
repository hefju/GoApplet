package main

import (
	"fmt"
	"github.com/hefju/GoApplet/BackupFile/config"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {

	files := config.LocalConfig.Copyfiles       //需要复制的文件名
	out_folder := config.LocalConfig.OutputPath //输出的文件夹.暂时不会自动建立.
	if !CheckFileExist(out_folder) {
		log.Fatal("无法找到输出文件夹.")
	}
	fmt.Println("copying...")
	for _, src_path := range files {
		BackupFile(src_path, out_folder)
	}

	fmt.Println("end")
}

//备份文件,默认当前程序的文件夹下,每天生产一个文件夹
func BackupFile(src_path, out_folder string) {
	if CheckFileExist(src_path) { //检查文件是否存在.
		var foldername string = out_folder + "/" + time.Now().AddDate(0, 0, 0).Format("2006_01_02") //创建日期文件夹,每天一个文件夹
		if !CheckFileExist(foldername) {                                                            //检查文件夹是否存在
			err := os.Mkdir(foldername, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}
		filename := filepath.Base(src_path)
		newPath := foldername + "/" + filename //新文件路径
		b, err := CopyFile(newPath, src_path)
		if err != nil {
			fmt.Println("cao:38")
			log.Fatal(err)
		}
		fmt.Println("copy bytes:", b, " ", src_path)
	}
}

//复制文件
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open error.", err)
		return
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer src.Close()
	defer dst.Close()
	return io.Copy(dst, src)
}

//1.检查目标文件或者文件夹是否存在。
func CheckFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
