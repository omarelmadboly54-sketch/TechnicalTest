package models



type Product struct{
	ID			int			`json:"id"`
	Title		string		`json:"title"`
	Price		float64		`json:"price"`
	Rating		float64		`json:"rating"`
	Stock		int			`json:"stock"`
	Thumbnail	string		`json:"thumbnail"`

}