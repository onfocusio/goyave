(window.webpackJsonp=window.webpackJsonp||[]).push([[17],{389:function(t,s,a){"use strict";a.r(s);var n=a(25),e=Object(n.a)({},(function(){var t=this,s=t.$createElement,a=t._self._c||s;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"controllers"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#controllers"}},[t._v("#")]),t._v(" Controllers")]),t._v(" "),a("p"),a("div",{staticClass:"table-of-contents"},[a("ul",[a("li",[a("a",{attrs:{href:"#defining-controllers"}},[t._v("Defining controllers")])]),a("li",[a("a",{attrs:{href:"#handlers"}},[t._v("Handlers")])]),a("li",[a("a",{attrs:{href:"#naming-conventions"}},[t._v("Naming conventions")])])])]),a("p"),t._v(" "),a("h2",{attrs:{id:"defining-controllers"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#defining-controllers"}},[t._v("#")]),t._v(" Defining controllers")]),t._v(" "),a("p",[t._v("Controllers are files containing a collection of Handlers related to a specific feature. Each feature should have its own package. For example, if you have a controller handling user registration, user profiles, etc, you should create a "),a("code",[t._v("http/controller/user")]),t._v(" package. Creating a package for each feature has the advantage of cleaning up route definitions a lot and helps keeping a clean structure for your project.")]),t._v(" "),a("p",[t._v("Let's take a very simple CRUD as an example for a controller definition:\n"),a("strong",[t._v("http/controller/product/product.go")]),t._v(":")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Index")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("response "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\tproducts "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("model"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\tresult "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" database"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Conn")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Find")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("&")]),t._v("products"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("HandleDatabaseError")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("result"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\tresponse"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("JSON")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("http"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("StatusOK"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" products"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Show")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("response "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\tproduct "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" model"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\tresult "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" database"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Conn")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("First")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("&")]),t._v("product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Params"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("HandleDatabaseError")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("result"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\tresponse"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("JSON")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("http"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("StatusOK"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Store")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("response "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\tproduct "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" model"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\tName"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(":")]),t._v("  request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("String")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"name"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n\t\tPrice"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(":")]),t._v(" request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Numeric")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"price"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" database"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Conn")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Create")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("&")]),t._v("product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Error"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(";")]),t._v(" err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("!=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("nil")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\tresponse"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Error")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("err"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("else")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\tresponse"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("JSON")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("http"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("StatusCreated"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("map")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("string")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token builtin"}},[t._v("uint")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(":")]),t._v(" product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("ID"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Update")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("response "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\tproduct "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" model"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\tdb "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" database"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Conn")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\tresult "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" db"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Select")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("First")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("&")]),t._v("product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Params"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("HandleDatabaseError")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("result"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\t"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" db"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Model")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("&")]),t._v("product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Update")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"name"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("String")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"name"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Error"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(";")]),t._v(" err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("!=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("nil")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\t\tresponse"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Error")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("err"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Destroy")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("response "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\tproduct "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" model"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\tdb "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" database"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Conn")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\tresult "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" db"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Select")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("First")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("&")]),t._v("product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Params"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"id"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("HandleDatabaseError")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("result"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\t"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("if")]),t._v(" err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" db"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Delete")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("&")]),t._v("product"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Error"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(";")]),t._v(" err "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("!=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("nil")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n\t\t\tresponse"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Error")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("err"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n\t\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n\t"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("div",{staticClass:"custom-block tip"},[a("p",{staticClass:"custom-block-title"},[t._v("TIP")]),t._v(" "),a("ul",[a("li",[t._v("Learn how to handle database errors "),a("a",{attrs:{href:"https://gorm.io/docs/error_handling.html",target:"_blank",rel:"noopener noreferrer"}},[t._v("here"),a("OutboundLink")],1),t._v(".")]),t._v(" "),a("li",[t._v("It is not necessary to add "),a("code",[t._v("response.Status(http.StatusNoContent)")]),t._v(" at the end of "),a("code",[t._v("Update")]),t._v(" and "),a("code",[t._v("Destroy")]),t._v(" because the framework automatically sets the response status to 204 if its body is empty and no status has been set.")])])]),t._v(" "),a("h2",{attrs:{id:"handlers"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#handlers"}},[t._v("#")]),t._v(" Handlers")]),t._v(" "),a("p",[t._v("A "),a("code",[t._v("Handler")]),t._v(" is a "),a("code",[t._v("func(*goyave.Response, *goyave.Request)")]),t._v(". The first parameter lets you write a response, and the second contains all the information extracted from the raw incoming request.")]),t._v(" "),a("p",[t._v("Read about the available request information in the "),a("RouterLink",{attrs:{to:"/guide/basics/requests.html"}},[t._v("Requests")]),t._v(" section.")],1),t._v(" "),a("p",[t._v("Controller handlers contain the business logic of your application. They should be concise and focused on what matters for this particular feature in your application. For example, if you develop a service manipulating images, the image processing code shouldn't be written in controller handlers. In that case, the controller handler would simply pass the correct parameters to your image processor and write a response.")]),t._v(" "),a("div",{staticClass:"language-go extra-class"},[a("pre",{pre:!0,attrs:{class:"language-go"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("// This handler receives an image, optimizes it and sends the result back.")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token keyword"}},[t._v("func")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("OptimizeImage")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("response "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" request "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v("*")]),t._v("goyave"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("Request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v("\n    optimizedImg "),a("span",{pre:!0,attrs:{class:"token operator"}},[t._v(":=")]),t._v(" processing"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("OptimizeImage")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("request"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("File")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"image"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token number"}},[t._v("0")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n    response"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token function"}},[t._v("Write")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("(")]),t._v("http"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("StatusOK"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(",")]),t._v(" optimizedImg"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(")")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v("\n")])])]),a("div",{staticClass:"custom-block tip"},[a("p",{staticClass:"custom-block-title"},[t._v("TIP")]),t._v(" "),a("p",[t._v("Setting the "),a("code",[t._v("Content-Type")]),t._v(" header is not necessary. "),a("code",[t._v("response.Write")]),t._v(" automatically detects the content type and sets the header accordingly, if the latter has not been defined already.")])]),t._v(" "),a("h2",{attrs:{id:"naming-conventions"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#naming-conventions"}},[t._v("#")]),t._v(" Naming conventions")]),t._v(" "),a("ul",[a("li",[t._v("Controller packages are named after the model they are mostly using, in a singular form. For example a controller for a "),a("code",[t._v("Product")]),t._v(" model would be called "),a("code",[t._v("http/controller/product")]),t._v(". If a controller isn't related to a model, then give it an expressive name.")]),t._v(" "),a("li",[t._v("Controller handlers are always "),a("strong",[t._v("exported")]),t._v(" so they can be used when registering routes. All functions which aren't handlers "),a("strong",[t._v("must be unexported")]),t._v(".")]),t._v(" "),a("li",[t._v("CRUD operations naming and routing:")])]),t._v(" "),a("table",[a("thead",[a("tr",[a("th",[t._v("Method")]),t._v(" "),a("th",[t._v("URI")]),t._v(" "),a("th",[t._v("Handler name")]),t._v(" "),a("th",[t._v("Description")])])]),t._v(" "),a("tbody",[a("tr",[a("td",[a("code",[t._v("GET")])]),t._v(" "),a("td",[a("code",[t._v("/product")])]),t._v(" "),a("td",[a("code",[t._v("Index()")])]),t._v(" "),a("td",[t._v("Get the products list")])]),t._v(" "),a("tr",[a("td",[a("code",[t._v("POST")])]),t._v(" "),a("td",[a("code",[t._v("/product")])]),t._v(" "),a("td",[a("code",[t._v("Store()")])]),t._v(" "),a("td",[t._v("Create a product")])]),t._v(" "),a("tr",[a("td",[a("code",[t._v("GET")])]),t._v(" "),a("td",[a("code",[t._v("/product/{id}")])]),t._v(" "),a("td",[a("code",[t._v("Show()")])]),t._v(" "),a("td",[t._v("Show a product")])]),t._v(" "),a("tr",[a("td",[a("code",[t._v("PUT")]),t._v(" or "),a("code",[t._v("PATCH")])]),t._v(" "),a("td",[a("code",[t._v("/product/{id}")])]),t._v(" "),a("td",[a("code",[t._v("Update()")])]),t._v(" "),a("td",[t._v("Update a product")])]),t._v(" "),a("tr",[a("td",[a("code",[t._v("DELETE")])]),t._v(" "),a("td",[a("code",[t._v("/product/{id}")])]),t._v(" "),a("td",[a("code",[t._v("Destroy()")])]),t._v(" "),a("td",[t._v("Delete a product")])])])])])}),[],!1,null,null,null);s.default=e.exports}}]);