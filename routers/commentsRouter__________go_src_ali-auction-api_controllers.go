package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:AuctionTargetController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronHostController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronLoginLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronSettingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskHostController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronTaskLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"] = append(beego.GlobalControllerRouter["ali-auction-api/controllers:GocronUserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
