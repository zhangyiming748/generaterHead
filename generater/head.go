package generater

import (
	"fmt"
	"strings"
	"time"
)

func Head(fname, title, subtitle, author, catalog, tags string) (string, []string) {
	md := make([]string, 0)

	delimiter := "---"
	md = append(md, delimiter)
	layout := strings.Join([]string{"layout:", "post", "# 使用的布局(不需要改)"}, "\t")
	md = append(md, layout)
	title = strings.Join([]string{"title:", title, "标题"}, "\t")
	md = append(md, title)
	subtitle = strings.Join([]string{"subtitle:", subtitle, "# 副标题"}, "\t")
	md = append(md, subtitle)
	t := strings.Join([]string{time.Now().Format("2006-01-02 15:04:05"), "GMT+0800"}, " ")
	date := strings.Join([]string{"date:", t, "# 时间"}, "\t")
	md = append(md, date)
	author = strings.Join([]string{"author:", author, "# 作者"}, "\t")
	md = append(md, author)
	headimg := strings.Join([]string{"header-img:", "img/photo/birdAngle.webp", " # 这篇文章标题背景图片"}, "\t")
	md = append(md, headimg)
	catalog = strings.Join([]string{"catalog:", catalog, "# 是否归档"}, "\t")
	md = append(md, catalog)
	tag := strings.Join([]string{"tags:", "# 标签"}, "\t")
	md = append(md, tag)
	tags = strings.Join([]string{"-", tags}, " ")
	md = append(md, tags)
	md = append(md, delimiter)
	fmt.Println(t)
	titletime := time.Now().Format("2006-01-02")
	fname = strings.Join([]string{titletime, "-", fname, ".md"}, "")
	return fname, md
}
