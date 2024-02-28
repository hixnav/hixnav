"use strict";(self["webpackChunkHixNav"]=self["webpackChunkHixNav"]||[]).push([[443],{9647:function(t,e,a){a.r(e),a.d(e,{default:function(){return d}});var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"account"}},[a("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),a("div",{staticClass:"content"},[a("el-container",[a("el-container",[a("el-main",[a("el-button",{attrs:{type:"primary",size:"mini"},on:{click:t.openAddDrawer}},[a("i",{staticClass:"el-icon-edit"}),t._v("添加账号 ")]),a("el-table",{staticStyle:{width:"100%","margin-top":"20px"},attrs:{data:t.accountList,"row-class-name":t.tableRowClassName}},[a("el-table-column",{attrs:{prop:"Sitename",label:"站点",width:"140"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("a",{attrs:{href:e.row.Siteurl,target:"_blank"}},[t._v(t._s(e.row.Sitename))])]}}])}),a("el-table-column",{attrs:{prop:"Name",label:"账号",width:"160"}}),a("el-table-column",{attrs:{prop:"Password",label:"密码",width:"140"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("span",[t._v("******")])]}}])}),a("el-table-column",{attrs:{prop:"CreateTime",label:"添加时间"}}),a("el-table-column",{attrs:{fixed:"right",prop:"",label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{size:"mini"},nativeOn:{click:function(a){return a.preventDefault(),t.viewSecret(e.$index,t.accountList)}}},[a("i",{staticClass:"el-icon-view"}),t._v("查看密码 ")]),a("el-button",{attrs:{size:"mini"},nativeOn:{click:function(a){return a.preventDefault(),t.openEditDrawer(e.$index,t.accountList)}}},[a("i",{staticClass:"el-icon-edit"}),t._v("编辑 ")]),a("el-button",{attrs:{size:"mini",type:"danger"},nativeOn:{click:function(a){return a.preventDefault(),t.deleteRow(e.$index,t.accountList)}}},[a("i",{staticClass:"el-icon-delete"}),t._v("删除 ")])]}}])})],1)],1)],1)],1)],1),a("div",{staticClass:"drawer"},[a("el-drawer",{ref:"drawer",attrs:{title:t.addShow?"添加账号":"编辑账号",visible:t.dialog,direction:"rtl","custom-class":"demo-drawer"},on:{"update:visible":function(e){t.dialog=e}}},[a("div",{staticClass:"demo-drawer__content",staticStyle:{padding:"30px"}},[a("el-form",{attrs:{model:t.form}},[a("el-form-item",{attrs:{label:"名称","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_label",autocomplete:"off"},model:{value:t.form.sitename,callback:function(e){t.$set(t.form,"sitename",e)},expression:"form.sitename"}})],1),a("el-form-item",{attrs:{label:"地址","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_label",autocomplete:"off"},model:{value:t.form.siteurl,callback:function(e){t.$set(t.form,"siteurl",e)},expression:"form.siteurl"}})],1),a("el-form-item",{attrs:{label:"账号","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_name",autocomplete:"off"},model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}})],1),a("el-form-item",{attrs:{label:"秘钥","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_url",type:"password"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1)],1),a("div",{staticClass:"demo-drawer__footer",staticStyle:{float:"right"}},[a("el-button",{on:{click:t.cancelForm}},[t._v("取 消")]),a("el-button",{directives:[{name:"show",rawName:"v-show",value:t.addShow,expression:"addShow"}],attrs:{type:"primary"},on:{click:t.addAccount}},[t._v("确 定")]),a("el-button",{directives:[{name:"show",rawName:"v-show",value:t.editShow,expression:"editShow"}],attrs:{type:"primary"},on:{click:t.editAccount}},[t._v("确 定")])],1)],1)])],1)],1)},s=[],l=(a(8309),a(561),a(8862),a(8932)),o={name:"AccountView",components:{HeadBar:l.Z},data:function(){return{searchVal:"",activeIndex:"5",accountList:[],dialog:!1,addShow:!0,editShow:!1,form:{id:0,sitename:"",siteurl:"",name:"",password:""}}},methods:{openAddDrawer:function(){console.log("drawer"),this.addShow=!0,this.editShow=!1,this.dialog=!0,this.form={}},openEditDrawer:function(t,e){console.log(e[t]),this.addShow=!1,this.editShow=!0,this.dialog=!0,this.form.id=e[t].ID,this.form.sitename=e[t].Sitename,this.form.siteurl=e[t].Siteurl,this.form.name=e[t].Name},viewSecret:function(t,e){var a=this;this.$store.dispatch("account/accountScrete","id="+e[t].ID).then((function(t){a.$alert(t.data.Password,"您的秘钥",{confirmButtonText:"确定"})})).catch((function(t){console.log(t),a.$message.error("失败")}))},cancelForm:function(){this.dialog=!1},deleteRow:function(t,e){var a=this;console.log(e[t]),this.$confirm("此操作将永久删除该数据, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){a.deleteAccountData(e[t].ID),e.splice(t,1)})).catch((function(){}))},editAccount:function(){var t=this;console.log(this.form);var e=this;this.$store.dispatch("account/editAccount",JSON.stringify(this.form)).then((function(a){console.log(a),e.dialog=!1,e.$notify({title:"成功",message:"",type:"success"}),t.getData()})).catch((function(e){console.log(e),t.$message.error("失败")}))},addAccount:function(){var t=this;console.log(this.form);var e=this;this.$store.dispatch("account/addAccount",JSON.stringify(this.form)).then((function(a){console.log(a),e.dialog=!1,e.$notify({title:"成功",message:"",type:"success"}),t.getData()})).catch((function(e){console.log(e),t.$message.error("失败")}))},deleteAccountData:function(t){var e=this,a=this;this.$store.dispatch("account/delAccount","id="+t).then((function(t){console.log(t),a.$notify({title:"成功",message:"",type:"success"}),a.getData()})).catch((function(t){console.log(t),e.$message.error("失败")}))},getData:function(){var t=this;this.$store.dispatch("account/account",JSON.stringify(this.ruleForm)).then((function(e){console.log(e),t.accountList=e.data.accounts})).catch((function(t){console.log(t)}))}},created:function(){this.getData()}},n=o,r=a(1001),c=(0,r.Z)(n,i,s,!1,null,"860e0ae0",null),d=c.exports},969:function(t,e,a){a.r(e),a.d(e,{default:function(){return m}});var i=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{attrs:{id:"article"}},[i("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),i("div",{staticClass:"bg-banner"},[i("el-row",{staticStyle:{"margin-left":"0px","margin-right":"0px"},attrs:{gutter:16}})],1),i("div",{staticClass:"main"},[i("el-row",{staticClass:"el-row",attrs:{gutter:20}},[i("section",{},[i("section",{staticStyle:{display:"flex","flex-direction":"row","justify-content":"flex-start","align-items":"center"}},[i("img",{staticStyle:{display:"inline-block",width:"10px",height:"20px",padding:"0 10px 0","margin-top":"2px"},attrs:{src:a(3174),alt:""}}),i("div",{staticStyle:{"font-weight":"700","font-size":"18px"}},[t._v("文档")])])])]),i("el-row",{attrs:{gutter:20}},[i("div",{staticClass:"main_card"},[i("el-descriptions",{staticClass:"margin-top",attrs:{title:"",colon:!1,column:"4"}},[i("template",{slot:"extra"}),t._l(t.docLinks,(function(e){return i("el-descriptions-item",{key:e},[i("el-link",{attrs:{href:e.Url,target:"_blank"}},[i("el-avatar",{staticStyle:{width:"20px",height:"20px",margin:"0 10px","background-color":"#fff"},attrs:{src:e.Logo,size:"small"}}),t._v(" "+t._s(e.Name))],1),t.showBtn?i("el-button-group",[i("el-button",{staticClass:"operator-btn",attrs:{type:"text",icon:"el-icon-edit",size:"mini",circle:""},on:{click:function(a){return t.openDrawer(e)}}}),i("el-button",{staticClass:"operator-btn",attrs:{type:"text",icon:"el-icon-delete",size:"mini",circle:""},on:{click:function(a){return t.deleteNav(e.Id)}}})],1):t._e()],1)}))],2)],1)]),i("el-row",{staticClass:"el-row",attrs:{gutter:20}},[i("el-col",{attrs:{span:2}},[i("section",{staticStyle:{display:"flex","flex-direction":"row","justify-content":"flex-start","align-items":"center","padding-top":"2rem"}},[i("img",{staticStyle:{display:"inline-block",width:"10px",height:"20px",padding:"0 10px 0","margin-top":"2px"},attrs:{src:a(3174),alt:""}}),i("div",{staticStyle:{"font-weight":"700","font-size":"18px"}},[t._v("链接")])])]),i("el-col",{staticStyle:{padding:"20px 0"},attrs:{span:20}},[i("el-row",{attrs:{gutter:20}},t._l(t.linkCates,(function(e){return i("el-button",{key:e,staticStyle:{display:"inline",margin:"0 20px"},attrs:{size:"small",round:""}},[t._v(t._s(e.Catename))])})),1)],1)],1),i("el-row",{attrs:{gutter:20}},[i("div",{staticClass:"main_card"},[i("el-descriptions",{staticClass:"margin-top",attrs:{title:"",colon:!1,column:"4"}},[i("template",{slot:"extra"}),t._l(t.commonLinks,(function(e){return i("el-descriptions-item",{key:e},[i("el-link",{attrs:{href:e.Url,target:"_blank"}},[i("el-avatar",{staticStyle:{width:"20px",height:"20px",margin:"0 10px","background-color":"#fff"},attrs:{src:e.Logo,size:"small"}}),t._v(" "+t._s(e.Name))],1),t.showBtn?i("el-button-group",[i("el-button",{staticClass:"operator-btn",attrs:{type:"text",icon:"el-icon-edit",size:"mini",circle:""},on:{click:function(a){return t.openDrawer(e)}}}),i("el-button",{staticClass:"operator-btn",attrs:{type:"text",icon:"el-icon-delete",size:"mini",circle:""},on:{click:function(a){return t.deleteLink(e.Id)}}})],1):t._e()],1)}))],2)],1)])],1),i("el-button",{staticStyle:{position:"fixed",right:"40px",bottom:"108px","z-index":"9999"},attrs:{type:"primary",icon:"el-icon-plus",circle:""},on:{click:t.openAddDrawer}}),i("div",{staticStyle:{position:"fixed",right:"40px",bottom:"60px"}},[i("el-button",{attrs:{type:"primary",icon:"el-icon-edit-outline",circle:""},on:{click:t.handOffBtn}})],1),i("div",{staticClass:"drawer"},[i("el-drawer",{ref:"drawer",attrs:{title:"添加文链",visible:t.dialog,direction:"rtl","custom-class":"demo-drawer"},on:{"update:visible":function(e){t.dialog=e}}},[i("div",{staticClass:"demo-drawer__content",staticStyle:{padding:"30px"}},[i("el-form",{attrs:{model:t.form}},[i("el-form-item",{attrs:{label:"类别","label-width":t.formLabelWidth}},[i("el-select",{attrs:{placeholder:"类别"},model:{value:t.form.type,callback:function(e){t.$set(t.form,"type",e)},expression:"form.type"}},[i("el-option",{attrs:{label:"链接",value:"1"}}),i("el-option",{attrs:{label:"文档",value:"2"}})],1)],1),i("el-form-item",{attrs:{label:"标签","label-width":t.formLabelWidth}},[i("el-input",{attrs:{id:"linked_label",autocomplete:"off"},model:{value:t.form.catename,callback:function(e){t.$set(t.form,"catename",e)},expression:"form.catename"}})],1),i("el-form-item",{attrs:{label:"名称","label-width":t.formLabelWidth}},[i("el-input",{attrs:{id:"linked_name",autocomplete:"off"},model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}})],1),i("el-form-item",{attrs:{label:"地址","label-width":t.formLabelWidth}},[i("el-input",{attrs:{id:"linked_url"},model:{value:t.form.url,callback:function(e){t.$set(t.form,"url",e)},expression:"form.url"}})],1)],1),i("div",{staticClass:"demo-drawer__footer",staticStyle:{float:"right"}},[i("el-button",{on:{click:t.cancelForm}},[t._v("取 消")]),i("el-button",{attrs:{type:"primary"},on:{click:t.linkedSubmit}},[t._v("确 定")])],1)],1)])],1),i("FootBar")],1)},s=[],l=(a(8309),a(4916),a(3123),a(8862),a(8932)),o=a(8268),n=a(1753),r={name:"ArticleView",components:{HeadBar:l.Z,FootBar:o.Z},data:function(){return{searchVal:"",activeIndex:"2",showBtn:!1,linkCates:[],docLinks:[],commonLinks:[],dialog:!1,formLabelWidth:"80px",form:{name:"",type:"1",url:""}}},methods:{openDrawer:function(t){this.form.id=t.Id,this.form.name=t.Name,this.form.type=t.Type,this.form.url=t.Url,this.form.Logo=t.Logo,this.form.catename=t.Catename,this.dialog=!0},deleteLink:function(t){var e=this;this.$confirm("此操作将永久删除该数据, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){e.deleteLinkData(t)})).catch((function(){}))},deleteLinkData:function(t){var e=this;self=this,this.$store.dispatch("link/delArticleLink","id="+t).then((function(t){e.getLinks(),console.log(t),self.$message({message:"成功",type:"success"}),self.getData()})).catch((function(t){console.log(t)}))},openAddDrawer:function(){console.log("drawer"),this.dialog=!0},handOffBtn:function(){var t=(0,n.LP)();t||this.$router.push("signin"),0==this.showBtn?this.showBtn=!0:this.showBtn=!1},cancelForm:function(){this.dialog=!1},linkedSubmit:function(){var t=this,e=this.form.url,a=e.split("/"),i=a[2];this.form.logo=a[0]+"//"+i+"/favicon.ico",this.form.type=parseInt(this.form.type),console.log(this.form),self=this,this.$store.dispatch("link/addArticleLink",JSON.stringify(this.form)).then((function(e){t.getLinks(),console.log(e),self.dialog=!1,self.$notify({title:"成功",message:"手动刷新页面查看",type:"success"})})).catch((function(e){console.log(e),t.$message.error("失败")}))},getLinks:function(){self=this,this.$store.dispatch("link/article",{Type:2}).then((function(t){self.docLinks=t.data.links})).catch((function(t){console.log(t)})),this.$store.dispatch("link/article",{Type:1,Catename:""}).then((function(t){self.commonLinks=t.data.links})).catch((function(t){console.log(t)}))}},created:function(){this.getLinks()}},c=r,d=a(1001),u=(0,d.Z)(c,i,s,!1,null,"01c278a0",null),m=u.exports},2672:function(t,e,a){a.r(e),a.d(e,{default:function(){return u}});var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"cloud-hos"}},[a("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),a("el-container",{staticClass:"el-container"},[a("el-aside",{attrs:{width:"200px"}},[a("el-menu",{staticClass:"el-menu-vertical-demo",attrs:{"default-active":t.param.active},on:{select:t.handleSelect}},[a("el-menu-item",{attrs:{index:"1"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("文件")])]),a("el-menu-item",{attrs:{index:"2"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("笔记")])]),a("el-menu-item",{attrs:{index:"3"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("脑图")])]),a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("视频")])]),a("el-menu-item",{attrs:{index:"5"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("图片")])])],1)],1),a("el-container",{staticClass:"main"},[a("el-header",[a("el-upload",{staticClass:"upload-file",attrs:{name:t.file,action:"/api/uploadIO",data:t.param,limit:1,"show-file-list":"false","auto-upload":!0,"http-request":t.uploadIO,"with-credentials":!0}},[a("el-button",{attrs:{size:"small",type:"primary"}},[t._v("点击上传")])],1)],1),a("el-main",[a("el-card",{staticClass:"box-card",attrs:{shadow:"never"}},[[a("el-table",{staticStyle:{width:"100%",height:"100%"},attrs:{data:t.fileLists}},[a("el-table-column",{attrs:{prop:"key",label:"文件名",width:"280"}}),a("el-table-column",{attrs:{prop:"size",label:"文件大小"}}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"120"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{type:"text",size:"small"},on:{click:function(a){return t.downFile(e.row)}}},[t._v("下载")]),a("el-button",{attrs:{type:"text",size:"small"},on:{click:function(a){return t.delFile(e.row)}}},[t._v("删除")])]}}])})],1)],a("el-pagination",{staticClass:"main-page",attrs:{background:"",layout:"prev, pager, next",total:120}})],2)],1)],1)],1)],1)},s=[],l=(a(5505),a(7714),a(2801),a(1174),a(1539),a(2472),a(8675),a(2990),a(8927),a(3105),a(5035),a(4345),a(7174),a(2846),a(4731),a(7209),a(6319),a(8867),a(7789),a(3739),a(9368),a(4483),a(2056),a(3462),a(678),a(7462),a(3824),a(5021),a(2974),a(5016),a(8783),a(3948),a(285),a(1637),a(4916),a(3123),a(8932)),o=a(8268),n=(a(9719).Base64,{name:"CloundHos",components:{HeadBar:l.Z,FootBar:o.Z},data:function(){return{searchVal:"",activeIndex:"3",file:"file",param:{active:"1"},fileLists:[]}},methods:{downFile:function(t){t.active=this.param.active,console.log(t),this.$store.dispatch("cloud/downIO",t).then((function(e){console.log(e);for(var a=window.atob(e.data.file),i=e.data.contentLength,s=new Uint8Array(i),l=0;l<i;++l)s[l]=a.charCodeAt(l);var o=window.URL.createObjectURL(new Blob([s],{type:e.data.contentType})),n=document.createElement("a");n.href=o,n.download=t.key,n.style.display="none",n.click(),n.remove()})).catch((function(t){console.log(t)}))},base64ToBlob:function(t){for(var e=t.split(";base64,"),a=e[0].split(":")[1],i=window.atob(e[1]),s=i.length,l=new Uint8Array(s),o=0;o<s;++o)l[o]=i.charCodeAt(o);return new Blob([l],{type:a})},delFile:function(t){var e=this;t.active=this.param.active,console.log(t),this.$store.dispatch("cloud/delIO",t).then((function(t){console.log(t),e.$notify({title:"成功",message:"操作成功",type:"success"}),location.reload()})).catch((function(t){console.log(t),e.$message.error("操作失败")}))},handleSelect:function(t,e){console.log(t,e),this.param.active=t,this.listFile()},listFile:function(){var t=this;this.$store.dispatch("cloud/listIO",this.param).then((function(e){console.log(e),t.fileLists=e.data})).catch((function(t){console.log(t)}))},uploadIO:function(t){var e=this,a=this,i=t.file,s=new window.FormData;s.append(a.file,i),s.append("active",this.param.active),this.$store.dispatch("cloud/uploadIO",s).then((function(t){console.log(t),a.$notify({title:"成功",message:"上传成功",type:"success"}),location.reload()})).catch((function(t){console.log(t),e.$message.error("上传失败")}))}},created:function(){this.listFile()}}),r=n,c=a(1001),d=(0,c.Z)(r,i,s,!1,null,"2dcbfc42",null),u=d.exports},1422:function(t,e,a){a.r(e),a.d(e,{default:function(){return u}});var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"cloud-image"}},[a("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),a("section",{staticClass:"ci-main"},[a("el-row",[a("el-col",{attrs:{span:8,offset:8}},[a("el-upload",{ref:"upload",staticClass:"upload-demo",attrs:{drag:"","auto-upload":!1,"file-list":t.fileList,data:t.param,"list-type":"picture",multiple:"",limit:3,"http-request":t.uploadIO,"with-credentials":!0,"on-success":t.onSuccess}},[a("i",{staticClass:"el-icon-upload"}),a("div",{staticClass:"el-upload__text"},[t._v(" 将文件拖到此处，或"),a("em",[t._v("点击上传")])]),a("div",{staticClass:"el-upload__tip",attrs:{slot:"tip"},slot:"tip"},[a("span",[t._v("只能上传jpg/png文件，且不超过500kb")]),a("el-button",{staticStyle:{"margin-left":"10px"},attrs:{size:"small",type:"primary"},on:{click:t.submitUpload}},[t._v(" 开始上传 ")])],1)]),a("div",{staticStyle:{"margin-top":"15px"}},t._l(t.cloudImages,(function(e,i){return a("el-input",{key:e,ref:"demoInput",refInFor:!0,attrs:{placeholder:"http://"},model:{value:e.file,callback:function(a){t.$set(e,"file",a)},expression:"item.file"}},[a("el-button",{attrs:{slot:"append",icon:"el-icon-document-copy"},on:{click:function(e){return t.copyClicked(i)}},slot:"append"})],1)})),1)],1)],1)],1)],1)},s=[],l=a(8932),o=a(8268),n={name:"CloudImage",components:{HeadBar:l.Z,FootBar:o.Z},data:function(){return{searchVal:"",activeIndex:"4",file:"file",param:{active:"1"},fileList:[],cloudImages:[]}},methods:{submitUpload:function(){this.$refs.upload.submit()},handleRemove:function(t,e){console.log(t,e)},handlePreview:function(t){console.log(t)},onSuccess:function(t,e){console.log(t,e)},copyClicked:function(t){this.$refs.demoInput[t].select(),document.execCommand("copy")},uploadIO:function(t){var e=this,a=this,i=t.file,s=new window.FormData;s.append(a.file,i),this.$store.dispatch("cloud/uploadImg",s).then((function(t){e.cloudImages.push(t.data),a.$notify({title:"成功",message:"上传成功",type:"success"})})).catch((function(t){console.log(t),e.$message.error("上传失败")}))}}},r=n,c=a(1001),d=(0,c.Z)(r,i,s,!1,null,"6b1cde28",null),u=d.exports},8405:function(t,e,a){a.r(e),a.d(e,{default:function(){return u}});var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"docs"}},[a("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),a("el-container",[a("el-aside",{attrs:{width:"200px"}},[a("el-menu",{staticClass:"el-menu-vertical-demo",attrs:{"default-active":"1-4-1",collapse:t.isCollapse},on:{open:t.handleOpen,close:t.handleClose}},[a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("Golang笔记")])]),a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("Linux笔记")])]),a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("开源项目")])]),a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("必会宝典")])])],1)],1),a("el-aside",{attrs:{width:"200px"}},[a("el-menu",{staticClass:"el-menu-vertical-demo",attrs:{"default-active":"1-4-1",collapse:t.isCollapse},on:{open:t.handleOpen,close:t.handleClose}},[a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("Golang笔记")])]),a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("Linux笔记")])]),a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("开源项目")])]),a("el-menu-item",{attrs:{index:"4"}},[a("i",{staticClass:"el-icon-folder"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("必会宝典")])])],1)],1),a("el-container",[a("el-main")],1)],1),a("FootBar")],1)},s=[],l=a(8932),o=a(8268),n={name:"Docs",components:{HeadBar:l.Z,FootBar:o.Z},data:function(){return[]},method:{}},r=n,c=a(1001),d=(0,c.Z)(r,i,s,!1,null,"3bbf943a",null),u=d.exports},9994:function(t,e,a){a.r(e),a.d(e,{default:function(){return d}});var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"setting"}},[a("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),a("div",{staticClass:"content"},[a("el-container",[a("el-aside",{attrs:{width:"200px"}},[a("el-menu",{staticClass:"el-menu-vertical",attrs:{"default-active":"1-4-1",collapse:t.isCollapse},on:{open:t.handleOpen,close:t.handleClose,select:t.handleSelect}},[a("el-menu-item",{attrs:{index:"1"}},[a("i",{staticClass:"el-icon-user"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("多用户")])]),a("el-menu-item",{attrs:{index:"2"}},[a("i",{staticClass:"el-icon-coin"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("数据库配置")])]),a("el-menu-item",{attrs:{index:"3"}},[a("i",{staticClass:"el-icon-box"}),a("span",{attrs:{slot:"title"},slot:"title"},[t._v("存储配置")])])],1)],1),a("el-container",[1==t.showMain?a("el-main",[a("el-button",{attrs:{type:"primary",size:"small"},on:{click:t.openMultiUserDrawer}},[a("i",{staticClass:"el-icon-plus"}),t._v("添加用户 ")]),a("el-table",{staticStyle:{width:"100%","margin-top":"20px"},attrs:{data:t.userList,"row-class-name":t.tableRowClassName}},[a("el-table-column",{attrs:{prop:"Id",label:"序号",width:"180"}}),a("el-table-column",{attrs:{prop:"Account",label:"账号",width:"180"}}),a("el-table-column",{attrs:{prop:"CreateTime",label:"添加时间"}}),a("el-table-column",{attrs:{fixed:"right",prop:"",label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),t.deleteUserRow(e.$index,t.userList)}}},[a("i",{staticClass:"el-icon-delete"})])]}}],null,!1,659340457)})],1)],1):t._e(),a("div",{staticClass:"drawer"},[a("el-drawer",{ref:"drawer",attrs:{title:"添加用户",visible:t.multiUserDialog,direction:"rtl","custom-class":"demo-drawer"},on:{"update:visible":function(e){t.multiUserDialog=e}}},[a("div",{staticClass:"demo-drawer__content",staticStyle:{padding:"30px"}},[a("el-form",{attrs:{model:t.mutilUserForm}},[a("el-form-item",{attrs:{label:"账号","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_name",autocomplete:"off"},model:{value:t.mutilUserForm.account,callback:function(e){t.$set(t.mutilUserForm,"account",e)},expression:"mutilUserForm.account"}})],1),a("el-form-item",{attrs:{label:"密码","label-width":t.formLabelWidth}},[a("el-input",{attrs:{type:"password",id:"linked_url"},model:{value:t.mutilUserForm.passwd,callback:function(e){t.$set(t.mutilUserForm,"passwd",e)},expression:"mutilUserForm.passwd"}})],1)],1),a("div",{staticClass:"demo-drawer__footer",staticStyle:{float:"right"}},[a("el-button",{on:{click:t.cancelMultiUserDrawer}},[t._v("取 消")]),a("el-button",{attrs:{type:"primary"},on:{click:t.multiUserSubmit}},[t._v("确 定")])],1)],1)])],1),2==t.showMain?a("el-main",[a("h3",[t._v("如何完成数据迁移")]),a("el-steps",{staticStyle:{width:"500px","margin-bottom":"30px"},attrs:{active:t.active,"finish-status":"success"}},[a("el-step",{attrs:{title:"导出数据",description:"将现在的数据导出到本地"}}),a("el-step",{attrs:{title:"切换数据库",description:"更换新配置数据库连接信息并且初始化数据库"}}),a("el-step",{attrs:{title:"导入数据",description:"将本地的数据导入新的数据库"}})],1),a("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"将现在的数据导出到本地",placement:"bottom-end"}},[a("el-button",{attrs:{type:"primary",size:"small"},on:{click:t.exportDBData}},[t._v("导出数据库")])],1),a("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"更换新配置数据库连接信息并且初始化数据库",placement:"bottom-end"}},[a("router-link",{staticStyle:{margin:"0 10px"},attrs:{to:"/install"}},[a("el-button",{attrs:{type:"primary",size:"small"},on:{click:t.exportDBData}},[t._v("更改数据库配置")])],1)],1),a("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"将本地的数据导入新的数据库",placement:"bottom-end"}},[a("el-button",{attrs:{type:"primary",size:"small"},on:{click:t.exportDBData}},[t._v("导入历史数据")])],1)],1):t._e(),3==t.showMain?a("el-main",[a("el-form",{ref:"form",attrs:{model:t.config,"label-width":"100px"}},[a("el-form-item",{staticStyle:{width:"60%"},attrs:{label:"存储地址"}},[a("el-input",{attrs:{placeholder:"请输入地址，https://%s-***.cos.ap-***.myqcloud.com/"},model:{value:t.config.cos,callback:function(e){t.$set(t.config,"cos",e)},expression:"config.cos"}})],1),a("el-form-item",{staticStyle:{width:"60%"},attrs:{label:"AccessKey"}},[a("el-input",{attrs:{placeholder:"请输入SecretID",autocomplete:"off"},model:{value:t.config.cosSecretID,callback:function(e){t.$set(t.config,"cosSecretID",e)},expression:"config.cosSecretID"}})],1),a("el-form-item",{staticStyle:{width:"60%"},attrs:{label:"SecretKey"}},[a("el-input",{attrs:{type:"password",placeholder:"请输入SecretKey",autocomplete:"off"},model:{value:t.config.cosSecretKey,callback:function(e){t.$set(t.config,"cosSecretKey",e)},expression:"config.cosSecretKey"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary",size:"small"},on:{click:t.onSubmit}},[t._v("保存并初始化数据")])],1)],1)],1):t._e()],1)],1)],1)],1)},s=[],l=(a(561),a(8862),a(8932)),o={name:"SettingView",components:{HeadBar:l.Z},data:function(){return{searchVal:"",activeIndex:"99",showMain:1,multiUserDialog:!1,mutilUserForm:{account:"",passwd:""},userList:[{date:"2016-05-02",account:"wangxiaohu",uid:"1"},{date:"2016-05-04",account:"wangxiaohu",uid:"2"},{date:"2016-05-01",account:"wangxiaohu",uid:"3"},{date:"2016-05-03",account:"wangxiaohu",uid:"4"}],config:{}}},methods:{tableRowClassName:function(t){t.row;var e=t.rowIndex;return 1===e?"warning-row":3===e?"success-row":""},handleSelect:function(t,e){this.showMain=t},deleteRow:function(t,e){e.splice(t,1)},exportDBData:function(){console.log("export db data")},onSubmit:function(){var t=this;this.$store.dispatch("config/cosMigrate",JSON.stringify(this.config)).then((function(e){console.log(e),0==e.code&&t.$message.success("success")})).catch((function(t){console.log(t)}))},listUser:function(){var t=this;this.$store.dispatch("user/listUser",{}).then((function(e){console.log(e),t.userList=e.data.users})).catch((function(t){console.log(t)}))},deleteUserRow:function(t,e){var a=this;console.log(t,e),this.$store.dispatch("user/delUser","id="+e[t].Id).then((function(t){0==t.code?(a.multiUserDialog=!1,a.$message({message:"用户删除成功",type:"success"}),a.listUser()):a.$message({message:"用户删除失败",type:"warning"})})).catch((function(){a.$message({message:"用户删除失败",type:"warning"})}))},openMultiUserDrawer:function(){this.multiUserDialog=!0},cancelMultiUserDrawer:function(){this.multiUserDialog=!1},multiUserSubmit:function(){var t=this;this.$store.dispatch("user/addUser",JSON.stringify(this.mutilUserForm)).then((function(e){console.log(e),0==e.code?(t.multiUserDialog=!1,t.$message({message:"用户添加成功",type:"success"}),t.listUser()):t.$message({message:"用户添加失败",type:"warning"})})).catch((function(){t.$message({message:"用户添加失败",type:"warning"})}))}},created:function(){this.listUser()}},n=o,r=a(1001),c=(0,r.Z)(n,i,s,!1,null,"53c0ba8b",null),d=c.exports},2283:function(t,e,a){a.r(e),a.d(e,{default:function(){return c}});var i=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{staticClass:"signin"},[i("el-card",{staticClass:"signin_card"},[i("div",{staticClass:"signin_logo"},[i("router-link",{attrs:{to:"/"}},[i("img",{staticClass:"signin_logo",attrs:{src:a(6949),alt:""}})])],1),i("div",{staticClass:"signin_title"},[t._v("Sign in with your account")]),i("div",{staticClass:"signin_tip"},[t._v(" Please enter the correct account and password and log in ")]),i("div",[i("el-form",{ref:"loginForm",staticClass:"login-form",attrs:{model:t.loginForm,rules:t.loginRules,"auto-complete":"on","label-position":"left"}},[i("el-form-item",{attrs:{prop:"username"}},[i("el-input",{ref:"username",attrs:{type:"text",name:"username",placeholder:"请输入账号"},model:{value:t.loginForm.username,callback:function(e){t.$set(t.loginForm,"username",e)},expression:"loginForm.username"}})],1),i("el-form-item",{attrs:{prop:"password"}},[i("el-input",{key:t.passwordType,ref:"password",attrs:{name:"password",type:t.passwordType,placeholder:"请输入密码"},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.handleLogin.apply(null,arguments)}},model:{value:t.loginForm.password,callback:function(e){t.$set(t.loginForm,"password",e)},expression:"loginForm.password"}},[i("template",{slot:"suffix"},[i("i",{staticClass:"el-icon-view",on:{click:t.showPwd}})])],2)],1)],1)],1),i("div",{staticStyle:{"padding-top":"12px"}},[i("el-button",{staticStyle:{width:"100%"},attrs:{loading:t.loading,type:"primary"},nativeOn:{click:function(e){return e.preventDefault(),t.handleLogin.apply(null,arguments)}}},[t._v("Sign In")])],1),i("div",{staticClass:"back"},[i("router-link",{attrs:{to:"/"}},[i("i",{staticClass:"el-icon-s-home"}),t._v(" return to home page")])],1)])],1)},s=[],l=(a(1703),{name:"signin",data:function(){var t=function(t,e,a){e.length<1?a(new Error("Please enter the correct user name")):a()},e=function(t,e,a){e.length<6?a(new Error("The password can not be less than 6 digits")):a()};return{loginForm:{username:"",password:""},loginRules:{username:[{required:!0,trigger:"blur",validator:t}],password:[{required:!0,trigger:"blur",validator:e}]},loading:!1,passwordType:"password",redirect:void 0}},methods:{showPwd:function(){var t=this;"password"===this.passwordType?this.passwordType="":this.passwordType="password",this.$nextTick((function(){t.$refs.password.focus()}))},handleLogin:function(){var t=this;this.$refs.loginForm.validate((function(e){if(!e)return t.$message.error("登陆失败"),!1;t.loading=!0,t.$store.dispatch("user/login",t.loginForm).then((function(e){console.log(t.redirect),t.$router.push({path:t.redirect||"/"}),t.loading=!1})).catch((function(e){console.log(e),t.loading=!1}))}))}},watch:{$route:{handler:function(t){this.redirect=t.query&&t.query.redirect},immediate:!0}}}),o=l,n=a(1001),r=(0,n.Z)(o,i,s,!1,null,null,null),c=r.exports},6949:function(t,e,a){t.exports=a.p+"assets/img/logo.8fc44113.png"}}]);
//# sourceMappingURL=about-legacy.30432dcb.js.map