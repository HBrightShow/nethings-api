package types


type  Album struct {
	Id 				uint64 			`json:"id"`
	Uid				uint64			`json:"uid"`
	Phone			string			`json:"phone"`
	Name			string			`json:"name"`
	Zipcode			string			`json:"zipcode"`
	Address			string			`json:"address"`
	Default_address 		uint64 			`json:"default_address"`
	Add_time 		uint64 			`json:"add_time"`

}