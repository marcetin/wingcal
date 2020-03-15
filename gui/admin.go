package main

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/marcetin/wingcal/model"
	"github.com/marcetin/wingcal/pkg/gel"
)

var (
	adminMenuList = &layout.List{
		Axis: layout.Vertical,
	}
	list = &layout.List{
		Axis: layout.Vertical,
	}
	thingEditTitle   = new(gel.Editor)
	thingEditContent = new(gel.Editor)
	post             = new(model.DuoCMSpost)
	dodajDugme       = new(gel.Button)
	kolicina         = &gel.DuoUIcounter{
		Value:        1,
		OperateValue: 1,
		From:         1,
		To:           100,
		CounterInput: &gel.Editor{
			Alignment:  text.Middle,
			SingleLine: true,
		},
		CounterIncrease: new(gel.Button),
		CounterDecrease: new(gel.Button),
		CounterReset:    new(gel.Button),
	}
)

type DuoCMSadmin struct {
	Menu *DuoCMSmenu
}

type DuoCMSmenu struct {
	Title string
	Items map[string]DuoCMSmenuItem
}

type DuoCMSmenuItem struct {
	Title       string
	Description string
	Icon        string
	Link        func()
	subItems    map[string]DuoCMSmenuItem
}

func (w *WingCal) admin() func() {
	return func() {

		//post := &model.DuoCMSpost{
		//	Title:    "",
		//	Subtitle: "",
		//}

		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(w.Context,
			layout.Rigid(func() {
				w.Theme.DuoUIitem(0, w.Theme.Colors["Primary"]).Layout(w.Context, layout.Center, func() {
					w.Theme.H3("Header").Layout(w.Context)
				})
			}),
			layout.Flexed(1, func() {
				layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(w.Context,
					layout.Rigid(func() {
						w.Theme.DuoUIitem(0, w.Theme.Colors["Warning"]).Layout(w.Context, layout.Center, func() {

							list.Layout(w.Context, len(w.Transfered.Elementi), func(i int) {
								element := w.Transfered.Elementi[i]
								layout.UniformInset(unit.Dp(16)).Layout(w.Context, func() {
									if w.Context.Constraints.Width.Max > 500 {
										w.Context.Constraints.Width.Max = 500
									}
									layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
										layout.Rigid(func() {

											btn := w.Theme.Button(element.Naziv)

											for w.ElementsButtons[i].Clicked(w.Context) {
												w.PrikazaniElement = &element
											}

											btn.Layout(w.Context, w.ElementsButtons[i])
										}),
									)

								})
							})

						})
					}),
					layout.Flexed(1, func() {
						layout.Flex{
							Axis: layout.Vertical,
						}.Layout(w.Context,
							layout.Rigid(func() {
								w.Theme.DuoUIitem(0, w.Theme.Colors["Danger"]).Layout(w.Context, layout.Center, func() {
									w.Theme.H3("Content header").Layout(w.Context)
								})
							}),
							layout.Flexed(1, func() {
								w.Theme.DuoUIitem(0, w.Theme.Colors["Secondary"]).Layout(w.Context, layout.Center, func() {
									//thingEdit := []func(){
									//	func() {
									//		w.Theme.DuoUIitem(0, w.Theme.Colors["Light"]).Layout(w.Context, layout.Center, func() {
									//			w.Theme.H6("Title").Layout(w.Context)
									//		})
									//	},
									//	func() {
									//		w.Theme.DuoUIitem(0, w.Theme.Colors["Light"]).Layout(w.Context, layout.Center, func() {
									//			w.Theme.DuoUIeditor("").Layout(w.Context, thingEditTitle)
									//			for _, e := range thingEditTitle.Events(w.Context) {
									//				switch e.(type) {
									//				case gel.ChangeEvent:
									//					post.Title = thingEditTitle.Text()
									//				}
									//			}
									//		})
									//	},
									//	func() {
									//		w.Theme.DuoUIitem(0, w.Theme.Colors["Light"]).Layout(w.Context, layout.Center, func() {
									//			w.Theme.H6("Content").Layout(w.Context)
									//		})
									//	},
									//	func() {
									//		w.Theme.DuoUIitem(0, w.Theme.Colors["Light"]).Layout(w.Context, layout.Center, func() {
									//			w.Theme.DuoUIeditor("").Layout(w.Context, thingEditContent)
									//			for _, e := range thingEditContent.Events(w.Context) {
									//				switch e.(type) {
									//				case gel.ChangeEvent:
									//					post.Subtitle = thingEditContent.Text()
									//				}
									//			}
									//		})
									//	},
									//	func() {
									//		w.Theme.DuoUIitem(0, w.Theme.Colors["Light"]).Layout(w.Context, layout.Center, func() {
									//			for btn.Clicked(w.Context) {
									//				//marshalled, _ := js.Marshal(WingCalElement{
									//				//	//Data: &post,
									//				//})
									//
									//				//fmt.Println("post:",post)
									//				//connection.Write(marshalled)
									//			}
									//			w.Theme.Button("Click me!").Layout(w.Context, btn)
									//		})
									//	},
									//}

									//list.Layout(w.Context, len(thingEdit), func(i int) {
									//	layout.UniformInset(unit.Dp(16)).Layout(w.Context, thingEdit[i])
									//})
									sumaCena := float64(kolicina.Value) * w.PrikazaniElement.Cena
									layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
										layout.Rigid(func() {
											w.Theme.H5(w.PrikazaniElement.Naziv).Layout(w.Context)
										}),
										layout.Rigid(func() {
											w.Theme.Body1(w.PrikazaniElement.Opis).Layout(w.Context)
										}),
										layout.Rigid(func() {
											w.Theme.Caption(w.PrikazaniElement.Obracun).Layout(w.Context)
										}),
										layout.Rigid(func() {
											w.Theme.H6("Cena:" + fmt.Sprint(w.PrikazaniElement.Cena)).Layout(w.Context)
										}),
										layout.Rigid(func() {
											w.Theme.DuoUIcounter(func() {}).Layout(w.Context, kolicina, "KOLICINA", fmt.Sprint(kolicina.Value))
										}),

										layout.Rigid(func() {
											btn := w.Theme.Button("DODAJ")

											suma := WingIzabraniElement{
												Kolicina: kolicina.Value,
												SumaCena: sumaCena,
												Element:  *w.PrikazaniElement,
											}
											for dodajDugme.Clicked(w.Context) {
												w.Suma.Elementi = append(w.Suma.Elementi, suma)
											}

											btn.Layout(w.Context, dodajDugme)
										}),

										layout.Rigid(func() {
											w.Theme.H6("Suma:" + fmt.Sprint(sumaCena)).Layout(w.Context)
										}),
									)

								})

							}),
							layout.Rigid(func() {
								w.Theme.DuoUIitem(0, w.Theme.Colors["Primary"]).Layout(w.Context, layout.Center, func() {
									w.Theme.H3("Content footer").Layout(w.Context)
								})
							}))
					}),
					layout.Rigid(func() {
						w.Theme.DuoUIitem(0, w.Theme.Colors["Danger"]).Layout(w.Context, layout.Center, func() {
							var sumaSumarum float64
							layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
								layout.Rigid(func() {
									list.Layout(w.Context, len(w.Suma.Elementi), func(i int) {
										element := w.Suma.Elementi[i]
										sumaSumarum = sumaSumarum + element.SumaCena
										layout.UniformInset(unit.Dp(16)).Layout(w.Context, func() {
											if w.Context.Constraints.Width.Max > 500 {
												w.Context.Constraints.Width.Max = 500
											}
											layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
												layout.Rigid(func() {
													layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(w.Context,
														layout.Rigid(func() {
															w.Theme.Body1(element.Element.Naziv).Layout(w.Context)
														}),
														layout.Rigid(func() {
															w.Theme.H6(fmt.Sprint(element.Kolicina)).Layout(w.Context)
														}),
														layout.Rigid(func() {
															w.Theme.H5(fmt.Sprint(element.SumaCena)).Layout(w.Context)
														}))
												}),
											)

										})
									})
								}),
								layout.Rigid(func() {
									suma := w.Theme.H5("Suma: " + fmt.Sprint(sumaSumarum))
									suma.Alignment = text.End
									suma.Layout(w.Context)
								}),
							)

						})
					}))
			}),
			layout.Rigid(func() {
				w.Theme.DuoUIitem(0, w.Theme.Colors["DarkGray"]).Layout(w.Context, layout.Center, func() {
					w.Theme.H3("Footer").Layout(w.Context)
				})
			}))
	}
}

