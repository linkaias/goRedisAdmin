(function(){"use strict";var e={5263:function(e,t,n){var r={};n.r(r),n.d(r,{reqAddVal:function(){return Z},reqDelKey:function(){return q},reqFlush:function(){return I},reqGetDbList:function(){return T},reqGetKeys:function(){return P},reqLogin:function(){return L}});var o=n(6369),i=function(){var e=this,t=e._self._c;return t("div",[t("el-container",[t("el-header",{staticClass:"el-header"},[t("HeaderItem")],1),t("el-main",[t("router-view")],1)],1)],1)},a=[],u=function(){var e=this,t=e._self._c;return t("div",[t("el-row",{attrs:{gutter:10}},[t("el-col",{attrs:{span:2,offset:6}},[t("div",{staticStyle:{"line-height":"68px","text-align":"center","font-size":"22px","font-weight":"bold"}},[t("span",[e._v(" GoRedisAdmin")])])]),t("el-col",{attrs:{span:10,offset:1}},[t("el-menu",{attrs:{"default-active":"1",mode:"horizontal"},on:{select:e.handleSelect}},[t("el-menu-item",{attrs:{index:"1"}},[e._v("数据中心")]),t("el-menu-item",{attrs:{index:"3"}},[e._v("设置")]),t("el-menu-item",{attrs:{index:"4"}},[t("svg",{staticClass:"icon",staticStyle:{height:"18px",width:"18px"},attrs:{t:"1666160259230",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"3458",width:"128",height:"128"}},[t("path",{attrs:{d:"M511.6 76.3C264.3 76.2 64 276.4 64 523.5 64 718.9 189.3 885 363.8 946c23.5 5.9 19.9-10.8 19.9-22.2v-77.5c-135.7 15.9-141.2-73.9-150.3-88.9C215 726 171.5 718 184.5 703c30.9-15.9 62.4 4 98.9 57.9 26.4 39.1 77.9 32.5 104 26 5.7-23.5 17.9-44.5 34.7-60.8-140.6-25.2-199.2-111-199.2-213 0-49.5 16.3-95 48.3-131.7-20.4-60.5 1.9-112.3 4.9-120 58.1-5.2 118.5 41.6 123.2 45.3 33-8.9 70.7-13.6 112.9-13.6 42.4 0 80.2 4.9 113.5 13.9 11.3-8.6 67.3-48.8 121.3-43.9 2.9 7.7 24.7 58.3 5.5 118 32.4 36.8 48.9 82.7 48.9 132.3 0 102.2-59 188.1-200 212.9 23.5 23.2 38.1 55.4 38.1 91v112.5c0.8 9 0 17.9 15 17.9 177.1-59.7 304.6-227 304.6-424.1 0-247.2-200.4-447.3-447.5-447.3z","p-id":"3459"}})]),t("a",{attrs:{href:"https://github.com/linkaias/goRedisAdmin",target:"_blank"}},[e._v(" Github")])]),t("el-menu-item",{attrs:{index:"4"}},[t("a",{attrs:{href:"https://github.com/linkaias/goRedisAdmin",target:"_blank"}},[e._v("文档")])]),t("el-menu-item",{attrs:{index:"5"}},[t("a",{attrs:{target:"_blank"},on:{click:e.logout}},[e._v(" 退出登录")])])],1)],1)],1)],1)},l=[],s=n(3505),c={name:"HeaderItem",methods:{handleSelect(e,t){console.log(e,t)},logout(){(0,s.FS)(),location.reload()}}},d=c,f=n(1001),p=(0,f.Z)(d,u,l,!1,null,"7dc7050e",null),h=p.exports,m={name:"App",components:{HeaderItem:h}},g=m,v=(0,f.Z)(g,i,a,!1,null,null,null),b=v.exports,y=n(2631),w=[{path:"*",redirect:"/home"},{path:"/home",component:()=>n.e(608).then(n.bind(n,608))},{path:"/login",component:()=>n.e(902).then(n.bind(n,5902))}];o["default"].use(y.Z);let k=y.Z.prototype.push,_=y.Z.prototype.replace;y.Z.prototype.push=function(e,t,n){t&&n?k.call(this,e,t,n):k.call(this,e,(()=>{}))},y.Z.prototype.replace=function(e,t,n){t&&n?_.call(this,e,t,n):_.call(this,e,(()=>{}),(()=>{}))};let x=new y.Z({routes:w,scrollBehavior(e,t,n){return{y:0}}});x.beforeEach((async function(e,t,n){let r=(0,s.YT)();console.log("token:",r),r?("/login"===e.path&&n("/"),n()):("/login"!==e.path&&n("/login"),n())}));var S=x,A=n(70),C=n(8499),E=n.n(C);let O=A.ZP.create({baseURL:"/api/v1/",timeout:5e3});O.interceptors.request.use((e=>(e.headers["Authorization"]="Bearer "+(0,s.YT)(),e))),O.interceptors.response.use((e=>{let t=e.data;return 0!==t.code&&200!==t.code?((0,C.Message)({message:t.message||"Error",type:"error",duration:5e3}),Promise.reject(new Error(t.message||"Error"))):t}),(e=>(console.log("err"+e),(0,C.Message)({message:e.message,type:"error",duration:5e3}),Promise.reject(e))));var j=O;const T=()=>j({url:"/db/db_list",method:"get"}),P=e=>j({url:`/db/get_keys?db_num=${e}`,method:"get"}),q=(e,t)=>j({url:`/db/key?db_num=${e}&key=${t}`,method:"delete"}),Z=e=>j({url:"/db/key",method:"post",data:e}),I=(e,t)=>j({url:`/db/flush?type=${e}&db_num=${t}`,method:"delete"}),L=e=>j({url:"/user/login",method:"post",data:e});var F={dbApi:r},N=n(407);o["default"].config.productionTip=!1,o["default"].prototype.$API=F,o["default"].use(E(),{locale:N["default"]}),new o["default"]({render:e=>e(b),router:S}).$mount("#app")},3505:function(e,t,n){n.d(t,{D0:function(){return r},FS:function(){return i},YT:function(){return o}});const r=e=>{localStorage.setItem("token",e)},o=()=>localStorage.getItem("token"),i=()=>{localStorage.removeItem("token")}}},t={};function n(r){var o=t[r];if(void 0!==o)return o.exports;var i=t[r]={id:r,loaded:!1,exports:{}};return e[r].call(i.exports,i,i.exports,n),i.loaded=!0,i.exports}n.m=e,function(){var e=[];n.O=function(t,r,o,i){if(!r){var a=1/0;for(c=0;c<e.length;c++){r=e[c][0],o=e[c][1],i=e[c][2];for(var u=!0,l=0;l<r.length;l++)(!1&i||a>=i)&&Object.keys(n.O).every((function(e){return n.O[e](r[l])}))?r.splice(l--,1):(u=!1,i<a&&(a=i));if(u){e.splice(c--,1);var s=o();void 0!==s&&(t=s)}}return t}i=i||0;for(var c=e.length;c>0&&e[c-1][2]>i;c--)e[c]=e[c-1];e[c]=[r,o,i]}}(),function(){n.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return n.d(t,{a:t}),t}}(),function(){n.d=function(e,t){for(var r in t)n.o(t,r)&&!n.o(e,r)&&Object.defineProperty(e,r,{enumerable:!0,get:t[r]})}}(),function(){n.f={},n.e=function(e){return Promise.all(Object.keys(n.f).reduce((function(t,r){return n.f[r](e,t),t}),[]))}}(),function(){n.u=function(e){return"js/"+e+"."+{608:"3005c261",902:"f73fe9a3"}[e]+".js"}}(),function(){n.miniCssF=function(e){return"css/"+e+"."+{608:"57b47934",902:"99e1412a"}[e]+".css"}}(),function(){n.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)}}(),function(){var e={},t="web:";n.l=function(r,o,i,a){if(e[r])e[r].push(o);else{var u,l;if(void 0!==i)for(var s=document.getElementsByTagName("script"),c=0;c<s.length;c++){var d=s[c];if(d.getAttribute("src")==r||d.getAttribute("data-webpack")==t+i){u=d;break}}u||(l=!0,u=document.createElement("script"),u.charset="utf-8",u.timeout=120,n.nc&&u.setAttribute("nonce",n.nc),u.setAttribute("data-webpack",t+i),u.src=r),e[r]=[o];var f=function(t,n){u.onerror=u.onload=null,clearTimeout(p);var o=e[r];if(delete e[r],u.parentNode&&u.parentNode.removeChild(u),o&&o.forEach((function(e){return e(n)})),t)return t(n)},p=setTimeout(f.bind(null,void 0,{type:"timeout",target:u}),12e4);u.onerror=f.bind(null,u.onerror),u.onload=f.bind(null,u.onload),l&&document.head.appendChild(u)}}}(),function(){n.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})}}(),function(){n.nmd=function(e){return e.paths=[],e.children||(e.children=[]),e}}(),function(){n.p="/"}(),function(){var e=function(e,t,n,r){var o=document.createElement("link");o.rel="stylesheet",o.type="text/css";var i=function(i){if(o.onerror=o.onload=null,"load"===i.type)n();else{var a=i&&("load"===i.type?"missing":i.type),u=i&&i.target&&i.target.href||t,l=new Error("Loading CSS chunk "+e+" failed.\n("+u+")");l.code="CSS_CHUNK_LOAD_FAILED",l.type=a,l.request=u,o.parentNode.removeChild(o),r(l)}};return o.onerror=o.onload=i,o.href=t,document.head.appendChild(o),o},t=function(e,t){for(var n=document.getElementsByTagName("link"),r=0;r<n.length;r++){var o=n[r],i=o.getAttribute("data-href")||o.getAttribute("href");if("stylesheet"===o.rel&&(i===e||i===t))return o}var a=document.getElementsByTagName("style");for(r=0;r<a.length;r++){o=a[r],i=o.getAttribute("data-href");if(i===e||i===t)return o}},r=function(r){return new Promise((function(o,i){var a=n.miniCssF(r),u=n.p+a;if(t(a,u))return o();e(r,u,o,i)}))},o={143:0};n.f.miniCss=function(e,t){var n={608:1,902:1};o[e]?t.push(o[e]):0!==o[e]&&n[e]&&t.push(o[e]=r(e).then((function(){o[e]=0}),(function(t){throw delete o[e],t})))}}(),function(){var e={143:0};n.f.j=function(t,r){var o=n.o(e,t)?e[t]:void 0;if(0!==o)if(o)r.push(o[2]);else{var i=new Promise((function(n,r){o=e[t]=[n,r]}));r.push(o[2]=i);var a=n.p+n.u(t),u=new Error,l=function(r){if(n.o(e,t)&&(o=e[t],0!==o&&(e[t]=void 0),o)){var i=r&&("load"===r.type?"missing":r.type),a=r&&r.target&&r.target.src;u.message="Loading chunk "+t+" failed.\n("+i+": "+a+")",u.name="ChunkLoadError",u.type=i,u.request=a,o[1](u)}};n.l(a,l,"chunk-"+t,t)}},n.O.j=function(t){return 0===e[t]};var t=function(t,r){var o,i,a=r[0],u=r[1],l=r[2],s=0;if(a.some((function(t){return 0!==e[t]}))){for(o in u)n.o(u,o)&&(n.m[o]=u[o]);if(l)var c=l(n)}for(t&&t(r);s<a.length;s++)i=a[s],n.o(e,i)&&e[i]&&e[i][0](),e[i]=0;return n.O(c)},r=self["webpackChunkweb"]=self["webpackChunkweb"]||[];r.forEach(t.bind(null,0)),r.push=t.bind(null,r.push.bind(r))}();var r=n.O(void 0,[998],(function(){return n(5263)}));r=n.O(r)})();