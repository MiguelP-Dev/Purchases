package templates

import (
	"html/template"
)

// Funciones personalizadas para las plantillas
var Functions = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"mul": func(a, b int) int {
		return a * b
	},
	"div": func(a, b int) int {
		if b == 0 {
			return 0
		}
		return a / b
	},
} 