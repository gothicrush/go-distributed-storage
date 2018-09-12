package objects

import (
	"io"
	"net/http"
	"fmt"
	"os"
	"strings"
)

// 处理器
func Handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method


	// 判断是否通过 GET 方法调用 get 接口
	if method == http.MethodGet {
		get(w, r)
		return
	}

	// 判断是否通过 PUT 方法调用 put 接口
	if method == http.MethodPut {
		put(w, r)
		return
	}

	// 当调用方法既不是 GET 又不是 PUT 时，报错

	w.WriteHeader(http.StatusMethodNotAllowed)
}

// PUT 方法
func put(w http.ResponseWriter, r *http.Request) {

	// 我们要求<object_name>部分不能含有 '/'，如果含有 '/'，则'/'之后部分会被丢弃
	// 真实环境中，<object_name>在放入URL之前，都已经经过了转义，即<object_name>不含有 '/'

	// 使用 r.URL对象的 EscapePath 获取 url 转义后的字符串
	// 然后对字符串进行切割，取出 <object_name>
	object_name := strings.Split(r.URL.EscapedPath(), "/")[2]

	// 创建新的文件
	file, err := os.Create("./objects_storage/" + object_name)

	// 如果创建文件失败
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError) // 则返回服务器错误
		return
	}

	// 如果文件创建成功，则延时关闭
	defer file.Close()

	io.Copy(file, r.Body) // 将 PUT 进来的新数据进行保存
}

// GET 方法
func get(w http.ResponseWriter, r *http.Request) {

	// 我们要求<object_name>部分不能含有 '/'，如果含有 '/'，则'/'之后部分会被丢弃
	// 真实环境中，<object_name>在放入URL之前，都已经经过了转义，即<object_name>不含有 '/'

	// 使用 r.URL对象的 EscapePath 获取 url 转义后的字符串
	// 然后对字符串进行切割，取出 <object_name>
	object_name := strings.Split(r.URL.EscapedPath(), "/")[2]

	file, err := os.Open("./objects_storage/" + object_name)

	// 如果无法找对对应的资源
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound) // 返回找不到资源
		return 
	}

	defer file.Close() // 如果找到资源，则延时关闭

	io.Copy(w, file)
}