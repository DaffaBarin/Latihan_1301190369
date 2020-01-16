package main
import(
	 "fmt"
	 "os"
)
type (
	pengunjung struct {
		plat string
		rodaberapa string
		waktu clock
		hari calender

	}
	clock struct {
		hour int
		minute int
		second int
	}
	calender struct {
		date string
		month string
		year string
	}
	petak struct{
		T [4] int
		N [4] int
		M [4] int
	}
) 
var vehicle petak
var datamotor [10]pengunjung
var datamobil [10]pengunjung
var datatruk [10]pengunjung
var nmotor,nmobil,ntruk int
var in,out,earn int
var berapa int
var brick,quit bool
var kendaraan,nomorkendaraan  string
var cariplat int
func main() {
	var user string
	var ngapainuser,ngapainadmin string
	

	fmt.Println("Selamat datang di parkiran Mall Telkom.")
	fmt.Println("Parkir gratis pada 10 menit pertam.")
	fmt.Println("Parkir motor seharga 1500/jam, mobil 3000/jam, truk 5000/jam")
	fmt.Println("Ukuran parkir motor sebanyak 4 petak dan mobil/truk sebanyak 8 petak.")
	fmt.Println("Motor dan mobil memerlukan 1 petak, sedangkan truk memerlukan 2 petak.")
	fmt.Println("Apakah anda admin?")
	fmt.Println("Tekan '1' jika anda admin.")
	fmt.Println("Tekan '2' jika bukan.")
	fmt.Println("Tekan berapapun untuk keluar program.")
	fmt.Scanln(&user)
	for user == "1" || user == "2" {
		berapa=0
		if user == "1" {
			fmt.Println("Tekan '1' untuk mencetak kartu parkir kendaraan berdasarkan nomor kendaraan.")
			fmt.Println("Tekan '2' untuk melakukan tinjau lahan parkir dan mengurutkan kendaraan berdasarkan nomor kendaraan.")
			fmt.Println("Tekan '3' untuk mencetak statistik hari ini.")
			fmt.Println("Tekan '4' untuk mencetak pendapatan harian.")
			fmt.Scanln(&ngapainadmin)
			if ngapainadmin == "1" {
				//search kendaraan
				searchparkir()							
			}else if ngapainadmin == "2" {
				//sorting motor
				sortparkir()
			}else if ngapainadmin == "3" {
				okupansi()
			}else if ngapainadmin == "4" {
				pendapatan()
				
			}
			fmt.Println("Selamat datang di parkiran Mall Telkom.")
			fmt.Println("Apakah anda admin?")
			fmt.Println("Tekan '1' jika anda admin.")
			fmt.Println("Tekan '2' jika bukan.")
			fmt.Println("Tekan berapapun untuk keluar program.")
			fmt.Scanln(&user)

		}else if user == "2" {
				fmt.Println("Tekan '1' untuk masuk parkir")
				fmt.Println("Tekan '2' untuk keluar parkir")
				fmt.Scanln(&ngapainuser)
				if ngapainuser == "1" {
					masukparkir()
				}else if ngapainuser == "2" {
					keluarparkir()
				}	
				fmt.Println("Selamat datang di parkiran Mall Telkom.")
				fmt.Println("Apakah anda admin?")
				fmt.Println("Tekan '1' jika anda admin.")
				fmt.Println("Tekan '2' jika bukan.")
				fmt.Println("Tekan berapapun untuk keluar program.")
				fmt.Scanln(&user)
		}	
	}
	fmt.Print("Terimakasih telah memakai aplikasi Daffa!")
}


