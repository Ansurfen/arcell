package test

import "testing"

const in = `{
	-h: bool,
	-v: bool,
	{
		xx.lua
	},
    clone {
        {clone.js},
    },
    push {
        {
            push.lua
        },
        -a: bool,
        -r: string,
    },
    save {
        save.perl
    }, 
	run {
		push {
			push.lua
		}
	}
}`

// gits -h -v
// 得判断是否为bool
// 是bool就不需要传值
// gits push -a -r github
// 主要是激活函数 @ 需要flag触发
// 如果不是激活函数直接序列化
func TestParser(t *testing.T) {

}
