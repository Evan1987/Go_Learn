package counter

type alertCounter int  // 未公开类型，首字母小写

// function New 是公开的，返回一个alertCounter类型
func New(value int) alertCounter {
	return alertCounter(value)
}