func searchparkir() {
	//fungsi 
	quit := false
	fmt.Println("Jenis kendaraan? motor mobil truk")
	fmt.Scanln(&kendaraan)
	for kendaraan != "motor" && kendaraan != "mobil" && kendaraan != "truk" {
		fmt.Println("Masukkan kendaraan yang valid.")
		fmt.Scanln(&kendaraan)
	}
	fmt.Println("Nomor kendaraan? Angka positif.")
	fmt.Scanln(&nomorkendaraan)
	for nomorkendaraan < "0" {
		fmt.Println("Masukan tidak valid.")
		fmt.Println("Nomor kendaraan?")
		fmt.Scanln(&nomorkendaraan)
	}
	if kendaraan == "motor" {
		for quit == false {
			if datamotor[berapa].plat == nomorkendaraan {
				fmt.Println("=========================================================")
				fmt.Println("		Ini karcis anda!")
				fmt.Println("		Jenis kendaraan:",datamotor[berapa].rodaberapa)
				fmt.Println("		Nomor kendaraan:",datamotor[berapa].plat)
				fmt.Println("		Tanggal masuk:", datamotor[berapa].hari.date," ", datamotor[berapa].hari.month, " ", datamotor[berapa].hari.year)
				fmt.Println("		Jam masuk:",datamotor[berapa].waktu.hour," ",datamotor[berapa].waktu.minute," ",datamotor[berapa].waktu.second)
				fmt.Println("=========================================================")
				quit = true
			}
			berapa++
			if berapa == 3 && datamotor[berapa].plat != nomorkendaraan {
				fmt.Println("Motor tidak ada pada parkiran.")
				quit = true
			}
		}
	}else if kendaraan == "mobil" {
		for quit == false {
			if datamobil[berapa].plat == nomorkendaraan {
				fmt.Println("=========================================================")
				fmt.Println("		Ini karcis anda!")
				fmt.Println("		Jenis kendaraan:",datamobil[berapa].rodaberapa)
				fmt.Println("		Nomor kendaraan:",datamobil[berapa].plat)
				fmt.Println("		Tanggal masuk:", datamobil[berapa].hari.date," ", datamobil[berapa].hari.month, " ", datamobil[berapa].hari.year)
				fmt.Println("		Jam masuk:",datamobil[berapa].waktu.hour," ",datamobil[berapa].waktu.minute," ",datamobil[berapa].waktu.second)
				fmt.Println("=========================================================")
				quit = true
			}
			berapa++
			if berapa == 3 && datamobil[berapa].plat != nomorkendaraan {
				fmt.Print("Mobil tidak ada pada parkiran.")
				quit = true
			}
		}
	}else if kendaraan == "truk" {
		for quit == false {
			if datatruk[berapa].plat == nomorkendaraan {
				fmt.Println("=========================================================")
				fmt.Println("		Ini karcis anda!")
				fmt.Println("		Jenis kendaraan:",datatruk[berapa].rodaberapa)
				fmt.Println("		Nomor kendaraan:",datatruk[berapa].plat)
				fmt.Println("		Tanggal masuk:", datatruk[berapa].hari.date," ", datatruk[berapa].hari.month, " ", datatruk[berapa].hari.year)
				fmt.Println("		Jam masuk:",datatruk[berapa].waktu.hour," ",datatruk[berapa].waktu.minute," ",datatruk[berapa].waktu.second)
				fmt.Println("=========================================================")
				quit = true
			}
			berapa++
			if berapa == 3 && datatruk[berapa].plat != nomorkendaraan {
				fmt.Print("Truk tidak ada pada parkiran.")
				quit = true
			}
		}
	}
}

func sortparkir() {
	fmt.Println("Nomor motor")
	a:=datamotor
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 {
			if a[j-1].plat > a[j].plat {
				a[j-1].plat, a[j].plat = a[j].plat, a[j-1].plat
			}
			j = j - 1
		}
	}
	for berapa < len(a){
		if a[berapa].plat != "0" {
			fmt.Print(a[berapa].plat," ")
		}
		berapa++
	}
	fmt.Println("")
	//sorting mobil
	fmt.Println("Nomor mobil")
	b:=datamobil
	for i := 1; i < len(b); i++ {
		j := i
		for j > 0 {
			if b[j-1].plat > b[j].plat {
				b[j-1].plat, b[j].plat = b[j].plat, b[j-1].plat
			}
			j = j - 1
		}
	}
	berapa = 0
	for berapa < len(b){
		if b[berapa].plat != "0" {
			fmt.Print(b[berapa].plat," ")
		}
		berapa++
	}
	fmt.Println("")
	//sorting truk
	fmt.Println("Nomor truk")
	c:=datatruk
	for i := 1; i < len(c); i++ {
		n := c[i].plat
		j := i
		for j > 0 && c[j-1].plat > n {
			c[j].plat = c[j-1].plat
			j--
		}
		c[j].plat = n
	}
	berapa = 0
	for berapa < len(c){
		if c[berapa].plat != "0" {
			fmt.Print(c[berapa].plat," ")
		}
		berapa++
	}
	fmt.Println("")
}

