package updates

type Market struct{
	ID			int			`json:"id"`
	OwnerID		int			`json:"owner_id"`
	Title		string		`json:"title"`
	Description	string		`json:"description"`
	Price		Price		`json:"price"`
	Dimensions	Dimensions	`json:"dimensions"`
	Weight		int			`json:"weight"`
	Category	Category	`json:"category"`
	Cover		string		`json:"thumb_photo"`
	Date		int			`json:"date"`
	Access 		int			`json:"availability"`
	Favorite 	bool		`json:"is_favorite"`
	Vendor 		string		`json:"sku"`

	Photos 		[]interface{}`json:"photos,omitempty"`
	Comment 	int			`json:"can_comment,omitempty"`
	Respost 	int			`json:"can_repost,omitempty"`
	Likes 		[]Like		`json:"likes,omitempty"`
	URL 		string		`json:"url,omitempty"`
	Button 		string		`json:"button_title,omitempty"`
}

type Price struct{
	Amount			string		`json:"amount,omitempty"`
	Currency		Currency	`json:"currency,omitempty"`
	Undiscounted	string		`json:"old_amount,omitempty"`
	Text			string		`json:"text,omitempty"`
}
type Currency struct{
	ID		int		`json:"id,omitempty"`
	Name	string	`json:"name,omitempty"`
}

type Dimensions struct{
	Width	int	`json:"width,omitempty"`
	Height	int	`json:"height,omitempty"`
	Length	int	`json:"length,omitempty"`
}

type Category struct{
	ID		int		`json:"id,omitempty"`
	Name	int		`json:"name,omitempty"`
	Section	Section	`json:"section,omitempty"`
}
type Section struct{
	ID		int		`json:"id,omitempty"`
	Name	string	`json:"name,omitempty"`
}

type Like struct{
	Liked	int `json:"user_likes,omitempty"`
	Count	int `json:"count,omitempty"`
}