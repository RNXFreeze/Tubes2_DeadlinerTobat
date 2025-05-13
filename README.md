# Tubes2_DeadlinerTobat
IF2211 - Strategi Algoritma - Tugas Besar (Tubes) 2 - Tahun 2024/2025

## About The Project
<p align = "justify">
Program ini adalah program implementasi untuk menyelesaikan permainan sederhana recipe finder game Little Alchemy 2 menggunakan algoritma pencarian BFS & DFS & Bidirectional secara murni tanpa heuristik apapun selain optimalisasi. Permainan ini dilakukan dengan mencari sebuah rute yang dapat membuat elemen yang didapatkan dengan menggabungkan berbagai elemen yang awal mulanya dari 4 elemen dasar (water, air, fire, dan earth). Program ini ditulis dengan bahasa Golang untuk backend, dan dengan menggunakan bahasa Javascript dengan framework Next.js untuk frontend.</p>

## About The Algorithm
<p align = "justify">
Dalam program ini, terdapat 3 algoritma pencarian yang digunakan, yaitu algoritma BFS (Breadth-First Search), DFS (Depth-First Search), dan BDR (Bidirectional Search). Algoritma BFS merupakan algoritma yang melakukan pencarian secara melebar terlebih dahulu yang diimplementasikan menggunakan struktur data queue dengan level-order traversal dan optimalisasi signature untuk mencegah duplikasi. Algoritma DFS merupakan algoritma yang melakukan pencarian secara mendalam terlebih dahulu yang diimplementasikan menggunakan struktur data stack dengan backtracking dan optimalisasi memoization dan cycle-safe untuk mencegah duplikasi dan mempercepat pencarian. Algoritma BDR merupakan algoritma yang melakukan pencarian secara enumerasi dua arah dari top dan bottom dengan tujuan bertemu di tengah (mid-tier) yang dapat fleksibel yang diimplementasikan menggunakan struktur data map dengan pendekatan meet-in-the-middle dan optimalisasi memoization dan signature untuk mencegah duplikasi dan mempercepat pencarian.</p>

## Project Feature
- Recipe Finder Game Little Alchemy 2
- Single & Multiple Recipe Result With Node Counter
- Algoritma BFS & DFS & Bidirectional (BDR)
- Terminal CLI Interaktif & Website Visual Menarik
- Optimalisasi Program Dengan Multithreading
- Live Update Searching Tree Algorithm
- Containerize With Docker
- Deployed On Internet

## Program Requirements
- Version Go v1.24.3+
- Version Node JS v22.14.0+
- Version NPM v11.3.0+

## How To Clone
Untuk clone, langsung buka terminal dan masukkan 1 command berikut:
```bash
git clone https://github.com/RNXFreeze/Tubes2_DeadlinerTobat.git
```

## How To Compile - Setup go.mod & go.sum
Untuk compile, langsung buka terminal dan masukkan 3 command berikut:
```bash
go mod init github.com/RNXFreeze/Tubes2_DeadlinerTobat
go mod tidy
go build ./...
```

## How To Run - CLI Terminal
Untuk run, langsung buka terminal dan masukkan 2 command berikut:
```bash
cd src
go run main.go
```

## How To Run - Website Local
Untuk run, langsung buka terminal 1 dan masukkan 2 command berikut:
```bash
cd src
go run api/api.go
```
Lalu, langsung buka terminal 2 dan masukkan 3 command berikut:
```bash
cd src/frontend
npm install
npm run dev
```
Lalu, klik link local host yang muncul di terminal seperti berikut:
```bash
https://localhost:3000/
```

## How To Compile & Run - Website Public
Untuk run, langsung buka website dengan link berikut:
```bash
https://DeadlinerTobat_LittleAlchemy2RecipeFinder.com/
```

## Author
Author 1 : Muhammad Raihan Nazhim Oktana (K01 / 13523021)
<br>
Author 2 : Mayla Yaffa Ludmilla (K01 / 13523050)
<br>
Author 3 : Anella Utari Gunadi (K02 / 13523078)
<br>
Instansi : Teknik Informatika (IF-G) ~ Institut Teknologi Bandung (ITB)
