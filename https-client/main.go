package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main1() {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}



func mainhttsclient() {
	pool := x509.NewCertPool()
	caCertPath := "ca.pem"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://127.0.0.1:8080")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main221() {
	client := http.Client{
		// Skip TLS dial
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err := client.Get("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal(fmt.Errorf("error making request: %v", err))
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
}

func mainh2c(){
	client := http.Client{

		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("请求失败: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("读取响应失败: %s", err)
	}

	fmt.Printf("获取响应 %d: %s\n", resp.StatusCode, string(body))


}

func http2client(url string)(res string) {

	tr := &http2.Transport{
	//	AllowHTTP: true, //充许非加密的链接
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	httpClient := http.Client{Transport: tr}

	ctx, cancel := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancel()
	})

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	req = req.WithContext(ctx)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp StatusCode:", resp.StatusCode)
		return
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	res =  string(body)
	return res
}


func main(){

	//for {
	//	a:=http2client("https://localhost:8080")
	//	fmt.Println(a)
	//	time.Sleep(time.Millisecond*10)
	//}
	for ; ;  {
		a,e:=http2Do()
		fmt.Println(a,e)
		time.Sleep(time.Second*5)

	}

	return
	m := make(map[string]struct{},0)
	var mutex sync.Mutex

	//for i:=0;i<1;i++{
		go func() {
			for{
				//a:=http2client("https://localhost:8080")
				a,_:=http2Do()
				mutex.Lock()
				m[a]= struct{}{}
				mutex.Unlock()
				fmt.Println(a,len(m))
				time.Sleep(time.Millisecond*10)
			}
		}()
//	}


	select {

	}


}
func httpDo() (res string,err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil{
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
		// handle error
	}

	//fmt.Println(string(body))
	res = string(body)
	return
}
var  trans *http2.Transport
var httptrans *http.Transport
func init(){
  trans = &http2.Transport{
	TLSClientConfig:            &tls.Config{
		InsecureSkipVerify:true,
	},
	}
	httptrans = &http.Transport{
		TLSClientConfig:&tls.Config{
			InsecureSkipVerify:          true,
		},
	}
	http2.ConfigureTransport(httptrans)
}
func http2Do() (res string,err error) {



	//client := &http.Client{
	//	Transport:trans,
	//}
	client := &http.Client{
		Transport:httptrans,
	}
	req, err := http.NewRequest("GET", "https://127.0.0.1:8080", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
		return
	}

	client.CloseIdleConnections()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

//	client.CloseIdleConnections()
	if err != nil{
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
		// handle error
	}

	//fmt.Println(string(body))
	res = string(body)
	return
}