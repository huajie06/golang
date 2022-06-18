package douban

const localDBname = "./ify_movie_base.json"
const localNewMovieList = "./newMovie.json"
const localDuonao = "./iyf.json"

type MovieInfo struct {
	Region, Language, AddTime, DNRate string
}

type Iyf struct {
	Ret           int    `json:"ret"`
	Data          Data   `json:"data"`
	Msg           string `json:"msg"`
	IsSpecialArea int    `json:"isSpecialArea"`
}

// =========================below is from duonao=========================

type duonaoMovieInfoRet struct {
	Title, Region, Language, AddTime, DNRate string
}

type duonaoMovieInfo struct {
	Ret           int    `json:"ret"`
	Data          Data   `json:"data"`
	Msg           string `json:"msg"`
	IsSpecialArea int    `json:"isSpecialArea"`
}
type Result struct {
	// the key and lastkey, one of them can be used to create link
	AtypeName      string      `json:"atypeName"`
	VideoClassID   string      `json:"videoClassID"`
	Image          string      `json:"image"`
	Key            string      `json:"key"`
	Lang           string      `json:"lang"`
	Cid            string      `json:"cid"`
	LastName       string      `json:"lastName"`
	IsShowTodayNum bool        `json:"isShowTodayNum"`
	Title          string      `json:"title"`
	Hot            int         `json:"hot"`
	Rating         string      `json:"rating"`
	Year           int         `json:"year"`
	Regional       string      `json:"regional"`
	AddTime        string      `json:"addTime"`
	Directed       string      `json:"directed"`
	Starring       string      `json:"starring"`
	ShareCount     int         `json:"shareCount"`
	Dd             int         `json:"dd"`
	Dc             int         `json:"dc"`
	Comments       int         `json:"comments"`
	FavoriteCount  int         `json:"favoriteCount"`
	Contxt         string      `json:"contxt"`
	IsSerial       bool        `json:"isSerial"`
	Updateweekly   string      `json:"updateweekly"`
	CidMapper      string      `json:"cidMapper"`
	LastKey        string      `json:"lastKey"`
	Recommended    bool        `json:"recommended"`
	Updates        int         `json:"updates"`
	Tags           interface{} `json:"tags"`
	IsFilm         bool        `json:"isFilm"`
	IsDocumentry   bool        `json:"isDocumentry"`
	Labels         string      `json:"labels"`
	Charge         int         `json:"charge"`
}
type Info struct {
	Recordcount int      `json:"recordcount"`
	Result      []Result `json:"result"`
}
type Data struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info []Info `json:"info"`
}

// =========================above is from duonao=========================

type doubanIndividualMovie struct {
	SearchedTitle   string
	ReturnReason    string
	Name            string
	Url             string
	DatePublished   string          `json:datePublished`
	Genre           []string        `json:genre`
	Duration        string          `json:duration`
	AggregateRating AggregateRating `json:aggregateRating`
	QueryDateTime   string
}

type AggregateRating struct {
	Type        string `json:"@type"`
	RatingCount string `json:"ratingCount"`
	BestRating  string `json:"bestRating"`
	WorstRating string `json:"worstRating"`
	RatingValue string `json:"ratingValue"`
}
