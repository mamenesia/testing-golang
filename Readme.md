<h1 align='center'>Testing Golang</h1>

<p align='center'>
    <img width="400" src='https://www.vertica.com/wp-content/uploads/2019/07/Golang.png' />
</p>

<br>
<br>

## Introduction
[![Golang](https://img.shields.io/badge/Go-v1.13-blue)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-v6.60-orange)]()


<p align='justify'>This repository is for answering the questions for hiring process</p>
<p align='justify'><strong>Note : </strong>I had no experience with Go before, so I am asking for a privilege for additional 3 days to learning Go first.</p>


## Requirements
1. <a href="https://golang.org/">Golang</a>

## How to Start
1. Download this Project or you can type ``` git clone https://github.com/mamenesia/testing-golang.git ```
2. Open app's directory in CMD or Terminal
3. Folder `0. Count` for solution of question no. 0.
   Before run this program, make sure CMD or terminal is in this folder directory.
   To run the program, write in CMD or terminal `go run main.go count.go`
   To test the program, write in CMD or terminal `go test` 
4. Folder `1. Sort` for solution of question no. 1.
   Before run this program, make sure CMD or terminal is in this folder directory.
   To run the program, write in CMD or terminal `go run main.go sort.go`
   To test the program, write in CMD or terminal `go test` 
5. Folder `2. Webservice` for solution of question no. 2. 
   Before run this program, make sure CMD or terminal is in this folder directory.
   To run the program, write in CMD or terminal `go run main.go handler.go count.go sort.go`, the server runs in `localhost:9000`.
   To test the program, write in CMD or terminal `go test`
   This webservice has 3 routes:  Index `"/"`, Count `"/count"`, and Sort `"/sort"` 
6. Folder `3. ServiceDocker` for solution of question no. 3


## Screenshot
### Folder 0. Count
<p align='center'>
  <span>
      <p>Running the program</p>
      <image width="500" src="./screenshot/count_run.png" />
      <p>Testing the program</p>
      <image width="500" src="./screenshot/count_test.png" />   
  </span>
</p>

### Folder 1. Sort
<p align='center'>
  <span>
      <p>Running the program</p>
      <image width="500" src="./screenshot/sort_run.png" />
      <p>Testing the program</p>
      <image width="500" src="./screenshot/sort_test.png" />   
  </span>
</p>

### Folder 2. Webservice
<p align='center'>
  <span>
      <p>Running the program</p>
      <image width="500" src="./screenshot/webservice_run.png" />
      <p>Testing the program with Postman</p>
      <p>- Route Home/Index, method GET</p>
      <image width="800" src="./screenshot/webservice_getIndex.png" />   
      <p>- Route Count, method GET</p>
      <image width="800" src="./screenshot/webservice_getCount.png" />   
      <p>- Route Count, method POST with input = "osama"</p>
      <image width="800" src="./screenshot/webservice_postCount.png" />   
      <p>- Route Sort, method GET</p>
      <image width="800" src="./screenshot/webservice_getSort.png" />   
      <p>- Route Sort, method POST with input = "osama"</p>
      <image width="800" src="./screenshot/webservice_postSort.png" />   
  </span>
</p>