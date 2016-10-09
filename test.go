package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //����������Ĭ���ǲ��������
	fmt.Println(r.Form) //��Щ��Ϣ��������������˵Ĵ�ӡ��Ϣ
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
	fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello as
	taxie!") //���д�뵽 w ����������ͻ��˵�
}

func main() {
	http.HandleFunc("/", sayhelloName) //���÷��ʵ�·��
	err := http.ListenAndServe(":9090", nil) //���ü����Ķ˿�
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}