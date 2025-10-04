# Word Counter in Go  

![Made with Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Beginner Friendly](https://img.shields.io/badge/Level-Beginner-brightgreen?style=for-the-badge)
![Open Source](https://img.shields.io/badge/Open%20Source-%E2%9D%A4-red?style=for-the-badge)

---

## Overview  
The **Word Counter** is a beginner-friendly Go program that counts how many times each word appears in a sentence and shows the **total number of words**.  
It’s perfect for learning about **maps**, **loops**, and **string handling** in Go.  

---

## Concepts Used  
- Variables & Constants  
- For Loops  
- Maps  
- String Manipulation (`strings` package)  
- Input / Output  
- Go Syntax Basics  

---

## How It Works  
1. Takes a sentence input from the user.  
2. Converts the sentence to lowercase for case-insensitive comparison.  
3. Splits it into words using `strings.Fields()`.  
4. Counts each word using a map.  
5. Prints each word’s frequency and the total number of words.  

---

## Example  

**Input:**  
Go is Great and Go is Fun

**Output:**  
Word Frequency:
go : 2
is : 2
great : 1
and : 1
fun : 1

Total Words: 7
