


function syscall_after_1(fd){
	
	return {}
}

function syscall_before_59(){
	
	const argv = hooks.getFdInfo()
	
	hooks.print(JSON.stringify(argv))		
	
	return {"ret": -1, "errno": 11}
}
