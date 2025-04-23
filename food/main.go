package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, "小明")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

type Product struct {
	Name  string
	Price float64
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		tmpl := template.Must(template.New("index").Parse(`
// 			<!DOCTYPE html>
// 			<html>
// 			<head>
// 				<title>Product List</title>
// 				<style>
// 					.product-grid-container {
// 						overflow-x: auto; /* 横向滚动 */
// 						margin: 20px 0;
// 					}
// 					.product-grid {
// 						display: grid;
// 						grid-template-columns: repeat(6, 1fr); /* 固定6列 */
// 						gap: 20px; /* 列间距 */
// 						min-width: max-content; /* 确保容器宽度足够 */
// 					}
// 					.product-card {
// 						border: 1px solid #ddd;
// 						padding: 15px;
// 						min-width: 150px; /* 卡片最小宽度 */
// 					}
// 				</style>
// 			</head>
// 			<body>
// 				<h1>Products</h1>
// 				<div class="product-grid-container">
// 					<div class="product-grid">
// 						{{range .}}
// 							<div class="product-card">
// 								<h3>{{.Name}}</h3>
// 								<p>Price: ${{.Price}}</p>
// 							</div>
// 						{{end}}
// 					</div>
// 				</div>
// 			</body>
// 			</html>
// 		`))

// 		products := []Product{
// 			{Name: "Laptop", Price: 999.99},
// 			{Name: "Phone", Price: 699.99},
// 			{Name: "Tablet", Price: 399.99},
// 			{Name: "Headphones", Price: 199.99},
// 			{Name: "Keyboard", Price: 49.99},
// 			{Name: "Mouse", Price: 29.99},
// 			{Name: "Monitor", Price: 199.99},
// 			{Name: "Printer", Price: 149.99},
// 			{Name: "Scanner", Price: 99.99},
// 			{Name: "Router", Price: 79.99},
// 			{Name: "SSD", Price: 59.99},
// 		}

// 		tmpl.Execute(w, products)
// 	})

// 	http.ListenAndServe(":8080", nil)
// }
