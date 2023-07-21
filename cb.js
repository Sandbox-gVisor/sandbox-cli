

function syscall_after_0(a, b, c, d, e, f){
	
	const ret = hooks.getFdInfo(a); 
	
	//if(hooks.)
	//hooks.print(JSON.stringify(ret));  
	//const err = {}
	//throw err
	
	
	return{};
}


function syscall_before_0(fd){
	
	const ret = hooks.getFdInfo(fd); 
	//hooks.print(fd, typeof(fd))
	
	//if(hooks.)
	const argv = hooks.getArgv()
	

	if(fd > 2){
		//hooks.print(JSON.stringify(argv))
		//hooks.print(JSON.stringify(ret));
	}  
	//const err = {}
	//throw err
	
	
	return{};
}


function syscall_after_59(a, b, c, d, e, f){
	
	const ret = hooks.getPidInfo(); 
	
	//hooks.print(JSON.stringify(args))
	//hooks.print("after")
	
	return {};
}


function syscall_before_59(a, b, c, d, e, f){
	
	const ret = hooks.getPidInfo(); 
	//hooks.print(JSON.stringify(args))
	hooks.print("before exxxxx")
	
	return {"ret": -1, "errn": 22};
}