func okupansi() {
	fmt.Println("Jumlah motor yang sedang parkir sekarang : ")
	fmt.Print(nmotor)
	fmt.Println("Persentase okupansi parkir motor: ")
	fmt.Println((nmotor*100/len(vehicle.T)),"%")
	
	fmt.Println("Jumlah mobil yang sedang parkir sekarang : ")
	fmt.Println(nmobil)
	fmt.Println("Jumlah truk yang sedang parkir sekarang : ")
	fmt.Println(ntruk)	
	fmt.Println("Persentase okupansi parkir mobil dan truk: ")
	fmt.Println(((nmobil+ntruk)*100/(len(vehicle.N)+len(vehicle.M))),"%")
}
func pendapatan() {
	fmt.Println("Kendaraan yang telah masuk : ")
	fmt.Println(in)
	fmt.Println("Kendaraan yang telah keluar : ")
	fmt.Println(out)
	fmt.Println("Pendapatan sejauh ini sebanyak")
	fmt.Println(earn)

}

func masukparkir() {
	var tanggal,bulan,tahun string
	var jam,menit,detik int
	fmt.Println("Jenis kendaraan anda?")
	fmt.Println("Jika motor tekan '1'.")
	fmt.Println("Jika mobil kecil tekan '2'.")
	fmt.Println("Jika truk tekan '3'.")
	fmt.Scanln(&kendaraan)

	if kendaraan == "1" {
		
		for vehicle.T[berapa] != 0 && berapa < 3 {
			berapa++
		}
		if vehicle.T[berapa] != 1{
			datamotor[berapa].rodaberapa = "motor"							
			fmt.Println("Parkiran motor tersedia.")
			fmt.Println("Masukkan plat nomor anda. 123")
			fmt.Scanln(&nomorkendaraan)
			datamotor[berapa].plat = nomorkendaraan
			fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
			fmt.Scanln(&tanggal,&bulan,&tahun)
			for tanggal < "0" || tanggal > "31" || bulan < "0" || bulan > "12" {
				fmt.Println("Masukkan tidak valid. Masukkan kembali.")
				fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
				fmt.Scanln(&tanggal,&bulan,&tahun)
			}
			datamotor[berapa].hari.date = tanggal
			datamotor[berapa].hari.month = bulan
			datamotor[berapa].hari.year = tahun
			fmt.Println("Pukul berapa anda masuk?")
			fmt.Scanln(&jam,&menit,&detik)
			for jam < 9 || jam >= 22 || menit > 59 || menit < 0 || detik > 59 || detik < 0 {
				fmt.Println("Jam tidak valid, masukkan kembali.")
				fmt.Println("Pukul berapa anda masuk?")
				fmt.Scanln(&jam,&menit,&detik)
			}
			datamotor[berapa].waktu.hour = jam
			datamotor[berapa].waktu.minute = menit
			datamotor[berapa].waktu.second = detik
			fmt.Println("=========================================================")
			fmt.Println("		KARCIS PARKIR MALL TELKOM")
			fmt.Println("		Jenis kendaraan:",datamotor[berapa].rodaberapa)
			fmt.Println("		Nomor kendaraan:",datamotor[berapa].plat)
			fmt.Println("		Tanggal masuk:", datamotor[berapa].hari.date," ", datamotor[berapa].hari.month, " ", datamotor[berapa].hari.year)
			fmt.Println("		Jam masuk:",datamotor[berapa].waktu.hour," ",datamotor[berapa].waktu.minute," ",datamotor[berapa].waktu.second)
			fmt.Println("=========================================================")		
			vehicle.T[berapa] = 1
			nmotor++
			in++
									
		}else{
			if vehicle.T[berapa] != 0 {
				fmt.Println("Maaf parkir motor penuh.")
					}
				}
			
		
		//bawah mobil dan truk di kendaraan 2 dan 3 bawah

	}else if kendaraan == "2"  {
		brick = false
		for vehicle.N[berapa] != 0 && berapa < 3 && brick == false {
			berapa++
			if berapa == 3 && vehicle.N[berapa] == 1 {
				berapa = 0 
				brick = true
				for vehicle.M[berapa] != 0 && berapa < 3 {
					berapa++
				}
			}
		}
		if vehicle.N[berapa] != 1 && brick == false {
			datamobil[berapa].rodaberapa = "mobil"						
			fmt.Println("Parkiran mobil tersedia.")
			fmt.Println("Masukkan plat nomor anda. 123")
			fmt.Scanln(&nomorkendaraan)
			datamobil[berapa].plat = nomorkendaraan
			fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
			fmt.Scanln(&tanggal,&bulan,&tahun)
			for tanggal < "0" || tanggal > "31" || bulan < "0" || bulan > "12" {
				fmt.Println("Masukkan tidak valid. Masukkan kembali.")
				fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
				fmt.Scanln(&tanggal,&bulan,&tahun)
			}
			datamobil[berapa].hari.date = tanggal
			datamobil[berapa].hari.month = bulan
			datamobil[berapa].hari.year = tahun
			fmt.Println("Pukul berapa anda masuk?")
			fmt.Scanln(&jam,&menit,&detik)
			for jam < 9 || jam >= 22 || menit > 59 || menit < 0 || detik > 59 || detik < 0 {
				fmt.Println("Jam tidak valid, masukkan kembali.")
				fmt.Println("Pukul berapa anda masuk?")
				fmt.Scanln(&jam,&menit,&detik)
			}
			datamobil[berapa].waktu.hour = jam
			datamobil[berapa].waktu.minute = menit
			datamobil[berapa].waktu.second = detik
			fmt.Println("=========================================================")
			fmt.Println("		KARCIS PARKIR MALL TELKOM")
			fmt.Println("		Jenis kendaraan:",datamobil[berapa].rodaberapa)
			fmt.Println("		Nomor kendaraan:",datamobil[berapa].plat)
			fmt.Println("		Tanggal masuk:", datamobil[berapa].hari.date," ", datamobil[berapa].hari.month, " ", datamobil[berapa].hari.year)
			fmt.Println("		Jam masuk:",datamobil[berapa].waktu.hour," ",datamobil[berapa].waktu.minute," ",datamobil[berapa].waktu.second)
			fmt.Println("=========================================================")		
			vehicle.N[berapa] = 1
			nmobil++
			in++
			
		}else if vehicle.M[berapa] != 1 && brick == true {
			datamobil[berapa].rodaberapa = "mobil"							
			fmt.Println("Parkiran mobil tersedia.")
			fmt.Println("Masukkan plat nomor anda. 123")
			fmt.Scanln(&nomorkendaraan)
			datamobil[berapa].plat = nomorkendaraan
			fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
			fmt.Scanln(&tanggal,&bulan,&tahun)
			for tanggal < "0" || tanggal > "31" || bulan < "0" || bulan > "12" {
				fmt.Println("Masukkan tidak valid. Masukkan kembali.")
				fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
				fmt.Scanln(&tanggal,&bulan,&tahun)
			}
			datamobil[berapa].hari.date = tanggal
			datamobil[berapa].hari.month = bulan
			datamobil[berapa].hari.year = tahun
			fmt.Println("Pukul berapa anda masuk?")
			fmt.Scanln(&jam,&menit,&detik)
			for jam < 9 || jam >= 22 || menit > 59 || menit < 0 || detik > 59 || detik < 0 {
				fmt.Println("Jam tidak valid, masukkan kembali.")
				fmt.Println("Pukul berapa anda masuk?")
				fmt.Scanln(&jam,&menit,&detik)
			}
			datamobil[berapa].waktu.hour = jam
			datamobil[berapa].waktu.minute = menit
			datamobil[berapa].waktu.second = detik
			fmt.Println("=========================================================")
			fmt.Println("		KARCIS PARKIR MALL TELKOM")
			fmt.Println("Jenis kendaraan:",datamobil[berapa].rodaberapa)
			fmt.Println("Nomor kendaraan:",datamobil[berapa].plat)
			fmt.Println("Tanggal masuk:", datamobil[berapa].hari.date," ", datamobil[berapa].hari.month, " ", datamobil[berapa].hari.year)
			fmt.Println("Jam masuk:",datamobil[berapa].waktu.hour," ",datamobil[berapa].waktu.minute," ",datamobil[berapa].waktu.second)
			fmt.Println("=========================================================")		
			vehicle.M[berapa] = 1
			nmobil++
			in++
		

		}else if brick == true || brick == false{
			
				fmt.Println("Maaf lahan parkir mobil dan truk penuh.")
					
				}

	}else if kendaraan == "3" {
		brick = false
		for vehicle.N[berapa] != 0 && berapa < 3 && brick == false {
			berapa++
			if berapa == 2 && vehicle.N[berapa] == 1 {
				berapa = 0 
				brick = true
				for vehicle.M[berapa] != 0 && berapa < 2 {
					berapa++
				}
			}
		}
		if vehicle.N[berapa] != 1 && vehicle.N[berapa+1] != 1 && brick == false {
			datatruk[berapa].rodaberapa = "truk"
			fmt.Println("Parkiran truk tersedia.")
			fmt.Println("Masukkan plat nomor anda. 123")
			fmt.Scanln(&nomorkendaraan)
			datatruk[berapa].plat = nomorkendaraan
			fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
			fmt.Scanln(&tanggal,&bulan,&tahun)
			for tanggal < "0" || tanggal > "31" || bulan < "0" || bulan > "12" {
				fmt.Println("Masukkan tidak valid. Masukkan kembali.")
				fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
				fmt.Scanln(&tanggal,&bulan,&tahun)
			}
			datatruk[berapa].hari.date = tanggal
			datatruk[berapa].hari.month = bulan
			datatruk[berapa].hari.year = tahun
			fmt.Println("Pukul berapa anda masuk?")
			fmt.Scanln(&jam,&menit,&detik)
			for jam < 9 || jam >= 22 || menit > 59 || menit < 0 || detik > 59 || detik < 0 {
				fmt.Println("Jam tidak valid, masukkan kembali.")
				fmt.Println("Pukul berapa anda masuk?")
				fmt.Scanln(&jam,&menit,&detik)
			}
			datatruk[berapa].waktu.hour = jam
			datatruk[berapa].waktu.minute = menit
			datatruk[berapa].waktu.second = detik
			fmt.Println("=========================================================")
			fmt.Println("		KARCIS PARKIR MALL TELKOM")
			fmt.Println("		Jenis kendaraan:",datatruk[berapa].rodaberapa)
			fmt.Println("		Nomor kendaraan:",datatruk[berapa].plat)
			fmt.Println("		Tanggal masuk:", datatruk[berapa].hari.date," ", datatruk[berapa].hari.month, " ", datatruk[berapa].hari.year)
			fmt.Println("		Jam masuk:",datatruk[berapa].waktu.hour," ",datatruk[berapa].waktu.minute," ",datatruk[berapa].waktu.second)
			fmt.Println("=========================================================")	
			vehicle.N[berapa] = 1
			vehicle.N[berapa+1] = 1	
			ntruk++
			in++						
		}else if vehicle.M[berapa] != 1 && vehicle.M[berapa+1] != 1 && brick == true {
			add:=3
			datatruk[add+berapa].rodaberapa = "truk"
			fmt.Println("Parkiran truk tersedia.")
			fmt.Println("Masukkan plat nomor anda. 123")
			fmt.Scanln(&nomorkendaraan)
			datatruk[add+berapa].plat = nomorkendaraan
			fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
			fmt.Scanln(&tanggal,&bulan,&tahun)
			for tanggal < "0" || tanggal > "31" || bulan < "0" || bulan > "12" {
				fmt.Println("Masukkan tidak valid. Masukkan kembali.")
				fmt.Println("Tanggal bulan tahun? contoh 26 03 2019")
				fmt.Scanln(&tanggal,&bulan,&tahun)
			}
			datatruk[add+berapa].hari.date = tanggal
			datatruk[add+berapa].hari.month = bulan
			datatruk[add+berapa].hari.year = tahun
			fmt.Println("Pukul berapa anda masuk?")
			fmt.Scanln(&jam,&menit,&detik)
			for jam < 9 || jam >= 22 || menit > 59 || menit < 0 || detik > 59 || detik < 0 {
				fmt.Println("Jam tidak valid, masukkan kembali.")
				fmt.Println("Pukul berapa anda masuk?")
				fmt.Scanln(&jam,&menit,&detik)
			}
			datatruk[add+berapa].waktu.hour = jam
			datatruk[add+berapa].waktu.minute = menit
			datatruk[add+berapa].waktu.second = detik
			fmt.Println("=========================================================")
			fmt.Println("		KARCIS PARKIR MALL TELKOM")
			fmt.Println("		Jenis kendaraan:",datatruk[add+berapa].rodaberapa)
			fmt.Println("		Nomor kendaraan:",datatruk[add+berapa].plat)
			fmt.Println("		Tanggal masuk:", datatruk[add+berapa].hari.date," ", datatruk[add+berapa].hari.month, " ", datatruk[add+berapa].hari.year)
			fmt.Println("		Jam masuk:",datatruk[add+berapa].waktu.hour," ",datatruk[add+berapa].waktu.minute," ",datatruk[add+berapa].waktu.second)
			fmt.Println("=========================================================")		
			vehicle.M[berapa] = 1
			vehicle.M[berapa+1] = 1	
			ntruk++	
			in++											
		}else{
			if vehicle.N[berapa] == 1 || vehicle.M[berapa] == 1 || brick == true || brick == false{
				fmt.Println("Maaf lahan parkir mobil dan truk penuh.")
					}
				}
	}

}

