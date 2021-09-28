package filemanager

import (
	"github.com/GoAdminGroup/filemanager/controller"
	"github.com/GoAdminGroup/filemanager/guard"
	language2 "github.com/GoAdminGroup/filemanager/modules/language"
	"github.com/GoAdminGroup/filemanager/modules/permission"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
)

const cFm = "fm"

func (f *FileManager) InitPlugin2(admin *admin.Admin, srv service.List) {
	if srv == nil { srv = admin.Base.Services }
	f.InitBase(srv, cFm)

	p := permission.Permission{
		AllowUpload:    f.allowUpload,
		AllowCreateDir: f.allowCreateDir,
		AllowDelete:    f.allowDelete,
		AllowMove:      f.allowMove,
		AllowRename:    f.allowRename,
		AllowDownload:  f.allowDownload,
	}

	f.handler = controller.NewHandler(f.roots, p)
	f.guard = guard.New(f.roots, f.Conn, p)
	f.App = f.initRouter2(admin)
	f.handler.HTML = f.HTMLMenu

	//language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)
}

func (f *FileManager) initRouter2(admin *admin.Admin) *context.App {
	app := context.NewApp()
	authRoute := app.Group("/", admin.GlobalErrorHandler(), auth.Middleware(f.Conn))

	authRoute.GET("/", f.guard.Files, f.handler.ListFiles)
	authRoute.GET("/:__prefix/list", f.guard.Files, f.handler.ListFiles)
	authRoute.GET("/:__prefix/download", f.handler.Download)
	authRoute.POST("/:__prefix/upload", f.guard.Upload, f.handler.Upload)
	authRoute.POST("/:__prefix/create/dir/popup", f.handler.CreateDirPopUp)
	authRoute.POST("/:__prefix/create/dir", f.guard.CreateDir, f.handler.CreateDir)
	authRoute.POST("/:__prefix/delete", f.guard.Delete, f.handler.Delete)
	authRoute.POST("/:__prefix/move/popup", f.handler.MovePopup)
	authRoute.POST("/:__prefix/move", f.guard.Move, f.handler.Move)
	authRoute.GET("/:__prefix/preview", f.guard.Preview, f.handler.Preview)
	authRoute.POST("/:__prefix/rename/popup", f.handler.RenamePopUp)
	authRoute.POST("/:__prefix/rename", f.guard.Rename, f.handler.Rename)

	return app
}
