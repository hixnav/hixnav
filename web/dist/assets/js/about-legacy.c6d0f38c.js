(self["webpackChunkHixNav"]=self["webpackChunkHixNav"]||[]).push([[443],{2984:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return u}});var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"account"}},[n("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),n("div",[t._v("开发中...")])],1)},r=[],o=(n(561),n(2372)),i={name:"AccountView",components:{HeadBar:o.Z},data:function(){return{searchVal:"",activeIndex:"5",tableData:[{site:"2016-05-02",account:"wangxiaohu",uid:"1"}]}},methods:{deleteRow:function(t,e){e.splice(t,1)}}},l=i,s=n(1001),c=(0,s.Z)(l,a,r,!1,null,"f6ca449c",null),u=c.exports},7105:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return d}});var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"add-link"}},[n("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),n("div",{staticStyle:{"padding-top":"20px"}},[n("el-form",{ref:"ruleForm",staticClass:"demo-ruleForm",attrs:{model:t.ruleForm,"status-icon":"",rules:t.rules,"label-width":"100px"}},[n("el-form-item",{attrs:{label:"标题",prop:"name"}},[n("el-input",{attrs:{type:"input",autocomplete:"off"},model:{value:t.ruleForm.name,callback:function(e){t.$set(t.ruleForm,"name",e)},expression:"ruleForm.name"}})],1),n("el-form-item",{attrs:{label:"图标",prop:"logo"}},[n("el-input",{attrs:{type:"input",autocomplete:"off"},model:{value:t.ruleForm.logo,callback:function(e){t.$set(t.ruleForm,"logo",e)},expression:"ruleForm.logo"}})],1),n("el-form-item",{attrs:{label:"简述",prop:"desc"}},[n("el-input",{attrs:{type:"input",autocomplete:"off"},model:{value:t.ruleForm.desc,callback:function(e){t.$set(t.ruleForm,"desc",e)},expression:"ruleForm.desc"}})],1),n("el-form-item",{attrs:{label:"地址",prop:"url"}},[n("el-input",{attrs:{type:"input",autocomplete:"off"},model:{value:t.ruleForm.url,callback:function(e){t.$set(t.ruleForm,"url",e)},expression:"ruleForm.url"}})],1),n("el-form-item",{attrs:{label:"分类ID",prop:"cate"}},[n("el-input",{attrs:{type:"number",autocomplete:"off"},model:{value:t.ruleForm.cate,callback:function(e){t.$set(t.ruleForm,"cate",e)},expression:"ruleForm.cate"}})],1),n("el-form-item",{attrs:{label:"分类名称",prop:"catename"}},[n("el-input",{attrs:{type:"input",autocomplete:"off"},model:{value:t.ruleForm.catename,callback:function(e){t.$set(t.ruleForm,"catename",e)},expression:"ruleForm.catename"}})],1),n("el-form-item",[n("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.submitForm("ruleForm")}}},[t._v("提交")]),n("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.submitFormWd("ruleForm")}}},[t._v("提交并查看")]),n("el-button",{on:{click:function(e){return t.resetForm("ruleForm")}}},[t._v("重置")])],1)],1)],1)],1)},r=[],o=(n(1703),n(3161),n(9653),n(8862),n(2372)),i=n(3932),l={name:"AddLink",components:{HeadBar:o.Z,FootBar:i.Z},data:function(){var t=this,e=function(t,e,n){if(!e)return n(new Error("年龄不能为空"));setTimeout((function(){Number.isInteger(e)?e<18?n(new Error("必须年满18岁")):n():n(new Error("请输入数字值"))}),1e3)},n=function(e,n,a){""===n?a(new Error("请输入密码")):(""!==t.ruleForm.checkPass&&t.$refs.ruleForm.validateField("checkPass"),a())},a=function(e,n,a){""===n?a(new Error("请再次输入密码")):n!==t.ruleForm.pass?a(new Error("两次输入密码不一致!")):a()};return{searchVal:"",activeIndex:"1s",ruleForm:{name:"",desc:"",url:"",logo:"",cate:0,catename:""},rules:{pass:[{validator:n,trigger:"blur"}],checkPass:[{validator:a,trigger:"blur"}],age:[{validator:e,trigger:"blur"}]}}},methods:{submitForm:function(t){var e=this;self=this,this.$refs[t].validate((function(t){if(!t)return console.log("error submit!!"),!1;console.log(e.ruleForm),e.ruleForm.cate=parseInt(e.ruleForm.cate),e.$store.dispatch("nav/addLink",JSON.stringify(e.ruleForm)).then((function(t){console.log(t),self.$notify({title:"成功",message:"手动刷新页面查看",type:"success"})})).catch((function(t){console.log(t),e.$message.error("失败")}))}))},submitFormWd:function(t){this.submitForm(t),this.$router.push("/")},resetForm:function(t){this.$refs[t].resetFields()}}},s=l,c=n(1001),u=(0,c.Z)(s,a,r,!1,null,"5f40f184",null),d=u.exports},2921:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return d}});var a=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"article"}},[a("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),a("div",{staticClass:"bg-banner"},[a("el-row",{staticStyle:{"margin-left":"0px","margin-right":"0px"},attrs:{gutter:16}})],1),a("div",{staticClass:"main"},[a("el-row",{staticClass:"el-row",attrs:{gutter:20}},[a("section",{},[a("section",{staticStyle:{display:"flex","flex-direction":"row","justify-content":"flex-start","align-items":"center"}},[a("img",{staticStyle:{display:"inline-block",width:"10px",height:"20px",padding:"0 10px 0","margin-top":"2px"},attrs:{src:n(3174),alt:""}}),a("div",{staticStyle:{"font-weight":"700","font-size":"18px"}},[t._v("文档")])])])]),a("el-row",{attrs:{gutter:20}},[a("div",{staticClass:"main_card"},[a("el-descriptions",{staticClass:"margin-top",attrs:{title:"",colon:!1,column:"4"}},[a("template",{slot:"extra"}),t._l(t.docLinks,(function(e){return a("el-descriptions-item",{key:e},[a("el-link",{attrs:{href:e.Url,target:"_blank"}},[a("el-avatar",{staticStyle:{width:"20px",height:"20px",margin:"0 10px","background-color":"#fff"},attrs:{src:e.Logo,size:"small"}}),t._v(" "+t._s(e.Name))],1)],1)}))],2)],1)]),a("el-row",{staticClass:"el-row",attrs:{gutter:20}},[a("el-col",{attrs:{span:2}},[a("section",{staticStyle:{display:"flex","flex-direction":"row","justify-content":"flex-start","align-items":"center","padding-top":"2rem"}},[a("img",{staticStyle:{display:"inline-block",width:"10px",height:"20px",padding:"0 10px 0","margin-top":"2px"},attrs:{src:n(3174),alt:""}}),a("div",{staticStyle:{"font-weight":"700","font-size":"18px"}},[t._v("链接")])])]),a("el-col",{staticStyle:{padding:"20px 0"},attrs:{span:20}},[a("el-row",{attrs:{gutter:20}},t._l(t.linkCates,(function(e){return a("el-button",{key:e,staticStyle:{display:"inline",margin:"0 20px"},attrs:{size:"small",round:""}},[t._v(t._s(e.Catename))])})),1)],1)],1),a("el-row",{attrs:{gutter:20}},[a("div",{staticClass:"main_card"},[a("el-descriptions",{staticClass:"margin-top",attrs:{title:"",colon:!1,column:"4"}},[a("template",{slot:"extra"}),t._l(t.commonLinks,(function(e){return a("el-descriptions-item",{key:e},[a("el-link",{attrs:{href:e.Url,target:"_blank"}},[a("el-avatar",{staticStyle:{width:"20px",height:"20px",margin:"0 10px","background-color":"#fff"},attrs:{src:e.Logo,size:"small"}}),t._v(" "+t._s(e.Name))],1)],1)}))],2)],1)])],1),a("el-button",{staticStyle:{position:"fixed",right:"40px",bottom:"108px","z-index":"9999"},attrs:{type:"primary",icon:"el-icon-edit",circle:""},on:{click:t.openAddDrawer}}),a("div",{staticClass:"drawer"},[a("el-drawer",{ref:"drawer",attrs:{title:"添加文链",visible:t.dialog,direction:"rtl","custom-class":"demo-drawer"},on:{"update:visible":function(e){t.dialog=e}}},[a("div",{staticClass:"demo-drawer__content",staticStyle:{padding:"30px"}},[a("el-form",{attrs:{model:t.form}},[a("el-form-item",{attrs:{label:"类别","label-width":t.formLabelWidth}},[a("el-select",{attrs:{placeholder:"类别"},model:{value:t.form.type,callback:function(e){t.$set(t.form,"type",e)},expression:"form.type"}},[a("el-option",{attrs:{label:"文档",value:"2"}}),a("el-option",{attrs:{label:"链接",value:"1"}})],1)],1),a("el-form-item",{attrs:{label:"标签","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_label",autocomplete:"off"},model:{value:t.form.catename,callback:function(e){t.$set(t.form,"catename",e)},expression:"form.catename"}})],1),a("el-form-item",{attrs:{label:"名称","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_name",autocomplete:"off"},model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}})],1),a("el-form-item",{attrs:{label:"地址","label-width":t.formLabelWidth}},[a("el-input",{attrs:{id:"linked_url"},model:{value:t.form.url,callback:function(e){t.$set(t.form,"url",e)},expression:"form.url"}})],1)],1),a("div",{staticClass:"demo-drawer__footer",staticStyle:{float:"right"}},[a("el-button",{on:{click:t.cancelForm}},[t._v("取 消")]),a("el-button",{attrs:{type:"primary"},on:{click:t.linkedSubmit}},[t._v("确 定")])],1)],1)])],1),a("FootBar")],1)},r=[],o=(n(4916),n(3123),n(8862),n(2372)),i=n(3932),l={name:"ArticleView",components:{HeadBar:o.Z,FootBar:i.Z},data:function(){return{searchVal:"",activeIndex:"2",linkCates:[],docLinks:[],commonLinks:[],dialog:!1,formLabelWidth:"80px",form:{name:"",type:1,url:""}}},methods:{openAddDrawer:function(){console.log("drawer"),this.dialog=!0},cancelForm:function(){this.dialog=!1},linkedSubmit:function(){var t=this,e=this.form.url,n=e.split("/"),a=n[2];this.form.logo=n[0]+"//"+a+"/favicon.ico",this.form.type=parseInt(this.form.type),console.log(this.form),self=this,this.$store.dispatch("link/addArticleLink",JSON.stringify(this.form)).then((function(t){console.log(t),self.dialog=!1,self.$notify({title:"成功",message:"手动刷新页面查看",type:"success"})})).catch((function(e){console.log(e),t.$message.error("失败")}))}},created:function(){self=this,this.$store.dispatch("link/article",{Type:2}).then((function(t){self.docLinks=t.data.links})).catch((function(t){console.log(t)})),this.$store.dispatch("link/article",{Type:1,Catename:""}).then((function(t){self.commonLinks=t.data.links})).catch((function(t){console.log(t)}))}},s=l,c=n(1001),u=(0,c.Z)(s,a,r,!1,null,"ba5ce0f0",null),d=u.exports},371:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return d}});var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"cloud-hos"}},[n("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),n("el-container",{staticClass:"el-container"},[n("el-aside",{attrs:{width:"200px"}},[n("el-menu",{staticClass:"el-menu-vertical-demo",attrs:{"default-active":t.param.active},on:{select:t.handleSelect}},[n("el-menu-item",{attrs:{index:"1"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("文件")])]),n("el-menu-item",{attrs:{index:"2"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("笔记")])]),n("el-menu-item",{attrs:{index:"3"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("脑图")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("视频")])]),n("el-menu-item",{attrs:{index:"5"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("图片")])])],1)],1),n("el-container",{staticClass:"main"},[n("el-header",[n("el-upload",{staticClass:"upload-file",attrs:{name:t.file,action:"/api/uploadIO",data:t.param,limit:1,"show-file-list":"false","auto-upload":!0,"http-request":t.uploadIO,"with-credentials":!0}},[n("el-button",{attrs:{size:"small",type:"primary"}},[t._v("点击上传")])],1)],1),n("el-main",[n("el-card",{staticClass:"box-card",attrs:{shadow:"never"}},[[n("el-table",{staticStyle:{width:"100%",height:"100%"},attrs:{data:t.fileLists}},[n("el-table-column",{attrs:{prop:"key",label:"文件名",width:"280"}}),n("el-table-column",{attrs:{prop:"size",label:"文件大小"}}),n("el-table-column",{attrs:{fixed:"right",label:"操作",width:"120"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"text",size:"small"},on:{click:function(n){return t.downFile(e.row)}}},[t._v("下载")]),n("el-button",{attrs:{type:"text",size:"small"},on:{click:function(n){return t.delFile(e.row)}}},[t._v("删除")])]}}])})],1)],n("el-pagination",{staticClass:"main-page",attrs:{background:"",layout:"prev, pager, next",total:120}})],2)],1)],1)],1)],1)},r=[],o=(n(5505),n(7714),n(2801),n(1174),n(1539),n(2472),n(8675),n(2990),n(8927),n(3105),n(5035),n(4345),n(7174),n(2846),n(4731),n(7209),n(6319),n(8867),n(7789),n(3739),n(9368),n(4483),n(2056),n(3462),n(678),n(7462),n(3824),n(5021),n(2974),n(5016),n(8783),n(3948),n(285),n(1637),n(4916),n(3123),n(2372)),i=n(3932),l=(n(9719).Base64,{name:"CloundHos",components:{HeadBar:o.Z,FootBar:i.Z},data:function(){return{searchVal:"",activeIndex:"3",file:"file",param:{active:"1"},fileLists:[]}},methods:{downFile:function(t){t.active=this.param.active,console.log(t),this.$store.dispatch("cloud/downIO",t).then((function(e){console.log(e);for(var n=window.atob(e.data.file),a=e.data.contentLength,r=new Uint8Array(a),o=0;o<a;++o)r[o]=n.charCodeAt(o);var i=window.URL.createObjectURL(new Blob([r],{type:e.data.contentType})),l=document.createElement("a");l.href=i,l.download=t.key,l.style.display="none",l.click(),l.remove()})).catch((function(t){console.log(t)}))},base64ToBlob:function(t){for(var e=t.split(";base64,"),n=e[0].split(":")[1],a=window.atob(e[1]),r=a.length,o=new Uint8Array(r),i=0;i<r;++i)o[i]=a.charCodeAt(i);return new Blob([o],{type:n})},delFile:function(t){var e=this;t.active=this.param.active,console.log(t),this.$store.dispatch("cloud/delIO",t).then((function(t){console.log(t),e.$notify({title:"成功",message:"操作成功",type:"success"}),location.reload()})).catch((function(t){console.log(t),e.$message.error("操作失败")}))},handleSelect:function(t,e){console.log(t,e),this.param.active=t,this.listFile()},listFile:function(){var t=this;this.$store.dispatch("cloud/listIO",this.param).then((function(e){console.log(e),t.fileLists=e.data.files})).catch((function(t){console.log(t)}))},uploadIO:function(t){var e=this,n=this,a=t.file,r=new window.FormData;r.append(n.file,a),r.append("active",this.param.active),this.$store.dispatch("cloud/uploadIO",r).then((function(t){console.log(t),n.$notify({title:"成功",message:"上传成功",type:"success"}),location.reload()})).catch((function(t){console.log(t),e.$message.error("上传失败")}))}},created:function(){this.listFile()}}),s=l,c=n(1001),u=(0,c.Z)(s,a,r,!1,null,"40bf894b",null),d=u.exports},4692:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return d}});var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"cloud-image"}},[n("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),n("section",{staticClass:"ci-main"},[n("el-row",[n("el-col",{attrs:{span:8,offset:8}},[n("el-upload",{ref:"upload",staticClass:"upload-demo",attrs:{drag:"","auto-upload":!1,"file-list":t.fileList,"list-type":"picture",action:"/api/upload",multiple:"",limit:3,"on-success":t.onSuccess}},[n("i",{staticClass:"el-icon-upload"}),n("div",{staticClass:"el-upload__text"},[t._v(" 将文件拖到此处，或"),n("em",[t._v("点击上传")])]),n("div",{staticClass:"el-upload__tip",attrs:{slot:"tip"},slot:"tip"},[n("span",[t._v("只能上传jpg/png文件，且不超过500kb")]),n("el-button",{staticStyle:{"margin-left":"10px"},attrs:{size:"small",type:"primary"},on:{click:t.submitUpload}},[t._v(" 开始上传 ")])],1)]),n("div",{staticStyle:{"margin-top":"15px"}},t._l(t.cloudImages,(function(e,a){return n("el-input",{key:e,ref:"demoInput",refInFor:!0,attrs:{placeholder:"http://"},model:{value:e.file,callback:function(n){t.$set(e,"file",n)},expression:"item.file"}},[n("el-button",{attrs:{slot:"append",icon:"el-icon-document-copy"},on:{click:function(e){return t.copyClicked(a)}},slot:"append"})],1)})),1)],1)],1)],1)],1)},r=[],o=n(2372),i=n(3932),l={name:"CloudImage",components:{HeadBar:o.Z,FootBar:i.Z},data:function(){return{searchVal:"",activeIndex:"4",fileList:[],cloudImages:[]}},methods:{submitUpload:function(){this.$refs.upload.submit()},handleRemove:function(t,e){console.log(t,e)},handlePreview:function(t){console.log(t)},onSuccess:function(t,e){console.log(t,e),this.cloudImages.push(t)},copyClicked:function(t){this.$refs.demoInput[t].select(),document.execCommand("copy")}}},s=l,c=n(1001),u=(0,c.Z)(s,a,r,!1,null,"6af08516",null),d=u.exports},8405:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return d}});var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"docs"}},[n("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),n("el-container",[n("el-aside",{attrs:{width:"200px"}},[n("el-menu",{staticClass:"el-menu-vertical-demo",attrs:{"default-active":"1-4-1",collapse:t.isCollapse},on:{open:t.handleOpen,close:t.handleClose}},[n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("Golang笔记")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("Linux笔记")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("开源项目")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("必会宝典")])])],1)],1),n("el-aside",{attrs:{width:"200px"}},[n("el-menu",{staticClass:"el-menu-vertical-demo",attrs:{"default-active":"1-4-1",collapse:t.isCollapse},on:{open:t.handleOpen,close:t.handleClose}},[n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("Golang笔记")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("Linux笔记")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("开源项目")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-folder"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("必会宝典")])])],1)],1),n("el-container",[n("el-main")],1)],1),n("FootBar")],1)},r=[],o=n(2372),i=n(3932),l={name:"Docs",components:{HeadBar:o.Z,FootBar:i.Z},data:function(){return[]},method:{}},s=l,c=n(1001),u=(0,c.Z)(s,a,r,!1,null,"3bbf943a",null),d=u.exports},3399:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return u}});var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"setting"}},[n("HeadBar",{attrs:{title:"海芯导航",searchVal:t.searchVal,activeIndex:t.activeIndex}}),n("div",{staticClass:"content"},[n("el-container",[n("el-aside",{attrs:{width:"200px"}},[n("el-menu",{staticClass:"el-menu-vertical",attrs:{"default-active":"1-4-1",collapse:t.isCollapse},on:{open:t.handleOpen,close:t.handleClose}},[n("el-menu-item",{attrs:{index:"1"}},[n("i",{staticClass:"el-icon-user-solid"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("多用户")])]),n("el-menu-item",{attrs:{index:"2"}},[n("i",{staticClass:"el-icon-user-solid"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("数据库配置")])]),n("el-menu-item",{attrs:{index:"3"}},[n("i",{staticClass:"el-icon-user-solid"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("存储配置")])]),n("el-menu-item",{attrs:{index:"4"}},[n("i",{staticClass:"el-icon-user-solid"}),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("数据迁移")])])],1)],1),n("el-container",[n("el-main",[n("el-button",{attrs:{type:"primary"}},[n("i",{staticClass:"el-icon-edit"}),t._v("添加用户")]),n("el-table",{staticStyle:{width:"100%","margin-top":"20px"},attrs:{data:t.tableData,"row-class-name":t.tableRowClassName}},[n("el-table-column",{attrs:{prop:"uid",label:"序号",width:"180"}}),n("el-table-column",{attrs:{prop:"account",label:"账号",width:"180"}}),n("el-table-column",{attrs:{prop:"date",label:"添加时间"}}),n("el-table-column",{attrs:{fixed:"right",prop:"",label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(n){return n.preventDefault(),t.deleteRow(e.$index,t.tableData)}}},[n("i",{staticClass:"el-icon-delete"})])]}}])})],1)],1)],1)],1)],1)],1)},r=[],o=(n(561),n(2372)),i={name:"SettingView",components:{HeadBar:o.Z},data:function(){return{searchVal:"",activeIndex:"99",tableData:[{date:"2016-05-02",account:"wangxiaohu",uid:"1"},{date:"2016-05-04",account:"wangxiaohu",uid:"2"},{date:"2016-05-01",account:"wangxiaohu",uid:"3"},{date:"2016-05-03",account:"wangxiaohu",uid:"4"}]}},methods:{tableRowClassName:function(t){t.row;var e=t.rowIndex;return 1===e?"warning-row":3===e?"success-row":""},deleteRow:function(t,e){e.splice(t,1)}}},l=i,s=n(1001),c=(0,s.Z)(l,a,r,!1,null,"53c09e3e",null),u=c.exports},2283:function(t,e,n){"use strict";n.r(e),n.d(e,{default:function(){return c}});var a=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"signin"},[a("el-card",{staticClass:"signin_card"},[a("div",{staticClass:"signin_logo"},[a("router-link",{attrs:{to:"/"}},[a("img",{staticClass:"signin_logo",attrs:{src:n(6949),alt:""}})])],1),a("div",{staticClass:"signin_title"},[t._v("Sign in with your account")]),a("div",{staticClass:"signin_tip"},[t._v(" Please enter the correct account and password and log in ")]),a("div",[a("el-form",{ref:"loginForm",staticClass:"login-form",attrs:{model:t.loginForm,rules:t.loginRules,"auto-complete":"on","label-position":"left"}},[a("el-form-item",{attrs:{prop:"username"}},[a("el-input",{ref:"username",attrs:{type:"text",name:"username",placeholder:"请输入账号"},model:{value:t.loginForm.username,callback:function(e){t.$set(t.loginForm,"username",e)},expression:"loginForm.username"}})],1),a("el-form-item",{attrs:{prop:"password"}},[a("el-input",{key:t.passwordType,ref:"password",attrs:{name:"password",type:t.passwordType,placeholder:"请输入密码"},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.handleLogin.apply(null,arguments)}},model:{value:t.loginForm.password,callback:function(e){t.$set(t.loginForm,"password",e)},expression:"loginForm.password"}},[a("template",{slot:"suffix"},[a("i",{staticClass:"el-icon-view",on:{click:t.showPwd}})])],2)],1)],1)],1),a("div",{staticStyle:{"padding-top":"12px"}},[a("el-button",{staticStyle:{width:"100%"},attrs:{loading:t.loading,type:"primary"},nativeOn:{click:function(e){return e.preventDefault(),t.handleLogin.apply(null,arguments)}}},[t._v("Sign In")])],1),a("div",{staticClass:"back"},[a("router-link",{attrs:{to:"/"}},[a("i",{staticClass:"el-icon-s-home"}),t._v(" return to home page")])],1)])],1)},r=[],o=(n(1703),{name:"signin",data:function(){var t=function(t,e,n){e.length<1?n(new Error("Please enter the correct user name")):n()},e=function(t,e,n){e.length<6?n(new Error("The password can not be less than 6 digits")):n()};return{loginForm:{username:"",password:""},loginRules:{username:[{required:!0,trigger:"blur",validator:t}],password:[{required:!0,trigger:"blur",validator:e}]},loading:!1,passwordType:"password",redirect:void 0}},methods:{showPwd:function(){var t=this;"password"===this.passwordType?this.passwordType="":this.passwordType="password",this.$nextTick((function(){t.$refs.password.focus()}))},handleLogin:function(){var t=this;this.$refs.loginForm.validate((function(e){if(!e)return t.$message.error("登陆失败"),!1;t.loading=!0,t.$store.dispatch("user/login",t.loginForm).then((function(e){console.log(t.redirect),t.$router.push({path:t.redirect||"/"}),t.loading=!1})).catch((function(e){console.log(e),t.loading=!1}))}))}},watch:{$route:{handler:function(t){this.redirect=t.query&&t.query.redirect},immediate:!0}}}),i=o,l=n(1001),s=(0,l.Z)(i,a,r,!1,null,null,null),c=s.exports},4170:function(t){for(var e="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",n={},a=0;a<66;a++)n[e.charAt(a)]=a;t.exports={itoc:e,ctoi:n}},3161:function(t,e,n){var a=n(2109),r=n(5988);a({target:"Number",stat:!0},{isInteger:r})},8145:function(t,e,n){"use strict";var a=n(3832),r=n(260).exportTypedArrayStaticMethod,o=n(7321);r("from",o,a)},2472:function(t,e,n){var a=n(9843);a("Uint8",(function(t){return function(e,n,a){return t(this,e,n,a)}}))},5505:function(t,e,n){var a=n(2109),r=n(5005),o=n(1702),i=n(7293),l=n(1340),s=n(2597),c=n(8053),u=n(4170).ctoi,d=/[^\d+/a-z]/i,f=/[\t\n\f\r ]+/g,p=/[=]+$/,m=r("atob"),h=String.fromCharCode,g=o("".charAt),v=o("".replace),b=o(d.exec),x=i((function(){return""!==atob(" ")})),y=!x&&!i((function(){m()}));a({global:!0,enumerable:!0,forced:x||y},{atob:function(t){if(c(arguments.length,1),y)return m(t);var e,n,a=v(l(t),f,""),o="",i=0,x=0;if(a.length%4==0&&(a=v(a,p,"")),a.length%4==1||b(d,a))throw new(r("DOMException"))("The string is not correctly encoded","InvalidCharacterError");while(e=g(a,i++))s(u,e)&&(n=x%4?64*n+u[e]:u[e],x++%4&&(o+=h(255&n>>(-2*x&6))));return o}})},9719:function(t,e,n){var a,r,o=n(7964)["default"];n(5505),n(7714),n(2801),n(1174),n(7042),n(1539),n(4747),n(8145),n(2472),n(8675),n(2990),n(8927),n(3105),n(5035),n(4345),n(7174),n(2846),n(4731),n(7209),n(6319),n(8867),n(7789),n(3739),n(9368),n(4483),n(2056),n(3462),n(678),n(7462),n(3824),n(5021),n(2974),n(5016),n(1249),n(4916),n(5306),n(1703),n(9714),n(9600),n(7601),n(7941),function(i,l){"object"===o(e)?t.exports=l():(a=l,r="function"===typeof a?a.call(e,n,e,t):a,void 0===r||(t.exports=r))}("undefined"!==typeof self?self:"undefined"!==typeof window?window:"undefined"!==typeof n.g&&n.g,(function(){"use strict";var t="3.7.2",e=t,n="function"===typeof atob,a="function"===typeof btoa,r="function"===typeof Buffer,o="function"===typeof TextDecoder?new TextDecoder:void 0,i="function"===typeof TextEncoder?new TextEncoder:void 0,l="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",s=Array.prototype.slice.call(l),c=function(t){var e={};return t.forEach((function(t,n){return e[t]=n})),e}(s),u=/^(?:[A-Za-z\d+\/]{4})*?(?:[A-Za-z\d+\/]{2}(?:==)?|[A-Za-z\d+\/]{3}=?)?$/,d=String.fromCharCode.bind(String),f="function"===typeof Uint8Array.from?Uint8Array.from.bind(Uint8Array):function(t,e){return void 0===e&&(e=function(t){return t}),new Uint8Array(Array.prototype.slice.call(t,0).map(e))},p=function(t){return t.replace(/=/g,"").replace(/[+\/]/g,(function(t){return"+"==t?"-":"_"}))},m=function(t){return t.replace(/[^A-Za-z0-9\+\/]/g,"")},h=function(t){for(var e,n,a,r,o="",i=t.length%3,l=0;l<t.length;){if((n=t.charCodeAt(l++))>255||(a=t.charCodeAt(l++))>255||(r=t.charCodeAt(l++))>255)throw new TypeError("invalid character found");e=n<<16|a<<8|r,o+=s[e>>18&63]+s[e>>12&63]+s[e>>6&63]+s[63&e]}return i?o.slice(0,i-3)+"===".substring(i):o},g=a?function(t){return btoa(t)}:r?function(t){return Buffer.from(t,"binary").toString("base64")}:h,v=r?function(t){return Buffer.from(t).toString("base64")}:function(t){for(var e=4096,n=[],a=0,r=t.length;a<r;a+=e)n.push(d.apply(null,t.subarray(a,a+e)));return g(n.join(""))},b=function(t,e){return void 0===e&&(e=!1),e?p(v(t)):v(t)},x=function(t){if(t.length<2){var e=t.charCodeAt(0);return e<128?t:e<2048?d(192|e>>>6)+d(128|63&e):d(224|e>>>12&15)+d(128|e>>>6&63)+d(128|63&e)}e=65536+1024*(t.charCodeAt(0)-55296)+(t.charCodeAt(1)-56320);return d(240|e>>>18&7)+d(128|e>>>12&63)+d(128|e>>>6&63)+d(128|63&e)},y=/[\uD800-\uDBFF][\uDC00-\uDFFFF]|[^\x00-\x7F]/g,w=function(t){return t.replace(y,x)},C=r?function(t){return Buffer.from(t,"utf8").toString("base64")}:i?function(t){return v(i.encode(t))}:function(t){return g(w(t))},k=function(t,e){return void 0===e&&(e=!1),e?p(C(t)):C(t)},_=function(t){return k(t,!0)},F=/[\xC0-\xDF][\x80-\xBF]|[\xE0-\xEF][\x80-\xBF]{2}|[\xF0-\xF7][\x80-\xBF]{3}/g,$=function(t){switch(t.length){case 4:var e=(7&t.charCodeAt(0))<<18|(63&t.charCodeAt(1))<<12|(63&t.charCodeAt(2))<<6|63&t.charCodeAt(3),n=e-65536;return d(55296+(n>>>10))+d(56320+(1023&n));case 3:return d((15&t.charCodeAt(0))<<12|(63&t.charCodeAt(1))<<6|63&t.charCodeAt(2));default:return d((31&t.charCodeAt(0))<<6|63&t.charCodeAt(1))}},A=function(t){return t.replace(F,$)},B=function(t){if(t=t.replace(/\s+/g,""),!u.test(t))throw new TypeError("malformed base64.");t+="==".slice(2-(3&t.length));for(var e,n,a,r="",o=0;o<t.length;)e=c[t.charAt(o++)]<<18|c[t.charAt(o++)]<<12|(n=c[t.charAt(o++)])<<6|(a=c[t.charAt(o++)]),r+=64===n?d(e>>16&255):64===a?d(e>>16&255,e>>8&255):d(e>>16&255,e>>8&255,255&e);return r},I=n?function(t){return atob(m(t))}:r?function(t){return Buffer.from(t,"base64").toString("binary")}:B,S=r?function(t){return f(Buffer.from(t,"base64"))}:function(t){return f(I(t),(function(t){return t.charCodeAt(0)}))},L=function(t){return S(Z(t))},E=r?function(t){return Buffer.from(t,"base64").toString("utf8")}:o?function(t){return o.decode(S(t))}:function(t){return A(I(t))},Z=function(t){return m(t.replace(/[-_]/g,(function(t){return"-"==t?"+":"/"})))},V=function(t){return E(Z(t))},U=function(t){if("string"!==typeof t)return!1;var e=t.replace(/\s+/g,"").replace(/={0,2}$/,"");return!/[^\s0-9a-zA-Z\+/]/.test(e)||!/[^\s0-9a-zA-Z\-_]/.test(e)},O=function(t){return{value:t,enumerable:!1,writable:!0,configurable:!0}},T=function(){var t=function(t,e){return Object.defineProperty(String.prototype,t,O(e))};t("fromBase64",(function(){return V(this)})),t("toBase64",(function(t){return k(this,t)})),t("toBase64URI",(function(){return k(this,!0)})),t("toBase64URL",(function(){return k(this,!0)})),t("toUint8Array",(function(){return L(this)}))},z=function(){var t=function(t,e){return Object.defineProperty(Uint8Array.prototype,t,O(e))};t("toBase64",(function(t){return b(this,t)})),t("toBase64URI",(function(){return b(this,!0)})),t("toBase64URL",(function(){return b(this,!0)}))},D=function(){T(),z()},H={version:t,VERSION:e,atob:I,atobPolyfill:B,btoa:g,btoaPolyfill:h,fromBase64:V,toBase64:k,encode:k,encodeURI:_,encodeURL:_,utob:w,btou:A,decode:V,isValid:U,fromUint8Array:b,toUint8Array:L,extendString:T,extendUint8Array:z,extendBuiltins:D,Base64:{}};return Object.keys(H).forEach((function(t){return H.Base64[t]=H[t]})),H}))},6949:function(t,e,n){"use strict";t.exports=n.p+"assets/img/logo.8fc44113.png"}}]);
//# sourceMappingURL=about-legacy.c6d0f38c.js.map