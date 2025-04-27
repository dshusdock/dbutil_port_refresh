package file_upload

import (
	
	"encoding/csv"
	"io"
	"log"
	"mime/multipart"
)

type LabDataRow struct {
	Cab string
	U string
	ISO string
	Name string
	SerialNbr string
	IP string
	VIP string
	IdracIP string
	SWVer string
	ServerType string
	Enterprise string
	Role string
	Comments string
	VMLabServerHostIP string
}

func ProcessLabInfo(f multipart.File) ([]LabDataRow) {
	var result []LabDataRow
	reader := csv.NewReader(f)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var x = LabDataRow{}
		for i := 0; i < len(record); i++ {
			switch(i) {
			case 0:
				x.Cab = record[i]
			case 1:
				x.U = record[i]
			case 2:
				x.ISO = record[i]
			case 3:
				x.Name = record[i]
			case 4:
				x.SerialNbr = record[i]
			case 5:
				x.IP = record[i]
			case 6:
				x.VIP = record[i]
			case 7:
				x.IdracIP = record[i]
			case 8:
				x.SWVer = record[i]
			case 9:
				x.ServerType= record[i]
			case 10:
				x.Enterprise = record[i]
			case 11:
				x.Role = record[i]
			case 12:
				x.Comments = record[i]
			case 13:
				x.VMLabServerHostIP = record[i]				
			}
		}
		var sqlStr = "INSERT into LabSystem values('" + x.Cab + "', '" + x.U + "', '" + x.ISO + "', '" + x.Name + "', '" + x.SerialNbr + "', '" + x.IP + "', '" + x.VIP + "', '" + x.IdracIP + "', '" + x.SWVer + "', '" + x.ServerType + "', '" + x.Enterprise + "', '" + x.Role + "', '" + x.Comments + "', '" + x.VMLabServerHostIP + "')"
		log.Println(sqlStr)							
		//d.WriteLocalDB(sqlStr)

		result = append(result, x)
	}
	return result
}
