package routers

import (
	"github.com/astaxie/beego"
	"net/http"
	"star/controllers"
)

func init() {
	//首页
	beego.Router("/", &controllers.MainController{})

	// 能力地图
	beego.Router("/capability", &controllers.CapabilityController{})
	beego.Router("/capability/update", &controllers.CapabilityUpdateController{})
	beego.Router("/capability/addcapability", &controllers.CapabilityAddController{})

	/* 培训交流 */
	beego.Router("/training", &controllers.TrainingController{})
	//
	beego.Router("/training/publish", &controllers.TrainingPublishController{})
	//beego.Router("/training/collect", &controllers.TrainingController{})

	// 新员工交流
	beego.Router("/newemployee", &controllers.CommunicateController{})
	beego.Router("/newemployee/blog", &controllers.CommunicateController{})
	beego.Router("/newemployee/article/:ident", &controllers.CommunicateController{}, "get:Read")
	beego.Router("/newemployee/catalog/:ident", &controllers.CommunicateController{}, "get:ListByCatalog")
	//新员工交流 版主
	beego.Router("/newemployee/admin", &controllers.NAdminController{}, "get:Default")
	beego.Router("/newemployee/admin/catalog/add", &controllers.NCatalogController{}, "get:Add;post:DoAdd")
	beego.Router("/newemployee/admin/catalog/edit", &controllers.NCatalogController{}, "get:Edit;post:DoEdit")
	beego.Router("/newemployee/admin/catalog/del", &controllers.NCatalogController{}, "get:Del")
	beego.Router("/newemployee/admin/article/add", &controllers.NArticleController{}, "get:Add;post:DoAdd")
	beego.Router("/newemployee/admin/article/edit", &controllers.NArticleController{}, "get:Edit;post:DoEdit")
	beego.Router("/newemployee/admin/article/del", &controllers.NArticleController{}, "get:Del")
	beego.Router("/newemployee/admin/article/draft", &controllers.NArticleController{}, "get:Draft")

	// 登录与注册
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	//上传下载文件
	beego.Router("/file", &controllers.FileController{})
	beego.Router("/file/upload", &controllers.FileUploadController{})

	//blog
	beego.AutoRouter(&controllers.ApiController{})

	beego.Router("/blog", &controllers.MainController{})
	beego.Router("/article/:ident", &controllers.MainController{}, "get:Read")
	beego.Router("/catalog/:ident", &controllers.MainController{}, "get:ListByCatalog")

	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")

	//管理员
	beego.Router("/admin", &controllers.AdminController{}, "get:Default")
	beego.Router("/admin/catalog/add", &controllers.CatalogController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/catalog/edit", &controllers.CatalogController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/catalog/del", &controllers.CatalogController{}, "get:Del")
	beego.Router("/admin/article/add", &controllers.ArticleController{}, "get:Add;post:DoAdd")
	beego.Router("/admin/article/edit", &controllers.ArticleController{}, "get:Edit;post:DoEdit")
	beego.Router("/admin/article/del", &controllers.ArticleController{}, "get:Del")
	beego.Router("/admin/article/draft", &controllers.ArticleController{}, "get:Draft")

	beego.Router("/ueditor/go/controller", &controllers.UeditorController{})
	http.Handle("/static/", http.FileServer(http.Dir(".")))

}