func keluarparkir() {
	var jamkeluar,menitkeluar,detikkeluar int
	var detikawal,detikakhir int
	var jumlahdetik int
	var harga int

	fmt.Println("Jenis kendaraan anda? (motor,mobil,truk)")
	fmt.Scanln(&kendaraan)
	fmt.Println("Plat nomor anda?")
	fmt.Scanln(&nomorkendaraan)
	fmt.Println("Jam anda keluar? (jam,menit,detik)")
	fmt.Scanln(&jamkeluar,&menitkeluar,&detikkeluar)

	cariplat := 0
	if kendaraan == "motor" {
		//search plat motor
		gone:=false
		for datamotor[cariplat].plat != nomorkendaraan && cariplat < 3 {
			cariplat++
			if datatruk[cariplat].plat == nomorkendaraan {
				gone=true
			}
		}
		if gone == true {
			catat()
			datamotor[cariplat].plat = "0"
			detikawal  =  (datamotor[cariplat].waktu.hour*3600) + (datamotor[cariplat].waktu.minute*60) + datamotor[cariplat].waktu.second
			detikakhir =  ((jamkeluar*3600 + (menitkeluar*60) + detikkeluar))
			jumlahdetik = detikakhir - detikawal
			
			if jumlahdetik > 0 {
				if jumlahdetik <= 600 {
					harga = 0
					fmt.Println("Parkir Gratis!")
				}else {
					harga = (((jumlahdetik-1)/3600)+1) * 1500
					fmt.Print("Biaya parkir :")
					fmt.Println(harga)
				}
				for out:=3 ; out >= 0 ; out-- {
					if vehicle.T[out] == 1 {
						vehicle.T[out] = 0
						out = -1
					}
				}
				earn = earn + harga
				nmotor--
				out++
			}else if jumlahdetik <= 0 {
				fmt.Println("Waktu keluar tidak valid. Negatif.")
				//	minta inputan waktu keluar lagi
			}
		} else {
			fmt.Print("Kendaraan tidak ada.")
		}
	}else if kendaraan == "mobil" {
		//search plat mobil
		gone:=false
		for datamobil[cariplat].plat != nomorkendaraan && cariplat < 3{
			cariplat++
			if datamobil[cariplat].plat == nomorkendaraan {
				gone=true
			}
		}
		if gone == true {
			catat()
			datamobil[cariplat].plat = "0"
			detikawal  =  (datamobil[cariplat].waktu.hour*3600) + (datamobil[cariplat].waktu.minute*60) + datamobil[cariplat].waktu.second
			detikakhir =  ((jamkeluar*3600 + (menitkeluar*60) + detikkeluar))
			jumlahdetik = detikakhir - detikawal
			
			if jumlahdetik > 0 {
				if jumlahdetik <= 600 {
					harga = 0
					fmt.Println("Parkir Gratis!")
				}else {
					harga = (((jumlahdetik-1)/3600)+1) * 3000
					fmt.Print("Biaya parkir :")
					fmt.Println(harga)
				}
				for out:=3 ; out >= 0 ; out-- {
					if vehicle.M[out] == 1 {
						vehicle.M[out] = 0
						out = -1
					}
				}
				for out:=3 ; out >= 0 ; out-- {
					if vehicle.N[out] == 1 {
						vehicle.N[out] = 0
						out = -1
					}
				}
				earn = earn + harga
				nmobil--
				out++
			}else if jumlahdetik <= 0 {
				fmt.Println("Waktu keluar tidak valid. Negatif.")
				//minta inputan waktu keluar lagi
				}
		}else {
			fmt.Print("Kendaraan tidak ada.")
		}
	}else if kendaraan == "truk" {
		//search plat truk
		gone:=true
		for datatruk[cariplat].plat != nomorkendaraan && cariplat < 3{
			cariplat++
			if datatruk[cariplat].plat == nomorkendaraan {
				gone=true
			}
		}
		if gone == true {
			catat()
			datatruk[cariplat].plat = "0"
			detikawal  =  (datatruk[cariplat].waktu.hour*3600) + (datatruk[cariplat].waktu.minute*60) + datatruk[cariplat].waktu.second
			detikakhir =  ((jamkeluar*3600 + (menitkeluar*60) + detikkeluar))
			jumlahdetik = detikakhir - detikawal
			
			if jumlahdetik > 0 {
				if jumlahdetik <= 600 {
					harga = 0
					fmt.Println("Parkir Gratis!")
				}else {
					harga = (((jumlahdetik-1)/3600)+1) * 5000
					fmt.Print("Biaya parkir :")
					fmt.Println(harga)
				}
				for out:=3 ; out >= 0 ; out-- {
					if vehicle.M[out] == 1 {
						vehicle.M[out] = 0
						vehicle.M[out-1] = 0
						out = -1
					}
				}
				for out:=3 ; out >= 0 ; out-- {
					if vehicle.N[out] == 1 {
						vehicle.N[out] = 0
						vehicle.M[out-1] = 0
						out = -1
					}
				}
				earn = earn + harga
				ntruk--
				out++
			}else if jumlahdetik <= 0 {
				fmt.Println("Waktu keluar tidak valid. Negatif.")
				//minta inputan waktu keluar lagi
			}
		} else {
			fmt.Print("Kendaraan tidak ada")
		}
	}
}

func catat() {
	f, err := os.Create("Plat.txt") // create
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("Jenis kendaraan : %s \n", kendaraan))
	_, err = f.WriteString(fmt.Sprintf("Nomor kendaraan : %s \n", nomorkendaraan))//write
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	


	//func main() {
	//	txt("D 3454 ASU") tes coba coba
	
}
