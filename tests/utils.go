package tests

// 没有泛型不太好 assert
// 这里的 want, got 类型不想用 interface
// 又不想每种类型都写一个 AssertTrue 方法，比如 AssertTrueInt(), AssertTrueString(), AssertTrueMap[int]int()..
//func AssertTrue(want, got ??)  {
////
////}

