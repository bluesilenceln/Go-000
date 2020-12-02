学习笔记

#### 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

#### 回答
个人认为dao层应该Wrap这个error，并且把sql.ErrNoRows转为统一的自定义错误码返回。因为dao层应该非透明向上抛错误，尽可能让上层和dao层更少依赖。