package main

type Container struct {
	Ok   int `json:"ok"`
	Data struct {
		UserInfo struct {
			ID              int64  `json:"id"`
			ScreenName      string `json:"screen_name"`
			ProfileImageURL string `json:"profile_image_url"`
			ProfileURL      string `json:"profile_url"`
		} `json:"userInfo"`
		TabsInfo struct {
			SelectedTab int `json:"selectedTab"`
			Tabs        []struct {
				ID          int    `json:"id"`
				TabKey      string `json:"tabKey"`
				MustShow    int    `json:"must_show"`
				Hidden      int    `json:"hidden"`
				Title       string `json:"title"`
				TabType     string `json:"tab_type"`
				Containerid string `json:"containerid"`
				Apipath     string `json:"apipath,omitempty"`
				URL         string `json:"url,omitempty"`
			} `json:"tabs"`
		} `json:"tabsInfo"`
	} `json:"data"`
	Msg string `json:"msg,omitempty"`
}

type Weibo struct {
	Ok   int `json:"ok"`
	Data struct {
		Cards []struct {
			CardType int    `json:"card_type"`
			Itemid   string `json:"itemid"`
			Scheme   string `json:"scheme"`
			Mblog    struct {
				CreatedAt                string `json:"created_at"`
				ID                       string `json:"id"`
				Idstr                    string `json:"idstr"`
				Mid                      string `json:"mid"`
				CanEdit                  bool   `json:"can_edit"`
				ShowAdditionalIndication int    `json:"show_additional_indication"`
				Text                     string `json:"text"`
				TextLength               int    `json:"textLength"`
				Source                   string `json:"source"`
				Favorited                bool   `json:"favorited"`
				PicTypes                 string `json:"pic_types"`
				IsPaid                   bool   `json:"is_paid"`
				MblogVipType             int    `json:"mblog_vip_type"`
				User                     struct {
					ID              int64  `json:"id"`
					ScreenName      string `json:"screen_name"`
					ProfileImageURL string `json:"profile_image_url"`
					ProfileURL      string `json:"profile_url"`
					StatusesCount   int    `json:"statuses_count"`
					Verified        bool   `json:"verified"`
					VerifiedType    int    `json:"verified_type"`
					CloseBlueV      bool   `json:"close_blue_v"`
					Description     string `json:"description"`
					Gender          string `json:"gender"`
					Mbtype          int    `json:"mbtype"`
					Urank           int    `json:"urank"`
					Mbrank          int    `json:"mbrank"`
					FollowMe        bool   `json:"follow_me"`
					Following       bool   `json:"following"`
					FollowersCount  int    `json:"followers_count"`
					FollowCount     int    `json:"follow_count"`
					CoverImagePhone string `json:"cover_image_phone"`
					AvatarHd        string `json:"avatar_hd"`
					Like            bool   `json:"like"`
					LikeMe          bool   `json:"like_me"`
					Badge           struct {
						UnreadPool          int `json:"unread_pool"`
						UnreadPoolExt       int `json:"unread_pool_ext"`
						UserNameCertificate int `json:"user_name_certificate"`
						Double112018        int `json:"double11_2018"`
						Hongbaofei2019      int `json:"hongbaofei_2019"`
						StatusVisible       int `json:"status_visible"`
					} `json:"badge"`
				} `json:"user"`
				RepostsCount         int  `json:"reposts_count"`
				CommentsCount        int  `json:"comments_count"`
				AttitudesCount       int  `json:"attitudes_count"`
				PendingApprovalCount int  `json:"pending_approval_count"`
				IsLongText           bool `json:"isLongText"`
				RewardExhibitionType int  `json:"reward_exhibition_type"`
				HideFlag             int  `json:"hide_flag"`
				Visible              struct {
					Type   int `json:"type"`
					ListID int `json:"list_id"`
				} `json:"visible"`
				Mblogtype             int `json:"mblogtype"`
				MoreInfoType          int `json:"more_info_type"`
				ExternSafe            int `json:"extern_safe"`
				NumberDisplayStrategy struct {
					ApplyScenarioFlag    int    `json:"apply_scenario_flag"`
					DisplayTextMinNumber int    `json:"display_text_min_number"`
					DisplayText          string `json:"display_text"`
				} `json:"number_display_strategy"`
				ContentAuth int `json:"content_auth"`
				EditConfig  struct {
					Edited bool `json:"edited"`
				} `json:"edit_config"`
				IsTop         int    `json:"isTop"`
				WeiboPosition int    `json:"weibo_position"`
				Bid           string `json:"bid"`
				Title         struct {
					Text      string `json:"text"`
					BaseColor int    `json:"base_color"`
				} `json:"title"`
			} `json:"mblog"`
		} `json:"cards"`
	} `json:"data"`
}

type Notifier struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type WeiboId struct {
	Uid         string `json:"uid"`
	ContainerId string `json:"container_id, omitempty"`
}

type Config struct {
	WeiboId  []WeiboId `json:"weibo_id"`
	Notifier Notifier  `json:"notifier"`
}

type SaveWeibo struct {
	Timestamp int64
	WeiboId   string
	Content   string
}
