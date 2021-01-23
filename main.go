package main

import "html/template"

var tpl *template.Template

fun init(){
   tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
