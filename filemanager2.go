package filemanager

import (
	"github.com/GoAdminGroup/filemanager/controller"
	"github.com/GoAdminGroup/filemanager/guard"
	language2 "github.com/GoAdminGroup/filemanager/modules/language"
	"github.com/GoAdminGroup/filemanager/modules/permission"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
)

const cFm = "fm"

func (f *FileManager) InitPlugin2(base *plugins.Base, srv service.List) {
	if srv == nil { srv = base.Services }
	f.InitBase(srv, "fm")
	f.UI = base.UI

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
	f.App = f.initRouter(srv)
	f.handler.HTML = base.HTMLMenu

	//language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)
}
