毕绅博客 GO-MVC-BLOG
===================


初学GO语言，这是我的第一个GO项目，可能有很多不足的地方，但如果你也是初学者，也许对你我都会有帮助。
这个项目基本实现了一个WEB编程的mvc架构。数据库使用了sqlite3，做博客方便迁移。视图使用了JADE，本人从NODE来学来的习惯，而且也不喜欢GO官方的模板。增加了一个ckeditor在线编辑器，支持GO的代码高亮。实现了简单的COOKIE权限验证。
核心框架：[martini](https://github.com/go-martini/martini)
配置文件：[ozzo-config](https://github.com/go-ozzo/ozzo-config)
数据ORM：[GORM](https://github.com/jinzhu/gorm)
视图VIEW：[GOJADE](https://github.com/zdebeer99/gojade)
 
DEMO演示：
[http://blog.bishen.org](http://blog.bishen.org)