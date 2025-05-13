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
/* Deskripsi  : F05 - Multithreading Optimization (Half Max CPU / Max CPU - 4)   */
/* PIC F05    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana                   */

package backend;

import (
	"runtime";
	"sync";
)

var (
    taskChan chan func();
    wg       sync.WaitGroup;
    onceInit sync.Once;
)

func Multithreading() {
	onceInit.Do(func() {
		runtime.GOMAXPROCS(max(runtime.NumCPU() / 2 , runtime.NumCPU() - 4));
		workers := runtime.NumCPU();
		taskChan = make(chan func() , workers);
		for i := 0 ; i < workers ; i++ {
			wg.Add(1);
			go func() {
				defer wg.Done();
				for task := range taskChan {
					if (task != nil) {
						task();
					}
				}
			}();
		}
	});
}

func SubmitJob(job func()) {
    if (taskChan == nil) {
		Multithreading();
	}
    taskChan <- job;
}

func WaitJobs() {
	close(taskChan);
	wg.Wait();
	taskChan = nil;
}