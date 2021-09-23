(this.webpackJsonptodo=this.webpackJsonptodo||[]).push([[0],{34:function(t,e,n){},44:function(t,e,n){},64:function(t,e,n){"use strict";n.r(e);var r=n(1),c=n.n(r),a=n(11),o=n.n(a),i=(n(34),n(35),n(36),n(37),n(7)),s=function(t){return{type:"SET_DATA",data:t}},u=n(6),l=n.n(u),p=n(13),f=n(18),d=n(4);function h(){var t=Object(r.useState)(""),e=Object(f.a)(t,2),n=e[0],c=e[1],a=Object(i.b)(),o=Object(r.useRef)(null),u=Object(r.useRef)(null),h=Object(i.c)((function(t){return t.todoDataReducer})),v=null,b=Object(r.useState)([]),j=Object(f.a)(b,2),m=j[0],g=j[1];Object(r.useEffect)((function(){h.forEach((function(t){g(t)}))}),[h]),Object(r.useEffect)((function(){!function(t){switch(t){case"all":var e=[];h.forEach((function(t){t.forEach((function(t){!0===t.status?e.unshift(t):e.push(t)}))})),g(e);break;case"completed":var n=[];h.forEach((function(t){t.forEach((function(t){!1===t.status&&n.push(t)}))})),g(n);break;case"uncompleted":var r=[];h.forEach((function(t){t.forEach((function(t){!0===t.status&&r.push(t)}))})),g(r)}}(n)}),[n]);var O=function(){if(""!==o.current.value&&""!==u.current.value){var t={createdAt:(new Date).toString(),taskDescription:o.current.value,taskTitle:u.current.value,status:!0};!function(t){A.apply(this,arguments)}(t),g(m.concat(t)),o.current.value="",u.current.value=""}},x=function(t){t.target.style.opacity=.5,v=t.target,t.dataTransfer.effectAllowed="move"},k=function(t){t.preventDefault(),t.dataTransfer.dropEffect="move"},y=function(t){t.target.classList.add("over")},E=function(t){t.target.classList.remove("over")},D=function(t){if(t.stopPropagation(),v!==t.target){var e=m.filter((function(t,e){return e.toString()!==v.id})),n=m.filter((function(t,e){return e.toString()===v.id}))[0],r=Number(t.target.id),c=[];if(r>=e.length)c=e.slice(0).concat(n),g(c),t.target.classList.remove("over");else if(r<e.length){var a=(c=e.slice(0,r).concat(n)).concat(e.slice(r));g(a),t.target.classList.remove("over")}}else console.log("nothing happened");t.target.classList.remove("over")},T=function(t){t.target.style.opacity=1};function w(t){return t.split("").map((function(t){return t+"\u0336"})).join("")}function A(){return(A=Object(p.a)(l.a.mark((function t(e){var n;return l.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,fetch("http://6146ecde65467e00173849b9.mockapi.io/todoApi/task",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(e)});case 3:n=t.sent,a({type:"GET_DATA"}),console.log(n),t.next=11;break;case 8:t.prev=8,t.t0=t.catch(0),console.log(t.t0);case 11:case"end":return t.stop()}}),t,null,[[0,8]])})))).apply(this,arguments)}function N(){return(N=Object(p.a)(l.a.mark((function t(e,n){var r;return l.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,fetch("http://6146ecde65467e00173849b9.mockapi.io/todoApi/task/"+n,{method:"PUT",headers:{"Content-Type":"application/json"},body:JSON.stringify({status:e})});case 3:r=t.sent,console.log(r),t.next=10;break;case 7:t.prev=7,t.t0=t.catch(0),console.log(t.t0);case 10:case"end":return t.stop()}}),t,null,[[0,7]])})))).apply(this,arguments)}function S(){return(S=Object(p.a)(l.a.mark((function t(e){var n;return l.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,fetch("http://6146ecde65467e00173849b9.mockapi.io/todoApi/task/"+e,{method:"DELETE"});case 3:n=t.sent,console.log(n),t.next=10;break;case 7:t.prev=7,t.t0=t.catch(0),console.log(t.t0);case 10:case"end":return t.stop()}}),t,null,[[0,7]])})))).apply(this,arguments)}return Object(d.jsxs)("div",{className:"container",children:[Object(d.jsx)("h1",{style:{color:"white",textAlign:"center"},children:"Todo-List"}),Object(d.jsxs)("div",{id:"setInput",className:"form-group",children:[Object(d.jsx)("label",{style:{color:"white",fontWeight:"bolder"},htmlFor:"formGroupExampleInput",children:"SET TITLE"}),Object(d.jsx)("input",{type:"text",className:"form-control",id:"formGroupExampleInput",ref:o}),Object(d.jsx)("label",{style:{color:"white",fontWeight:"bolder"},htmlFor:"formGroupExampleInput",children:"SET DESCRIPTION"}),Object(d.jsx)("input",{type:"text",className:"form-control",id:"formGroupExampleInput",ref:u}),Object(d.jsx)("button",{id:"addbutton",className:"addButton",onClick:function(){O()},children:"ADD"})]}),Object(d.jsxs)("select",{id:"sort",onClick:function(t){return c(t.target.value)},children:[Object(d.jsx)("option",{id:"t",value:"all",children:"all"}),Object(d.jsx)("option",{value:"completed",children:"completed"}),Object(d.jsx)("option",{value:"uncompleted",children:"uncompleted"})]}),m.map((function(t,e){return Object(d.jsxs)("div",{className:"todo-list",children:[Object(d.jsx)("input",{className:"radio",type:"radio",value:t.status,onClick:function(){t.status=!1,function(t,e){N.apply(this,arguments)}(!1,t.id),w(t.taskDescription),a(s(m))}}),Object(d.jsx)("li",{id:e,type:"text",className:"input-item",draggable:"true",onDragStart:x,onDragOver:k,onDragEnter:y,onDragLeave:E,onDrop:D,onDragEnd:T,value:t.taskDescription,children:!1===t.status?w(t.taskDescription+" :- "+t.taskTitle):t.taskDescription+" :- "+t.taskTitle}),Object(d.jsx)("div",{id:e,className:"delButton",onClick:function(e){!function(t){t.preventDefault();var e=m.filter((function(e,n){return n!==Number(t.target.id)}));g(e)}(e),function(t){S.apply(this,arguments)}(t.id)},children:"X"})]},e)}))]})}n(44);var v=function(){var t=Object(i.b)();return Object(r.useEffect)((function(){t({type:"GET_DATA"})}),[t]),Object(d.jsx)("div",{className:"App",children:Object(d.jsx)(h,{})})},b=function(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:[],e=arguments.length>1?arguments[1]:void 0;switch(e.type){case"SET_DATA":return[e.data];case"ADD_DATA":default:return t}},j=n(8),m=Object(j.b)({todoDataReducer:b}),g=n(29),O=n(27),x=n(10),k=n(28),y=n.n(k);function E(){return y.a.request({method:"get",url:"http://6146ecde65467e00173849b9.mockapi.io/todoApi/task"})}var D=l.a.mark(T);function T(t){var e,n;return l.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,Object(x.a)(E);case 3:return e=t.sent,n=e,t.next=7,Object(x.b)(s(n.data));case 7:t.next=12;break;case 9:t.prev=9,t.t0=t.catch(0),console.log(t.t0);case 12:case"end":return t.stop()}}),D,null,[[0,9]])}var w=l.a.mark(A);function A(){return l.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,Object(O.a)("GET_DATA",T);case 2:case"end":return t.stop()}}),w)}var N=Object(g.a)(),S=[N],I=Object(j.d)(m,{},j.a.apply(void 0,S));N.run(A);var L=I;Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));o.a.render(Object(d.jsx)(c.a.StrictMode,{children:Object(d.jsx)(i.a,{store:L,children:Object(d.jsx)(v,{})})}),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(t){t.unregister()}))}},[[64,1,2]]]);
//# sourceMappingURL=main.0a4ce920.chunk.js.map