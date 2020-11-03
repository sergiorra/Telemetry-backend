package models

const (
	TimeLayout = "15:04:05.000"
)

// DataResponse representation of dataResponse into struct
type DataResponse struct {
	Kind string 		`json:"kind"`
	Data InternalData 	`json:"data"`
}

type InternalData struct {
	Time  	string 		`json:"time"`
	Gear  	string 		`json:"gear"`
	Rpm   	int    		`json:"rpm"`
	Speed 	int    		`json:"speed"`
}

// NewDataResponse initialize struct DataResponse
func NewDataResponse(data Data) (d *DataResponse) {
	d = &DataResponse{
		Kind: "data",
		Data: InternalData{
			Time: data.Time.Format(TimeLayout),
			Gear: data.Gear,
			Rpm: data.Rpm,
			Speed: data.Speed,
		},
	}
	return
}
