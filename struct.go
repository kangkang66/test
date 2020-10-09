package main

import (
	"encoding/json"
	"log"
)

type User struct {
	ID               int         `json:"id"`
	UserID           int         `json:"userId"`
	Nickname         string      `json:"nickname"`
	Avatar           string      `json:"avatar"`
	Gender           interface{} `json:"gender"`
	Province         string      `json:"province"`
	City             string      `json:"city"`
	Country          string      `json:"country"`
	Status           int         `json:"status"`
	LastHiddenTime   interface{} `json:"lastHiddenTime"`
	Remark           string      `json:"remark"`
	RemarkSelf       string      `json:"remarkSelf"`
	IsNew            int         `json:"isNew"`
	CertStatus       int         `json:"certStatus"`
	Applytimes       int         `json:"applytimes"`
	CompletionDegree int         `json:"completionDegree"`
	DataLevel        int         `json:"dataLevel"`
	VipLevel         int         `json:"vipLevel"`
	HiddenFields     string      `json:"hiddenFields"`
	Updatetime       int         `json:"updatetime"`
	HideReason       interface{} `json:"hideReason"`
	Basicinfo        struct {
		ID               int         `json:"id"`
		LoverUserID      int         `json:"loverUserId"`
		UID              string      `json:"uid"`
		Phone            string      `json:"phone"`
		Nickname         string      `json:"nickname"`
		Wechatid         string      `json:"wechatid"`
		Avatar           string      `json:"avatar"`
		Photos           string      `json:"photos"`
		Gender           int         `json:"gender"`
		Birthday         string      `json:"birthday"`
		Constellation    string      `json:"constellation"`
		Height           int         `json:"height"`
		Weight           int         `json:"weight"`
		Company          string      `json:"company"`
		CompanyID        interface{} `json:"companyId"`
		Position         string      `json:"position"`
		Isabroad         int         `json:"isabroad"`
		Abroadcountry    string      `json:"abroadcountry"`
		HomelandProvince string      `json:"homelandProvince"`
		HomelandCity     string      `json:"homelandCity"`
		LivingProvince   string      `json:"livingProvince"`
		LivingCity       string      `json:"livingCity"`
		Highestschool    string      `json:"highestschool"`
		SchoolID         interface{} `json:"schoolId"`
		Highestdegree    int         `json:"highestdegree"`
		Annualincome     int         `json:"annualincome"`
		Marital          int         `json:"marital"`
		Createtime       interface{} `json:"createtime"`
		Updatetime       interface{} `json:"updatetime"`
		Record           interface{} `json:"record"`
		DataDeleted      int         `json:"dataDeleted"`
	} `json:"basicinfo"`
	Card []struct {
		ID            int         `json:"id"`
		LoverUserID   int         `json:"loverUserId"`
		Type          string      `json:"type"`
		Placeholder   string      `json:"placeholder"`
		Topic         string      `json:"topic"`
		Title         string      `json:"title"`
		Contenttext   string      `json:"contenttext"`
		Contentimages string      `json:"contentimages"`
		Isshow        int         `json:"isshow"`
		Isdefault     int         `json:"isdefault"`
		Weigh         int         `json:"weigh"`
		Createtime    int         `json:"createtime"`
		Updatetime    interface{} `json:"updatetime"`
		Record        interface{} `json:"record"`
		DataDeleted   int         `json:"dataDeleted"`
		CardStatus    interface{} `json:"cardStatus"`
	} `json:"card"`
	PhotoStatus   interface{} `json:"photoStatus"`
	TagStatus     interface{} `json:"tagStatus"`
	CustomCard    interface{} `json:"customCard"`
	Certification []struct {
		Type         string      `json:"type"`
		Status       int         `json:"status"`
		Remark       string      `json:"remark"`
		HasReAuth    bool        `json:"hasReAuth"`
		RejectReason string      `json:"rejectReason"`
		RejectTime   interface{} `json:"rejectTime"`
	} `json:"certification"`
	Tag                      interface{} `json:"tag"`
	Prefer                   interface{} `json:"prefer"`
	Tags                     interface{} `json:"tags"`
	Friend                   interface{} `json:"friend"`
	RongcloudToken           string      `json:"rongcloudToken"`
	UserUUID                 string      `json:"userUuid"`
	BlockSchool              interface{} `json:"blockSchool"`
	BlockCompany             bool        `json:"blockCompany"`
	Collected                interface{} `json:"collected"`
	LoverBlocked             interface{} `json:"loverBlocked"`
	TargetBlocked            interface{} `json:"targetBlocked"`
	SignCount                interface{} `json:"signCount"`
	ShareTitle               interface{} `json:"shareTitle"`
	UserAssets               interface{} `json:"userAssets"`
	CustomTag                interface{} `json:"customTag"`
	Signed                   interface{} `json:"signed"`
	FriendStatus             interface{} `json:"friendStatus"`
	FriendStatusV2           interface{} `json:"friendStatusV2"`
	FinishFollowTask         interface{} `json:"finishFollowTask"`
	DynamicNum               int         `json:"dynamicNum"`
	NeedHideRemark           bool        `json:"needHideRemark"`
	Prerogative              string      `json:"prerogative"`
	UnlockStudentPrerogative interface{} `json:"unlockStudentPrerogative"`
	Role                     int         `json:"role"`
	Identity                 string      `json:"identity"`
	Lbs                      struct {
		District interface{} `json:"district"`
		Distance string      `json:"distance"`
	} `json:"lbs"`
	NewTag []struct {
		SortID  int    `json:"sortId"`
		Sort    string `json:"sort"`
		Order   int    `json:"order"`
		TagList []struct {
			Idx  int    `json:"idx"`
			Name string `json:"name"`
		} `json:"tagList"`
	} `json:"newTag"`
	HasUnionID             interface{} `json:"hasUnionId"`
	GroupID                interface{} `json:"groupId"`
	ExtendInfo             interface{} `json:"extendInfo"`
	IsJoinedGroup          interface{} `json:"isJoinedGroup"`
	IsCollectGroup         interface{} `json:"isCollectGroup"`
	GroupOwnerCenterStatus interface{} `json:"groupOwnerCenterStatus"`
	Createtime             int         `json:"createtime"`
	GetNewUserGift         interface{} `json:"getNewUserGift"`
	AvatarAuth             interface{} `json:"avatarAuth"`
}

