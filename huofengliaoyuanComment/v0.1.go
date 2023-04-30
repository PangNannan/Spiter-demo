package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*爬取火凤燎原的评论
动态爬取
*/

func main() {

	Spiter()
}

func Spiter() {
	//1、客户端请求
	client := http.Client{}
	url := "https://api.bilibili.com/x/v2/reply/main?mode=3&next=0&oid=613031046&plat=1&type=1"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("request err :", err)
		return
	}
	//设置请求头
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("accept", "application/json, text/plain, */*")
	request.Header.Set("sec-ch-ua", "Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99")
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	request.Header.Set("sec-fetch-mode", "")

	//发送请求
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("client request err :", err)
		return
	}
	//2、解析body，对json反序列化
	bodyTestByte, err := io.ReadAll(resp.Body)
	var fengJson FenghuoliaoyuanJsonName

	err1 := json.Unmarshal(bodyTestByte, &fengJson)
	if err1 != nil {
		fmt.Println("json Unmarshal err :", err)
	}

	//3打印评论
	replies := fengJson.Data.Replies
	for _, reply := range replies {
		message := reply.Content.Message
		fmt.Println(message)

	}
}

type FenghuoliaoyuanJsonName struct {
	Code int64 `json:"code"`
	Data struct {
		Replies []struct {
			Action    int64 `json:"action"`
			Assist    int64 `json:"assist"`
			Attr      int64 `json:"attr"`
			CardLabel []struct {
				Background       string `json:"background"`
				BackgroundHeight int64  `json:"background_height"`
				BackgroundWidth  int64  `json:"background_width"`
				Effect           int64  `json:"effect"`
				EffectStartTime  int64  `json:"effect_start_time"`
				Image            string `json:"image"`
				JumpURL          string `json:"jump_url"`
				LabelColorDay    string `json:"label_color_day"`
				LabelColorNight  string `json:"label_color_night"`
				Rpid             int64  `json:"rpid"`
				TextColorDay     string `json:"text_color_day"`
				TextColorNight   string `json:"text_color_night"`
				TextContent      string `json:"text_content"`
				Type             int64  `json:"type"`
			} `json:"card_label"`
			Content struct {
				Emote struct {
					Tv色 struct {
						Attr      int64  `json:"attr"`
						GifURL    string `json:"gif_url"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[tv_色]"`
					出游季比心 struct {
						Attr      int64  `json:"attr"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[出游季_比心]"`
					微笑 struct {
						Attr      int64  `json:"attr"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[微笑]"`
					恋爱告急爱你 struct {
						Attr      int64  `json:"attr"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[恋爱告急_爱你]"`
					火凤燎原冲鸭 struct {
						Attr      int64  `json:"attr"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[火凤燎原_冲鸭]"`
					火凤燎原当心了 struct {
						Attr      int64  `json:"attr"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[火凤燎原_当心了]"`
					火凤燎原我有新点子了 struct {
						Attr      int64  `json:"attr"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[火凤燎原_我有新点子了]"`
					火凤燎原三国群英集袁方_微笑 struct {
						Attr      int64  `json:"attr"`
						ID        int64  `json:"id"`
						JumpTitle string `json:"jump_title"`
						Meta      struct {
							Size int64 `json:"size"`
						} `json:"meta"`
						Mtime     int64  `json:"mtime"`
						PackageID int64  `json:"package_id"`
						State     int64  `json:"state"`
						Text      string `json:"text"`
						Type      int64  `json:"type"`
						URL       string `json:"url"`
					} `json:"[火凤燎原三国群英集_袁方 微笑]"`
				} `json:"emote"`
				JumpURL struct {
					泰酷辣 struct {
						AppName        string `json:"app_name"`
						AppPackageName string `json:"app_package_name"`
						AppURLSchema   string `json:"app_url_schema"`
						ClickReport    string `json:"click_report"`
						ExposureReport string `json:"exposure_report"`
						Extra          struct {
							GoodsClickReport    string `json:"goods_click_report"`
							GoodsCmControl      int64  `json:"goods_cm_control"`
							GoodsExposureReport string `json:"goods_exposure_report"`
							GoodsShowType       int64  `json:"goods_show_type"`
							IsWordSearch        bool   `json:"is_word_search"`
						} `json:"extra"`
						IconPosition int64  `json:"icon_position"`
						IsHalfScreen bool   `json:"is_half_screen"`
						MatchOnce    bool   `json:"match_once"`
						PcURL        string `json:"pc_url"`
						PrefixIcon   string `json:"prefix_icon"`
						State        int64  `json:"state"`
						Title        string `json:"title"`
						Underline    bool   `json:"underline"`
					} `json:"泰酷辣！"`
					火凤撩 struct {
						AppName        string `json:"app_name"`
						AppPackageName string `json:"app_package_name"`
						AppURLSchema   string `json:"app_url_schema"`
						ClickReport    string `json:"click_report"`
						ExposureReport string `json:"exposure_report"`
						Extra          struct {
							GoodsClickReport    string `json:"goods_click_report"`
							GoodsCmControl      int64  `json:"goods_cm_control"`
							GoodsExposureReport string `json:"goods_exposure_report"`
							GoodsShowType       int64  `json:"goods_show_type"`
							IsWordSearch        bool   `json:"is_word_search"`
						} `json:"extra"`
						IconPosition int64  `json:"icon_position"`
						IsHalfScreen bool   `json:"is_half_screen"`
						MatchOnce    bool   `json:"match_once"`
						PcURL        string `json:"pc_url"`
						PrefixIcon   string `json:"prefix_icon"`
						State        int64  `json:"state"`
						Title        string `json:"title"`
						Underline    bool   `json:"underline"`
					} `json:"火凤撩"`
					燎原火 struct {
						AppName        string `json:"app_name"`
						AppPackageName string `json:"app_package_name"`
						AppURLSchema   string `json:"app_url_schema"`
						ClickReport    string `json:"click_report"`
						ExposureReport string `json:"exposure_report"`
						Extra          struct {
							GoodsClickReport    string `json:"goods_click_report"`
							GoodsCmControl      int64  `json:"goods_cm_control"`
							GoodsExposureReport string `json:"goods_exposure_report"`
							GoodsShowType       int64  `json:"goods_show_type"`
							IsWordSearch        bool   `json:"is_word_search"`
						} `json:"extra"`
						IconPosition int64  `json:"icon_position"`
						IsHalfScreen bool   `json:"is_half_screen"`
						MatchOnce    bool   `json:"match_once"`
						PcURL        string `json:"pc_url"`
						PrefixIcon   string `json:"prefix_icon"`
						State        int64  `json:"state"`
						Title        string `json:"title"`
						Underline    bool   `json:"underline"`
					} `json:"燎原火"`
				} `json:"jump_url"`
				MaxLine  int64         `json:"max_line"`
				Members  []interface{} `json:"members"`
				Message  string        `json:"message"`
				Pictures []struct {
					ImgHeight int64   `json:"img_height"`
					ImgSize   float64 `json:"img_size"`
					ImgSrc    string  `json:"img_src"`
					ImgWidth  int64   `json:"img_width"`
				} `json:"pictures"`
				Topics     []string `json:"topics"`
				TopicsMeta struct {
					火凤燎原 struct {
						ID  int64  `json:"id"`
						URI string `json:"uri"`
					} `json:"火凤燎原"`
					火凤燎原开播 struct {
						ID  int64  `json:"id"`
						URI string `json:"uri"`
					} `json:"火凤燎原开播"`
				} `json:"topics_meta"`
				TopicsURI struct {
					火凤燎原 string `json:"火凤燎原"`
				} `json:"topics_uri"`
			} `json:"content"`
			Count        int64  `json:"count"`
			Ctime        int64  `json:"ctime"`
			Dialog       int64  `json:"dialog"`
			DynamicID    int64  `json:"dynamic_id"`
			DynamicIDStr string `json:"dynamic_id_str"`
			Fansgrade    int64  `json:"fansgrade"`
			Folder       struct {
				HasFolded bool   `json:"has_folded"`
				IsFolded  bool   `json:"is_folded"`
				Rule      string `json:"rule"`
			} `json:"folder"`
			Invisible bool  `json:"invisible"`
			Like      int64 `json:"like"`
			Member    struct {
				Avatar     string `json:"avatar"`
				AvatarItem struct {
					ContainerSize struct {
						Height float64 `json:"height"`
						Width  float64 `json:"width"`
					} `json:"container_size"`
					FallbackLayers struct {
						IsCriticalGroup bool `json:"is_critical_group"`
						Layers          []struct {
							GeneralSpec struct {
								PosSpec struct {
									AxisX         float64 `json:"axis_x"`
									AxisY         float64 `json:"axis_y"`
									CoordinatePos int64   `json:"coordinate_pos"`
								} `json:"pos_spec"`
								RenderSpec struct {
									Opacity int64 `json:"opacity"`
								} `json:"render_spec"`
								SizeSpec struct {
									Height int64 `json:"height"`
									Width  int64 `json:"width"`
								} `json:"size_spec"`
							} `json:"general_spec"`
							LayerConfig struct {
								IsCritical bool `json:"is_critical"`
								LayerMask  struct {
									GeneralSpec struct {
										PosSpec struct {
											AxisX         float64 `json:"axis_x"`
											AxisY         float64 `json:"axis_y"`
											CoordinatePos int64   `json:"coordinate_pos"`
										} `json:"pos_spec"`
										RenderSpec struct {
											Opacity int64 `json:"opacity"`
										} `json:"render_spec"`
										SizeSpec struct {
											Height int64 `json:"height"`
											Width  int64 `json:"width"`
										} `json:"size_spec"`
									} `json:"general_spec"`
									MaskSrc struct {
										Draw struct {
											ColorConfig struct {
												Day struct {
													Argb string `json:"argb"`
												} `json:"day"`
											} `json:"color_config"`
											DrawType int64 `json:"draw_type"`
											FillMode int64 `json:"fill_mode"`
										} `json:"draw"`
										SrcType int64 `json:"src_type"`
									} `json:"mask_src"`
								} `json:"layer_mask"`
								Tags struct {
									AvatarLayer  struct{} `json:"AVATAR_LAYER"`
									IconLayer    struct{} `json:"ICON_LAYER"`
									PendentLayer struct{} `json:"PENDENT_LAYER"`
								} `json:"tags"`
							} `json:"layer_config"`
							Resource struct {
								ResImage struct {
									ImageSrc struct {
										Local       int64 `json:"local"`
										Placeholder int64 `json:"placeholder"`
										Remote      struct {
											BfsStyle string `json:"bfs_style"`
											URL      string `json:"url"`
										} `json:"remote"`
										SrcType int64 `json:"src_type"`
									} `json:"image_src"`
								} `json:"res_image"`
								ResNativeDraw struct {
									DrawSrc struct {
										Draw struct {
											ColorConfig struct {
												Day struct {
													Argb string `json:"argb"`
												} `json:"day"`
												IsDarkModeAware bool `json:"is_dark_mode_aware"`
												Night           struct {
													Argb string `json:"argb"`
												} `json:"night"`
											} `json:"color_config"`
											DrawType int64 `json:"draw_type"`
											FillMode int64 `json:"fill_mode"`
										} `json:"draw"`
										SrcType int64 `json:"src_type"`
									} `json:"draw_src"`
								} `json:"res_native_draw"`
								ResType int64 `json:"res_type"`
							} `json:"resource"`
							Visible bool `json:"visible"`
						} `json:"layers"`
					} `json:"fallback_layers"`
					Mid string `json:"mid"`
				} `json:"avatar_item"`
				ContractDesc string `json:"contract_desc"`
				FaceNftNew   int64  `json:"face_nft_new"`
				FansDetail   struct {
					GuardIcon         string `json:"guard_icon"`
					GuardLevel        int64  `json:"guard_level"`
					HonorIcon         string `json:"honor_icon"`
					Intimacy          int64  `json:"intimacy"`
					IsReceive         int64  `json:"is_receive"`
					Level             int64  `json:"level"`
					MasterStatus      int64  `json:"master_status"`
					MedalColor        int64  `json:"medal_color"`
					MedalColorBorder  int64  `json:"medal_color_border"`
					MedalColorEnd     int64  `json:"medal_color_end"`
					MedalColorLevel   int64  `json:"medal_color_level"`
					MedalColorName    int64  `json:"medal_color_name"`
					MedalID           int64  `json:"medal_id"`
					MedalLevelBgColor int64  `json:"medal_level_bg_color"`
					MedalName         string `json:"medal_name"`
					Score             int64  `json:"score"`
					UID               int64  `json:"uid"`
				} `json:"fans_detail"`
				IsContractor   bool  `json:"is_contractor"`
				IsSeniorMember int64 `json:"is_senior_member"`
				LevelInfo      struct {
					CurrentExp   int64 `json:"current_exp"`
					CurrentLevel int64 `json:"current_level"`
					CurrentMin   int64 `json:"current_min"`
					NextExp      int64 `json:"next_exp"`
				} `json:"level_info"`
				Mid       string `json:"mid"`
				Nameplate struct {
					Condition  string `json:"condition"`
					Image      string `json:"image"`
					ImageSmall string `json:"image_small"`
					Level      string `json:"level"`
					Name       string `json:"name"`
					Nid        int64  `json:"nid"`
				} `json:"nameplate"`
				NftInteraction interface{} `json:"nft_interaction"`
				OfficialVerify struct {
					Desc string `json:"desc"`
					Type int64  `json:"type"`
				} `json:"official_verify"`
				Pendant struct {
					Expire            int64  `json:"expire"`
					Image             string `json:"image"`
					ImageEnhance      string `json:"image_enhance"`
					ImageEnhanceFrame string `json:"image_enhance_frame"`
					Name              string `json:"name"`
					Pid               int64  `json:"pid"`
				} `json:"pendant"`
				Rank   string `json:"rank"`
				Senior struct {
					Status int64 `json:"status"`
				} `json:"senior"`
				Sex         string `json:"sex"`
				Sign        string `json:"sign"`
				Uname       string `json:"uname"`
				UserSailing struct {
					Cardbg struct {
						Fan struct {
							Color   string `json:"color"`
							IsFan   int64  `json:"is_fan"`
							Name    string `json:"name"`
							NumDesc string `json:"num_desc"`
							Number  int64  `json:"number"`
						} `json:"fan"`
						ID      int64  `json:"id"`
						Image   string `json:"image"`
						JumpURL string `json:"jump_url"`
						Name    string `json:"name"`
						Type    string `json:"type"`
					} `json:"cardbg"`
					CardbgWithFocus interface{} `json:"cardbg_with_focus"`
					Pendant         struct {
						ID                int64  `json:"id"`
						Image             string `json:"image"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
						JumpURL           string `json:"jump_url"`
						Name              string `json:"name"`
						Type              string `json:"type"`
					} `json:"pendant"`
				} `json:"user_sailing"`
				Vip struct {
					AccessStatus    int64  `json:"accessStatus"`
					AvatarSubscript int64  `json:"avatar_subscript"`
					DueRemark       string `json:"dueRemark"`
					Label           struct {
						BgColor               string `json:"bg_color"`
						BgStyle               int64  `json:"bg_style"`
						BorderColor           string `json:"border_color"`
						ImgLabelURIHans       string `json:"img_label_uri_hans"`
						ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
						ImgLabelURIHant       string `json:"img_label_uri_hant"`
						ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
						LabelTheme            string `json:"label_theme"`
						Path                  string `json:"path"`
						Text                  string `json:"text"`
						TextColor             string `json:"text_color"`
						UseImgLabel           bool   `json:"use_img_label"`
					} `json:"label"`
					NicknameColor string `json:"nickname_color"`
					ThemeType     int64  `json:"themeType"`
					VipDueDate    int64  `json:"vipDueDate"`
					VipStatus     int64  `json:"vipStatus"`
					VipStatusWarn string `json:"vipStatusWarn"`
					VipType       int64  `json:"vipType"`
				} `json:"vip"`
			} `json:"member"`
			Mid       int64  `json:"mid"`
			Oid       int64  `json:"oid"`
			Parent    int64  `json:"parent"`
			ParentStr string `json:"parent_str"`
			Rcount    int64  `json:"rcount"`
			Replies   []struct {
				Action  int64 `json:"action"`
				Assist  int64 `json:"assist"`
				Attr    int64 `json:"attr"`
				Content struct {
					AtNameToMid struct {
						E班小渚   int64 `json:"E班小渚"`
						一只黄花   int64 `json:"一只黄花"`
						包桑爱绫波丽 int64 `json:"包桑爱绫波丽"`
						媚东阳    int64 `json:"媚东阳"`
						废浅年华   int64 `json:"废浅年华"`
						狮子玉佩   int64 `json:"狮子玉佩"`
						贫穷会员号  int64 `json:"贫穷会员号"`
					} `json:"at_name_to_mid"`
					Emote struct {
						Tv坏笑 struct {
							Attr      int64  `json:"attr"`
							GifURL    string `json:"gif_url"`
							ID        int64  `json:"id"`
							JumpTitle string `json:"jump_title"`
							Meta      struct {
								Size int64 `json:"size"`
							} `json:"meta"`
							Mtime     int64  `json:"mtime"`
							PackageID int64  `json:"package_id"`
							State     int64  `json:"state"`
							Text      string `json:"text"`
							Type      int64  `json:"type"`
							URL       string `json:"url"`
						} `json:"[tv_坏笑]"`
						Tv大哭 struct {
							Attr      int64  `json:"attr"`
							GifURL    string `json:"gif_url"`
							ID        int64  `json:"id"`
							JumpTitle string `json:"jump_title"`
							Meta      struct {
								Size int64 `json:"size"`
							} `json:"meta"`
							Mtime     int64  `json:"mtime"`
							PackageID int64  `json:"package_id"`
							State     int64  `json:"state"`
							Text      string `json:"text"`
							Type      int64  `json:"type"`
							URL       string `json:"url"`
						} `json:"[tv_大哭]"`
						兔年 struct {
							Attr      int64  `json:"attr"`
							ID        int64  `json:"id"`
							JumpTitle string `json:"jump_title"`
							Meta      struct {
								Size int64 `json:"size"`
							} `json:"meta"`
							Mtime     int64  `json:"mtime"`
							PackageID int64  `json:"package_id"`
							State     int64  `json:"state"`
							Text      string `json:"text"`
							Type      int64  `json:"type"`
							URL       string `json:"url"`
						} `json:"[兔年]"`
						大哭 struct {
							Attr      int64  `json:"attr"`
							ID        int64  `json:"id"`
							JumpTitle string `json:"jump_title"`
							Meta      struct {
								Size int64 `json:"size"`
							} `json:"meta"`
							Mtime     int64  `json:"mtime"`
							PackageID int64  `json:"package_id"`
							State     int64  `json:"state"`
							Text      string `json:"text"`
							Type      int64  `json:"type"`
							URL       string `json:"url"`
						} `json:"[大哭]"`
						笑哭 struct {
							Attr      int64  `json:"attr"`
							ID        int64  `json:"id"`
							JumpTitle string `json:"jump_title"`
							Meta      struct {
								Size int64 `json:"size"`
							} `json:"meta"`
							Mtime     int64  `json:"mtime"`
							PackageID int64  `json:"package_id"`
							State     int64  `json:"state"`
							Text      string `json:"text"`
							Type      int64  `json:"type"`
							URL       string `json:"url"`
						} `json:"[笑哭]"`
						给心心 struct {
							Attr      int64  `json:"attr"`
							ID        int64  `json:"id"`
							JumpTitle string `json:"jump_title"`
							Meta      struct {
								Size int64 `json:"size"`
							} `json:"meta"`
							Mtime     int64  `json:"mtime"`
							PackageID int64  `json:"package_id"`
							State     int64  `json:"state"`
							Text      string `json:"text"`
							Type      int64  `json:"type"`
							URL       string `json:"url"`
						} `json:"[给心心]"`
						辣眼睛 struct {
							Attr      int64  `json:"attr"`
							ID        int64  `json:"id"`
							JumpTitle string `json:"jump_title"`
							Meta      struct {
								Size int64 `json:"size"`
							} `json:"meta"`
							Mtime     int64  `json:"mtime"`
							PackageID int64  `json:"package_id"`
							State     int64  `json:"state"`
							Text      string `json:"text"`
							Type      int64  `json:"type"`
							URL       string `json:"url"`
						} `json:"[辣眼睛]"`
					} `json:"emote"`
					JumpURL struct{} `json:"jump_url"`
					MaxLine int64    `json:"max_line"`
					Members []struct {
						Avatar         string `json:"avatar"`
						FaceNftNew     int64  `json:"face_nft_new"`
						IsSeniorMember int64  `json:"is_senior_member"`
						LevelInfo      struct {
							CurrentExp   int64 `json:"current_exp"`
							CurrentLevel int64 `json:"current_level"`
							CurrentMin   int64 `json:"current_min"`
							NextExp      int64 `json:"next_exp"`
						} `json:"level_info"`
						Mid       string `json:"mid"`
						Nameplate struct {
							Condition  string `json:"condition"`
							Image      string `json:"image"`
							ImageSmall string `json:"image_small"`
							Level      string `json:"level"`
							Name       string `json:"name"`
							Nid        int64  `json:"nid"`
						} `json:"nameplate"`
						OfficialVerify struct {
							Desc string `json:"desc"`
							Type int64  `json:"type"`
						} `json:"official_verify"`
						Pendant struct {
							Expire            int64  `json:"expire"`
							Image             string `json:"image"`
							ImageEnhance      string `json:"image_enhance"`
							ImageEnhanceFrame string `json:"image_enhance_frame"`
							Name              string `json:"name"`
							Pid               int64  `json:"pid"`
						} `json:"pendant"`
						Rank   string   `json:"rank"`
						Senior struct{} `json:"senior"`
						Sex    string   `json:"sex"`
						Sign   string   `json:"sign"`
						Uname  string   `json:"uname"`
						Vip    struct {
							AccessStatus    int64  `json:"accessStatus"`
							AvatarSubscript int64  `json:"avatar_subscript"`
							DueRemark       string `json:"dueRemark"`
							Label           struct {
								BgColor               string `json:"bg_color"`
								BgStyle               int64  `json:"bg_style"`
								BorderColor           string `json:"border_color"`
								ImgLabelURIHans       string `json:"img_label_uri_hans"`
								ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
								ImgLabelURIHant       string `json:"img_label_uri_hant"`
								ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
								LabelTheme            string `json:"label_theme"`
								Path                  string `json:"path"`
								Text                  string `json:"text"`
								TextColor             string `json:"text_color"`
								UseImgLabel           bool   `json:"use_img_label"`
							} `json:"label"`
							NicknameColor string `json:"nickname_color"`
							ThemeType     int64  `json:"themeType"`
							VipDueDate    int64  `json:"vipDueDate"`
							VipStatus     int64  `json:"vipStatus"`
							VipStatusWarn string `json:"vipStatusWarn"`
							VipType       int64  `json:"vipType"`
						} `json:"vip"`
					} `json:"members"`
					Message string `json:"message"`
				} `json:"content"`
				Count        int64  `json:"count"`
				Ctime        int64  `json:"ctime"`
				Dialog       int64  `json:"dialog"`
				DynamicIDStr string `json:"dynamic_id_str"`
				Fansgrade    int64  `json:"fansgrade"`
				Folder       struct {
					HasFolded bool   `json:"has_folded"`
					IsFolded  bool   `json:"is_folded"`
					Rule      string `json:"rule"`
				} `json:"folder"`
				Invisible bool  `json:"invisible"`
				Like      int64 `json:"like"`
				Member    struct {
					Avatar     string `json:"avatar"`
					AvatarItem struct {
						ContainerSize struct {
							Height float64 `json:"height"`
							Width  float64 `json:"width"`
						} `json:"container_size"`
						FallbackLayers struct {
							IsCriticalGroup bool `json:"is_critical_group"`
							Layers          []struct {
								GeneralSpec struct {
									PosSpec struct {
										AxisX         float64 `json:"axis_x"`
										AxisY         float64 `json:"axis_y"`
										CoordinatePos int64   `json:"coordinate_pos"`
									} `json:"pos_spec"`
									RenderSpec struct {
										Opacity int64 `json:"opacity"`
									} `json:"render_spec"`
									SizeSpec struct {
										Height int64 `json:"height"`
										Width  int64 `json:"width"`
									} `json:"size_spec"`
								} `json:"general_spec"`
								LayerConfig struct {
									IsCritical bool `json:"is_critical"`
									LayerMask  struct {
										GeneralSpec struct {
											PosSpec struct {
												AxisX         float64 `json:"axis_x"`
												AxisY         float64 `json:"axis_y"`
												CoordinatePos int64   `json:"coordinate_pos"`
											} `json:"pos_spec"`
											RenderSpec struct {
												Opacity int64 `json:"opacity"`
											} `json:"render_spec"`
											SizeSpec struct {
												Height int64 `json:"height"`
												Width  int64 `json:"width"`
											} `json:"size_spec"`
										} `json:"general_spec"`
										MaskSrc struct {
											Draw struct {
												ColorConfig struct {
													Day struct {
														Argb string `json:"argb"`
													} `json:"day"`
												} `json:"color_config"`
												DrawType int64 `json:"draw_type"`
												FillMode int64 `json:"fill_mode"`
											} `json:"draw"`
											SrcType int64 `json:"src_type"`
										} `json:"mask_src"`
									} `json:"layer_mask"`
									Tags struct {
										AvatarLayer  struct{} `json:"AVATAR_LAYER"`
										IconLayer    struct{} `json:"ICON_LAYER"`
										PendentLayer struct{} `json:"PENDENT_LAYER"`
									} `json:"tags"`
								} `json:"layer_config"`
								Resource struct {
									ResAnimation struct {
										WebpSrc struct {
											Placeholder int64 `json:"placeholder"`
											Remote      struct {
												BfsStyle string `json:"bfs_style"`
												URL      string `json:"url"`
											} `json:"remote"`
											SrcType int64 `json:"src_type"`
										} `json:"webp_src"`
									} `json:"res_animation"`
									ResImage struct {
										ImageSrc struct {
											Local       int64 `json:"local"`
											Placeholder int64 `json:"placeholder"`
											Remote      struct {
												BfsStyle string `json:"bfs_style"`
												URL      string `json:"url"`
											} `json:"remote"`
											SrcType int64 `json:"src_type"`
										} `json:"image_src"`
									} `json:"res_image"`
									ResNativeDraw struct {
										DrawSrc struct {
											Draw struct {
												ColorConfig struct {
													Day struct {
														Argb string `json:"argb"`
													} `json:"day"`
													IsDarkModeAware bool `json:"is_dark_mode_aware"`
													Night           struct {
														Argb string `json:"argb"`
													} `json:"night"`
												} `json:"color_config"`
												DrawType int64 `json:"draw_type"`
												FillMode int64 `json:"fill_mode"`
											} `json:"draw"`
											SrcType int64 `json:"src_type"`
										} `json:"draw_src"`
									} `json:"res_native_draw"`
									ResType int64 `json:"res_type"`
								} `json:"resource"`
								Visible bool `json:"visible"`
							} `json:"layers"`
						} `json:"fallback_layers"`
						Mid string `json:"mid"`
					} `json:"avatar_item"`
					ContractDesc string `json:"contract_desc"`
					FaceNftNew   int64  `json:"face_nft_new"`
					FansDetail   struct {
						GuardIcon         string `json:"guard_icon"`
						GuardLevel        int64  `json:"guard_level"`
						HonorIcon         string `json:"honor_icon"`
						Intimacy          int64  `json:"intimacy"`
						IsReceive         int64  `json:"is_receive"`
						Level             int64  `json:"level"`
						MasterStatus      int64  `json:"master_status"`
						MedalColor        int64  `json:"medal_color"`
						MedalColorBorder  int64  `json:"medal_color_border"`
						MedalColorEnd     int64  `json:"medal_color_end"`
						MedalColorLevel   int64  `json:"medal_color_level"`
						MedalColorName    int64  `json:"medal_color_name"`
						MedalID           int64  `json:"medal_id"`
						MedalLevelBgColor int64  `json:"medal_level_bg_color"`
						MedalName         string `json:"medal_name"`
						Score             int64  `json:"score"`
						UID               int64  `json:"uid"`
					} `json:"fans_detail"`
					IsContractor   bool  `json:"is_contractor"`
					IsSeniorMember int64 `json:"is_senior_member"`
					LevelInfo      struct {
						CurrentExp   int64 `json:"current_exp"`
						CurrentLevel int64 `json:"current_level"`
						CurrentMin   int64 `json:"current_min"`
						NextExp      int64 `json:"next_exp"`
					} `json:"level_info"`
					Mid       string `json:"mid"`
					Nameplate struct {
						Condition  string `json:"condition"`
						Image      string `json:"image"`
						ImageSmall string `json:"image_small"`
						Level      string `json:"level"`
						Name       string `json:"name"`
						Nid        int64  `json:"nid"`
					} `json:"nameplate"`
					NftInteraction struct {
						Region struct {
							Icon       string `json:"icon"`
							ShowStatus int64  `json:"show_status"`
							Type       int64  `json:"type"`
						} `json:"region"`
					} `json:"nft_interaction"`
					OfficialVerify struct {
						Desc string `json:"desc"`
						Type int64  `json:"type"`
					} `json:"official_verify"`
					Pendant struct {
						Expire            int64  `json:"expire"`
						Image             string `json:"image"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
						Name              string `json:"name"`
						Pid               int64  `json:"pid"`
					} `json:"pendant"`
					Rank   string `json:"rank"`
					Senior struct {
						Status int64 `json:"status"`
					} `json:"senior"`
					Sex         string `json:"sex"`
					Sign        string `json:"sign"`
					Uname       string `json:"uname"`
					UserSailing struct {
						Cardbg struct {
							Fan struct {
								Color   string `json:"color"`
								IsFan   int64  `json:"is_fan"`
								Name    string `json:"name"`
								NumDesc string `json:"num_desc"`
								Number  int64  `json:"number"`
							} `json:"fan"`
							ID      int64  `json:"id"`
							Image   string `json:"image"`
							JumpURL string `json:"jump_url"`
							Name    string `json:"name"`
							Type    string `json:"type"`
						} `json:"cardbg"`
						CardbgWithFocus interface{} `json:"cardbg_with_focus"`
						Pendant         struct {
							ID                int64  `json:"id"`
							Image             string `json:"image"`
							ImageEnhance      string `json:"image_enhance"`
							ImageEnhanceFrame string `json:"image_enhance_frame"`
							JumpURL           string `json:"jump_url"`
							Name              string `json:"name"`
							Type              string `json:"type"`
						} `json:"pendant"`
					} `json:"user_sailing"`
					Vip struct {
						AccessStatus    int64  `json:"accessStatus"`
						AvatarSubscript int64  `json:"avatar_subscript"`
						DueRemark       string `json:"dueRemark"`
						Label           struct {
							BgColor               string `json:"bg_color"`
							BgStyle               int64  `json:"bg_style"`
							BorderColor           string `json:"border_color"`
							ImgLabelURIHans       string `json:"img_label_uri_hans"`
							ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
							ImgLabelURIHant       string `json:"img_label_uri_hant"`
							ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
							LabelTheme            string `json:"label_theme"`
							Path                  string `json:"path"`
							Text                  string `json:"text"`
							TextColor             string `json:"text_color"`
							UseImgLabel           bool   `json:"use_img_label"`
						} `json:"label"`
						NicknameColor string `json:"nickname_color"`
						ThemeType     int64  `json:"themeType"`
						VipDueDate    int64  `json:"vipDueDate"`
						VipStatus     int64  `json:"vipStatus"`
						VipStatusWarn string `json:"vipStatusWarn"`
						VipType       int64  `json:"vipType"`
					} `json:"vip"`
				} `json:"member"`
				Mid          int64       `json:"mid"`
				Oid          int64       `json:"oid"`
				Parent       int64       `json:"parent"`
				ParentStr    string      `json:"parent_str"`
				Rcount       int64       `json:"rcount"`
				Replies      interface{} `json:"replies"`
				ReplyControl struct {
					MaxLine  int64  `json:"max_line"`
					TimeDesc string `json:"time_desc"`
				} `json:"reply_control"`
				Root     int64  `json:"root"`
				RootStr  string `json:"root_str"`
				Rpid     int64  `json:"rpid"`
				RpidStr  string `json:"rpid_str"`
				State    int64  `json:"state"`
				Type     int64  `json:"type"`
				UpAction struct {
					Like  bool `json:"like"`
					Reply bool `json:"reply"`
				} `json:"up_action"`
			} `json:"replies"`
			ReplyControl struct {
				BizScene          string `json:"biz_scene"`
				IsNoteV2          bool   `json:"is_note_v2"`
				MaxLine           int64  `json:"max_line"`
				SubReplyEntryText string `json:"sub_reply_entry_text"`
				SubReplyTitleText string `json:"sub_reply_title_text"`
				TimeDesc          string `json:"time_desc"`
			} `json:"reply_control"`
			Root     int64  `json:"root"`
			RootStr  string `json:"root_str"`
			Rpid     int64  `json:"rpid"`
			RpidStr  string `json:"rpid_str"`
			State    int64  `json:"state"`
			Type     int64  `json:"type"`
			UpAction struct {
				Like  bool `json:"like"`
				Reply bool `json:"reply"`
			} `json:"up_action"`
		} `json:"replies"`
	} `json:"data"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
}
