package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tumblr/tumblr.go"
	"github.com/tumblr/tumblrclient.go"
)

var cli *tumblrclient.Client
func main() {

	cli = client()

	blogs, err := tumblr.GetBlogInfo(cli, "mmqilysm.tumblr.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(blogs.Name)
	//获取头像缩略图
	avatar, err := blogs.GetAvatar()
	if err != nil {
		panic(err)
	}

	fmt.Println(avatar)

	//通过type指定获取的Post类型
	//https://www.tumblr.com/docs/en/api/v2#posts--retrieve-published-posts
	values := make(map[string][]string)
	values["type"] = []string{"photo"}
	values["limit"]=[]string{"20"}
	posts, err := blogs.GetPosts(values)
	if err != nil {
		panic(err)
	}

	allPosts, err := posts.All()
	if err != nil {
		panic(err)
	}

	for _, p := range allPosts {
		buffer := strings.NewReader(p.GetSelf().Body)
		doc, err := goquery.NewDocumentFromReader(buffer)
		if err != nil{
			fmt.Println(err.Error())
			continue
		}

		//获取图片地址和原始尺寸
		img := doc.Find("img")
		for _, g := range img.Nodes{
			fmt.Println(g.Attr)
			for _, a := range g.Attr{
				if a.Key == "src"{
					fmt.Println("src: "+a.Val)
				}
				if a.Key == "data-orig-height"{
					fmt.Println("data-orig-height: "+a.Val)
				}
				if a.Key == "data-orig-width"{
					fmt.Println("data-orig-width: "+a.Val)
				}
			}
		}

		fmt.Println("-------------")
	}
}


func client()*tumblrclient.Client{
	return tumblrclient.NewClient(os.Getenv("KEY"), os.Getenv("SEC"))
}
