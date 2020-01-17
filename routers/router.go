// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ali-auction-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/gocron_user",
			beego.NSInclude(
				&controllers.GocronUserController{},
			),
		),

		beego.NSNamespace("/gocron_task",
			beego.NSInclude(
				&controllers.GocronTaskController{},
			),
		),

		beego.NSNamespace("/gocron_task_log",
			beego.NSInclude(
				&controllers.GocronTaskLogController{},
			),
		),

		beego.NSNamespace("/gocron_host",
			beego.NSInclude(
				&controllers.GocronHostController{},
			),
		),

		beego.NSNamespace("/gocron_setting",
			beego.NSInclude(
				&controllers.GocronSettingController{},
			),
		),

		beego.NSNamespace("/gocron_login_log",
			beego.NSInclude(
				&controllers.GocronLoginLogController{},
			),
		),

		beego.NSNamespace("/gocron_task_host",
			beego.NSInclude(
				&controllers.GocronTaskHostController{},
			),
		),

		beego.NSNamespace("/auction_item",
			beego.NSInclude(
				&controllers.AuctionItemController{},
			),
		),

		beego.NSNamespace("/auction_target",
			beego.NSInclude(
				&controllers.AuctionTargetController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