//func navMenuButton(rc *rcd.RcVar, w.Context *layout.Context, th *gelook.DuoUItheme, page *gelook.DuoUIpage, title, icon string, navButton *gel.Button) func() {
//	return func() {
//		layout.UniformInset(unit.Dp(0)).Layout(w.Context, func() {
//			var menuItem gelook.DuoUIbutton
//			menuItem = w.Theme.DuoUIbutton(w.Theme.Fonts["Secondary"], title, w.Theme.Colors["Dark"], w.Theme.Colors["LightGrayII"], w.Theme.Colors["LightGrayII"], w.Theme.Colors["Dark"], icon, CurrentCurrentPageColor(rc.ShowPage, title, navItemIconColor, w.Theme.Colors["Primary"]), navItemTextSize, navItemTconSize, navItemWidth, navItemHeight, navItemPaddingVertical, navItemPaddingHorizontal)
//			for navButton.Clicked(w.Context) {
//				//rc.ShowPage = title
//				//page.Command()
//				//SetPage(rc, page)
//			}
//			menuItem.MenuLayout(w.Context, navButton)
//		})
//	}
//}

//func navButtons(rc *rcd.RcVar, w.Context *layout.Context, th *gelook.DuoUItheme, allPages *model.DuoUIpages) []func() {
//	return []func(){
//		navMenuButton(rc, w.Context, th, allPages.Theme["OVERVIEW"], "OVERVIEW", "overviewIcon", navButtonOverview),
//		w.Theme.DuoUIline(w.Context, "LightGrayIII"),
//		navMenuButton(rc, w.Context, th, allPages.Theme["SEND"], "SEND", "sendIcon", navButtonSend),
//		w.Theme.DuoUIline(w.Context, "LightGrayIII"),
//		navMenuButton(rc, w.Context, th, allPages.Theme["RECEIVE"], "RECEIVE", "receiveIcon", navButtonReceive),
//		w.Theme.DuoUIline(w.Context, "LightGrayIII"),
//		navMenuButton(rc, w.Context, th, allPages.Theme["ADDRESSBOOK"], "ADDRESSBOOK", "addressBookIcon", navButtonAddressBook),
//		w.Theme.DuoUIline(w.Context, "LightGrayIII"),
//		navMenuButton(rc, w.Context, th, allPages.Theme["HISTORY"], "HISTORY", "historyIcon", navButtonHistory),
//	}
//}
//
//func adminMenu(w.Context *layout.Context) {
//	adminMenuList.Layout(w.Context, len(navButtons(rc, w.Context, th, allPages)), func(i int) {
//		layout.UniformInset(unit.Dp(0)).Layout(w.Context, navButtons(rc, w.Context, th, allPages)[i])
//	})
//}
