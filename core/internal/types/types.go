// Code generated by goctl. DO NOT EDIT.
package types

type Loginreq struct {
	Name     string `form:"name"`
	PassWord string `form:"password"`
}

type Loginres struct {
	Token string `form:"token"`
}

type Userregisterreq struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"mail"`
	Code     string `form:"code"`
}

type Userregisterres struct {
}

type MailCodeSendreq struct {
	Email string `form:"email"`
}

type MailCodeSendres struct {
	Code string `form:"code"`
}

type Userdetailreq struct {
	Identity string `path:"identity"`
}

type Userdetailres struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type FileUploadreq struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadres struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRepositorySavereq struct {
	ParentId           int    `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveres struct {
	Identity string `json:"identity"`
}

type UserFileListreq struct {
	Id   int `json:"id,optional"`   //parent_id
	Page int `json:"page,optional"` //第几页
	Size int `json:"size,optional"` //每页数量
}

type UserFileListres struct {
	List  []*UserFile `json:"list"`
	Count int         `json:"count"`
}

type UserFile struct {
	Id                 int    `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int    `json:"size"`
}

type UserFileNameUpdatereq struct {
	Identity string `json:"identity"` //(repository_pool -> repository_identity)
	Name     string `json:"name"`     //新name
}

type UserFileNameUpdateres struct {
}

type UserFolderCreatereq struct {
	ParentId int    `json:"parent_id"` //该 文件夹的 父级文件夹 的id(顶层为0)
	Name     string `json:"name"`
}

type UserFolderCreateres struct {
	Identity string `json:"identity"`
}

type UserFolderDeletereq struct {
	Identity string `path:"identity"`
}

type UserFolderDeleteres struct {
}
