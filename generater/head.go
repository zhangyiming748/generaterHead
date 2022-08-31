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
	title = strings.Join([]string{"title:", title, "# 标题"}, "\t")
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
	tags = strings.Join([]string{"\t-", tags}, " ")
	md = append(md, tags)
	md = append(md, delimiter)

	titletime := time.Now().Format("2006-01-02")
	fname = strings.Join([]string{titletime, "-", fname, ".md"}, "")
	return fname, md
}
func MarketingGenerater(subject, event, otherword string) string {
	s := strings.Join([]string{subject, event, "是怎么回事呢?", subject, "相信大家都很熟悉,但是", subject, event, "是怎么回事呢,下面就让小编带大家一起了解吧<br>", subject, event, ",其实就是", otherword, ",大家可能会很惊讶", subject, "怎么会", event, "呢?但事实就是这样,小编也感到非常惊讶<br>这就是关于", subject, event, "的事情了,大家有什么想法呢,欢迎在评论区告诉小编一起讨论哦!"}, "")
	fmt.Println(s)
	return s
}