func (u *User) ToJson() (jb []byte) {
	var err error
	jb,err = json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

type ResponseRecommendList struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    []*User `json:"data"`
}

type ResponseDynamicList struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    []struct {
		UID             string        `json:"uid"`
		Nickname        string        `json:"nickname"`
		City            string        `json:"city"`
		HomelandCity    string        `json:"homelandCity"`
		Birthday        string        `json:"birthday"`
		Avatar          string        `json:"avatar"`
		Gender          int           `json:"gender"`
		TopicID         int           `json:"topicId"`
		Position        string        `json:"position"`
		TopicContent    string        `json:"topicContent"`
		DynamicImages   string        `json:"dynamicImages"`
		DynamicContent  string        `json:"dynamicContent"`
		DynamicType     int           `json:"dynamicType"`
		DynamicID       int           `json:"dynamicId"`
		CreateTime      string        `json:"createTime"`
		DynamicAreaDesc string        `json:"dynamicAreaDesc"`
		Longitude       string        `json:"longitude"`
		Latitude        string        `json:"latitude"`
		WatchNum        int           `json:"watchNum"`
		LikeNum         int           `json:"likeNum"`
		LikeUserAvatar  interface{}   `json:"likeUserAvatar"`
		CommentNum      int           `json:"commentNum"`
		Tag             string        `json:"tag"`
		IsOfficial      bool          `json:"isOfficial"`
		IsStick         bool          `json:"isStick"`
		IsOpen          bool          `json:"isOpen"`
		IsMask          bool          `json:"isMask"`
		Width           int           `json:"width"`
		Height          int           `json:"height"`
		OwnerLike       bool          `json:"ownerLike"`
		UserDeleted     bool          `json:"userDeleted"`
		ShareNum        int           `json:"shareNum"`
		ShareImg        interface{}   `json:"shareImg"`
		LinkDatas       []interface{} `json:"linkDatas"`
		Voteinfo        interface{}   `json:"voteinfo"`
		Videoinfo       interface{}   `json:"videoinfo"`
		VoiceData       interface{}   `json:"voiceData"`
		Comments        interface{}   `json:"comments"`
		ExposeNum       int           `json:"exposeNum"`
	} `json:"data"`
}

type ResponseUserInfo struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		UserInfo *User `json:"userInfo"`
		Type int `json:"type"`
	} `json:"data"`
}

type ResponseVistors struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		TodayVisitorNums int `json:"todayVisitorNums"`
		TotalVisitorNums int `json:"totalVisitorNums"`
		Visitors         []struct {
			ID            int    `json:"id"`
			Nickname      string `json:"nickname"`
			Avatar        string `json:"avatar"`
			UID           string `json:"uid"`
			Gender        int    `json:"gender"`
			LastVisitTime int64  `json:"lastVisitTime"`
			VisitTimes    int    `json:"visitTimes"`
			Similarity    string `json:"similarity"`
			Mask          bool   `json:"mask"`
			UnlockCost    int    `json:"unlockCost"`
			LivingCity    string `json:"livingCity"`
			Birthday      string `json:"birthday"`
		} `json:"visitors"`
	} `json:"data"`
}


type ResponseCollect struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type ResponseCollectList struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		List []struct {
			LoverUserID   int    `json:"loverUserId"`
			UID           string `json:"uid"`
			Nickname      string `json:"nickname"`
			Avatar        string `json:"avatar"`
			Birthday      string `json:"birthday"`
			LivingCity    string `json:"livingCity"`
			Height        int    `json:"height"`
			Gender        int    `json:"gender"`
			Highestdegree string `json:"highestdegree"`
			CertStatus    int    `json:"certStatus"`
			ID            int    `json:"id"`
			Unlocked      bool   `json:"unlocked"`
			HappenTime    string `json:"happenTime"`
		} `json:"list"`
		UnReadCount int `json:"unReadCount"`
	} `json:"data"`
}


type ResponseIp struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}