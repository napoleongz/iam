//package main
//
//import "fmt"
//
//// 定义抽象结构体 Algorithm，包含三个抽象方法
//type Algorithm interface {
//	TemplateMethod()
//	Step1()
//	Step2()
//}
//
//// 具体实现结构体A
//type ConcreteAlgorithmA struct{}
//
//// 实现 Algorithm 接口的 TemplateMethod 方法
//func (a *ConcreteAlgorithmA) TemplateMethod() {
//	fmt.Println("ConcreteAlgorithmA TemplateMethod")
//	a.Step1() // 调用 Step1 方法
//	a.Step2() // 调用 Step2 方法
//}
//
//// 实现 Algorithm 接口的 Step1 方法
//func (a *ConcreteAlgorithmA) Step1() {
//	fmt.Println("ConcreteAlgorithmA Step1")
//}
//
//// 实现 Algorithm 接口的 Step2 方法
//func (a *ConcreteAlgorithmA) Step2() {
//	fmt.Println("ConcreteAlgorithmA Step2")
//}
//
//// 具体实现结构体B
//type ConcreteAlgorithmB struct{}
//
//// 实现 Algorithm 接口的 TemplateMethod 方法
//func (b *ConcreteAlgorithmB) TemplateMethod() {
//	fmt.Println("ConcreteAlgorithmB TemplateMethod")
//	b.Step1() // 调用 Step1 方法
//	b.Step2() // 调用 Step2 方法
//}
//
//// 实现 Algorithm 接口的 Step1 方法
//func (b *ConcreteAlgorithmB) Step1() {
//	fmt.Println("ConcreteAlgorithmB Step1")
//}
//
//// 实现 Algorithm 接口的 Step2 方法
//func (b *ConcreteAlgorithmB) Step2() {
//	fmt.Println("ConcreteAlgorithmB Step2")
//}
//
//func main() {
//	// 使用具体实现结构体A
//	algorithmA := &ConcreteAlgorithmA{} // 初始化具体结构体A的实例
//	algorithmA.TemplateMethod()         // 调用 ConcreteAlgorithmA 的 TemplateMethod 方法
//
//	fmt.Println()
//
//	// 使用具体实现结构体B
//	algorithmB := &ConcreteAlgorithmB{} // 初始化具体结构体B的实例
//	algorithmB.TemplateMethod()         // 调用 ConcreteAlgorithmB 的 TemplateMethod 方法
//}
