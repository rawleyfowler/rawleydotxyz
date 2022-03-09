package utils

/*
Copyright (C) 2022 Rawley Fowler

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.Rawley Fowler, 2022
*/

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServePage(t string) gin.HandlerFunc {
	title := strings.Split(t, ".")[0]
	if title == "index" {
		title = ""
	} else {
		title = " | " + title
	}
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, t, struct{ Title string }{Title: title})
	}
}
