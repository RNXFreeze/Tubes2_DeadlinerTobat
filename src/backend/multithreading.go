/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/multithreading.go              */
/* Deskripsi  : F05 - Multithreading Optimization (Half Max CPU || Max CPU - 4)  */
/* PIC F05    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana                   */

package backend;

import "runtime";

func EnableMultithreading() {
	runtime.GOMAXPROCS(max(runtime.NumCPU() / 2 , runtime.NumCPU() - 4));
}

func DisableMultithreading() {
    runtime.GOMAXPROCS(1);
}