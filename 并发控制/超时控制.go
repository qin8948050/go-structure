//方法一 time.NewTicker
func method1(){
	close:=make(chan int,1)
	cleanup:=time.NewTicker(2*time.Second)
	defer cleanup.Stop()
	go func(){
		a:=func () int {
			return 1
		}()
		time.Sleep(3*time.Second)
		close<-a
	}()
	select {
		case result:=<-close:
			fmt.Println(result)
		case <-cleanup.C:
			fmt.Println("timeout")
	}
}


//方法二 context.WithTimeout
func method1(){
	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()
	//为了防止超时后导致goroutine不受控制内存泄漏情况
	c := make(chan int,1)   
	go func() {
		time.Sleep(3*time.Second)
		c <- 1
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	case res := <-c:
		fmt.Println("aaaa",res)
	}
}


//扩展
func job(){
	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()
	c := make(chan int,0)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case c<-rpc():
			}
		}


	}(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	case res := <-c:
		fmt.Println("aaaa",res)
	}
}


