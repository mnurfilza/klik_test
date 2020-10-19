# klik_test

var (
	user     = "root"
	password = "mypass123"  // adjust with your mysql configuration
	host     = "172.17.0.2" //adjust with your configuration
	port     = "3306"
	dbname   = "klik_test"
)

Sesuaikan configuration diatas dengan configuration anda
endpoint:
Method Post : klikdaily/adjustment
             {
              id:..
              product:...
              adjustment:... 
             }
             
Method Get : klikdaily/logs?location_id=....
Method Get : klikdaily/stocks
